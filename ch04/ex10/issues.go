package main

import (
	"os"

	"fmt"
	"log"

	"time"

	"math"

	"gopl.io/ch4/github"
)

type filterType int

const (
	overAnYear filterType = iota
	inAnYear
	inAMonth
)

const (
	anYearInSeconds float64 = 31536000
	monthInSeconds  float64 = 2592000
)

func main() {
	now := time.Now()

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	filteredInAMonth := filter(inAMonth, result.Items, now)
	filteredInAnYear := filter(inAnYear, result.Items, now)
	filteredOverAnYear := filter(overAnYear, result.Items, now)

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("\n===== in a month =====\n")
	for _, item := range filteredInAMonth {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("\n===== in an year =====\n")
	for _, item := range filteredInAnYear {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("\n===== over an year =====\n")
	for _, item := range filteredOverAnYear {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func filter(fType filterType, items []*github.Issue, now time.Time) []github.Issue {
	res := make([]github.Issue, 0)

	for _, item := range items {

		diff := diff(fType, now, item.CreatedAt)
		if diff {
			res = append(res, *item)
		}
	}

	return res
}

func diff(fType filterType, now time.Time, createdAt time.Time) bool {
	switch fType {
	case overAnYear:
		return math.Abs(createdAt.Sub(now).Seconds()) >= anYearInSeconds
	case inAnYear:
		return math.Abs(createdAt.Sub(now).Seconds()) <= anYearInSeconds
	case inAMonth:
		return math.Abs(createdAt.Sub(now).Seconds()) <= monthInSeconds
	default:
		return math.Abs(createdAt.Sub(now).Seconds()) >= anYearInSeconds
	}
}
