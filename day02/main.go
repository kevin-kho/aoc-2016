package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Dir struct {
	X int
	Y int
}

func createGrid() [][]int {
	var res [][]int
	res = append(res, []int{1, 2, 3})
	res = append(res, []int{4, 5, 6})
	res = append(res, []int{7, 8, 9})
	return res
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
	grid := createGrid()
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

}
