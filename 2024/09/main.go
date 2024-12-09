package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("example.txt")
	handleError(err)
	defer file.Close()

	disc := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		id := 0
		for i, r := range text {
			size, err := strconv.Atoi(string(r))
			handleError(err)

			if i%2 == 0 {
				for i := 0; i < size; i++ {
					disc += strconv.Itoa(id)
				}
				id++
			} else {
				for i := 0; i < size; i++ {
					disc += "."
				}
			}
		}
	}

	fmt.Println(disc)
}
