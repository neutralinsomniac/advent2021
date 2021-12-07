package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func calcFuel(crabs []int, pos int) int {
	fuel := 0

	for _, v := range crabs {
		// avoid abs
		if pos > v {
			fuel += pos - v
		} else {
			fuel += v - pos
		}
	}

	return fuel
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

	scanner.Scan()

	row := scanner.Text()
	arrStr := strings.Split(row, ",")
	crabs := make([]int, len(arrStr))

	for i, s := range arrStr {
		v, _ := strconv.Atoi(s)
		crabs[i] = v
	}

	sort.Ints(crabs)

	/* median works here */
	pos := crabs[500]

	fuel := calcFuel(crabs, pos)

	fmt.Println(fuel)
	fmt.Println("time:", time.Since(start))
}
