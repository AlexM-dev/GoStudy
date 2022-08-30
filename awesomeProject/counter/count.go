package counter

import (
	"sync"
)

type counter struct {
	maxGo    int
	activeGo int
	total    int
	C        chan int
	n        int
	Wg       *sync.WaitGroup
}

func NewCounter(n int, total int, c chan int, wg *sync.WaitGroup) *counter {
	return &counter{maxGo: n, total: total, C: c, activeGo: 0, n: 0, Wg: wg}
}

func (c *counter) NewGo() {
	if c.activeGo >= c.maxGo {
		return
	}
	c.activeGo++
	go c.increment()
}

func (c *counter) increment() {
	var f bool
	for {
		_, f = <-c.C
		if f == false {
			c.Wg.Done()
			return
		}
		c.n++
		if c.n == c.total {
			close(c.C)
		} else {
			c.C <- c.n
		}
	}
}

func (c *counter) Num() int {
	return c.n
}

func (c *counter) Total() int {
	return c.total
}
