package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Ip struct {
	Nets      [][]byte
	Hypernets [][]byte
}

func createIps(data []byte) []Ip {

	// Assumes each entry has valid bracket configs
	var ips []Ip
	for entry := range bytes.SplitSeq(data, []byte{10}) {

		var nets [][]byte
		var hypernets [][]byte
		var curr []byte
		for _, char := range entry {

			if char == 91 {
				nets = append(nets, curr)
				curr = []byte{}
				continue
			}

			if char == 93 {
				hypernets = append(hypernets, curr)
				curr = []byte{}
				continue
			}

			curr = append(curr, char)
		}
		if len(curr) > 0 {
			nets = append(nets, curr)
		}

		ips = append(ips, Ip{
			Nets:      nets,
			Hypernets: hypernets,
		})

	}

	return ips

}

func containsAbba(net []byte) bool {
	if len(net) < 4 {
		return false
	}

	for i := 3; i < len(net); i++ {
		if net[i-3] != net[i-2] && net[i-3] == net[i] && net[i-2] == net[i-1] {
			return true
		}
	}
	return false
}

func solvePartOne(ips []Ip) int {
	var count int
	for _, ip := range ips {
		if slices.ContainsFunc(ip.Nets, containsAbba) && !slices.ContainsFunc(ip.Hypernets, containsAbba) {
			count++
		}

	}

	return count
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	ips := createIps(data)

	res := solvePartOne(ips)
	fmt.Println(res)

}
