package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var file, err = os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var scanner = bufio.NewScanner(file)
	var number int = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var num = ""

		for _, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)
				break
			}
		}

		var runeLine = []rune(line)
		var length = len(runeLine) - 1

		for i := range runeLine {
			var r = runeLine[length-i]
			if unicode.IsDigit(r) {
				num += string(r)
				break
			}
		}

		var i, _ = strconv.Atoi(num)
		number += i
	}

	fmt.Println(number)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
