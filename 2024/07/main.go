package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func sArrToiArr(sArr []string) []int {
	res := []int{}

	for _, s := range sArr {
		n, err := strconv.Atoi(s)
		handleError(err)
		res = append(res, n)
	}

	return res
}

type CallibrationEquations struct {
	sum     int
	numbers []int
}

type MathOperator func(int, int) int

func (ce *CallibrationEquations) sumPossible(operators []MathOperator, i, prevSum int) bool {
	if i == len(ce.numbers) {
		return prevSum == ce.sum
	}

	for _, operator := range operators {
		sum := operator(prevSum, ce.numbers[i])
		if sum > ce.sum {
			continue
		}
		if ce.sumPossible(operators, i+1, sum) {
			return true
		}
	}

	return false
}

func (ce *CallibrationEquations) SumPossible(operators []MathOperator) bool {
	return ce.sumPossible(operators, 1, ce.numbers[0])
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	ceArr := []CallibrationEquations{}
	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		split := strings.Split(text, ": ")
		sNumbers := strings.Split(split[1], " ")

		sum, err := strconv.Atoi(split[0])
		handleError(err)
		numbers := sArrToiArr(sNumbers)

		ceArr = append(ceArr, CallibrationEquations{sum, numbers})
	}

	for _, ce := range ceArr {
		if ce.SumPossible([]MathOperator{add, multiply}) {
			sum += ce.sum
		}
	}

	fmt.Println("Sum", sum)
}
