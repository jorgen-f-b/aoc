package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var mul = "mul("

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		inst := ""
		sTall1 := ""
		sTall2 := ""
		isSecondNumber := false

		for _, r := range text {
			if len(inst) != len(mul) && r == rune(mul[len(inst)]) {
				inst += string(r)
				continue
			}
			if inst == mul {
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
