package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func processEntry(arr []byte) string {
	var sb strings.Builder

	i := 0
	j := 0

	for i < len(arr) && j < len(arr) {

		dupe := 1

		// Compression
		if arr[i] == 40 {
			i++
			for arr[j] != 41 {
				j++
			}

			// Get chars to dupe
			comp := strings.Split(string(arr[i:j]), "x")
			chars, _ := strconv.Atoi(comp[0])
			dupe, _ = strconv.Atoi(comp[len(comp)-1])

			// Move pointers to chars to dupe
			i = j
			i++
			for range chars {
				j++
			}

		}

		// Write to string builder
		for range dupe {
			sb.Write(arr[i : j+1])
		}

		// Reset pointers
		i = j
		i++
		j++

	}

	return sb.String()

}

func solvePartOne(data []byte) int {

	var count int

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		count += len(processEntry(entry))
	}

	return count

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)
	res := solvePartOne(data)
	fmt.Println(res)

}
