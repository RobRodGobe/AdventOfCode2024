package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(day1a())
	fmt.Println(day1b())
}

func day1a() int {
	file := readDayFile(1)

	var list1, list2 []int

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		pairs := strings.Fields(line)
		x, err := strconv.Atoi(pairs[0])
		if err != nil {
			panic("Error parsing x")
		}
		y, err := strconv.Atoi(pairs[1])
		if err != nil {
			panic("Error parsing y")
		}
		list1 = append(list1, x)
		list2 = append(list2, y)
	}

	diff := 0

	slices.SortFunc(list1, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	slices.SortFunc(list2, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	for i, _ := range list1 {
		diff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return diff
}

func day1b() int {
	file := readDayFile(1)

	var list1, list2 []int

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		pairs := strings.Fields(line)
		x, err := strconv.Atoi(pairs[0])
		if err != nil {
			panic("Error parsing x")
		}
		y, err := strconv.Atoi(pairs[1])
		if err != nil {
			panic("Error parsing y")
		}
		list1 = append(list1, x)
		list2 = append(list2, y)
	}

	similar := 0

	for i, left := range list1 {
		count := 0
		for _, right := range list2 {
			if right == left {
				count++
			}
		}
		similar += list1[i] * count
	}

	return similar
}

func readDayFile(day int32) string {
	filePath := fmt.Sprintf("../AoC_Files/%d.txt", day)

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(content)
	return fileContents
}
