package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

type clientInfo struct {
	c    client
	name string
}

var (
	entering = make(chan clientInfo)
	leaving  = make(chan clientInfo)
	messages = make(chan string) // すべてのクライアントから送信されるメッセージ
)

func broadcaster() {
	// クライアントのチャンネルと名前
	clients := make(map[client]string)
	for {
		select {
		// メッセージがきたら client チャンネルに送信する
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		// entering に クライアントがきたらマップに追加する
		case info := <-entering:
			clients[info.c] = info.name
			// 現在のクライアントリストを送信
			info.c <- "current clients:\n"
			for _, name := range clients {
				info.c <- name + "\n"
			}
		// leaving に クライアントがきたらマップから削除する
		// クライアントチャンネルをクローズする
		case info := <-leaving:
			delete(clients, info.c)
			close(info.c)
		}
	}
}

// チャンネルからメッセージを受け取ったら net.Conn に送信する
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func watchConn(conn net.Conn, tickerChan <-chan time.Time, info clientInfo) {
	for {
		select {
		case <-tickerChan:
			leaving <- info
			messages <- info.name + " has left\n"
		}
	}
}

// コネクションごとにクライアントのチャンネルを作成する
func handleConn(conn net.Conn) {
	// クライアントのチャンネル
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	info := clientInfo{ch, who}
	ch <- "You are " + who + "\n"
	messages <- who + " has arrived\n"
	entering <- info

	// 1行ごとにconnから読み込み
	input := bufio.NewScanner(conn)
	// TODO: for input.Scan がどこからどこまでなのか
	tickerChan := time.Tick(5 * time.Minute)
	go watchConn(conn, tickerChan, info)
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}

	leaving <- info
	messages <- who + " has left\n"
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
