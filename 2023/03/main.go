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

type Symbol struct {
	val     rune
	numbers []int
}

func addToGear(symbol *Symbol, symbols map[Coordinate]Symbol, coordinates *Coordinate, val int) {
	if symbol.val != '*' {
		return
	}

	symbol.numbers = append(symbol.numbers, val)
	symbols[*coordinates] = *symbol
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := []Number{}
	symbols := make(map[Coordinate]Symbol)
	number := ""
	col := 0
	length := -1
	sum := 0
	gearSum := 0

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
				symbols[Coordinate{col, row}] = Symbol{char, []int{}}
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

		leftCoordinate := Coordinate{num.col, start}
		leftSymbol, leftOk := symbols[leftCoordinate]
		if leftOk {
			sum += num.val
			addToGear(&leftSymbol, symbols, &leftCoordinate, num.val)
			continue
		}

		rightCoordinate := Coordinate{num.col, end}
		rightSymbol, rightOk := symbols[rightCoordinate]
		if rightOk {
			sum += num.val
			addToGear(&rightSymbol, symbols, &rightCoordinate, num.val)
			continue
		}

		for i := start; i <= end; i++ {
			if num.col != 0 {
				coordinate := Coordinate{num.col - 1, i}
				symbol, ok := symbols[coordinate]
				if ok {
					sum += num.val
					addToGear(&symbol, symbols, &coordinate, num.val)
					break
				}
			}
			if num.col != length {
				coordinate := Coordinate{num.col + 1, i}
				symbol, ok := symbols[coordinate]
				if ok {
					sum += num.val
					addToGear(&symbol, symbols, &coordinate, num.val)
					break
				}
			}
		}
	}

	for _, symbol := range symbols {
		if len(symbol.numbers) != 2 {
			continue
		}

		gearSum += symbol.numbers[0] * symbol.numbers[1]
	}

	fmt.Println(sum)
	fmt.Println(gearSum)
}
