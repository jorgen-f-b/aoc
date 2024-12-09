package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Position struct {
	x int
	y int
}

type Guard struct {
	pos       Position
	posMap    map[Position][]rune
	direction rune
}

func printMap(m []string) {
	fmt.Println()
	for _, s := range m {
		fmt.Println(s)
	}
	fmt.Println()
}

func replaceAtIndex(s string, i int, c rune) string {
	out := []rune(s)
	out[i] = c
	return string(out)
}

func findAll(s string, c rune) []int {
	arr := []int{}
	for i, r := range s {
		if r == c {
			arr = append(arr, i)
		}
	}
	return arr
}

func copyArray(arr []string) []string {
	nArr := []string{}
	for _, s := range arr {
		nArr = append(nArr, s)
	}
	return nArr
}

func (g *Guard) rotate(m []string) bool {
	if g.direction == '^' && (m[g.pos.y-1][g.pos.x] == '#' || m[g.pos.y-1][g.pos.x] == 'O') {
		g.direction = '>'
		return true
	}
	if g.direction == 'v' && (m[g.pos.y+1][g.pos.x] == '#' || m[g.pos.y+1][g.pos.x] == 'O') {
		g.direction = '<'
		return true
	}
	if g.direction == '>' && (m[g.pos.y][g.pos.x+1] == '#' || m[g.pos.y][g.pos.x+1] == 'O') {
		g.direction = 'v'
		return true
	}
	if g.direction == '<' && (m[g.pos.y][g.pos.x-1] == '#' || m[g.pos.y][g.pos.x-1] == 'O') {
		g.direction = '^'
		return true
	}
	return false
}

func (g *Guard) markPosition() bool {
	if g.posMap == nil {
		g.posMap = make(map[Position][]rune)
	}

	positions, ok := g.posMap[g.pos]
	if ok {
		if slices.Contains(positions, g.direction) {
			return true
		}
		positions = append(positions, g.direction)
		g.posMap[g.pos] = positions
	} else {
		g.posMap[g.pos] = []rune{g.direction}
	}
	return false
}

func (g *Guard) walk(m []string) bool {
	if g.rotate(m) {
		return false
	}

	switch g.direction {
	case '^':
		g.pos.y--
	case 'v':
		g.pos.y++
	case '>':
		g.pos.x++
	case '<':
		g.pos.x--
	default:
		panic("walk: Not legal direction")
	}

	m[g.pos.y] = replaceAtIndex(m[g.pos.y], g.pos.x, 'X')

	if g.pos.x == 0 || g.pos.x == len(m[0])-1 || g.pos.y == 0 || g.pos.y == len(m)-1 {
		return true
	}

	return false
}

func (g *Guard) walkAndMark(m []string) bool {
	run := true
	for run {
		run = !g.walk(m)

		if g.markPosition() {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	m := []string{}
	tmpM := []string{}
	g := Guard{}
	run := true
	dp := []Position{}
	cl := 0
	sp := Position{}

	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		di := strings.Index(text, "^")
		if di != -1 {
			sp.x = di
			sp.y = row

			g.direction = '^'
			g.pos = sp

			m = append(m, replaceAtIndex(text, di, 'X'))
			tmpM = append(tmpM, replaceAtIndex(text, di, 'X'))
		} else {
			m = append(m, text)
			tmpM = append(tmpM, text)
		}
		row++
	}

	for run {
		run = !g.walk(m)
	}

	row = 0
	for _, s := range m {
		pos := findAll(s, 'X')
		for _, p := range pos {
			if sp.x == p && sp.y == row {
				continue
			}
			dp = append(dp, Position{p, row})
		}
		row++
	}

	for _, p := range dp {
		newM := copyArray(tmpM)
		newM[p.y] = replaceAtIndex(newM[p.y], p.x, 'O')

		g.direction = '^'
		g.pos = sp
		g.posMap = nil
		if g.walkAndMark(newM) {
			cl++
		}
	}

	fmt.Println("Distint positions", len(dp)+1)
	fmt.Println("Creates loop", cl)
}
