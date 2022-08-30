package main

import (
	"fmt"
	"sem/semaphore"
	"sync"
	"time"
)

func main() {
	s := semaphore.NewSemaphore(1, 5*time.Second)
	//да, я придумал плохой пример для теста, ведь он работает только если семафор двоичный
	//ну, зато видно, как он не работает, если семафор не двоичный, это ведь тоже по своему показательно. Да ведь?(
	i := 0
	wg := sync.WaitGroup{}
	for k := 0; k < 10000; k++ {
		wg.Add(1)
		go func() {
			if err := s.Acquire(); err != nil {
				wg.Done()
				return
			}
			i++
			if err := s.Release(); err != nil {
				wg.Done()
				return
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(i)
}
