package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func getTriangles(data []byte) ([][]int, error) {
	var res [][]int
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		var triangle []int
		entryStr := strings.TrimSpace(string(entry))
		for e := range strings.SplitSeq(entryStr, " ") {
			if len(e) == 0 {
				continue
			}
			val, err := strconv.Atoi(e)
			if err != nil {
				return res, err

			}
			triangle = append(triangle, val)

		}
		res = append(res, triangle)
	}
	return res, nil

}

func isValidTriangle(triangle []int) bool {
	a := triangle[0]+triangle[1] > triangle[2]
	b := triangle[1]+triangle[2] > triangle[0]
	c := triangle[0]+triangle[2] > triangle[1]

	return a && b && c
}

func solvePartOne(triangles [][]int) int {
	var count int
	for _, tri := range triangles {
		if isValidTriangle(tri) {
			count++
		}
	}

	return count
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)
	triangles, err := getTriangles(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(triangles)
	fmt.Println(res)

}
