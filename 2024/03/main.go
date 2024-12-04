package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sumMul(reg *regexp.Regexp, text string, en bool) (int, bool) {
	sum := 0
	enabled := en

	matches := reg.FindAllString(text, -1)
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else {
			if !enabled {
				continue
			}
			tall := strings.Split(match[4:len(match)-1], ",")
			tall1, err := strconv.Atoi(tall[0])
			handleError(err)
			tall2, err := strconv.Atoi(tall[1])
			handleError(err)
			sum += tall1 * tall2
		}
	}

	return sum, enabled
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	reg1, err := regexp.Compile(`mul\(\d+,\d+\)`)
	handleError(err)
	enabled1 := true
	sum1 := 0

	reg2, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	handleError(err)
	enabled2 := true
	sum2 := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		res, en := sumMul(reg1, text, enabled1)
		sum1 += res
		enabled1 = en

		res, en = sumMul(reg2, text, enabled2)
		sum2 += res
		enabled2 = en
	}

	fmt.Println("Sum1", sum1)
	fmt.Println("Sum2", sum2)
}
