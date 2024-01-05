package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// Record represents a single record in the time series JSON feed
type Record struct {
	Timestamp int64  `json:"timestamp"` //Used to represent timestamp value
	Value     string `json:"value"`     // Used to represent transaction record
}

// TransactionProcessor represents main structure used to drive entire funtionality
type TransactionProcessor struct {
	circularBuffer *CircularBuffer // Circular buffer to store values
	readers        int             // Number of reader goroutines
	wg             sync.WaitGroup  // WaitGroup to coordinate goroutines
}

// Function to instantiate Transaction processor
func NewTransactionProcessor(circularBuffer *CircularBuffer, numReaders int) *TransactionProcessor {
	return &TransactionProcessor{
		circularBuffer: circularBuffer,
		readers:        numReaders,
		wg:             sync.WaitGroup{},
	}
}

// Function to read, unmarshal JSON contents and write to buffer
func (tp *TransactionProcessor) WriteToBuffer(data []byte) error {

	var records []Record
	unmarshalErr := json.Unmarshal(data, &records) // Unmarshal JSON data into a slice of Record structs
	if unmarshalErr != nil {
		fmt.Println("Error in unmarshalling JSON contents")
		return unmarshalErr
	}

	for _, record := range records {
		tp.circularBuffer.Write(record.Value) // Write value to the circular buffer
	}

	return nil
}

// Function to read from buffer using multiple readers
func (tp *TransactionProcessor) ReadFromBuffer() {
	for i := 0; i < tp.readers; i++ {
		tp.wg.Add(1) // Add 1 to the WaitGroup for each reader goroutine
		go func(i int) {
			for {
				fmt.Println("Reader Id - ", i)   // Print the ID of the current reader goroutine
				data := tp.circularBuffer.Read() // Read data from the circular buffer

				if data == "Empty Buffer" { // If the buffer is empty, break the loop
					break
				}

				fmt.Println(" Data - ", data) // Print the data read from the buffer
			}

			// Tell the WaitGroup that this goroutine is done.
			defer tp.wg.Done()
		}(i) // Pass the current iteration value as the goroutine ID
	}

	// Wait for all the read goroutines to finish.
	tp.wg.Wait()
}

func main() {

	circularBufferSize := 10 // Size of the circular buffer
	numReaders := 3          // Number of reader goroutines

	circularBuffer := NewCircularBuffer(circularBufferSize)                     // Create a new circular buffer
	transactionProcessor := NewTransactionProcessor(circularBuffer, numReaders) // Create a new transaction processor

	file, err := os.Open("data.json") // Open the file
	if err != nil {
		fmt.Println("Error in opening file")
		return
	}
	defer file.Close() // Defer closing the file to ensure it's closed after processing

	data, err := ioutil.ReadAll(file) // Read the file contents
	if err != nil {
		fmt.Println("Error in reading file")
		return
	}

	writeErr := transactionProcessor.WriteToBuffer(data) // Process the file and populate the circular buffer
	if writeErr != nil {
		fmt.Println("Error in writing to Buffer")
		return
	}

	transactionProcessor.ReadFromBuffer() // Start reading from the circular buffer using reader goroutines
}
