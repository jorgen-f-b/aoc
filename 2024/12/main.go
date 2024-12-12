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

type Region struct {
	area      int
	perimeter int
}

func (r *Region) getCost() int {
	return r.area * r.perimeter
}

func markMap(m []string, x, y int) {
	tmp := []rune(m[y])
	tmp[x] = '.'
	m[y] = string(tmp)
}

func compStrRune(s string, r rune, x int) bool {
	return []rune(s)[x] == r
}

func getRegion(m []string, gm []string, region *Region, plot rune, x, y int) {
	markMap(m, x, y)
	markMap(gm, x, y)
	region.area++

	if y == 0 || !compStrRune(m[y-1], plot, x) {
		if y == 0 || !compStrRune(m[y-1], '.', x) {
			region.perimeter++
		}
	} else {
		getRegion(m, gm, region, plot, x, y-1)
	}

	if y == len(m)-1 || !compStrRune(m[y+1], plot, x) {
		if y == len(m)-1 || !compStrRune(m[y+1], '.', x) {
			region.perimeter++
		}
	} else {
		getRegion(m, gm, region, plot, x, y+1)
	}

	if x == 0 || !compStrRune(m[y], plot, x-1) {
		if x == 0 || !compStrRune(m[y], '.', x-1) {
			region.perimeter++
		}
	} else {
		getRegion(m, gm, region, plot, x-1, y)
	}

	if x == len(m[0])-1 || !compStrRune(m[y], plot, x+1) {
		if x == len(m[0])-1 || !compStrRune(m[y], '.', x+1) {
			region.perimeter++
		}
	} else {
		getRegion(m, gm, region, plot, x+1, y)
	}
}

func findFreeSpace(m []string) (int, int) {
	for i, s := range m {
		for j, r := range s {
			if r != '.' {
				return j, i
			}
		}
	}
	return -1, -1
}

func getFenceCost(m, gm []string) int {
	x, y := findFreeSpace(gm)
	if x == -1 {
		return 0
	}
	region := Region{}
	plot := []rune(gm[y])[x]
	cm := make([]string, len(gm))
	copy(cm, m)
	getRegion(cm, gm, &region, plot, x, y)
	return region.getCost() + getFenceCost(m, gm)
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	m := []string{}
	gm := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		m = append(m, text)
		gm = append(gm, text)
	}

	fmt.Println("Sum", getFenceCost(m, gm))
}
