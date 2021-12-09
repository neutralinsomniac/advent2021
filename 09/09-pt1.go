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

type Coord struct {
	x, y int
}

type World map[Coord]int

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	world := make(World)

	x := 0
	for scanner.Scan() {
		row := scanner.Text()

		for y := 0; y < len(row); y++ {
			coord := Coord{x, y}
			world[coord], _ = strconv.Atoi(string(row[y]))
		}
		x += 1
	}

	sum := 0
	for coord, check := range world {
		x := coord.x
		y := coord.y
		/* up */
		if adj, ok := world[Coord{x + 1, y}]; ok {
			if adj <= check {
				continue
			}
		}
		/* down */
		if adj, ok := world[Coord{x - 1, y}]; ok {
			if adj <= check {
				continue
			}
		}
		/* right */
		if adj, ok := world[Coord{x, y + 1}]; ok {
			if adj <= check {
				continue
			}
		}
		/* left */
		if adj, ok := world[Coord{x, y - 1}]; ok {
			if adj <= check {
				continue
			}
		}

		sum += 1 + check
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
