package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // すべてのクライアントから送信されるメッセージ
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		// メッセージがきたら client チャンネルに送信する
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		// entering に クライアントがきたらマップに追加する
		case cli := <-entering:
			clients[cli] = true
		// leaving に クライアントがきたらマップから削除する
		// クライアントチャンネルをクローズする
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// チャンネルからメッセージを受け取ったら net.Conn に送信する
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

// コネクションごとにクライアントのチャンネルを作成する
func handleConn(conn net.Conn) {
	// クライアントのチャンネル
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\n"
	messages <- who + " has arrived\n"
	entering <- ch

	// 1行ごとにconnから読み込み
	input := bufio.NewScanner(conn)
	// TODO: for input.Scan がどこからどこまでなのか
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}

	leaving <- ch
	messages <- who + "has left\n"
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
