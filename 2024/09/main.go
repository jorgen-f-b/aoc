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

func swapBlocks(arr []int, from, to, size int) {
	for i := 0; i < size; i++ {
		tmp := arr[from+i]
		arr[from+i] = arr[to+i]
		arr[to+i] = tmp
	}
}

func compatctIntArr(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	id := -1
	size := 0
	freeSize := 0

	for j := len(res) - 1; j > 0; j-- {
		if res[j] == -1 && id == -1 {
			continue
		}

		if id == -1 {
			id = res[j]
		}
		size++

		if res[j-1] != id {
			for i := 0; i < j; i++ {
				if res[i] == -1 {
					freeSize++
					if freeSize == size {
						swapBlocks(res, i-freeSize+1, j, size)
						break
					}
				} else {
					freeSize = 0
				}
			}
			freeSize = 0
			size = 0
			id = -1
		}
	}

	return res
}

func toIntArr(disc []Data) []int {
	res := []int{}
	for _, data := range disc {
		for i := 0; i < data.size; i++ {
			res = append(res, data.id)
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
	sum2 := 0

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

	for i, data := range compatctIntArr(toIntArr(disc)) {
		if data == -1 {
			continue
		}
		sum2 += data * i
	}

	fmt.Println("Sum", sum)
	fmt.Println("Sum2", sum2)
}
