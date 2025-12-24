package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Destination struct {
	Kind Type
	Num  int
}

type Type int

const (
	Bot Type = iota
	Output
)

func PopulateBotInventory(data []byte) (map[int][]int, error) {
	inventory := map[int][]int{}
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

func solvePartOne(inventory map[int][]int, instructions map[int][][2]Destination) {

	output := make(map[int][]int)

	var queue []int
	for bot, inv := range inventory {
		if len(inv) > 1 {
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
				if slices.Contains(botInventory, 17) && slices.Contains(botInventory, 61) {
					fmt.Println(bot)
				}

				var low int
				var high int
				if botInventory[0] > botInventory[1] {
					high = botInventory[0]
					low = botInventory[1]
				} else {
					high = botInventory[1]
					low = botInventory[0]
				}

				lowTargetKind := instruct[0].Kind
				lowTargetNum := instruct[0].Num
				if lowTargetKind == Bot {
					inventory[lowTargetNum] = append(inventory[lowTargetNum], low)
					if len(inventory[lowTargetNum]) > 1 {
						queue = append(queue, lowTargetNum)
					}
				} else {
					output[lowTargetNum] = append(output[lowTargetNum], low)

				}

				highTargetKind := instruct[1].Kind
				highTargetNum := instruct[1].Num
				if highTargetKind == Bot {
					inventory[highTargetNum] = append(inventory[highTargetNum], high)
					if len(inventory[highTargetNum]) > 1 {
						queue = append(queue, highTargetNum)
					}

				} else {
					output[highTargetNum] = append(output[highTargetNum], high)
				}

			}

		}

	}

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

	instructions, err := GetBotInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(inventory, instructions)

}
