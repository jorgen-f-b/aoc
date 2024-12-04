package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func stringToIntSlice(arr []string) []int {
	intArr := make([]int, 0, len(arr))

	for _, s := range arr {
		n, err := strconv.Atoi(s)
		handleError(err)

		intArr = append(intArr, n)
	}

	return intArr
}

func removeFromArr(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func checkLevel(arr []int, remove bool) bool {
	prevLevel := -1
	rising := 0
	faultIndex := -1
	safe := true

	for i, level := range arr {
		if prevLevel == -1 {
			prevLevel = level
			continue
		}

		diff := prevLevel - level

		if diff > 3 || diff < -3 {
			faultIndex = i
			safe = false
			break
		}

		if diff == 0 {
			faultIndex = i
			safe = false
			break
		} else if diff < 0 {
			if rising == 1 {
				faultIndex = i
				safe = false
				break
			}
			if rising == 0 {
				rising = -1
			}
		} else if diff > 0 {
			if rising == -1 {
				faultIndex = i
				safe = false
				break
			}
			if rising == 0 {
				rising = 1
			}
		}

		prevLevel = level
	}

	if !safe && remove {
		for i := faultIndex - 2; i <= faultIndex+2 && i < len(arr)-2; i++ {
			newArr := removeFromArr(arr, i)
			safe = checkLevel(newArr, false)
			fmt.Println(i, newArr)
			if safe {
				break
			}
		}
	}

	return safe
}

func main() {
	file, err := os.Open("example.txt")
	handleError(err)

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		list := stringToIntSlice(strings.Split(text, " "))

		safe := checkLevel(list, true)
		if safe {
			sum++
		}
	}

	fmt.Println("SafeAmount", sum)
}
