package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func trim(s string) string {
	return strings.Trim(s, " ")
}

type Bag = map[string]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bag := Bag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sumPossible := 0
	sumMinimum := 0

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		res := strings.Split(line, ":")
		rounds := strings.Split(res[1], ";")

		possible := true
		minAmountBag := make(Bag)

		for _, round := range rounds {
			for _, cube := range strings.Split(round, ",") {
				c := strings.Split(trim(cube), " ")
				color := c[1]
				amount, _ := strconv.Atoi(c[0])

				bagAmount, ok := minAmountBag[color]
				if (ok && bagAmount < amount) || !ok {
					minAmountBag[color] = amount
				}

				maxAmount, ok := bag[color]
				if possible && ok && amount > maxAmount {
					possible = false
				}
			}
		}

		power := 1
		for _, val := range minAmountBag {
			power *= val
		}
		sumMinimum += power

		if possible {
			game := strings.Split(res[0], " ")
			gameNum, _ := strconv.Atoi(game[1])
			sumPossible += gameNum
		}
	}

	fmt.Println(sumPossible)
	fmt.Println(sumMinimum)
}
