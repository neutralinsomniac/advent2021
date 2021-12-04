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

var bingoBoard map[Coord]bool

func ReadCalls(s string) []int {
	callStrs := strings.Split(s, ",")
	var calls []int

	for _, val := range callStrs {
		i, err := strconv.Atoi(val)
		checkErr(err)
		calls = append(calls, i)
	}

	return calls
}

type Slot struct {
	val    int
	scored bool
}

type Board [5][5]Slot

func printBoard(board Board) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			fmt.Printf("%3d", board[x][y].val)
			if board[x][y].scored {
				fmt.Printf(".")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func ReadBoards(scanner *bufio.Scanner) []Board {

	var boards []Board
	var board Board
	var x, y int

	for scanner.Scan() {
		row := scanner.Text()
		numsStr := strings.Fields(row)
		for y = 0; y < 5; y++ {
			board[x][y].val, _ = strconv.Atoi(numsStr[y])
		}
		x += 1
		if x == 5 {
			// new board!
			boards = append(boards, board)
			x = 0
			scanner.Scan()
		}
	}

	return boards
}

/* returns whether we marked the slot or not */
func markBoard(board *Board, call int) (bool, int, int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if board[x][y].val == call {
				board[x][y].scored = true
				return true, x, y
			}
		}
	}
	return false, 0, 0
}

/* x and y are hints to the last modified position */
/* returns sum of row or column that won */
func checkBoard(board Board, x, y int) (bool, int) {
	/* check row */
	var i int
	var won bool

	for i = 0; i < 5; i++ {
		if board[x][(y+i)%5].scored != true {
			break
		}
	}
	if i == 5 {
		won = true
	}

	if !won {
		/* check column */
		for i = 0; i < 5; i++ {
			if board[(x+i)%5][y].scored != true {
				break
			}
		}
		if i == 5 {
			won = true
		}
	}

	if won {
		return true, scoreBoard(board)
	}
	return false, 0
}

func scoreBoard(board Board) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if board[x][y].scored == false {
				sum += board[x][y].val
			}
		}
	}
	return sum
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

	/* snag the calls */
	scanner.Scan()
	calls := ReadCalls(scanner.Text())

	/* snag the boards */
	scanner.Scan()
	boards := ReadBoards(scanner)

	/* now run the calls against the boards */
	for _, call := range calls {
		var newBoardList []Board
		/* run the call against all the boards */
		for _, board := range boards {
			found, x, y := markBoard(&board, call)
			won := false
			if found {
				won, _ = checkBoard(board, x, y)
				/* ding ding ding */
				if won && len(boards) == 1 {
					fmt.Println(call * scoreBoard(board))
					fmt.Println("time:", time.Since(start))
					return
				}
			}
			/* didn't win; add this board to be considered next call */
			if !won {
				newBoardList = append(newBoardList, board)
			}
		}
		boards = newBoardList
	}
}
