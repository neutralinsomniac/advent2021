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

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	length := 0
	entries := 0
	var bits []int

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if length == 0 {
			length = len(s)
			bits = make([]int, length)
		}
		for i := 0; i < length; i++ {
			if s[i] == '1' {
				bits[i] += 1
			}
		}
		entries += 1
	}

	answer := 0

	for i, val := range bits {
		if val > entries/2 {
			answer += int(math.Pow(2, float64(length-i-1)))
		}
	}
	answer *= answer ^ (int(math.Pow(2, float64(length))) - 1)

	fmt.Println(answer)
	fmt.Println("time:", time.Since(start))
}
