package sync

import "sync"

type Counter struct {
	mutex sync.Mutex
	val   int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
