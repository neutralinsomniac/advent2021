package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Digit struct {
	a, b, c, d, e, f, g bool
}

func convertDigitStringToDigit(digitStr string) Digit {
	digit := Digit{}
	for _, c := range digitStr {
		switch c {
		case 'a':
			digit.a = true
			break
		case 'b':
			digit.b = true
			break
		case 'c':
			digit.c = true
			break
		case 'd':
			digit.d = true
			break
		case 'e':
			digit.e = true
			break
		case 'f':
			digit.f = true
			break
		case 'g':
			digit.g = true
			break
		}
	}

	return digit
}

func countLitSegs(digit Digit) int {
	lit := 0

	if digit.a {
		lit++
	}
	if digit.b {
		lit++
	}
	if digit.c {
		lit++
	}
	if digit.d {
		lit++
	}
	if digit.e {
		lit++
	}
	if digit.f {
		lit++
	}
	if digit.g {
		lit++
	}

	return lit
}

func countLitIntersections(digitA, digitB Digit) int {
	numIntersections := 0

	if digitA.a && digitB.a {
		numIntersections += 1
	}
	if digitA.b && digitB.b {
		numIntersections += 1
	}
	if digitA.c && digitB.c {
		numIntersections += 1
	}
	if digitA.d && digitB.d {
		numIntersections += 1
	}
	if digitA.e && digitB.e {
		numIntersections += 1
	}
	if digitA.f && digitB.f {
		numIntersections += 1
	}
	if digitA.g && digitB.g {
		numIntersections += 1
	}

	return numIntersections
}

type DigitMap map[Digit]int
type NumberMap map[int]Digit

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		digitMap := make(DigitMap, 0)
		numberMap := make(NumberMap, 0)
		row := scanner.Text()

		sp := strings.Split(row, "|")
		patterns := strings.Fields(sp[0])
		var digitsToSolve []Digit
		for _, digitStr := range patterns {
			digit := convertDigitStringToDigit(digitStr)
			switch len(digitStr) {
			case 2:
				digitMap[digit] = 1
				numberMap[1] = digit
				break
			case 3:
				digitMap[digit] = 7
				numberMap[7] = digit
				break
			case 4:
				digitMap[digit] = 4
				numberMap[4] = digit
				break
			case 7:
				digitMap[digit] = 8
				numberMap[8] = digit
				break
			default:
				digitsToSolve = append(digitsToSolve, digit)
				break
			}
		}

		/* deduce other digits here */
		/* '6' has 6 segs lit, and contains only one of the segs in '1' */
		for i, digit := range digitsToSolve {
			if countLitSegs(digit) == 6 && countLitIntersections(digit, numberMap[1]) == 1 {
				digitMap[digit] = 6
				numberMap[6] = digit
				// remove 6 from our list of digits we need to solve
				digitsToSolve[i] = digitsToSolve[len(digitsToSolve)-1]
				digitsToSolve = digitsToSolve[:len(digitsToSolve)-1]
				break
			}
		}
		/* '3' has 5 segs lit, and contains both of the segs in '1' */
		for i, digit := range digitsToSolve {
			if countLitSegs(digit) == 5 && countLitIntersections(digit, numberMap[1]) == 2 {
				digitMap[digit] = 3
				numberMap[3] = digit
				// remove 3 from our list of digits we need to solve
				digitsToSolve[i] = digitsToSolve[len(digitsToSolve)-1]
				digitsToSolve = digitsToSolve[:len(digitsToSolve)-1]
				break
			}
		}

		/* solve 5 using 6 */
		/* '5' has 5 segs lit, and contains 5 segs from '6' */
		for i, digit := range digitsToSolve {
			if countLitSegs(digit) == 5 && countLitIntersections(digit, numberMap[6]) == 5 {
				digitMap[digit] = 5
				numberMap[5] = digit
				// remove 5 from our list of digits we need to solve
				digitsToSolve[i] = digitsToSolve[len(digitsToSolve)-1]
				digitsToSolve = digitsToSolve[:len(digitsToSolve)-1]
				break
			}
		}
		/* solve 2 using number of segs */
		/* '2' has 5 segs lit */
		for i, digit := range digitsToSolve {
			if countLitSegs(digit) == 5 {
				digitMap[digit] = 2
				numberMap[2] = digit
				// remove 2 from our list of digits we need to solve
				digitsToSolve[i] = digitsToSolve[len(digitsToSolve)-1]
				digitsToSolve = digitsToSolve[:len(digitsToSolve)-1]
				break
			}
		}

		/* solve 0 using 5 */
		/* '0' contains 4 segs from '5' */
		for i, digit := range digitsToSolve {
			if countLitIntersections(digit, numberMap[5]) == 4 {
				digitMap[digit] = 0
				numberMap[0] = digit
				// remove 0 from our list of digits we need to solve
				digitsToSolve[i] = digitsToSolve[len(digitsToSolve)-1]
				digitsToSolve = digitsToSolve[:len(digitsToSolve)-1]
				break
			}
		}

		/* only '9' left */
		lastDigit := digitsToSolve[0]
		digitMap[lastDigit] = 9
		numberMap[9] = lastDigit

		/* now calculate the output */
		outputs := strings.Fields(sp[1])
		val := 0
		for x, digitStr := range outputs {
			digit := convertDigitStringToDigit(digitStr)
			val += (int(math.Pow(10, float64(3-x)))) * digitMap[digit]
		}
		sum += val
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
