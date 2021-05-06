package main

import (
	context "context"
	"fmt"
	"math"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Branch is an instance of a branch in the bank that handles
// customer requests as well as propogating changes to other branches
type Branch struct {
	ID               int                    // id of branch
	Port             int                    // port for the branch
	Balance          float64                // balance of all branches - eventually consistent
	BranchStubs      []BankClient           // stubs connecting to all the other branches
	branchConn       []*grpc.ClientConn     // branchConn to all the other branches, used for shutdown
	MessagesReceived []*MsgDeliveryResponse // all of the messages received by this process
	sync.RWMutex                            // lock to prevent concurrent access to balance
}

// NewBranch creates a new branch but does NOT spin up the process yet
func NewBranch(port int, id int, starting float64) *Branch {
	return &Branch{
		ID:               id,
		Port:             port,
		Balance:          starting,
		BranchStubs:      make([]BankClient, 0),
		MessagesReceived: make([]*MsgDeliveryResponse, 0),
	}

}

// AddBranch adds another branch to the list of branches to synchronize with
func (b *Branch) AddBranch(port int) error {
	b.Lock()
	defer b.Unlock()
	// connect with the branch on the specific port of that branch
	target := fmt.Sprintf("localhost:%d", port)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := NewBankClient(conn)

	// save stub
	b.BranchStubs = append(b.BranchStubs, client)
	// save connection for shutdown purposes
	b.branchConn = append(b.branchConn, conn)
	return nil

}

// Shutdown closes the connection between the different branches
func (b *Branch) Shutdown() {
	b.RLock()
	defer b.RUnlock()
	for _, c := range b.branchConn {
		c.Close()
	}
}

// GetMessages returns all the messages seen by the branch
func (b *Branch) GetMessages() Output {
	b.RLock()
	defer b.RUnlock()

	var messages []Message
	for i, m := range b.MessagesReceived {
		// just interested in the action performed
		// and the result (success/error)
		msg := Message{
			Interface: m.Interface.String(),
			Result:    m.Result.String(),
		}

		// only add the total balance to the last message
		if i+1 == len(b.MessagesReceived) {
			msg.Money = math.Round(b.Balance*100) / 100
			log.Debug(fmt.Sprintf("Server %d has final balance of $%.2f", b.ID, msg.Money))
		}

		messages = append(messages, msg)
	}

	result := Output{
		ID:       b.ID,
		Received: messages,
	}

	return result
}

// StartServerProcess starts the server in a separate goroutine and listens on a specified port
func (b *Branch) StartServerProcess() error {
	// listen and serve on a specific port
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", b.Port))
	if err != nil {
		return err
	}

	// startup gRPC server
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.MaxConcurrentStreams(20),
	)
	RegisterBankServer(s, b)

	// "go" starts the new goroutine, which will service the messages from the customer
	go s.Serve(server)
	log.Info(fmt.Sprintf("Server %d is serving on port %d with balance $%.2f", b.ID, b.Port, b.Balance))

	return nil
}

// MsgDelivery handles messages from both the clients
func (b *Branch) MsgDelivery(ctx context.Context, msg *MsgDeliveryRequest) (*MsgDeliveryResponse, error) {
	b.Lock()
	defer b.Unlock()
	result := Result_success
	var err error
	// only update the balance when it is a deposit or withdraw
	switch msg.Interface {
	case CustomerInterface_deposit:
		log.Debug(fmt.Sprintf("Server %d depositing $%.2f from a customer request", b.ID, msg.Money))
		b.deposit(msg.Money)
		// propogate deposit, but do this asynchronously since it can't error
		go b.propogateDeposit(msg.Money)
	case CustomerInterface_withdraw:
		log.Debug(fmt.Sprintf("Server %d withdrawing $%.2f from a customer request", b.ID, msg.Money))
		err = b.withdraw(msg.Money)
		if err != nil {
			// we can't withdraw that amount of money from this branch
			log.Error(fmt.Sprintf("Server %d could not withdraw %.2f from %.2f", b.ID, msg.Money, b.Balance))
			result = Result_error
		} else {
			// this branch can currently withdraw this kind of money,
			// since the branches are eventually consistent, it may not be okay with other branches
			// try to propogate withdraw and see if that is okay with the other servers
			err = b.propogateWithdraw(msg.Money)
			if err != nil {
				// could not propogate
				result = Result_error
			}
		}
	}

	// craft response
	response := &MsgDeliveryResponse{
		Result:    result,
		Interface: msg.Interface,
		Money:     b.Balance,
	}

	// save the response
	b.MessagesReceived = append(b.MessagesReceived, response)

	return response, err
}

// PropogateMsg handles messages from the other servers
func (b *Branch) PropogateMsg(ctx context.Context, msg *PropogateRequest) (*MsgDeliveryResponse, error) {
	b.Lock()
	defer b.Unlock()
	result := Result_success
	var err error
	switch msg.Interface {
	case BranchInterface_propogate_deposit:
		log.Debug(fmt.Sprintf("Server %d depositing $%.2f from a propogation request", b.ID, msg.Money))
		b.deposit(msg.Money)
	case BranchInterface_propogate_withdraw:
		// do not allow negative balance
		err = b.withdraw(msg.Money)
		if err != nil {
			log.Error(fmt.Sprintf("Server %d tried to withdraw $%.2f from a balance of $%.2f from a propogation request", b.ID, msg.Money, b.Balance))
			result = Result_error
		} else {
			log.Debug(fmt.Sprintf("Server %d withdrew $%.2f from a propogation request", b.ID, msg.Money))
		}
	}

	response := &MsgDeliveryResponse{
		Result: result,
		Money:  b.Balance,
	}
	return response, err
}

// withdraw withdraws money from the branch, propogates changes to other branches,
// and returns true on success
func (b *Branch) withdraw(money float64) error {
	if b.Balance-money < 0 {
		return status.Errorf(codes.InvalidArgument, "Not enough money to perform withdraw")
	}
	b.Balance = b.Balance - money
	return nil
}

// deposit adds money to the branch, propogates changes to other branches,
// and returns true on success
func (b *Branch) deposit(money float64) {
	b.Balance = b.Balance + money
}

// propogateWithdraw sends message to all the other branches regarding a
// withdraw that occurred on this branch
func (b *Branch) propogateWithdraw(money float64) error {
	// craft message to send to all other branches
	msg := &PropogateRequest{
		Money:     money,
		Interface: BranchInterface_propogate_withdraw,
	}

	for i, branch := range b.BranchStubs {
		// Send withdraw to all other branches
		err := b.sendPropogation(branch, msg)
		if err != nil {
			log.Error(fmt.Sprintf("Server %d had an error propogating withdraw, rolling back.", b.ID))
			b.deposit(money)
			// roll back all the previous withdraws we had done on other branches
			for h := 0; h < i; h++ {
				// deposit whatever we withdrew from other branches
				msg := &PropogateRequest{
					Money:     money,
					Interface: BranchInterface_propogate_deposit,
				}
				b.sendPropogation(b.BranchStubs[h], msg)
			}
			return err
		}
	}
	return nil
}

// propogateDeposit sends message to all the other branches regarding a
// deposit that occurred on this branch
func (b *Branch) propogateDeposit(money float64) error {
	// craft message to send to all other branches
	msg := &PropogateRequest{
		Money:     money,
		Interface: BranchInterface_propogate_deposit,
	}

	for _, branch := range b.BranchStubs {
		// Send deposit to all other branches
		b.sendPropogation(branch, msg)
	}

	return nil
}

// sendPropogation connects to the other branches using grpc and sends the update
func (b *Branch) sendPropogation(client BankClient, msg *PropogateRequest) error {
	// send message to other branches
	_, err := client.PropogateMsg(context.Background(), msg)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to propogate message: %v", err))
		return err
	}

	return nil
}
