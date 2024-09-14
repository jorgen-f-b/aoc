package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

type Number struct {
	val   int
	start int
	end   int
	col   int
}

type Coordinate struct {
	col int
	row int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := []Number{}
	symbols := make(map[Coordinate]string)
	number := ""
	col := 0
	length := -1
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := []rune(scanner.Text())
		start := -1

		if length == -1 {
			length = len(text) - 1
		}

		for row, char := range text {
			switch {
			case char == '.':
				if number != "" {
					num, _ := strconv.Atoi(number)
					numbers = append(numbers, Number{num, start, row - 1, col})
					number = ""
					start = -1
				}
			case unicode.IsDigit(char):
				number += string(char)
				if start == -1 {
					start = row
				}
			case isSymbol(char):
				symbols[Coordinate{col, row}] = string(char)
				if number != "" {
					num, _ := strconv.Atoi(number)
					numbers = append(numbers, Number{num, start, row - 1, col})
					number = ""
					start = -1
				}
			}
		}
		if number != "" {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, Number{num, start, length, col})
			number = ""
			start = -1
		}
		col++
	}
	file.Close()

	for _, num := range numbers {
		var start int
		if num.start == 0 {
			start = 0
		} else {
			start = num.start - 1
		}

		var end int
		if num.end == length {
			end = length
		} else {
			end = num.end + 1
		}

		_, leftOk := symbols[Coordinate{num.col, start}]
		if leftOk {
			sum += num.val
			continue
		}

		_, rightOk := symbols[Coordinate{num.col, end}]
		if rightOk {
			sum += num.val
			continue
		}

		for i := start; i <= end; i++ {
			if num.col != 0 {
				_, ok := symbols[Coordinate{num.col - 1, i}]
				if ok {
					sum += num.val
					break
				}
			}
			if num.col != length {
				_, ok := symbols[Coordinate{num.col + 1, i}]
				if ok {
					sum += num.val
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
