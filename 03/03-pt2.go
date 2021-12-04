package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func filterO2AtPosition(entries [][]int, pos int) [][]int {
	/* first count number of bits in this position */
	numOnes := 0
	var newEntries [][]int

	for _, entry := range entries {
		numOnes += entry[pos]
	}
	numZeroes := len(entries) - numOnes

	filter := 0
	if numOnes >= numZeroes {
		filter = 1
	}
	/* now filter */
	for _, entry := range entries {
		if entry[pos] == filter {
			newEntries = append(newEntries, entry)
		}
	}

	return newEntries
}

func filterCO2AtPosition(entries [][]int, pos int) [][]int {
	/* first count number of bits in this position */
	numOnes := 0
	var newEntries [][]int

	for _, entry := range entries {
		numOnes += entry[pos]
	}
	numZeroes := len(entries) - numOnes

	filter := 0
	if numOnes < numZeroes {
		filter = 1
	}
	/* now filter */
	for _, entry := range entries {
		if entry[pos] == filter {
			newEntries = append(newEntries, entry)
		}
	}

	return newEntries
}

func binaryArrayToInt(bin []int) int {
	result := 0
	for i, val := range bin {
		if val == 1 {
			result += int(math.Pow(2, float64(len(bin)-i-1)))
		}
	}

	return result
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

	width := 0

	var entries [][]int

	/* load all the numbers up first */
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if width == 0 {
			width = len(s)
		}
		entry := make([]int, width)

		for i := 0; i < width; i++ {
			if s[i] == '1' {
				entry[i] = 1
			}
		}
		entries = append(entries, entry)
	}

	entriesBackup := make([][]int, len(entries))
	copy(entriesBackup, entries)

	o2num := 0
	co2num := 0

	/* now start the o2 selection process */
	for pos := 0; pos < width; pos++ {
		entries = filterO2AtPosition(entries, pos)
		if len(entries) == 1 {
			o2num = binaryArrayToInt(entries[0])
			break
		}
	}

	entries = entriesBackup

	/* now start the co2 selection process */
	for pos := 0; pos < width; pos++ {
		entries = filterCO2AtPosition(entries, pos)
		if len(entries) == 1 {
			co2num = binaryArrayToInt(entries[0])
			break
		}
	}

	fmt.Println(o2num * co2num)
	fmt.Println("time:", time.Since(start))
}
