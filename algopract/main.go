package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := sync.Mutex{}
	_ = mu

	times := 5
	value := 0

	go func() {
		for {
			time.Sleep(time.Millisecond)
			go func() {
				for {
					time.Sleep(time.Millisecond)
					value++
				}
			}()
		}
	}()

	time.Sleep(time.Second * time.Duration(times))
	fmt.Println(value)
}
