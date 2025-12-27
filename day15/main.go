package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Disc struct {
	Number   int
	Pos      int
	TotalPos int
}

func (d Disc) CalcFinalPos(sec int) int {
	return (d.Number + d.Pos + sec) % d.TotalPos
}

func GetDiscs(data []byte) ([]Disc, error) {
	var discs []Disc
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		strArr := strings.Split(string(entry), " ")

		num := strArr[1]
		num = strings.TrimPrefix(num, "#")
		numInt, err := strconv.Atoi(num)
		if err != nil {
			return discs, err
		}

		totalPos := strArr[3]
		totalPosInt, err := strconv.Atoi(totalPos)
		if err != nil {
			return discs, err
		}

		pos := strArr[len(strArr)-1]
		pos = strings.TrimSuffix(pos, ".")
		posInt, err := strconv.Atoi(pos)
		if err != nil {
			return discs, err
		}

		discs = append(discs, Disc{
			Number:   numInt,
			Pos:      posInt,
			TotalPos: totalPosInt,
		})

	}

	return discs, nil
}

func main() {
	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	discs, err := GetDiscs(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(discs)

}
