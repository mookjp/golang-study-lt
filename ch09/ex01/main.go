package main

var deposits chan int        // 入金額を送信
var withdraws chan int       // 出金額を送信
var withdrawResult chan bool // 出金結果
var balances chan int        // 残高を受信

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawResult
}

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
		case amount := <-withdraws:
			tmpBalance := balance - amount
			//fmt.Printf("tmpBalance: %v\n", tmpBalance)
			if tmpBalance < 0 {
				withdrawResult <- false
			} else {
				balance = tmpBalance
				withdrawResult <- true
			}
		case balances <- balance:
		}
	}
}

func Init() {
	deposits = make(chan int)
	withdraws = make(chan int)
	withdrawResult = make(chan bool)
	balances = make(chan int)
	go teller()
}
