package main

import (
	"encoding/json"
	"testing"
)

func TestTransactionProcessor_WriteToBuffer_Success(t *testing.T) {
	circularBufferSize := 2
	numReaders := 1
	circularBuffer := NewCircularBuffer(circularBufferSize)
	transactionProcessor := NewTransactionProcessor(circularBuffer, numReaders)

	// Prepare test data
	records := []Record{
		{Timestamp: 1, Value: "data1"},
		{Timestamp: 2, Value: "data2"},
	}
	data, _ := json.Marshal(records)

	// Call the WriteToBuffer method
	err := transactionProcessor.WriteToBuffer(data)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the circular buffer contains the expected values
	expectedValues := []string{"data1", "data2"}
	for _, expected := range expectedValues {
		actual := circularBuffer.Read()
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}

func TestTransactionProcessor_WriteToBuffer_Error(t *testing.T) {
	circularBufferSize := 2
	numReaders := 1
	circularBuffer := NewCircularBuffer(circularBufferSize)
	transactionProcessor := NewTransactionProcessor(circularBuffer, numReaders)

	// Prepare invalid JSON data
	data := []byte("invalid-json-data")

	// Call the WriteToBuffer method
	err := transactionProcessor.WriteToBuffer(data)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestTransactionProcessor_ReadFromBuffer(t *testing.T) {
	circularBufferSize := 2
	numReaders := 1
	circularBuffer := NewCircularBuffer(circularBufferSize)
	transactionProcessor := NewTransactionProcessor(circularBuffer, numReaders)

	// Prepare test data
	records := []Record{
		{Timestamp: 1, Value: "data1"},
		{Timestamp: 2, Value: "data2"},
		{Timestamp: 2, Value: "data3"},
		{Timestamp: 2, Value: "data4"},
	}
	data, _ := json.Marshal(records)

	// Call the WriteToBuffer method
	err := transactionProcessor.WriteToBuffer(data)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Call the ReadFromBuffer method
	transactionProcessor.ReadFromBuffer()
}
