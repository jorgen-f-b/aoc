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
	newArr := make([]int, 0, len(arr)-1)
	if i == 0 {
		newArr = append(newArr, arr[i+1:]...)
	} else if i == len(arr)-1 {
		newArr = append(newArr, arr[:i]...)
	} else {
		newArr = append(newArr, arr[:i]...)
		newArr = append(newArr, arr[i+1:]...)
	}
	return newArr
}

func checkLevel(arr []int, remove bool) bool {
	rising := 0
	faultIndex := -1
	safe := true

	for i, level := range arr {
		if i == len(arr)-1 {
			break
		}

		diff := level - arr[i+1]

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
	}

	if !safe && remove {
		for i := faultIndex - 1; i <= faultIndex+1 && i < len(arr); i++ {
			if i < 0 {
				i = 0
			}
			newArr := removeFromArr(arr, i)
			safe = checkLevel(newArr, false)
			if safe {
				break
			}
		}
	}

	return safe
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	sum1 := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		list := stringToIntSlice(strings.Split(text, " "))

		safe1 := checkLevel(list, false)
		if safe1 {
			sum1++
		}

		safe2 := checkLevel(list, true)
		if safe2 {
			sum2++
		}
	}

	fmt.Println("SafeAmount1", sum1)
	fmt.Println("SafeAmount2", sum2)
}
