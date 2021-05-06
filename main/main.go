package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	sync "sync"
	"time"

	"github.com/sirupsen/logrus"
)

// logLevel determines the verbosity of the logs. 0 means just info/error logs, 1 means all logs
var logLevel int

// log initializes the logging preferences for the whole project
var log = logrus.New()

func init() {
	// initialize logging
	logLevel := flag.Int("logLevel", 0, "defines the log level. 0=production builds. 1=dev builds.")
	flag.Parse()
	log.Out = os.Stdout
	switch *logLevel {
	case 0:
		log.SetLevel(logrus.InfoLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	// set the number of cores the different goroutines can run on
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read in input file
	input, err := ReadInputFile("input.json")
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not read input file: %+v", err))
		return
	}

	// uses the input file to get the customer and their events as well
	// as the branches and their starting balance
	customers, branches := processInput(input)

	// each customer must be able to be assigned to its own branch
	if len(customers) < len(branches) {
		log.Fatal(fmt.Sprintf("Invalid input. There must be at least one branch per customer."))
		return
	}

	// start up the branches by alerting the branches of each other,
	// spinning up the goroutines for each branch that are serving on separate ports,
	//  and assigning a customer to the branch
	startUpBranches(branches, customers)

	// spin up goroutines for each customer to send events in parallel
	startUpCustomers(customers)

	// write the results to a file
	err = WriteToOutput("output.json", getOutput(branches))
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not write to output file: %+v", err))
	}

	log.Info("Nothing more to process. Shutting down.")

	// close all connections and shutdown servers to free up the ports
	shutdown(branches, customers)
}

// processInput returns the total number of customers and the total number of branches
// in an input file
func processInput(input []*Input) ([]*Customer, []*Branch) {
	foundCustomers := 0
	foundBranches := 0
	var branches []*Branch
	var customers []*Customer

	// loop through everything given in the input file
	for _, q := range input {
		switch q.Type {
		case "customer":
			foundCustomers++
			customer := NewCustomer(foundCustomers, q.Events)
			customers = append(customers, customer)
		case "branch":
			foundBranches++
			branch := NewBranch(9079+foundBranches, foundBranches, q.Balance)
			branches = append(branches, branch)
		}
	}

	return customers, branches
}

// startUpBranches starts the branches on different goroutines serving on separate ports,
// sets up inter-branch communication, and assigns the customers to a branch
func startUpBranches(branches []*Branch, customers []*Customer) {
	for i, b1 := range branches {
		// alert the branches about all other branches
		// before beginning to serve
		for _, b2 := range branches {
			if b1.Port != b2.Port {
				b1.AddBranch(b2.Port)
			}
		}
		// start the server process on a separate goroutine
		b1.StartServerProcess()
		// assign each customer a branch
		customers[i].AssignBranch(b1.Port)
	}
}

// startUpCustomers starts the customer processes that begin sending
// events to their branches in parallel
func startUpCustomers(customers []*Customer) {
	// use wait groups so the Goroutine can alert the main thread
	// when the customer has sent all the events
	var wg sync.WaitGroup
	for _, c := range customers {
		wg.Add(1)
		// spin up a process for the customer
		// that way all customers are running at the same time
		c.ExecuteEvents(&wg)
		log.Debug(fmt.Sprintf("Starting process to send customer %d's events", c.ID))
	}

	// wait for all the customers to finish sending events
	wg.Wait()

	// allow all the server processes to synchronize due to the branches being eventually consistent
	time.Sleep(3 * time.Second)

	log.Info("All events sent from customers to branches.")
}

// getOutput loops through the branches and gets all of the messages they served and the results
func getOutput(branches []*Branch) []Output {
	var messages []Output
	for _, b := range branches {
		msg := b.GetMessages()
		messages = append(messages, msg)
	}
	return messages
}

// shutdown closes all of the open connections (customer->branch and branch->branch communication)
func shutdown(branches []*Branch, customers []*Customer) {
	for _, c := range customers {
		c.Shutdown()
	}

	for _, b := range branches {
		b.Shutdown()
	}
}
