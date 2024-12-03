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
	fmt.Println(day2a())
	fmt.Println(day2b())
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

	for i := range list1 {
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

func day2a() int {
	file := strings.Split(readDayFile(2), "\n")
	safe := 0

	for _, line := range file {
		reports := parseLine(line)
		isAscending := true
		isDescending := true
		isSafe := true

		for i := 1; i < len(reports); i++ {
			diff := reports[i] - reports[i-1]

			if math.Abs(float64(diff)) > 3 {
				isAscending = false
				isDescending = false
				isSafe = false
				break
			}

			if diff < 0 {
				isAscending = false
			}
			if diff > 0 {
				isDescending = false
			}
			if diff == 0 {
				isAscending = false
				isDescending = false
			}

			if !isAscending && !isDescending {
				isSafe = false
				break
			}
		}

		if isSafe {
			safe++
		}
	}

	return safe
}

func day2b() int {
	file := strings.Split(readDayFile(2), "\n")
	safe := 0

	for _, line := range file {
		reports := parseLine(line)

		if isSafeReport(reports, true) || isSafeReport(reports, false) {
			safe++
			continue
		}

		isSafe := false
		for j := 0; j < len(reports); j++ {
			modifiedReports := append([]int{}, reports[:j]...)          // Copy the first part
			modifiedReports = append(modifiedReports, reports[j+1:]...) // Append the second part

			if isSafeReport(modifiedReports, true) || isSafeReport(modifiedReports, false) {
				isSafe = true
				break
			}
		}

		if isSafe {
			safe++
		}
	}

	return safe
}

func parseLine(line string) []int {
	parts := strings.Fields(line)
	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func isSafeReport(reports []int, ascending bool) bool {
	for i := 1; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]
		if ascending && diff < 0 {
			return false
		}
		if !ascending && diff > 0 {
			return false
		}
		if math.Abs(float64(diff)) > 3 || diff == 0 {
			return false
		}
	}
	return true
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
