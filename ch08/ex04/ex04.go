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
	"strings"
	"sync"
	"time"
)

func echo(c net.TCPConn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.TCPConn, wg *sync.WaitGroup) {
	fmt.Println("handleconn start")
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.CloseWrite()
	fmt.Println("handleconn closed")
	wg.Done()
}

//!-

func main() {
	serverAddr, addrErr := net.ResolveTCPAddr("tcp", "localhost:8000")
	if addrErr != nil {
		log.Fatal(addrErr)
	}
	l, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: カウンタがおかしくなる
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("waiting...")
		wg.Wait()
		fmt.Println("closing...")
		fmt.Println("closed")
	}()
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		wg.Add(1)
		wg.Done()
		go handleConn(conn, &wg)
	}
}
