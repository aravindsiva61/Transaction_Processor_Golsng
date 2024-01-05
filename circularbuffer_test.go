package main

import (
	"fmt"
	"testing"
)

func TestCircularBuffer_Write(t *testing.T) {
	buffer := NewCircularBuffer(10)

	// Write 10 strings to the buffer.
	for i := 0; i < 10; i++ {
		buffer.Write(fmt.Sprint(i))
	}

	// Read the data from the buffer.
	for i := 0; i < 10; i++ {
		expected := fmt.Sprint(i)
		actual := buffer.Read()
		if expected != actual {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}

	// Write one more string to the buffer. This should overwrite the oldest data.
	buffer.Write("11")

	// Check oldest data is overwritten
	expected := "11"
	buffer.rpos = 0
	actual := buffer.Read()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func TestCircularBuffer_Read(t *testing.T) {
	buffer := NewCircularBuffer(10)

	// Write 10 strings to the buffer.
	for i := 0; i < 10; i++ {
		buffer.Write(fmt.Sprint(i))
	}

	// Read the data from the buffer.
	for i := 0; i < 10; i++ {
		expected := fmt.Sprint(i)
		actual := buffer.Read()
		if expected != actual {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}

	// Check correct data is fetched
	expected := "Empty Buffer"
	actual := buffer.Read()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
