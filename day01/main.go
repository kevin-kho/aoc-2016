package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Cardinal int

const (
	North Cardinal = iota
	East
	South
	West
)

type Command struct {
	Turn   Direction
	Number int
}

type Pos struct {
	X    int
	Y    int
	Card Cardinal
}

func (p *Pos) Move(cmd Command) {
	switch p.Card {
	case North:
		if cmd.Turn == Left {
			p.Card = West
		} else {
			p.Card = East
		}
	case East:
		if cmd.Turn == Left {
			p.Card = North
		} else {
			p.Card = South
		}
	case South:
		if cmd.Turn == Left {
			p.Card = East
		} else {
			p.Card = West
		}
	case West:
		if cmd.Turn == Left {
			p.Card = South
		} else {
			p.Card = North
		}
	}

	var dx int
	var dy int
	switch p.Card {
	case North:
		dy = cmd.Number
	case East:
		dx = -cmd.Number
	case South:
		dy = -cmd.Number
	case West:
		dx = cmd.Number
	}

	p.X += dx
	p.Y += dy

}

func (p Pos) GetDistance() int {
	return common.IntAbs(p.X) + common.IntAbs(p.Y)
}

func getCommands(data []byte) ([]Command, error) {
	var res []Command

	for entry := range strings.SplitSeq(string(data), ",") {
		entry = strings.TrimSpace(entry)

		var d Direction

		switch string(entry[0]) {
		case "R":
			d = Right
		default:
			d = Left
		}

		n, err := strconv.Atoi(entry[1:])
		if err != nil {
			return res, err
		}

		res = append(res, Command{
			Turn:   d,
			Number: n,
		})

	}

	return res, nil

}

func solvePartOne(cmds []Command) int {
	var pos Pos
	for _, cmd := range cmds {
		pos.Move(cmd)
	}

	return pos.GetDistance()

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	cmds, err := getCommands(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(cmds)
	fmt.Println(res)

}
