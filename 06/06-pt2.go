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

func ReadInput(scanner *bufio.Scanner) []int {
	fish := make([]int, 9)
	scanner.Scan()
	row := scanner.Text()
	for _, s := range strings.Split(row, ",") {
		i, err := strconv.Atoi(s)
		checkErr(err)
		fish[i] += 1
	}

	return fish
}

func IterateFish(fish []int, numIterations int) []int {
	for i := 0; i < numIterations; i++ {
		swap := make([]int, 9)
		swap[8] = fish[0]
		swap[6] = fish[0]
		swap[7] = fish[8]
		swap[6] += fish[7]
		swap[5] = fish[6]
		swap[4] = fish[5]
		swap[3] = fish[4]
		swap[2] = fish[3]
		swap[1] = fish[2]
		swap[0] = fish[1]

		fish = swap
	}

	return fish
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

	fish := ReadInput(scanner)

	fish = IterateFish(fish, 256)

	sum := 0
	for _, i := range fish {
		sum += i
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))

}
