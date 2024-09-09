package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var letterNums = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func toLetterNum(letterNum string) (string, bool) {
	var numb, ok = letterNums[letterNum]
	return numb, ok
}

func inLetterNum(letterNum string) bool {
	for _, num := range numbers {
		if strings.Contains(num, letterNum) {
			return true
		}
	}
	return false
}

func removeRunes(letterNum string, reversed bool) string {
	for i := range letterNum {
		var newString string
		if reversed {
			newString = letterNum[:len(letterNum)-i]
		} else {
			newString = letterNum[i:]
		}
		if inLetterNum(newString) {
			return newString
		}
	}
	return ""
}

func main() {
	var file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var number = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var num = ""
		var letterNum = ""

		for _, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)
				letterNum = ""
				break
			}
			letterNum += string(char)
			if inLetterNum(letterNum) {
				var numb, ok = toLetterNum(letterNum)
				if ok {
					num += numb
					letterNum = ""
					break
				}
			} else {
				letterNum = removeRunes(letterNum, false)
			}
		}

		var runeLine = []rune(line)
		var length = len(runeLine) - 1

		for i := range runeLine {
			var r = runeLine[length-i]
			if unicode.IsDigit(r) {
				num += string(r)
				letterNum = ""
				break
			}
			letterNum = string(r) + letterNum
			if inLetterNum(letterNum) {
				var numb, ok = toLetterNum(letterNum)
				if ok {
					num += numb
					letterNum = ""
					break
				}
			} else {
				letterNum = removeRunes(letterNum, true)
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
