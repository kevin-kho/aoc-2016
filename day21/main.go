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

type Password struct {
	Word []byte
}

func (p *Password) SwapPos(x, y int) {
	p.Word[x], p.Word[y] = p.Word[y], p.Word[x]
}

func (p *Password) SwapLtr(x, y byte) {
	xIdx := slices.Index(p.Word, x)
	yIdx := slices.Index(p.Word, y)
	p.Word[xIdx], p.Word[yIdx] = p.Word[yIdx], p.Word[xIdx]
}

func (p *Password) ReversePos(x, y int) {
	l := min(x, y)
	r := max(x, y)
	for l < r {
		p.Word[l], p.Word[r] = p.Word[r], p.Word[l]
		l++
		r--
	}
}

func (p *Password) RotateLeft(x int) {
	orig := slices.Clone(p.Word)
	for i, char := range orig {
		newI := (i - x) % len(orig)
		if newI < 0 {
			newI += len(orig)
		}
		p.Word[newI] = char
	}

}

func (p *Password) RotateRight(x int) {
	orig := slices.Clone(p.Word)
	for i, char := range orig {
		newI := (i + x) % len(orig)
		p.Word[newI] = char
	}
}

func (p *Password) RotateByLtr(x byte) {
	xIdx := slices.Index(p.Word, x)
	n := 1 + xIdx
	if xIdx >= 4 {
		n += 1
	}

	p.RotateRight(n)
}

func (p *Password) MovePos(src, dst int) {
	char := p.Word[src]
	if dst == len(p.Word)-1 {
		p.Word = append(p.Word, char)
	} else {
		p.Word = slices.Insert(p.Word, dst, char)
	}

	remove := dst
	if src < dst {
		// Remove first occurance
		for i := 0; i < len(p.Word); i++ {
			if p.Word[i] == char {
				remove = i
				break
			}
		}
	} else {
		// Remove last occurance
		for i := len(p.Word) - 1; i >= 0; i-- {
			if p.Word[i] == char {
				remove = i
				break
			}
		}
	}

	p.Word = slices.Delete(p.Word, remove, remove+1)

}

func ConstructPassword(data []byte) Password {
	return Password{
		Word: data,
	}
}

func SolvePartOne(pwd Password, data []byte) {
	// TODO: handle error

	for cmd := range bytes.SplitSeq(data, []byte{'\n'}) {

		fmt.Print(string(pwd.Word), " ")
		cmdStrArr := strings.Split(string(cmd), " ")

		cmdType := strings.Join(cmdStrArr[:2], " ")
		fmt.Print(cmdStrArr, " ")
		switch cmdType {
		case "swap position":
			x, _ := strconv.Atoi(cmdStrArr[2])
			y, _ := strconv.Atoi(cmdStrArr[len(cmdStrArr)-1])
			pwd.SwapPos(x, y)
		case "swap letter":
			x := []byte(cmdStrArr[2])[0]
			y := []byte(cmdStrArr[len(cmdStrArr)-1])[0]
			pwd.SwapLtr(x, y)
		case "reverse positions":
			x, _ := strconv.Atoi(cmdStrArr[2])
			y, _ := strconv.Atoi(cmdStrArr[len(cmdStrArr)-1])
			pwd.ReversePos(x, y)
		case "rotate left":
			x, _ := strconv.Atoi(cmdStrArr[2])
			pwd.RotateLeft(x)
		case "rotate right":
			x, _ := strconv.Atoi(cmdStrArr[2])
			pwd.RotateRight(x)
		case "move position":
			x, _ := strconv.Atoi(cmdStrArr[2])
			y, _ := strconv.Atoi(cmdStrArr[len(cmdStrArr)-1])
			pwd.MovePos(x, y)
		case "rotate based":
			ltr := []byte(cmdStrArr[len(cmdStrArr)-1])[0]
			pwd.RotateByLtr(ltr)
		}

		fmt.Print(string(pwd.Word), "\n")

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

	pwd := ConstructPassword([]byte("abcde"))
	pwd = ConstructPassword([]byte("abcdefgh"))

	SolvePartOne(pwd, data)

}
