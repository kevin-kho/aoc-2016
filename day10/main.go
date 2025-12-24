package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Destination struct {
	Kind Type
	Num  int
}

type Type int

const (
	Bot Type = iota
	Output
)

func PopulateBotInventory(data []byte) (map[int]*IntHeap, error) {
	inventory := map[int]*IntHeap{}
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		if !strings.HasPrefix(string(entry), "value") {
			continue
		}

		entryStr := strings.Split(string(entry), " ")

		value, err := strconv.Atoi(entryStr[1])
		if err != nil {
			return inventory, err
		}
		bot, err := strconv.Atoi(entryStr[len(entryStr)-1])
		if err != nil {
			return inventory, err
		}
		if inventory[bot] == nil {
			inventory[bot] = &IntHeap{}
			heap.Init(inventory[bot])
		}

		heap.Push(inventory[bot], value)

	}

	return inventory, nil

}

func GetBotInstructions(data []byte) (map[int][][2]Destination, error) {
	instructions := map[int][][2]Destination{}
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		if !strings.HasPrefix(string(entry), "bot") {
			continue
		}

		entryStr := strings.Split(string(entry), " ")
		bot, err := strconv.Atoi(entryStr[1])
		if err != nil {
			return instructions, err
		}

		lowType := entryStr[5]
		low, err := strconv.Atoi(entryStr[6])
		if err != nil {
			return instructions, err
		}

		highType := entryStr[10]
		high, err := strconv.Atoi(entryStr[11])
		if err != nil {
			return instructions, err
		}

		var lowDest Destination
		var highDest Destination

		lowDest.Num = low
		if lowType == "bot" {
			lowDest.Kind = Bot
		} else {
			lowDest.Kind = Output
		}

		highDest.Num = high
		if highType == "bot" {
			highDest.Kind = Bot
		} else {
			highDest.Kind = Output
		}

		instructions[bot] = append(instructions[bot], [2]Destination{lowDest, highDest})

	}
	return instructions, nil
}

func HeapifyInventories(inv map[int]*IntHeap) {
	for _, h := range inv {
		heap.Init(h)
	}

}

func solvePartOne(inventory map[int]*IntHeap, instructions map[int][][2]Destination) {

	output := make(map[int][]int)

	var queue []int
	for bot, inv := range inventory {
		if inv.Len() > 1 {
			queue = append(queue, bot)
		}
	}

	// BFS with queue
	for len(queue) > 0 {
		for range len(queue) {
			bot := queue[0]
			queue = queue[1:]

			for _, instruct := range instructions[bot] {
				botInventory := inventory[bot]

				low := heap.Pop(botInventory).(int)
				lowTargetKind := instruct[0].Kind
				lowTargetNum := instruct[0].Num
				if lowTargetKind == Bot {
					heap.Push(inventory[lowTargetNum], low)
					if inventory[lowTargetNum].Len() > 1 {
						queue = append(queue, lowTargetNum)
					}
				} else {
					output[lowTargetNum] = append(output[lowTargetNum], low)

				}

				high := heap.Pop(botInventory).(int)
				fmt.Println(bot, low, high)
				highTargetKind := instruct[1].Kind
				highTargetNum := instruct[1].Num
				if highTargetKind == Bot {
					heap.Push(inventory[highTargetNum], high)
					if inventory[highTargetNum].Len() > 1 {
						queue = append(queue, highTargetNum)
					}

				} else {
					output[highTargetNum] = append(output[highTargetNum], high)
				}

			}

		}

	}

	// fmt.Println(inventory)
	// fmt.Println(output)

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	inventory, err := PopulateBotInventory(data)
	if err != nil {
		log.Fatal(err)
	}

	HeapifyInventories(inventory)

	instructions, err := GetBotInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(inventory, instructions)

}
