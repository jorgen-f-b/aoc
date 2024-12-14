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

type Room struct {
	w, h           int
	lu, ld, ru, rd int
}

func (room *Room) saftyFactor() int {
	return room.ld * room.lu * room.rd * room.ru
}

func (room *Room) add(robot *Robot) {
	halfX, halfY := room.w/2, room.h/2

	if robot.pos.x == halfX || robot.pos.y == halfY {
		return
	}

	if robot.pos.y < halfY {
		if robot.pos.x < halfX {
			room.lu++
		} else {
			room.ru++
		}
	} else {
		if robot.pos.x < halfX {
			room.ld++
		} else {
			room.rd++
		}
	}
}

type Vector struct{ x, y int }

type Robot struct {
	pos Vector
	vel Vector
}

func (r *Robot) print() {
	fmt.Println("Robot", "p", r.pos.x, r.pos.y, "v", r.vel.x, r.vel.y)
}

func (r *Robot) printOnMap(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == r.pos.y && x == r.pos.x {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (r *Robot) walk(w, h int) {
	r.pos.x += r.vel.x
	if r.pos.x >= w {
		r.pos.x -= w
	} else if r.pos.x < 0 {
		r.pos.x += w
	}

	r.pos.y += r.vel.y
	if r.pos.y >= h {
		r.pos.y -= h
	} else if r.pos.y < 0 {
		r.pos.y += h
	}
}

func main() {
	file, err := os.Open("input.txt")
	handleError(err)

	w := 0
	h := 0
	robots := []Robot{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		data := strings.Split(text, " ")

		if w == 0 {
			width, err := strconv.Atoi(data[0])
			handleError(err)
			w = width
			height, err := strconv.Atoi(data[1])
			h = height
			continue
		}

		sPos := strings.Split(data[0][2:], ",")
		px, err := strconv.Atoi(sPos[0])
		handleError(err)
		py, err := strconv.Atoi(sPos[1])
		handleError(err)

		sVec := strings.Split(data[1][2:], ",")
		vx, err := strconv.Atoi(sVec[0])
		handleError(err)
		vy, err := strconv.Atoi(sVec[1])
		handleError(err)

		robots = append(robots, Robot{Vector{px, py}, Vector{vx, vy}})
	}

	room := Room{w, h, 0, 0, 0, 0}

	for _, r := range robots {
		for i := 0; i < 100; i++ {
			r.walk(w, h)
		}
		room.add(&r)
	}

	fmt.Println("Safty Factor", room.saftyFactor())
}
