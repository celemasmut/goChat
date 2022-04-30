package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPAyment (p Payment) {
	p.Pay()
}

type BankPayment struct {}

func (BankPayment) Pay(bancAccount int)  {
	fmt.Printf("Paying using BancAccount %d\n",bancAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay()  {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main()  {
	cash := &CashPayment{}
	ProcessPAyment(cash)

	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPAyment(bpa)
}