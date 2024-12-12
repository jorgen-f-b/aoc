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
	area      [][]rune
	totalArea int
}

func makeRegion(height, width int) Region {
	return Region{makeArea(height, width), 0}
}

func (r *Region) addArea(x, y int) {
	r.area[y][x] = 'X'
	r.totalArea++
}

func (r *Region) hasArea(x, y int) bool {
	return r.area[y][x] == 'X'
}

func (r *Region) costDicount() int {
	perimeter := 0
	row := 0
	col := 0
	for y, arr := range r.area {
		isTop := false
		isBottom := false
		for x := range arr {
			if r.area[y][x] == '.' {
				if isTop {
					row++
					isTop = false
				}
				if isBottom {
					row++
					isBottom = false
				}
				continue
			}
			if y == 0 || !r.hasArea(x, y-1) {
				isTop = true
			} else if isTop {
				row++
				isTop = false
			}
			if y == len(r.area)-1 || !r.hasArea(x, y+1) {
				isBottom = true
			} else if isBottom {
				row++
				isBottom = false
			}
		}
		if isTop {
			row++
		}
		if isBottom {
			row++
		}
	}

	for x := 0; x < len(r.area[0]); x++ {
		isLeft := false
		isRight := false
		for y, arr := range r.area {
			if r.area[y][x] == '.' {
				if isLeft {
					col++
					isLeft = false
				}
				if isRight {
					col++
					isRight = false
				}
				continue
			}
			if x == 0 || !r.hasArea(x-1, y) {
				isLeft = true
			} else if isLeft {
				col++
				isLeft = false
			}
			if x == len(arr)-1 || !r.hasArea(x+1, y) {
				isRight = true
			} else if isRight {
				col++
				isRight = false
			}
		}
		if isLeft {
			col++
		}
		if isRight {
			col++
		}
	}

	perimeter = row + col
	return perimeter * r.totalArea
}

func (r *Region) cost() int {
	perimeter := 0
	for y, arr := range r.area {
		for x := range arr {
			if r.area[y][x] == '.' {
				continue
			}
			if y == 0 || !r.hasArea(x, y-1) {
				perimeter++
			}
			if x == 0 || !r.hasArea(x-1, y) {
				perimeter++
			}
			if y == len(r.area)-1 || !r.hasArea(x, y+1) {
				perimeter++
			}
			if x == len(arr)-1 || !r.hasArea(x+1, y) {
				perimeter++
			}
		}
	}
	return perimeter * r.totalArea
}

func makeRuneSlice(width int) []rune {
	slice := make([]rune, width)
	for i := range slice {
		slice[i] = '.'
	}
	return slice
}

func makeArea(height, width int) [][]rune {
	area := make([][]rune, height)
	for i := range area {
		area[i] = makeRuneSlice(width)
	}
	return area
}

func markMap(m [][]rune, x, y int) {
	m[y][x] = '.'
}

func getRegion(m [][]rune, region *Region, plot rune, x, y int) {
	markMap(m, x, y)
	region.addArea(x, y)

	if y > 0 && m[y-1][x] == plot {
		getRegion(m, region, plot, x, y-1)
	}

	if y < len(m)-1 && m[y+1][x] == plot {
		getRegion(m, region, plot, x, y+1)
	}

	if x > 0 && m[y][x-1] == plot {
		getRegion(m, region, plot, x-1, y)
	}

	if x < len(m[0])-1 && m[y][x+1] == plot {
		getRegion(m, region, plot, x+1, y)
	}
}

func findFreeSpace(m [][]rune) (int, int) {
	for i, s := range m {
		for j, r := range s {
			if r != '.' {
				return j, i
			}
		}
	}
	return -1, -1
}

func getFenceCost(m [][]rune) []Region {
	x, y := findFreeSpace(m)
	if x == -1 {
		return []Region{}
	}
	region := makeRegion(len(m), len(m[0]))
	plot := []rune(m[y])[x]
	getRegion(m, &region, plot, x, y)
	return append(getFenceCost(m), region)
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	m := [][]rune{}
	sumCost := 0
	sumDiscount := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		m = append(m, []rune(text))
	}

	rArr := getFenceCost(m)

	for _, region := range rArr {
		sumCost += region.cost()
	}

	for _, region := range rArr {
		sumDiscount += region.costDicount()
	}

	fmt.Println("SumCost", sumCost)
	fmt.Println("SumDiscount", sumDiscount)
}
