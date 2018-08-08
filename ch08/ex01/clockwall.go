package main

import (
	"fmt"
		"log"
	"net"
	"os"
	"regexp"
	"strconv"
			"bufio"
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
	timechan := make(chan string)
	for _, setting := range settings {
		fmt.Fprintf(os.Stdout, "Got setting: %v\n", setting)
		// TODO: sync.WaitGroup
		go connect(setting, timechan)
	}
	for {
		count := 1
		show(timechan, &count, len(settings))
	}
}

func show(timechan <-chan string, count *int, maxNum int) {
	time := <-timechan
	if (*count < maxNum) {
		fmt.Printf("%s ", time)
		fmt.Printf("(maxNum: %v, count: %v) ", maxNum, *count)
		*count = *count + 1
		fmt.Printf("(*count: %v)", *count)
	} else {
		fmt.Printf("%s\n", time)
		*count = 1
	}
}

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

func connect(setting clockSettings, timechan chan<- string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", setting.address, setting.port))
	fmt.Fprintln(os.Stdout, "connected to ", setting.address, setting.port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stdout, "error: %v", err)
		}
		//fmt.Println(string(line))
		timechan <- string(line)
	}
}
