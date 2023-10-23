package main 

import (
	"fmt"
)

type BankAccount interface {
	GetBalance() int 
	Deposit(amount int)
	WithDraw(amount int) error
}
