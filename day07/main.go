package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Ip struct {
	Supernets [][]byte
	Hypernets [][]byte
	SuperAba  map[string]bool
	HyperAba  map[string]bool
}

func (i *Ip) GetSuperAba() {

	superAbas := make(map[string]bool)

	for _, s := range i.Supernets {
		for i := 2; i < len(s); i++ {

			if (s[i-2] == s[i]) && (s[i-1] != s[i-2]) {
				superAbas[string(s[i-2:i+1])] = true
			}

		}
	}
	i.SuperAba = superAbas

}

func (i *Ip) GetHyperAba() {
	hyperAbas := make(map[string]bool)

	for _, s := range i.Hypernets {
		for i := 2; i < len(s); i++ {

			if (s[i-2] == s[i]) && (s[i-1] != s[i-2]) {
				hyperAbas[string(s[i-2:i+1])] = true
			}

		}
	}
	i.HyperAba = hyperAbas
}

func (i Ip) SupportsSsl() bool {

	for aba := range i.SuperAba {
		bab := string([]byte{aba[1], aba[0], aba[1]})
		if i.HyperAba[bab] {
			return true
		}
	}

	return false

}

func createIps(data []byte) []Ip {

	// Assumes each entry has valid bracket configs
	var ips []Ip
	for entry := range bytes.SplitSeq(data, []byte{10}) {

		var supernets [][]byte
		var hypernets [][]byte
		var curr []byte
		for _, char := range entry {

			if char == 91 {
				supernets = append(supernets, curr)
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
			supernets = append(supernets, curr)
		}

		ips = append(ips, Ip{
			Supernets: supernets,
			Hypernets: hypernets,
			SuperAba:  make(map[string]bool),
			HyperAba:  make(map[string]bool),
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
		if slices.ContainsFunc(ip.Supernets, containsAbba) && !slices.ContainsFunc(ip.Hypernets, containsAbba) {
			count++
		}

	}

	return count
}

func solvePartTwo(ips []Ip) int {
	var count int
	for _, ip := range ips {
		ip.GetSuperAba()
		ip.GetHyperAba()

		if ip.SupportsSsl() {
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

	res2 := solvePartTwo(ips)
	fmt.Println(res2)

}
