package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func correctWord(s string) bool {
	return s == "MAS" || s == "SAM"
}

func searchForX(arr []string) bool {
	word1 := string(arr[0][0]) + string(arr[1][1]) + string(arr[2][2])
	word2 := string(arr[0][2]) + string(arr[1][1]) + string(arr[2][0])
	return correctWord(word1) && correctWord(word2)
}

func searchForWord(arr []string, row int, col int, searchWord string) int {
	liggende := ""
	lf := 0
	staende := ""
	sf := 0
	skraH := ""
	shf := 0
	skraV := ""
	svf := 0

	line := arr[row]
	for i := col; i < len(line); i++ {
		liggende += string(line[i])
		if !strings.Contains(searchWord, liggende) {
			break
		}
		if searchWord == liggende {
			lf++
			break
		}
	}

	j := 0
	for i := row; i < len(arr) && i < row+len(searchWord); i++ {
		staende += string(arr[i][col])
		if col+j < len(arr[i]) {
			skraH += string(arr[i][col+j])
		}
		if col-j >= 0 {
			skraV += string(arr[i][col-j])
		}

		if !strings.Contains(searchWord, staende) && !strings.Contains(searchWord, skraH) && !strings.Contains(searchWord, skraV) {
			break
		}

		if sf == 0 && searchWord == staende {
			sf++
		}
		if shf == 0 && searchWord == skraH {
			shf++
		}
		if svf == 0 && searchWord == skraV {
			svf++
		}
		j++
	}

	return lf + sf + shf + svf
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	strArr := []string{}
	sum := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		strArr = append(strArr, text)
	}

	for i, linje := range strArr {
		for j, r := range linje {
			if r == 'X' {
				sum += searchForWord(strArr, i, j, "XMAS")
			}
			if r == 'S' {
				sum += searchForWord(strArr, i, j, "SAMX")
			}
			if (r == 'M' || r == 'S') && i+2 < len(strArr) && j+2 < len(linje) {
				if searchForX([]string{strArr[i][j : j+3], strArr[i+1][j : j+3], strArr[i+2][j : j+3]}) {
					sum2++
				}
			}
		}
	}

	fmt.Println("Sum", sum)
	fmt.Println("Sum2", sum2)
}
