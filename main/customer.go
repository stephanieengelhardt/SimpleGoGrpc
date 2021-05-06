package main

import (
	"context"
	"fmt"
	"math"
	sync "sync"

	"google.golang.org/grpc"
)

// Customer is an instance of a customer in the bank that sends
// messages to one specific branch
type Customer struct {
	ID               int                    // ID of the customer
	Events           []Event                // Events the customer is sending to the branch
	MessagesReceived []*MsgDeliveryResponse // Responses the customer got back from the branch
	Stub             BankClient             // Stub used to communicate with the branch
	conn             *grpc.ClientConn       // conn, used to shutdown
}

// NewCustomer creates a new customer with a given id and events to send to a branch
func NewCustomer(id int, events []Event) *Customer {
	return &Customer{
		ID:     id,
		Events: events,
	}
}

// AssignBranch assigns a branch to this customer and creates a stub to communicate on
func (c *Customer) AssignBranch(port int) {
	log.Info(fmt.Sprintf("Customer %d was assigned to server on port %d", c.ID, port))
	c.createStub(port)
}

// Shutdown closes down the connection with the branch
func (c *Customer) Shutdown() {
	c.conn.Close()
}

// ExecuteEvents spins off a goroutine for the the customer's events to be sent to a branch
func (c *Customer) ExecuteEvents(wg *sync.WaitGroup) {
	// "go" starts the goroutine which will send the events to the branch
	go c.startEvents(wg)
}

// helper methods
// startEvents is a helper method for ExecuteEvents and converts the input events into messages,
// sanitizes the input, and sends the message using the stub
func (c *Customer) startEvents(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, e := range c.Events {
		// convert from string in input file to interface used in grpc
		inter := CustomerInterface_unknown
		switch e.Interface {
		case "query":
			inter = CustomerInterface_query
		case "deposit":
			inter = CustomerInterface_deposit
		case "withdraw":
			inter = CustomerInterface_withdraw
		}

		// sanitize the input
		// money cannot be more than two decimals long
		money := e.Money
		moneyTemp := math.Round(money*100) / 100
		if moneyTemp != money {
			log.Warn(fmt.Sprintf("Attempted to send an invalid money amount ($%f), rounding to $%.2f", money, moneyTemp))
		}
		// if it is a negative number change the interface to the opposite and change money to positive
		moneyFinal := math.Abs(moneyTemp)
		if moneyFinal != moneyTemp {
			// if query, ignore this error
			switch inter {
			case CustomerInterface_deposit:
				inter = CustomerInterface_withdraw
				log.Warn(fmt.Sprintf("Attemped to use a negative number for money. Changing from deposit to withdraw."))
			case CustomerInterface_withdraw:
				inter = CustomerInterface_deposit
				log.Warn(fmt.Sprintf("Attemped to use a negative number for money. Changing from withdraw to deposit."))
			}
		}
		// craft message
		msg := &MsgDeliveryRequest{
			Money:     moneyFinal,
			Interface: inter,
		}
		// send message over the stub
		c.sendMsg(msg)
	}
}

// createStub is a helper method that creates the communication with the branch
func (c *Customer) createStub(port int) error {
	target := fmt.Sprintf("localhost:%d", port)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := NewBankClient(conn)

	// save stub
	c.Stub = client
	// also save connection to close at shutdown
	c.conn = conn
	return nil
}

// sendMsg uses the stub to send a message to the branch
func (c *Customer) sendMsg(msg *MsgDeliveryRequest) error {
	// send message to branch
	responseMessage, _ := c.Stub.MsgDelivery(context.Background(), msg)
	// save the response for debugging purposes
	c.MessagesReceived = append(c.MessagesReceived, responseMessage)

	return nil
}
