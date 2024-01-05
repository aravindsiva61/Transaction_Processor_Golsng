package main

import (
	"sync"
)

// Circular buffer representing fields required for implementing it
type CircularBuffer struct {
	data    []string      // data holding json contents
	size    int           // size of circular buffer
	rpos    int           // Read position
	wpos    int           // Write position
	rwmutex *sync.RWMutex // Read Write Mutex
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		data:    make([]string, size),
		size:    size,
		rpos:    0,
		wpos:    0,
		rwmutex: &sync.RWMutex{},
	}
}

func (c *CircularBuffer) Write(data string) {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()

	if c.wpos == c.size { // Overwrite the oldest data.
		c.wpos = 0
		c.data[c.wpos] = data
	} else { // Write the data to the buffer.
		c.data[c.wpos] = data
		c.wpos++
	}
}

func (c *CircularBuffer) Read() string {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()

	if c.rpos == c.size { // Empty buffer
		c.rpos = 0
		return "Empty Buffer"
	} else { // Empty buffer
		data := c.data[c.rpos]
		c.rpos++
		return data
	}
}
