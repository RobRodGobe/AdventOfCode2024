package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(day4a())
	fmt.Println(day4b())
}

// region Day1
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

// endregion

// region Day2
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

// endregion

// region Day3
func day3a() int {
	mult := 0
	file := readDayFile(3)
	pattern := `mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(file, -1)

	for i := range matches {
		numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(matches[i], "mul("), ")"), ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		mult += x * y
	}

	return mult
}

func day3b() int {
	mult := 0
	file := readDayFile(3)
	pattern := `mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)|do\(\)|don't\(\)`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(file, -1)

	multiply := true

	for i := range matches {
		if matches[i] == "do()" {
			multiply = true
		} else if matches[i] == "don't()" {
			multiply = false
		}

		if multiply && strings.HasPrefix(matches[i], "mul(") {
			numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(matches[i], "mul("), ")"), ",")
			x, _ := strconv.Atoi(numbers[0])
			y, _ := strconv.Atoi(numbers[1])
			mult += x * y
		}
	}

	return mult
}

// endregion

// region Day4
func day4a() int {
	file := strings.Split(readDayFile(4), "\n")
	word := "XMAS"
	rows := len(file)
	cols := len(file[0])
	wordLength := len(word)
	count := 0

	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{1, 1},   // Down-right
		{1, -1},  // Down-left
		{0, -1},  // Left
		{-1, 0},  // Up
		{-1, -1}, // Up-left
		{-1, 1},  // Up-right
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkWordBegin(x, y, dx, dy, wordLength, rows, cols, word, file) {
					count++
				}
			}
		}
	}

	return count
}

func day4b() int {
	file := strings.Split(readDayFile(4), "\n")
	rows := len(file)
	cols := len(file[0])
	count := 0

	for x := 1; x < rows-1; x++ {
		for y := 1; y < cols-1; y++ {
			if isXMasPattern(file, x, y) {
				count++
			}
		}
	}

	return count
}

func checkWordBegin(x, y, dx, dy, length, rows, cols int, word string, grid []string) bool {
	for i := 0; i < length; i++ {
		nx := x + i*dx
		ny := y + i*dy

		if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != word[i] {
			return false
		}
	}
	return true
}

func isXMasPattern(grid []string, x, y int) bool {
	topLeftToBottomRight := string(grid[x-1][y-1]) + string(grid[x][y]) + string(grid[x+1][y+1])
	topRightToBottomLeft := string(grid[x-1][y+1]) + string(grid[x][y]) + string(grid[x+1][y-1])

	return isValidMasPattern(topLeftToBottomRight) && isValidMasPattern(topRightToBottomLeft)
}

func isValidMasPattern(pattern string) bool {
	return pattern == "MAS" || pattern == "SAM"
}

// endregion

func readDayFile(day int32) string {
	filePath := fmt.Sprintf("../AoC_Files/%d.txt", day)

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(content)
	return fileContents
}
