package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Input is format of the input file
type Input struct {
	ID      int     `json:"id"`
	Type    string  `json:"type"`
	Events  []Event `json:"events"`
	Balance float64 `json:"balance"`
}

// Event is the input events that the customer must perform
type Event struct {
	ID        int64   `json: "id"`
	Interface string  `json:"interface"`
	Money     float64 `json:"money"`
}

// Output is the output for a branch, it's id and the messages it answered
type Output struct {
	ID       int       `json:"id"`
	Received []Message `json:"recv"`
}

// Message is the output "recv" messages
type Message struct {
	Interface string  `json:"interface"`
	Result    string  `json:"result"`
	Money     float64 `json:"money,omitempty"`
}

// ReadInputFile reads in a .json file and unmarshals it into the structs
func ReadInputFile(fileName string) ([]*Input, error) {
	var result []*Input
	// open the file
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file: %v", err))
		return result, err
	}
	defer jsonFile.Close()

	// read the contents of the file
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading file: %v", err))
		return result, err
	}

	// convert it to the struct
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Error(fmt.Sprintf("Error unmarshalling input: %v", err))
		return result, err
	}

	return result, nil
}

// WriteToOutput writes all of the messages for every branch to the given filename
func WriteToOutput(fileName string, messages []Output) error {
	// marshal the information
	file, err := json.MarshalIndent(messages, "", " ")
	if err != nil {
		log.Error(fmt.Sprintf("Error marshalling output: %v", err))
		return err
	}

	// write it to the file
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing to output file: %v", err))
		return err
	}

	log.Info("Successfully created output.json.")
	return nil
}
