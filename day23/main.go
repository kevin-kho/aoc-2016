package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Instruction struct {
	Command string
	Src     string
	Dst     string
	Value   int
	Offset  int
}

func GetToggledInstruction(in Instruction) Instruction {
	var cmd string
	switch in.Command {
	case "cpyReg":
		cmd = "jnzReg"
	case "cpyVal":
		cmd = "jnzVal"
	case "jnzVal":
		cmd = "cpyVal"
	case "jnzReg":
		cmd = "cpyReg"
	case "inc":
		cmd = "dec"
	case "dec":
		cmd = "inc"
	}

	in.Command = cmd

	return in
}

func CreateInstructions(data []byte) ([]Instruction, error) {
	var res []Instruction
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStrArr := strings.Split(string(entry), " ")
		cmd := entryStrArr[0]
		var src string
		var dst string
		var val int
		var offset int
		var err error
		switch cmd {
		case "cpy":
			val, err = strconv.Atoi(entryStrArr[1])
			cmd = "cpyVal"
			if err != nil {
				src = entryStrArr[1]
				cmd = "cpyReg"
			}
			dst = entryStrArr[len(entryStrArr)-1]
		case "dec":
			src = entryStrArr[len(entryStrArr)-1]
		case "inc":
			src = entryStrArr[len(entryStrArr)-1]
		case "tgl":
			offset, err = strconv.Atoi(entryStrArr[len(entryStrArr)-1])
			if err != nil {
				src = entryStrArr[len(entryStrArr)-1]
				cmd = "tglReg"
			}
		case "jnz":
			val, err = strconv.Atoi(entryStrArr[1])
			cmd = "jnzVal"
			if err != nil {
				src = entryStrArr[1]
				cmd = "jnzReg"
			}
			offset, err = strconv.Atoi(entryStrArr[len(entryStrArr)-1])
			if err != nil {
				src = entryStrArr[len(entryStrArr)-1]
			}
		}

		res = append(res, Instruction{
			Command: cmd,
			Src:     src,
			Dst:     dst,
			Value:   val,
			Offset:  offset,
		})

	}
	return res, nil
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	in, err := CreateInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range in {
		fmt.Printf("%+v\n", i)
	}

}
