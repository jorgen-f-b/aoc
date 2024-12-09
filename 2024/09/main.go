package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Data struct {
	id, size int
}

func compact(disc []Data) []Data {
	res := make([]Data, len(disc))
	copy(res, disc)

	i, j := 0, len(res)-1
	for {
		for res[i].id != -1 {
			i++
		}
		for res[j].id == -1 {
			j--
		}

		if i > j {
			break
		}

		free := res[i]
		data := res[j]

		if free.size == data.size {
			res[i] = data
			res[j] = free
		} else if free.size > data.size {
			res[i].size -= data.size
			res[j].id = -1
			res = slices.Insert(res, i, data)
		} else {
			res[i].id = res[j].id
			res[j].size = res[j].size - res[i].size
		}
	}

	return res
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()

	disc := []Data{}
	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		id := 0
		for i, r := range text {
			size, err := strconv.Atoi(string(r))
			handleError(err)

			if i%2 == 0 {
				if size > 0 {
					disc = append(disc, Data{id, size})
				}
				id++
			} else {
				if size > 0 {
					disc = append(disc, Data{-1, size})
				}
			}
		}
	}

	index := 0
	for _, data := range compact(disc) {
		for i := 0; i < data.size; i++ {
			if data.id != -1 {
				sum += index * data.id
			}
			index++
		}
	}

	fmt.Println("Sum", sum)
}
