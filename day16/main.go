package main

import (
	"fmt"
	"slices"
	"strings"
)

func Flip(strArr []string) string {
	var sb strings.Builder

	for _, char := range strArr {
		if char == "0" {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	return sb.String()

}

func GetHash(input string) string {
	var sb strings.Builder

	for i := 1; i < len(input); i += 2 {
		pair := input[i-1 : i+1]
		if pair == "00" || pair == "11" {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	return sb.String()
}

func SolvePartOne(input string, length int) string {

	// Generate the initial checksum
	for len(input) < length {
		b := strings.Split(input, "")
		slices.Reverse(b)
		input = input + "0" + Flip(b)

		if len(input) > length {
			input = input[:length+1]
		}

	}

	// Second checksum
	hash := ""
	for len(hash)%2 == 0 {
		hash = GetHash(input)
		input = hash
	}

	return hash

}

func main() {

	input := "11110010111001001"
	length := 272

	res := SolvePartOne(input, length)
	fmt.Println(res)

}
