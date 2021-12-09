package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var alreadyCounted map[Coord]bool

func getAdjacentsToConsider(world World, coord Coord) []Coord {
	adjacents := make([]Coord, 0)
	x := coord.x
	y := coord.y
	/* up */
	if adj, ok := world[Coord{x + 1, y}]; ok {
		if adj != 9 {
			if _, ok := alreadyCounted[Coord{x + 1, y}]; !ok {
				adjacents = append(adjacents, Coord{x + 1, y})
				alreadyCounted[Coord{x + 1, y}] = true
			}
		}
	}
	/* down */
	if adj, ok := world[Coord{x - 1, y}]; ok {
		if adj != 9 {
			if _, ok := alreadyCounted[Coord{x - 1, y}]; !ok {
				adjacents = append(adjacents, Coord{x - 1, y})
				alreadyCounted[Coord{x - 1, y}] = true
			}
		}
	}
	/* right */
	if adj, ok := world[Coord{x, y + 1}]; ok {
		if adj != 9 {
			if _, ok := alreadyCounted[Coord{x, y + 1}]; !ok {
				adjacents = append(adjacents, Coord{x, y + 1})
				alreadyCounted[Coord{x, y + 1}] = true
			}
		}
	}
	/* left */
	if adj, ok := world[Coord{x, y - 1}]; ok {
		if adj != 9 {
			if _, ok := alreadyCounted[Coord{x, y - 1}]; !ok {
				adjacents = append(adjacents, Coord{x, y - 1})
				alreadyCounted[Coord{x, y - 1}] = true
			}
		}
	}

	return adjacents
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

	world := make(World)
	alreadyCounted = make(map[Coord]bool)

	x := 0
	for scanner.Scan() {
		row := scanner.Text()

		for y := 0; y < len(row); y++ {
			coord := Coord{x, y}
			world[coord], _ = strconv.Atoi(string(row[y]))
		}
		x += 1
	}

	var basinSizes []int

	for {
		basinSize := 0
		var coordToConsider Coord
		/* find a non-9 coord first */
		for coord, check := range world {
			if check == 9 {
				continue
			}
			if _, ok := alreadyCounted[coord]; ok {
				continue
			}
			basinSize = 1
			alreadyCounted[coord] = true
			coordToConsider.x = coord.x
			coordToConsider.y = coord.y
			goto Parse
		}

		break

	Parse:
		adjacents := getAdjacentsToConsider(world, coordToConsider)
		/* remove these from the world */
		for len(adjacents) > 0 {
			adj := adjacents[0]
			basinSize += 1
			adjacents = append(adjacents, getAdjacentsToConsider(world, adj)...)
			adjacents = adjacents[1:]
		}
		basinSizes = append(basinSizes, basinSize)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
	fmt.Println("time:", time.Since(start))
}
