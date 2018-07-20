package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

// TODO: インプットのパーサをつくる Timezone=address:port
// TODO: ポート名で起動する
// TODO: 1つのタイムゾーンで動くようにする

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}

type serverSettings struct {
	timezone string
	address  string
	port     int
}

func parseArgs(args []string) []serverSettings {
	re, _ := regexp.Compile("(.+?)=(.+?):(\\d+)")
	settings := make([]serverSettings, 0)
	for _, arg := range args {
		matches := re.FindAllStringSubmatch(arg, -1)[0][1:]
		port, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		setting := serverSettings{matches[0], matches[1], port}
		settings = append(settings, setting)
	}
	return settings
}
