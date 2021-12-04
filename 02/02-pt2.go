package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	d := 0
	x := 0
	aim := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if s[0] == 'f' {
			tmp := 0
			tmp, err = strconv.Atoi(s[8:])
			checkErr(err)
			x += tmp
			d += aim * tmp
		} else if s[0] == 'd' {
			tmp := 0
			tmp, err = strconv.Atoi(s[5:])
			checkErr(err)
			aim += tmp
		} else if s[0] == 'u' {
			tmp := 0
			tmp, err = strconv.Atoi(s[3:])
			checkErr(err)
			aim -= tmp
		}
	}
	fmt.Println(d * x)
	fmt.Println("time:", time.Since(start))
}
