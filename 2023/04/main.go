package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func cleanArray(arr []string) []string {
	newStr := []string{}
	for _, str := range arr {
		if str == "" {
			continue
		}
		newStr = append(newStr, str)
	}
	return newStr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		card := strings.Split(text, ": ")
		numbers := strings.Split(card[1], " | ")
		winningNumbers := cleanArray(strings.Split(numbers[0], " "))
		choosenNumber := cleanArray(strings.Split(numbers[1], " "))
		points := 0

		for _, number := range choosenNumber {
			if slices.Contains(winningNumbers, number) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	fmt.Println(sum)
}
