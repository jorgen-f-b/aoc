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

func possibleGames(file *os.File) {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sumPossible := 0

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		res := strings.Split(line, ":")
		possible := true

		rounds := strings.Split(res[1], ";")
		for _, round := range rounds {
			for _, cube := range strings.Split(round, ",") {
				color := strings.Split(trim(cube), " ")
				maxAmount, ok := bag[color[1]]
				if ok {
					amount, _ := strconv.Atoi(color[0])
					if amount > maxAmount {
						possible = false
						break
					}
				} else {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			game := strings.Split(res[0], " ")
			gameNum, _ := strconv.Atoi(game[1])
			sumPossible += gameNum
		}
	}
	fmt.Println(sumPossible)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	possibleGames(file)
}
