package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Numbers struct {
	a, b, prize float64
}

type ClawMachine struct {
	x, y Numbers
}

func isInteger(x float64) bool {
	return math.Trunc(x) == x
}

func findMultiplier(x, y float64) []float64 {
	toX := float64(1)
	toY := x * toX / y
	for !isInteger(toY) {
		toX++
		toY = x * toX / y
	}
	return []float64{toX, toY}
}

func solution(cm ClawMachine) []float64 {
	multiplierse := findMultiplier(cm.x.a, cm.y.a)
	xa := cm.x.a * multiplierse[0]
	xb := cm.x.b * multiplierse[0]
	xPrize := cm.x.prize * multiplierse[0]

	yb := (cm.y.b * multiplierse[1]) - xb
	yPrize := ((cm.y.prize * multiplierse[1]) - xPrize) / yb
	if !isInteger(yPrize) {
		return []float64{-1, -1}
	}

	xb *= yPrize
	xPrize = (xPrize - xb) / xa
	if !isInteger(xPrize) {
		return []float64{-1, -1}
	}

	return []float64{xPrize, yPrize}
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	clawMachines := []ClawMachine{}
	total := float64(0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		cm := ClawMachine{}

		a := strings.Split(text, "+")
		x, err := strconv.Atoi(strings.Split(a[1], ",")[0])
		handleError(err)
		cm.x.a = float64(x)
		y, err := strconv.Atoi(a[2])
		handleError(err)
		cm.y.a = float64(y)

		scanner.Scan()
		text = scanner.Text()

		b := strings.Split(text, "+")
		x, err = strconv.Atoi(strings.Split(b[1], ",")[0])
		handleError(err)
		cm.x.b = float64(x)
		y, err = strconv.Atoi(b[2])
		handleError(err)
		cm.y.b = float64(y)

		scanner.Scan()
		text = scanner.Text()

		p := strings.Split(text, "=")
		x, err = strconv.Atoi(strings.Split(p[1], ",")[0])
		handleError(err)
		cm.x.prize = float64(x + 10000000000000)
		y, err = strconv.Atoi(p[2])
		handleError(err)
		cm.y.prize = float64(y + 10000000000000)

		clawMachines = append(clawMachines, cm)

		scanner.Scan()
	}

	for _, cm := range clawMachines {
		p := solution(cm)
		if p[0] == -1 {
			continue
		}
		total += (p[0] * 3) + p[1]
	}

	fmt.Println("Total", total)
}
