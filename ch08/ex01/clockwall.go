package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"text/tabwriter"
)

type clockSettings struct {
	timezone string
	address  string
	port     int
	time     string
}

func main() {
	if (len(os.Args)) < 2 {
		fmt.Println("You should specify server settings e.g. Asia/Tokyo=localhost:8010")
	}
	settings := parseArgs(os.Args[1:])
	for _, setting := range settings {
		fmt.Fprintf(os.Stdout, "Got setting: %v\n", setting)
		connect(&setting)
	}
	//for {
	//	showClock(settings)
	//	time.Sleep(1 * time.Second)
	//}
}

func showClock(settings []clockSettings) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	// Timezone
	for i, setting := range settings {
		if i != len(settings)-1 {
			fmt.Fprintf(w, "%s\t", setting.timezone)
		} else {
			fmt.Fprintf(w, "%s", setting.timezone)
		}
	}
	// Time
	for i, setting := range settings {
		if i != len(settings)-1 {
			fmt.Fprintf(w, "%s\t", setting.time)
		} else {
			fmt.Fprintf(w, "%s", setting.time)
		}
	}
	fmt.Fprintln(w)
	w.Flush()
}

// TODO: クライアントは受け取ったら時間を保持する

func parseArgs(args []string) []clockSettings {
	re, _ := regexp.Compile("(.+?)=(.+?):(\\d+)")
	settings := make([]clockSettings, 0)
	for _, arg := range args {
		matches := re.FindAllStringSubmatch(arg, -1)[0][1:]
		port, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		setting := clockSettings{matches[0], matches[1], port, ""}
		settings = append(settings, setting)
	}
	return settings
}

func connect(setting *clockSettings) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", setting.address, setting.port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// これは標準出力に流れる
	//mustCopy(os.Stdout, conn)

	// これは流れない
	// io.CopyがReaderがEOFを返すまで終了しないため動かない
	buf := new(bytes.Buffer)
	mustCopy(buf, conn)
	fmt.Println(buf.String())
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
