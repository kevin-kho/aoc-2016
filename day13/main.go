package main

import (
	"fmt"
	"math"
)

type Pos struct {
	X int
	Y int
}

func IsWall(x int, y int, favNum int) bool {
	num := (x * x) + (3 * x) + (2 * x * y) + y + (y * y)
	num += favNum

	// find binary rep
	var bits int
	for num > 0 {
		bits += num & 1
		num = num >> 1
	}

	if bits%2 == 1 {
		return true
	}

	return false

}

func IsOutOfBounds(x int, y int) bool {
	return x < 0 || y < 0
}

func PossibleDirs() []Pos {
	return []Pos{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

}

func solvePartOne(favNum int, dst Pos) int {
	dirs := PossibleDirs()

	seen := make(map[Pos]bool)
	res := math.MaxInt

	var move func(x, y, steps int)
	move = func(x, y, steps int) {
		// case: out of bounds
		if IsOutOfBounds(x, y) {
			return
		}

		// case: hit a wall
		if IsWall(x, y, favNum) {
			return
		}

		// case: already visited this cell
		if seen[Pos{X: x, Y: y}] {
			return
		}

		// case: hit goal
		if x == dst.X && y == dst.Y {
			res = min(steps, res)
			return
		}

		seen[Pos{X: x, Y: y}] = true
		for _, d := range dirs {
			move(x+d.X, y+d.Y, steps+1)
		}
		delete(seen, Pos{X: x, Y: y})

	}
	move(1, 1, 0)

	return res

}

func solvePartTwo(favNum int, maxSteps int) int {

	possible := make(map[Pos]bool)
	seen := make(map[Pos]bool)
	dirs := PossibleDirs()
	var move func(x, y, steps int)
	move = func(x, y, steps int) {
		// case: out of bounds
		if IsOutOfBounds(x, y) {
			return
		}

		// case: hit a wall
		if IsWall(x, y, favNum) {
			return
		}

		// case: already visited this cell
		if seen[Pos{X: x, Y: y}] {
			return
		}

		if steps > maxSteps {
			return
		}

		possible[Pos{X: x, Y: y}] = true
		seen[Pos{X: x, Y: y}] = true
		for _, d := range dirs {
			move(x+d.X, y+d.Y, steps+1)
		}
		delete(seen, Pos{X: x, Y: y})
	}
	move(1, 1, 0)

	return len(possible)
}

func main() {

	res := solvePartOne(1364, Pos{X: 31, Y: 39})
	fmt.Println(res)

	res2 := solvePartTwo(1364, 50)
	fmt.Println(res2)

}
