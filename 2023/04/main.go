package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Card struct {
	instances int
	matches   int
}

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
	cards := []Card{}
	total := 0
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		card := strings.Split(text, ": ")
		numbers := strings.Split(card[1], " | ")
		winningNumbers := cleanArray(strings.Split(numbers[0], " "))
		choosenNumber := cleanArray(strings.Split(numbers[1], " "))
		matches := 0
		points := 0

		for _, number := range choosenNumber {
			if slices.Contains(winningNumbers, number) {
				matches++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		cards = append(cards, Card{1, matches})
		sum += points
	}

	for i, card := range cards {
		for instance := 0; instance < card.instances; instance++ {
			for j := i + 1; j <= card.matches+i; j++ {
				if j >= len(cards) {
					break
				}
				cards[j].instances++
			}
		}
	}

	for _, card := range cards {
		total += card.instances
	}

	fmt.Println(sum)
	fmt.Println(total)
}
