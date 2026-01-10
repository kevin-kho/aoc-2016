package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func IsTrap(input [3]byte) bool {
	isTrap := map[[3]byte]bool{
		{'^', '^', '.'}: true,
		{'.', '^', '^'}: true,
		{'^', '.', '.'}: true,
		{'.', '.', '^'}: true,
	}

	return isTrap[input]
}

func GetNextRowValue(row []byte, i int) byte {

	left := i - 1
	right := i + 1

	var leftValue byte
	var rightValue byte
	if left < 0 {
		leftValue = '.'
	} else {
		leftValue = row[left]
	}

	if right >= len(row) {
		rightValue = '.'
	} else {
		rightValue = row[right]
	}

	tpl := [3]byte{leftValue, row[i], rightValue}
	if IsTrap(tpl) {
		return '^'
	}

	return '.'

}

func GetNextRow(row []byte) []byte {
	var nxtRow []byte
	for i := range row {
		nxtVal := GetNextRowValue(row, i)
		nxtRow = append(nxtRow, nxtVal)
	}

	return nxtRow

}

func CountSafe(row []byte) int {
	var count int
	for _, val := range row {
		if val == '.' {
			count++
		}
	}

	return count
}

func SolvePartOne(row []byte, rowCt int) int {

	var safeCount int
	for range rowCt - 1 {
		safeCount += CountSafe(row)
		row = GetNextRow(row)
	}

	safeCount += CountSafe(row)

	return safeCount

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	res := SolvePartOne(data, 40)
	fmt.Println(res)

	res2 := SolvePartOne(data, 400000)
	fmt.Println(res2)
}
