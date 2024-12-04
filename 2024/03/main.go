package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var mul = "mul("
var do = "do()"
var dont = "don't()"

func partOfValidInst(inst string, r rune) bool {
	return (strings.Contains(mul, inst) && len(inst) != len(mul) && r == rune(mul[len(inst)])) ||
		(strings.Contains(do, inst) && len(inst) != len(do) && r == rune(do[len(inst)])) ||
		(strings.Contains(dont, inst) && len(inst) != len(dont) && r == rune(dont[len(inst)]))
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	sum := 0
	enabled := true
	inst := ""
	sTall1 := ""
	sTall2 := ""
	isSecondNumber := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		for _, r := range text {
			if partOfValidInst(inst, r) {
				inst += string(r)
				continue
			}

			if inst == do {
				enabled = true
			} else if inst == dont {
				enabled = false
			} else if enabled && inst == mul {
				if unicode.IsDigit(r) {
					if isSecondNumber {
						sTall2 += string(r)
					} else {
						sTall1 += string(r)
					}
					continue
				}

				if r == ',' {
					isSecondNumber = true
					continue
				}

				if r == ')' && sTall1 != "" && sTall2 != "" {
					tall1, err := strconv.Atoi(sTall1)
					handleError(err)
					tall2, err := strconv.Atoi(sTall2)
					handleError(err)

					sum += tall1 * tall2
				}
			}

			inst = ""
			sTall1 = ""
			sTall2 = ""
			isSecondNumber = false
		}
	}

	fmt.Println("Sum", sum)
}
