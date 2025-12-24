package main

import (
	"bytes"
	"container/heap"
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

func PopulateBotInventory(data []byte) (map[int]IntHeap, error) {
	inventory := map[int]IntHeap{}
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

		inventory[bot] = append(inventory[bot], value)

	}

	return inventory, nil

}

func GetBotInstructions(data []byte) (map[int][][2]int, error) {
	instructions := map[int][][2]int{}
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		if !strings.HasPrefix(string(entry), "bot") {
			continue
		}

		entryStr := strings.Split(string(entry), " ")
		bot, err := strconv.Atoi(entryStr[1])
		if err != nil {
			return instructions, err
		}

		low, err := strconv.Atoi(entryStr[6])
		if err != nil {
			return instructions, err
		}
		high, err := strconv.Atoi(entryStr[11])
		if err != nil {
			return instructions, err
		}

		instructions[bot] = append(instructions[bot], [2]int{low, high})

	}
	return instructions, nil
}

func HeapifyInventories(inv map[int]IntHeap) {
	for _, h := range inv {
		heap.Init(&h)
	}

}

func solvePartOne(inventory map[int]IntHeap, instructions map[int][][2]int) {

	var queue []int
	for bot, inv := range inventory {
		if inv.Len() > 1 {
			queue = append(queue, bot)
		}
	}

	// BFS with queue
	// for len(queue) > 0 {
	// for range len(queue) {
	// bot := queue[0]
	// queue = queue[1:]

	// }

	// }

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
