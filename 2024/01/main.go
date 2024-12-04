package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sumDiff(slice1 []int, slice2 []int) int {
	sum := 0

	for i, tall1 := range slice1 {
		tall2 := slice2[i]

		if tall1 > tall2 {
			sum += tall1 - tall2
			continue
		}
		sum += tall2 - tall1
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	handleErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice1 := []int{}
	slice2 := []int{}

	for scanner.Scan() {
		tall := strings.Split(scanner.Text(), "   ")

		tall1, err := strconv.Atoi(tall[0])
		handleErr(err)
		tall2, err := strconv.Atoi(tall[1])
		handleErr(err)

		slice1 = append(slice1, tall1)
		slice2 = append(slice2, tall2)
	}

	slices.Sort(slice1)
	slices.Sort(slice2)

	fmt.Println(sumDiff(slice1, slice2))
}
