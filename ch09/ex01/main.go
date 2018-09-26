package main

var deposits chan int // 入金額を送信
var balances chan int // 残高を受信

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func Init() {
	deposits = make(chan int)
	balances = make(chan int)
	go teller()
}
