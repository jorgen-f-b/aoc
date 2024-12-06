package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Rule map[string][]string

func checkBackwards(rules Rule, page []string, firstIndex int, secondIndex int) bool {
	for i := firstIndex; i >= 0; i-- {
		first := page[i]
		second := page[secondIndex]

		list := rules[first]
		if slices.Contains(list, second) {
			return true
		}

		list = rules[second]
		if slices.Contains(list, first) {
			return false
		}
	}

	return true
}

func correctOrdering(rules Rule, page []string) bool {
	for i := 0; i < len(page)-1; i++ {
		first := page[i]
		second := page[i+1]

		list := rules[first]
		if slices.Contains(list, second) {
			continue
		}

		list = rules[second]
		if slices.Contains(list, first) {
			return false
		}

		if !checkBackwards(rules, page, i, i+1) {
			return false
		}
	}

	return true
}

func moveIndex(list []string, from, to int) {
	tmp := ""
	for i := to; i > from; i-- {
		tmp = list[i]
		list[i] = list[i-1]
		list[i-1] = tmp
	}
}

func fixBackwards(rules Rule, page []string, to int) {
	for i := 0; i < to; i++ {
		first := page[i]
		second := page[to]

		list := rules[second]
		if slices.Contains(list, first) {
			moveIndex(page, i, to)
		}
	}
}

func fixOrdering(rules Rule, page []string) {
	for i := 0; i < len(page)-1; i++ {
		first := page[i]
		second := page[i+1]

		list := rules[first]
		if slices.Contains(list, second) {
			continue
		}

		fixBackwards(rules, page, i+1)
	}
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	firstSection := true
	rules := make(Rule)
	pages := [][]string{}
	wrongOrderedPages := [][]string{}
	sum := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			firstSection = false
			continue
		}

		if firstSection {
			split := strings.Split(text, "|")
			numbers, ok := rules[split[0]]
			if ok {
				numbers = append(numbers, split[1])
				rules[split[0]] = numbers
			} else {
				rules[split[0]] = []string{split[1]}
			}
		} else {
			pages = append(pages, strings.Split(text, ","))
		}
	}

	for _, page := range pages {
		if correctOrdering(rules, page) {
			middleIndex := len(page) / 2
			num, err := strconv.Atoi(page[middleIndex])
			handleError(err)
			sum += num
		} else {
			wrongOrderedPages = append(wrongOrderedPages, page)
		}
	}

	for _, page := range wrongOrderedPages {
		fixOrdering(rules, page)
		middleIndex := len(page) / 2
		num, err := strconv.Atoi(page[middleIndex])
		handleError(err)
		sum2 += num
	}

	fmt.Println("Sum", sum)
	fmt.Println("Sum2", sum2)
}
