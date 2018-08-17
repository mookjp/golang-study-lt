// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	connectStatus := make(chan bool)
	go watch(c, connectStatus)

	for {
		if <-connectStatus == true {
			for input.Scan() {
				fmt.Fprint(os.Stdout, "got connect status true, echo\n")
				go echo(c, input.Text(), 1*time.Second)
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

// TODO: watch が timeout にいかない
func watch(c net.Conn, connectStatus chan bool) {
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Fprint(os.Stdout, "timeout\n")
			c.Close()
		default:
			fmt.Fprint(os.Stdout, "connect status true\n")
			connectStatus <- true
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
