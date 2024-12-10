package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Vec struct {
	x, y int
}

type Distinct map[Vec]int

func walkPaths(m [][]int, index Vec, set Distinct) {
	level := m[index.y][index.x]
	xL, xR, yU, yD := index.x-1, index.x+1, index.y-1, index.y+1

	if m[index.y][index.x] == 9 {
		rating, ok := set[index]
		if ok {
			set[index] = rating + 1
		} else {
			set[index] = 1
		}
		return
	}

	if yU >= 0 && m[yU][index.x]-level == 1 {
		walkPaths(m, Vec{index.x, yU}, set)
	}
	if yD < len(m) && m[yD][index.x]-level == 1 {
		walkPaths(m, Vec{index.x, yD}, set)
	}
	if xL >= 0 && m[index.y][xL]-level == 1 {
		walkPaths(m, Vec{xL, index.y}, set)
	}
	if xR < len(m[0]) && m[index.y][xR]-level == 1 {
		walkPaths(m, Vec{xR, index.y}, set)
	}
}

func findNumberOfPaths(m [][]int, index Vec) Distinct {
	distinct := make(Distinct)
	walkPaths(m, index, distinct)
	return distinct
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	m := [][]int{}
	score := 0
	rating := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		intArr := []int{}
		for _, r := range text {
			num, err := strconv.Atoi(string(r))
			handleError(err)
			intArr = append(intArr, num)
		}
		m = append(m, intArr)
	}

	for y, arr := range m {
		for x, n := range arr {
			if n == 0 {
				distinct := findNumberOfPaths(m, Vec{x, y})
				score += len(distinct)

				for _, r := range distinct {
					rating += r
				}
			}
		}
	}
	fmt.Println("Score", score)
	fmt.Println("Rating", rating)
}
