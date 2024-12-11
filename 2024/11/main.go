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

type Key struct {
	key    string
	amount int
}

var cache = make(map[Key]int)

func blink(stone string) []string {
	res := []string{}
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
	return res
}

func blinkCache(key Key) int {
	if key.amount == 0 {
		return 1
	}

	inCach, ok := cache[key]
	if ok {
		return inCach
	}

	res := 0
	for _, stone := range blink(key.key) {
		res += blinkCache(Key{stone, key.amount - 1})
	}

	cache[key] = res
	return res
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	arrangement := []string{}
	sum25 := 0
	sum75 := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		arrangement = strings.Split(text, " ")
	}

	for _, stone := range arrangement {
		sum25 += blinkCache(Key{stone, 25})
	}

	for _, stone := range arrangement {
		sum75 += blinkCache(Key{stone, 75})
	}

	fmt.Println("25 blinks", sum25)
	fmt.Println("75 blinks", sum75)
}
