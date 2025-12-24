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
	Target string
}

func CreateInstructions(data []byte) ([]Instruction, error) {
	var instructions []Instruction
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		entryStr := strings.Split(string(entry), " ")

		var action Action
		var value int
		var target string
		var err error
		switch entryStr[0] {
		case "cpy":
			action = CPY
			value, err = strconv.Atoi(entryStr[1])
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

		if err != nil {
			return instructions, err
		}

		instruct := Instruction{
			Action: action,
			Value:  value,
			Target: target,
		}
		instructions = append(instructions, instruct)

	}

	return instructions, nil
}

func main() {
	filePath := "./inputExample.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	instructions, err := CreateInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instructions)

}
