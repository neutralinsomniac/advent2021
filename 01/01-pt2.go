package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}

	numInc := 0
	for i := range arr {
		if i < len(arr)-3 {
			j := i + 1
			if arr[j]+arr[j+1]+arr[j+2] > arr[i]+arr[i+1]+arr[i+2] {
				numInc += 1
			}
		}
	}
	fmt.Println(numInc)
	fmt.Println("time:", time.Since(start))
}
