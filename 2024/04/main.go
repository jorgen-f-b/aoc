package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("example.txt")
	handleError(err)

	reg, err := regexp.Compile(`XMAS|SAMX`)
	handleError(err)
	reg2, err := regexp.Compile(`X.{10}M.{10}A.{10}S|S.{10}A.{10}M.{10}X`)
	handleError(err)

	sideway := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		sideway = append(sideway, text)
	}

	matches1 := reg.FindAllStringIndex(strings.Join(sideway, ","), -1)
	matches2 := reg2.FindAllStringIndex(strings.Join(sideway, ","), -1)

	// standing := []string{}
	// crossRight := []string{}

	// for i, linje := range sideway {
	// 	matches := reg.FindAllStringIndex(linje, -1)
	// 	count += len(matches)

	// 	for j, r := range linje {
	// 		if i == 0 {
	// 			standing = append(standing, string(r))
	// 			crossRight = append(crossRight, string(r))
	// 			continue
	// 		}

	// 		standing[j] = standing[j] + string(r)
	// 		if j == 0 {
	// 			crossRight = append(crossRight, string(r))
	// 		} else {
	// 			if j >= i {
	// 				crossRight[j-i] = crossRight[j-i] + string(r)
	// 			} else {
	// 				crossRight[i-j] = crossRight[i-j] + string(r)
	// 			}
	// 		}
	// 	}
	// }

	// for _, linje := range standing {
	// 	matches := reg.FindAllStringIndex(linje, -1)
	// 	count += len(matches)
	// }
	// for _, linje := range crossRight {
	// 	matches := reg.FindAllStringIndex(linje, -1)
	// 	count += len(matches)
	// }

	// fmt.Println(sideway)
	// fmt.Println(crossRight)
	fmt.Println("Count", len(matches1)+len(matches2))
}
