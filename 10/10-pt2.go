package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

type ByteStack struct {
	stack *list.List
}

func (s *ByteStack) Push(val rune) {
	s.stack.PushFront(val)
}

func (s *ByteStack) Pop() error {
	if s.stack.Len() > 0 {
		ele := s.stack.Front()
		s.stack.Remove(ele)
		return nil
	}
	return fmt.Errorf("stack empty")
}

func (s *ByteStack) Front() (rune, error) {
	if s.stack.Len() > 0 {
		if val, ok := s.stack.Front().Value.(rune); ok {
			return val, nil
		}
		return 0, fmt.Errorf("Peep Error: Stack Datatype is incorrect")
	}
	return 0, fmt.Errorf("Peep Error: Stack is empty")
}

func (s *ByteStack) Len() int {
	return s.stack.Len()
}

func (s *ByteStack) Init() {
	if s.stack == nil {
		s.stack = list.New()
	}
}

var parensMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scoreMap = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
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
	var scores []int

NextLine:
	for scanner.Scan() {
		row := scanner.Text()
		var s ByteStack
		s.Init()
		lineScore := 0
		for _, c := range row {
			/* opening char */
			if c == '(' || c == '[' || c == '{' || c == '<' {
				s.Push(c)
			} else {
				/* closing char */
				openingParen, err := s.Front()
				if err != nil {
					fmt.Println(err)
					continue NextLine
				}
				/* corruption; skip */
				if parensMap[openingParen] != c {
					continue NextLine
				}
				s.Pop()
			}
		}
		for s.Len() > 0 {
			c, _ := s.Front()
			lineScore *= 5
			lineScore += scoreMap[c]
			s.Pop()
		}
		scores = append(scores, lineScore)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
	fmt.Println("time:", time.Since(start))
}
