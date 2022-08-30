package semaphore

import (
	"fmt"
	"time"
)

type Semaphore struct {
	sem     chan int
	timeout time.Duration
}

func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- 0:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("не удалось захватить семафор")
	}
}

func (s *Semaphore) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("не удалось освободить семафор")
	}
}

func NewSemaphore(c int, timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan int, c),
		timeout: timeout,
	}
}
