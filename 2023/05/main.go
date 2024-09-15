package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	sourceStart      int
	destinationStart int
	length           int
}

func (m *Map) inMap(num int) bool {
	return num >= m.sourceStart && num < m.sourceStart+m.length
}

func (m *Map) destination(num int) (int, bool) {
	if m.inMap(num) {
		diff := num - m.sourceStart
		return m.destinationStart + diff, true
	}
	return -1, false
}

func destination(mArr *[]Map, num int) (int, bool) {
	ret := -1
	for _, m := range *mArr {
		newNum, ok := m.destination(num)
		if ok {
			ret = newNum
			break
		}
	}
	return ret, ret != -1
}

func toNumbers(arr []string) []int {
	intArr := []int{}
	for _, s := range arr {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		intArr = append(intArr, num)
	}
	return intArr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seeds []int
	mapList := make([][]Map, 0)
	isSeed := true
	location := -1
	location2 := -1

	for scanner.Scan() {
		text := scanner.Text()

		if isSeed {
			seeds = toNumbers(strings.Split(text, " "))
			isSeed = false
			continue
		} else if text == "" {
			mapList = append(mapList, make([]Map, 0))
			continue
		} else if strings.Contains(text, "map") {
			continue
		}

		numbers := toNumbers(strings.Split(text, " "))
		newMap := Map{numbers[1], numbers[0], numbers[2]}
		temp := mapList[len(mapList)-1]
		temp = append(temp, newMap)
		mapList[len(mapList)-1] = temp
	}

	for _, seed := range seeds {
		num := seed
		for _, mArr := range mapList {
			newNum, ok := destination(&mArr, num)
			if ok {
				num = newNum
			}
		}

		if location == -1 || location > num {
			location = num
		}
	}

	for i := 0; i < len(seeds); i++ {
		start := i
		i++
		end := seeds[i]

		for j := 0; j < end; j++ {
			seed := seeds[start] + j
			for _, mArr := range mapList {
				newSeed, ok := destination(&mArr, seed)
				if ok {
					seed = newSeed
				}
			}
			if location2 == -1 || location2 > seed {
				location2 = seed
			}
		}
	}

	fmt.Println(location)
	fmt.Println(location2)
}
