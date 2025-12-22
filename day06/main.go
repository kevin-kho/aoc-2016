package main

import (
	"bytes"
	"fmt"
	"log"
	"math"

	"github.com/kevin-kho/aoc-utilities/common"
)

func createFreqMap(data []byte) map[int]map[byte]int {
	freq := make(map[int]map[byte]int)
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		for i, char := range entry {
			if freq[i] == nil {
				freq[i] = make(map[byte]int)
			}
			freq[i][char]++
		}
	}

	return freq

}

func solvePartOne(freq map[int]map[byte]int) string {
	res := make([]byte, len(freq))
	for i, f := range freq {
		var currByte byte
		var currMax int
		for b, ct := range f {
			if ct > currMax {
				currMax = ct
				currByte = b
			}
		}

		res[i] = currByte

	}

	return string(res)

}

func solvePartTwo(freq map[int]map[byte]int) string {
	res := make([]byte, len(freq))
	for i, f := range freq {
		var currByte byte
		currMax := math.MaxInt
		for b, ct := range f {
			if ct < currMax {
				currMax = ct
				currByte = b
			}
		}

		res[i] = currByte

	}

	return string(res)

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	freqMap := createFreqMap(data)

	res := solvePartOne(freqMap)
	fmt.Println(res)

	res2 := solvePartTwo(freqMap)
	fmt.Println(res2)

}
