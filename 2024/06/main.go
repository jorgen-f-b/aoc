package main

import (
	"bufio"
	"fmt"
	"os"
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
	direction rune
}

func printMap(m []string) {
	for _, s := range m {
		fmt.Println(s)
	}
}

func replaceAtIndex(s string, i int, c rune) string {
	out := []rune(s)
	out[i] = c
	return string(out)
}

func (g *Guard) rotate(m []string) {
	if g.direction == '^' && m[g.pos.y-1][g.pos.x] == '#' {
		g.direction = '>'
	}
	if g.direction == 'v' && m[g.pos.y+1][g.pos.x] == '#' {
		g.direction = '<'
	}
	if g.direction == '>' && m[g.pos.y][g.pos.x+1] == '#' {
		g.direction = 'v'
	}
	if g.direction == '<' && m[g.pos.y][g.pos.x-1] == '#' {
		g.direction = '^'
	}
}

func (g *Guard) walk(m []string) bool {
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

	g.rotate(m)

	return false
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	m := []string{}
	g := Guard{}
	run := true
	dp := 0

	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		di := strings.Index(text, "^")
		if di != -1 {
			g.direction = '^'
			g.pos.x = di
			g.pos.y = row

			m = append(m, replaceAtIndex(text, di, 'X'))
		} else {
			m = append(m, text)
		}
		row++
	}

	for run {
		// printMap(m)
		// fmt.Println()
		// fmt.Println("----------")
		// fmt.Println()
		run = !g.walk(m)
	}

	for _, s := range m {
		dp += strings.Count(s, "X")
	}

	// printMap(m)
	// fmt.Println(g)
	fmt.Println("Distint positions", dp)
}
