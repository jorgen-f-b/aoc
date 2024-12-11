package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func blink(arrangement []string) []string {
	res := []string{}

	for _, stone := range arrangement {
		if stone == "0" {
			res = append(res, "1")
		} else if len(stone)%2 == 0 {
			num1, err := strconv.Atoi(stone[:len(stone)/2])
			handleError(err)
			res = append(res, strconv.Itoa(num1))

			num2, err := strconv.Atoi(stone[len(stone)/2:])
			handleError(err)
			res = append(res, strconv.Itoa(num2))
		} else {
			num, err := strconv.Atoi(stone)
			handleError(err)
			res = append(res, strconv.Itoa(num*2024))
		}
	}

	return res
}

func blinkSeries(initial []string, amount int) []string {
	res := make([]string, len(initial))
	copy(res, initial)

	for i := 0; i < amount; i++ {
		res = blink(res)
	}

	return res
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	arrangement := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		arrangement = strings.Split(text, " ")
	}

	blinks25 := blinkSeries(arrangement, 25)

	fmt.Println("25 blinks", len(blinks25))
}
