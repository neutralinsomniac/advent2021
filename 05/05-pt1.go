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

type Coord struct {
	x, y int
}

type VentMap map[Coord]int

func ReadInput(scanner *bufio.Scanner) VentMap {
	var ventMap VentMap
	ventMap = make(VentMap)

	for scanner.Scan() {
		row := scanner.Text()
		end := strings.Index(row, ",")
		startX, _ := strconv.Atoi(row[:end])
		row = row[end+1:]
		end = strings.Index(row, " ")
		startY, _ := strconv.Atoi(row[:end])
		row = row[end+4:]
		end = strings.Index(row, ",")
		endX, _ := strconv.Atoi(row[:end])
		row = row[end+1:]
		endY, _ := strconv.Atoi(row)

		/* skip diagonals */
		if startX != endX && startY != endY {
			continue
		}

		if startX > endX {
			startX, endX = endX, startX
		}
		if startY > endY {
			startY, endY = endY, startY
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				ventMap[Coord{x, y}] = ventMap[Coord{x, y}] + 1
			}
		}
	}

	return ventMap
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

	ventMap := ReadInput(scanner)

	numOverlaps := 0
	for _, val := range ventMap {
		if val >= 2 {
			numOverlaps += 1
		}
	}
	fmt.Println(numOverlaps)
	fmt.Println("time:", time.Since(start))

}
