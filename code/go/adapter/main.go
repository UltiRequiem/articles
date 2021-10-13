package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash...")
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using Bank Account %d...\n", bankAccount)
}

type BankPaymentAdapter struct {
	bankAccount int
	BankPayment *BankPayment
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	ProcessPayment(&CashPayment{})
	ProcessPayment(&BankPaymentAdapter{5, &BankPayment{}})
}
