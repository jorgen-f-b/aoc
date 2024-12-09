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

type Coordinate struct {
	x, y int
}

func (c1 *Coordinate) getAntinodes(c2 *Coordinate) []Coordinate {
	diffX := c1.x - c2.x
	diffY := c1.y - c2.y

	return []Coordinate{{c1.x + diffX, c1.y + diffY}, {c2.x - diffX, c2.y - diffY}}
}

func (c1 *Coordinate) getAntinodesHarmonics(c2 *Coordinate, maxX, maxY int) []Coordinate {
	diffX := c1.x - c2.x
	diffY := c1.y - c2.y
	coordinates := []Coordinate{}

	x := c1.x
	y := c1.y
	for x >= 0 && y >= 0 && x <= maxX && y <= maxY {
		coordinates = append(coordinates, Coordinate{x, y})
		x += diffX
		y += diffY
	}

	x = c2.x
	y = c2.y
	for x >= 0 && y >= 0 && x <= maxX && y <= maxY {
		coordinates = append(coordinates, Coordinate{x, y})
		x -= diffX
		y -= diffY
	}

	return coordinates
}

type Frequency map[rune][]Coordinate

type UniqueAntinode map[Coordinate]bool

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	m := []string{}
	frequencyMap := make(Frequency)
	uniqueAntinode := make(UniqueAntinode)
	uniqueAntinodeHarmonics := make(UniqueAntinode)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		m = append(m, text)
	}

	// Coordinate Frequencies
	for i, s := range m {
		for j, r := range s {
			if r == '.' {
				continue
			}

			coordinate := Coordinate{j, i}

			frequency, ok := frequencyMap[r]
			if ok {
				frequency = append(frequency, coordinate)
				frequencyMap[r] = frequency
			} else {
				frequencyMap[r] = []Coordinate{coordinate}
			}
		}
	}

	for _, cArr := range frequencyMap {
		for i := 0; i < len(cArr)-1; i++ {
			c1 := cArr[i]
			for j := i + 1; j < len(cArr); j++ {
				c2 := cArr[j]

				antinodes := c1.getAntinodes(&c2)
				antinode1 := antinodes[0]
				if antinode1.x >= 0 && antinode1.x < len(m[0]) && antinode1.y >= 0 && antinode1.y < len(m) {
					uniqueAntinode[antinode1] = true
				}
				antinode2 := antinodes[1]
				if antinode2.x >= 0 && antinode2.x < len(m[0]) && antinode2.y >= 0 && antinode2.y < len(m) {
					uniqueAntinode[antinode2] = true
				}

				antinodes = c1.getAntinodesHarmonics(&c2, len(m[0])-1, len(m)-1)
				for _, ah := range antinodes {
					uniqueAntinodeHarmonics[ah] = true
				}
			}
		}
	}

	fmt.Println("Sum", len(uniqueAntinode))
	fmt.Println("SumHarmonics", len(uniqueAntinodeHarmonics))
}
