package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Room struct {
	EncryptedName string
	SectorId      int
	CheckSum      string
}

type Freq struct {
	Char  rune
	Count int
}

func (r Room) IsValidRoom() bool {
	// create char frequency map
	freq := make(map[rune]int)
	for _, char := range r.EncryptedName {
		if char == 45 {
			continue
		}
		freq[char]++
	}

	// Build frequency array so it can be sorted
	var freqs []Freq
	for char, ct := range freq {
		freqs = append(freqs, Freq{Char: char, Count: ct})
	}

	// Sort by increasing count, if equal, sort by alphabetical order
	slices.SortFunc(freqs, func(a, b Freq) int {
		return cmp.Or(cmp.Compare(b.Count, a.Count), cmp.Compare(a.Char, b.Char))
	})

	if len(freqs) < len(r.CheckSum) {
		return false
	}

	i := 0
	for i < len(r.CheckSum) {
		if r.CheckSum[i] != byte(freqs[i].Char) {
			return false
		}
		i++
	}

	return true

}

func (r Room) DecipherName() string {
	shift := r.SectorId % 26
	var decipheredName []rune
	for _, char := range r.EncryptedName {
		if char == 45 {
			decipheredName = append(decipheredName, char)
			continue
		}
		newChar := char + rune(shift)
		if newChar > 122 {
			newChar = 97 + (newChar % 123)
		}
		decipheredName = append(decipheredName, newChar)
	}

	return string(decipheredName)

}

func createRooms(data []byte) ([]Room, error) {
	var rooms []Room
	for row := range bytes.SplitSeq(data, []byte{10}) {
		rowStrArr := strings.Split(string(row), "-")

		sectorId := strings.Split(rowStrArr[len(rowStrArr)-1], "[")[0]
		sectorIdInt, err := strconv.Atoi(sectorId)
		if err != nil {
			return rooms, err
		}

		checkSum := strings.Split(rowStrArr[len(rowStrArr)-1], "[")[1]
		checkSum = strings.TrimSuffix(checkSum, "]")

		room := Room{
			EncryptedName: strings.Join(rowStrArr[:len(rowStrArr)-1], "-"),
			SectorId:      sectorIdInt,
			CheckSum:      checkSum,
		}

		rooms = append(rooms, room)

	}

	return rooms, nil

}

func solvePartOne(rooms []Room) int {
	var sum int
	for _, rm := range rooms {
		if rm.IsValidRoom() {
			sum += rm.SectorId
		}
	}

	return sum

}

func solvePartTwo(rooms []Room) {
	for _, rm := range rooms {
		if rm.IsValidRoom() {
			decipheredName := rm.DecipherName()
			fmt.Println(decipheredName, rm.SectorId)
		}
	}
}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)
	rooms, err := createRooms(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(rooms)
	fmt.Println(res)

	solvePartTwo(rooms)

}
