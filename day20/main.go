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

type Interval struct {
	Start int
	End   int
}

func GetIntervals(data []byte) ([]Interval, error) {
	var intervals []Interval

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		entryStrArr := strings.Split(string(entry), "-")
		start, err := strconv.Atoi(entryStrArr[0])
		if err != nil {
			return intervals, err
		}
		end, err := strconv.Atoi(entryStrArr[len(entryStrArr)-1])
		if err != nil {
			return intervals, err
		}

		intervals = append(intervals, Interval{
			Start: start,
			End:   end,
		})
	}

	return intervals, nil

}

func GetSortedIntervals(intervals []Interval) []Interval {
	intervals = slices.Clone(intervals)
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Or(
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.End, b.End),
		)
	})

	return intervals
}

func MergeOverlappingIntervals(intervals []Interval) []Interval {
	var res []Interval
	sorted := GetSortedIntervals(intervals)
	curr := sorted[0]
	for i := 1; i < len(sorted); i++ {
		inter := sorted[i]
		if curr.Start <= inter.Start && inter.Start <= curr.End {
			curr.End = max(curr.End, inter.End)
			continue
		}
		res = append(res, curr)
		curr = inter
	}

	res = append(res, curr)

	return res
}

func MergeContiguousIntervals(intervals []Interval) []Interval {
	var res []Interval
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		inter := intervals[i]
		if curr.End+1 == inter.Start {
			curr.End = inter.End
			continue
		}
		res = append(res, curr)
		curr = inter

	}
	res = append(res, curr)

	return res

}

func solvePartOne(intervals []Interval) int {
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1].End+1 != intervals[i].Start {
			return intervals[i-1].End + 1
		}

	}

	return -1
}

func solvePartTwo(intervals []Interval) int {
	var count int
	for i := 1; i < len(intervals); i++ {
		count += intervals[i].Start - intervals[i-1].End - 1
	}

	return count
}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	intervals, err := GetIntervals(data)
	if err != nil {
		log.Fatal(err)
	}

	mergedIntervals := MergeOverlappingIntervals(intervals)
	res := solvePartOne(mergedIntervals)
	fmt.Println(res)

	res2 := solvePartTwo(mergedIntervals)
	fmt.Println(res2)

}
