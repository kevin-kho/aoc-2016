package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Coord struct {
	X int
	Y int
}

func CreateGrid(data []byte) [][]byte {
	return bytes.Split(data, []byte{'\n'})
}

func CollectNums(grid [][]byte) map[byte]bool {
	nums := make(map[byte]bool)
	for _, row := range grid {
		for _, col := range row {
			if col >= '0' && col <= '9' {
				nums[col] = true
			}
		}
	}
	return nums
}

func GetStartPoint(grid [][]byte) Coord {
	for y, row := range grid {
		for x, col := range row {
			if col == '0' {
				return Coord{
					X: x,
					Y: y,
				}
			}

		}
	}

	return Coord{}

}

func SolvePartOneDfs(grid [][]byte, nums map[byte]bool, start Coord) int {
	// Stack overflows
	DIRS := []Coord{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}
	Y := len(grid)
	X := len(grid[0])
	res := math.MaxInt
	seen := make(map[[2]int][]byte)
	var dfs func(x int, y int, steps int, curr []byte)
	dfs = func(x int, y int, steps int, curr []byte) {

		// case: we've visited all the nums
		if len(curr) == len(nums) {
			res = min(res, steps)
			return
		}

		// case: out of bounds
		if !(0 <= x && x < X) || !(0 <= y && y < Y) {
			return
		}

		// case: hit a wall
		if grid[y][x] == '#' {
			return
		}

		// case: already visited and the slice are equal
		if v, ok := seen[[2]int{x, y}]; ok {
			if slices.Equal(v, curr) {
				return
			}
		}

		// visit the square
		b := grid[y][x]
		if nums[b] {
			curr = append(curr, b)
		}
		seen[[2]int{x, y}] = slices.Clone(curr)

		// Go visit the other four directions
		for _, d := range DIRS {
			dfs(x+d.X, y+d.Y, steps+1, slices.Clone(curr))
		}

		// unvisit square
		if nums[b] {
			seen[[2]int{x, y}] = curr[:len(curr)-1]
		}

	}
	dfs(start.X, start.Y, 0, []byte{})

	return res
}

func SolvePartOneBfs(grid [][]byte, nums map[byte]bool, start Coord) int {

	DIRS := []Coord{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}
	res := math.MaxInt
	steps := 0
	Y := len(grid)
	X := len(grid[0])
	queue := []Coord{start}
	seen := make(map[Coord]bool)
	stepCount := make(map[byte]int)
	for len(queue) > 0 {
		for range queue {
			node := queue[0]
			queue = queue[1:]
			seen[node] = true
			b := grid[node.Y][node.X]
			if nums[b] {
				stepCount[b] = steps
			}
			// Add children
			for _, d := range DIRS {
				x := node.X + d.X
				y := node.Y + d.Y

				// case: out of bounds
				if !(0 <= x && x < X) || !(0 <= y && y < Y) {
					continue
				}

				// case: it's a wall
				if grid[y][x] == '#' {
					continue
				}

				c := Coord{
					X: x,
					Y: y,
				}

				// case: already seen
				if seen[c] {
					continue
				}

				queue = append(queue, c)

			}

		}
		steps += 1
	}

	fmt.Println(stepCount)

	return res

}

func main() {
	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	grid := CreateGrid(data)
	nums := CollectNums(grid)
	start := GetStartPoint(grid)

	res := SolvePartOneBfs(grid, nums, start)
	fmt.Println(res)
}
