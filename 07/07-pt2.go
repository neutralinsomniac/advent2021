package main

import (
	"bufio"
	"fmt"
	"os"
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
			fuel += ((pos - v) * ((pos - v) + 1)) / 2
		} else {
			fuel += ((v - pos) * ((v - pos) + 1)) / 2
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

	min := int(^uint(0) >> 1)
	max := 0
	for i, s := range arrStr {
		v, _ := strconv.Atoi(s)
		crabs[i] = v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	bestFuel := int(^uint(0) >> 1) // max int
	for pos := min; pos < max; pos++ {
		fuel := calcFuel(crabs, pos)
		if fuel < bestFuel {
			bestFuel = fuel
		}
	}

	fmt.Println(bestFuel)
	fmt.Println("time:", time.Since(start))
}
