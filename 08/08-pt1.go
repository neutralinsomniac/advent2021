package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Entry struct {
	pattern []string
	output  []string
}

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []Entry
	easyCount := 0

	for scanner.Scan() {
		row := scanner.Text()

		var entry Entry
		sp := strings.Split(row, "|")
		entry.pattern = strings.Fields(sp[0])
		entry.output = strings.Fields(sp[1])
		for _, out := range entry.output {
			if (len(out) >= 2 && len(out) <= 4) || (len(out) == 7) {
				easyCount += 1
			}
		}
		entries = append(entries, entry)
	}

	fmt.Println(easyCount)
	fmt.Println("time:", time.Since(start))
}
