package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func getTrianglesByRow(data []byte) ([][]int, error) {
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

func getTrianglesByColumn(data []byte) ([][]int, error) {
	var triangles [][]int
	col := make(map[int][]int)

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		var triple []int

		entryStr := strings.TrimSpace(string(entry))
		for e := range strings.SplitSeq(entryStr, " ") {
			if len(e) == 0 {
				continue
			}
			val, err := strconv.Atoi(e)
			if err != nil {
				return triangles, err
			}
			triple = append(triple, val)
		}

		for i, v := range triple {
			col[i] = append(col[i], v)
		}

	}

	var values []int
	for _, v := range col {
		values = append(values, v...)
	}

	for i := 2; i < len(values); i += 3 {
		triangles = append(triangles, []int{values[i-2], values[i-1], values[i]})
	}

	return triangles, nil

}

func isValidTriangle(triangle []int) bool {
	a := triangle[0]+triangle[1] > triangle[2]
	b := triangle[1]+triangle[2] > triangle[0]
	c := triangle[0]+triangle[2] > triangle[1]

	return a && b && c
}

func solve(triangles [][]int) int {
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
	triangleRows, err := getTrianglesByRow(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solve(triangleRows)
	fmt.Println(res)

	triangleCols, err := getTrianglesByColumn(data)
	if err != nil {
		log.Fatal(err)
	}

	res2 := solve(triangleCols)
	fmt.Println(res2)

}
