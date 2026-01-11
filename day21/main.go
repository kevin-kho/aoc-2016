package main

import (
	"fmt"
	"log"
	"slices"

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

func (p *Password) MovePos(x, y int) {
	char := p.Word[x]
	p.Word = slices.Insert(p.Word, y, char)

	remove := y
	if x < y {
		// Remove last occurance
		for i := 0; i < len(p.Word); i++ {
			if p.Word[i] == char {
				remove = i
				break
			}
		}
	} else {
		// Remove first occurance
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

func main() {
	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	pwd := ConstructPassword(data)

	pwd.MovePos(1, 0)
	fmt.Println(string(pwd.Word))

}
