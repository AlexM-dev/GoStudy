package bank

import (
	"fmt"
	"sync"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int, mu *sync.RWMutex)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int, mu *sync.RWMutex) error

	// Balance returns clients balance
	Balance(mu *sync.RWMutex) int
}

type bankAccount struct {
	balance int
}

func NewBankAccount() *bankAccount {
	return &bankAccount{}
}

func (b *bankAccount) Deposit(amount int, mu *sync.RWMutex) {
	mu.Lock()
	b.balance += amount
	mu.Unlock()
}

func (b *bankAccount) Withdrawal(amount int, mu *sync.RWMutex) error {
	if b.balance < amount {
		return fmt.Errorf("insufficient balance, you cant withdrawal more than %v", b.balance)
	}
	mu.Lock()
	b.balance -= amount
	mu.Unlock()
	return nil
}

func (b *bankAccount) Balance(mu *sync.RWMutex) int {
	return b.balance
}
