package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Node struct {
	X           int
	Y           int
	Size        int
	Used        int
	Available   int
	UsedPercent int
}

func CreateNodes(data []byte) ([]Node, error) {
	var nodes []Node
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		entryStrArr := strings.Split(string(entry), " ")
		// remove empty strings
		var entryStrArrPruned []string
		for _, e := range entryStrArr {
			if e != "" {
				entryStrArrPruned = append(entryStrArrPruned, e)
			}
		}

		nodePath := strings.TrimPrefix(entryStrArrPruned[0], "/dev/grid/node-")
		coords := strings.Split(nodePath, "-")
		x, err := strconv.Atoi(strings.TrimPrefix(coords[0], "x"))
		if err != nil {
			return nodes, err
		}
		y, err := strconv.Atoi(strings.TrimPrefix(coords[1], "y"))
		if err != nil {
			return nodes, err
		}

		size, err := strconv.Atoi(strings.TrimSuffix(entryStrArrPruned[1], "T"))
		if err != nil {
			return nodes, err
		}

		used, err := strconv.Atoi(strings.TrimSuffix(entryStrArrPruned[2], "T"))
		if err != nil {
			return nodes, err
		}

		avail, err := strconv.Atoi(strings.TrimSuffix(entryStrArrPruned[3], "T"))
		if err != nil {
			return nodes, err
		}

		usedPct, err := strconv.Atoi(strings.TrimSuffix(entryStrArrPruned[4], "%"))
		if err != nil {
			return nodes, err
		}

		node := Node{
			X:           x,
			Y:           y,
			Size:        size,
			Used:        used,
			Available:   avail,
			UsedPercent: usedPct,
		}

		nodes = append(nodes, node)

	}

	return nodes, nil

}

func SolvePartOne(nodes []Node) int {
	var count int
	for i, a := range nodes {
		if a.Used == 0 {
			continue
		}

		for j, b := range nodes {
			if i == j {
				continue
			}
			if b.Available >= a.Used {
				count++
			}
		}

	}

	return count

}

func main() {
	filePath := "input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	nodes, err := CreateNodes(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(nodes)
	fmt.Println(res)

}
