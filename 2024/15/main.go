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
			if r == '[' {
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

func copyMap(m [][]rune) [][]rune {
	cm := [][]rune{}
	for _, arr := range m {
		cArr := []rune{}
		cArr = append(cArr, arr...)
		cm = append(cm, cArr)
	}
	return cm
}

func move(m [][]rune, dir rune, x, y int) (bool, [][]rune) {
	cm := copyMap(m)

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

	if cm[nY][nX] == '#' {
		return false, cm
	}
	if cm[nY][nX] == '[' && nX == x {
		ok1, nm1 := move(m, dir, x, nY)
		if ok1 {
			ok2, nm2 := move(nm1, dir, x+1, nY)
			if !ok2 {
				return false, cm
			}
			cm = nm2
		} else {
			return false, cm
		}
	} else if cm[nY][nX] == ']' && nX == x {
		ok1, nm1 := move(m, dir, x, nY)
		if ok1 {
			ok2, nm2 := move(nm1, dir, x-1, nY)
			if !ok2 {
				return false, cm
			}
			cm = nm2
		} else {
			return false, cm
		}
	} else if cm[nY][nX] == 'O' || cm[nY][nX] == '[' || cm[nY][nX] == ']' {
		ok, nm := move(cm, dir, nX, nY)
		if !ok {
			return false, cm
		}
		cm = nm
	}

	cm[nY][nX] = cm[y][x]
	cm[y][x] = '.'

	return true, cm
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	getMap := true
	m := [][]rune{}
	em := [][]rune{}
	moves := []rune{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			getMap = false
		}

		if getMap {
			m = append(m, []rune(text))

			line := []rune{}
			for _, r := range text {
				if r == '@' {
					line = append(line, r, '.')
					continue
				}
				if r == 'O' {
					line = append(line, '[', ']')
					continue
				}
				line = append(line, r, r)
			}
			em = append(em, line)
		} else {
			moves = append(moves, []rune(text)...)
		}
	}

	for _, dir := range moves {
		_, m = move(m, dir, -1, -1)
	}

	for _, dir := range moves {
		_, em = move(em, dir, -1, -1)
	}

	fmt.Println("Sum", calculateGPSCoordinates(m))
	fmt.Println("SumExtended", calculateGPSCoordinates(em))
}
