package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func stringToIntSlice(arr []string) []int {
	intArr := make([]int, 0, len(arr))

	for _, s := range arr {
		n, err := strconv.Atoi(s)
		handleError(err)

		intArr = append(intArr, n)
	}

	return intArr
}

func checkLevel(arr []int) bool {
	prevLevel := -1
	rising := 0
	safe := true

	for _, level := range arr {
		if prevLevel == -1 {
			prevLevel = level
			continue
		}

		diff := prevLevel - level

		if diff > 3 || diff < -3 {
			safe = false
			break
		}

		if diff == 0 {
			safe = false
			break
		} else if diff < 0 {
			if rising == 1 {
				safe = false
				break
			}
			if rising == 0 {
				rising = -1
			}
		} else if diff > 0 {
			if rising == -1 {
				safe = false
				break
			}
			if rising == 0 {
				rising = 1
			}
		}

		prevLevel = level
	}

	return safe
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		list := stringToIntSlice(strings.Split(text, " "))

		safe := checkLevel(list)
		if safe {
			sum++
		}
	}

	fmt.Println("SafeAmount", sum)
}
