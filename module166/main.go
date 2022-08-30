package main

import (
	"fmt"
	"math/rand"
	"sync"
	"task16/bank"
	"time"
)

func main() {
	acc := bank.NewBankAccount()
	mu := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Duration(int64(500+rand.Intn(1500)) * int64(time.Microsecond)))
			acc.Deposit(1+rand.Intn(10), &mu)
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Duration(int64(500+rand.Intn(1500)) * int64(time.Microsecond)))
			err := acc.Withdrawal(1+rand.Intn(5), &mu)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	var s string
	var err error

	for {
		fmt.Println("Enter command:")
		_, err = fmt.Scanln(&s)
		catchErr(err)
		switch s {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "balance":
			fmt.Println(acc.Balance(&mu))
		case "withdrawal":
			fmt.Println("Enter amount:")
			var x int
			_, err = fmt.Scanln(&x)
			catchErr(err)
			err = acc.Withdrawal(x, &mu)
			catchErr(err)
			if err != nil {
				fmt.Println(err)
			}
		case "deposit":
			fmt.Println("Enter amount:")
			var x int
			_, err = fmt.Scanln(&x)
			catchErr(err)
			acc.Deposit(x, &mu)
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}
}

func catchErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
