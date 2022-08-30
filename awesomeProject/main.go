package main

import (
	"count/counter"
	"fmt"
	"sync"
)

var amountOfGoRoutines = 10
var total = 10100

func main() {
	ch := make(chan int, 1)
	ch <- 0
	var wg sync.WaitGroup
	wg.Add(amountOfGoRoutines)
	count := counter.NewCounter(amountOfGoRoutines, total, ch, &wg)

	for i := 0; i < amountOfGoRoutines; i++ {
		count.NewGo()
	}

	count.Wg.Wait()
	fmt.Println(count.Num())
}
