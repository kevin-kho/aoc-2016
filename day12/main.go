package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Action int

const (
	CPY Action = iota
	INC
	DEC
	JNZ
)

type Instruction struct {
	Action Action // replace with enum
	Value  int
	Source string // used only when copying from value of register
	Target string
}

func CreateInstructions(data []byte) ([]Instruction, error) {
	var instructions []Instruction
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		entryStr := strings.Split(string(entry), " ")

		var action Action
		var value int
		var source string
		var target string
		var err error
		switch entryStr[0] {
		case "cpy":
			action = CPY
			value, err = strconv.Atoi(entryStr[1])
			if err != nil {
				source = entryStr[1]
			}
			target = entryStr[len(entryStr)-1]

		case "inc":
			action = INC
			target = entryStr[len(entryStr)-1]
		case "dec":
			action = DEC
			target = entryStr[len(entryStr)-1]
		case "jnz":
			action = JNZ
			target = entryStr[1]
			value, err = strconv.Atoi(entryStr[len(entryStr)-1])
		}

		// if err != nil {
		// 	return instructions, err
		// }

		instruct := Instruction{
			Action: action,
			Value:  value,
			Source: source,
			Target: target,
		}
		instructions = append(instructions, instruct)

	}

	return instructions, nil
}

func solvePartOne(instructions []Instruction) {
	register := make(map[string]int)
	register["1"] = 1 // TODO: fix

	// Use pointer and while loop
	var i int
	for i < len(instructions) {
		cmd := instructions[i]
		switch cmd.Action {
		case CPY:
			if cmd.Source != "" {
				register[cmd.Target] = register[cmd.Source]
			} else {
				register[cmd.Target] = cmd.Value
			}
		case INC:
			register[cmd.Target]++
		case DEC:
			register[cmd.Target]--
		case JNZ:
			if register[cmd.Target] == 0 {
				break
			}
			i += cmd.Value
			continue

		}

		i++

	}
	fmt.Println(register)

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	instructions, err := CreateInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(instructions)

}
