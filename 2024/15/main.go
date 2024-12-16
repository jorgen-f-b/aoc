package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func printMap(m [][]rune) {
	for _, arr := range m {
		fmt.Println(string(arr))
	}
}

func calculateGPSCoordinates(m [][]rune) int {
	sum := 0
	for y, arr := range m {
		for x, r := range arr {
			if r == 'O' {
				sum += (100 * y) + x
			}
		}
	}
	return sum
}

func findSub(m [][]rune) (int, int) {
	for y, arr := range m {
		for x, r := range arr {
			if r == '@' {
				return x, y
			}
		}
	}
	panic("Sub gone")
}

func move(m [][]rune, dir rune, x, y int) bool {
	if x == -1 {
		x, y = findSub(m)
	}

	nX, nY := x, y
	switch dir {
	case '<':
		nX = x - 1
	case '>':
		nX = x + 1
	case '^':
		nY = y - 1
	case 'v':
		nY = y + 1
	}

	if m[nY][nX] == '#' {
		return false
	}
	if m[nY][nX] == 'O' {
		if !move(m, dir, nX, nY) {
			return false
		}
	}

	m[nY][nX] = m[y][x]
	m[y][x] = '.'

	return true
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	getMap := true
	m := [][]rune{}
	moves := []rune{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			getMap = false
		}

		if getMap {
			m = append(m, []rune(text))
		} else {
			moves = append(moves, []rune(text)...)
		}
	}

	for _, dir := range moves {
		move(m, dir, -1, -1)
	}
	fmt.Println("Sum", calculateGPSCoordinates(m))
}
