package counter

import "sync"

// Counter is a struct that takes a value and a mutex
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter is a wrapper of counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc is a method to imcrement the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value is a method to return the counter value
func (c *Counter) Value() int {
	return c.value
}
