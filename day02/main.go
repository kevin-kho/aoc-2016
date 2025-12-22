package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Dir struct {
	X int
	Y int
}

func createKeypadGrid() [][]int {
	var res [][]int
	res = append(res, []int{1, 2, 3})
	res = append(res, []int{4, 5, 6})
	res = append(res, []int{7, 8, 9})
	return res
}

func createSpecialGrid() [][]string {
	var grid [][]string
	grid = append(grid, []string{"X", "X", "1", "X", "X"})
	grid = append(grid, []string{"X", "2", "3", "4", "X"})
	grid = append(grid, []string{"5", "6", "7", "8", "9"})
	grid = append(grid, []string{"X", "A", "B", "C", "X"})
	grid = append(grid, []string{"X", "X", "D", "X", "X"})

	return grid
}

func createDirs(data []byte) [][]Dir {

	var res [][]Dir

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		var dirs []Dir
		for _, char := range entry {

			var d Dir
			switch string(char) {
			case "U":
				d.Y = -1
			case "D":
				d.Y = 1
			case "L":
				d.X = -1
			case "R":
				d.X = 1
			}
			dirs = append(dirs, d)
		}
		res = append(res, dirs)
	}

	return res

}

func solvePartOne(dirs [][]Dir) int {
	var numbers []int
	grid := createKeypadGrid()
	X := len(grid[0])
	Y := len(grid)
	curr := Dir{
		X: 1,
		Y: 1,
	}
	for _, dir := range dirs {
		for _, d := range dir {
			newX := curr.X + d.X
			newY := curr.Y + d.Y

			// case: out of bounds
			if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
				continue
			}
			curr.X = newX
			curr.Y = newY
		}
		numbers = append(numbers, grid[curr.Y][curr.X])

	}

	var res int
	for _, num := range numbers {
		res = res*10 + num
	}

	return res

}

func solvePartTwo(dirs [][]Dir) string {
	var numbers []string
	grid := createSpecialGrid()
	X := len(grid[0])
	Y := len(grid)
	curr := Dir{
		X: 0,
		Y: 2,
	}

	for _, dir := range dirs {
		for _, d := range dir {
			newX := curr.X + d.X
			newY := curr.Y + d.Y

			// case: out of bounds
			if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
				continue
			}

			// case: we hit an "X"
			if grid[newY][newX] == "X" {
				continue
			}

			curr.X = newX
			curr.Y = newY
		}

		numbers = append(numbers, grid[curr.Y][curr.X])
	}

	return strings.Join(numbers, "")

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	dirs := createDirs(data)
	res := solvePartOne(dirs)
	fmt.Println(res)

	res2 := solvePartTwo(dirs)
	fmt.Println(res2)

}
