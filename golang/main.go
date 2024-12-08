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
	fmt.Println(day7a())
	fmt.Println(day7b())
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

// region Day5
func day5a() int {
	rules, updates := parseFile()
	pages := 0

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			pages += getMiddlePage(update)
		}
	}

	return pages
}

func day5b() int {
	rules, updates := parseFile()
	pages := 0

	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			correctedUpdate := correctUpdate(update, rules)
			pages += getMiddlePage(correctedUpdate)
		}
	}

	return pages
}

func isUpdateValid(update []int, rules [][2]int) bool {
	pagePositions := make(map[int]int)
	for idx, page := range update {
		pagePositions[page] = idx
	}

	for _, rule := range rules {
		before, after := rule[0], rule[1]
		if posBefore, okBefore := pagePositions[before]; okBefore {
			if posAfter, okAfter := pagePositions[after]; okAfter {
				if posBefore >= posAfter {
					return false
				}
			}
		}
	}
	return true
}

func getMiddlePage(update []int) int {
	return update[len(update)/2]
}

func correctUpdate(update []int, rules [][2]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	for _, rule := range rules {
		before, after := rule[0], rule[1]
		if contains(update, before) && contains(update, after) {
			graph[before] = append(graph[before], after)
			inDegree[after]++
		}
	}

	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, item int) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func parseFile() ([][2]int, [][]int) {
	file := strings.Split(strings.TrimSpace(readDayFile(5)), "\n")
	var rules [][2]int
	var updates [][]int
	var dividerIndex int

	for i, line := range file {
		if line == "" {
			dividerIndex = i
			break
		}
	}

	for i := 0; i < dividerIndex; i++ {
		parts := strings.Split(file[i], "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, [2]int{before, after})
	}

	for i := dividerIndex + 1; i < len(file); i++ {
		var update []int
		for _, num := range strings.Split(file[i], ",") {
			val, _ := strconv.Atoi(num)
			update = append(update, val)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

// endregion

// region Day6
func day6a() int {
	file := strings.Split(readDayFile(6), "\n")
	rows := len(file)
	cols := len(file[0])

	directions := map[rune][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	turnRight := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	// Find the guard's starting position and direction
	var guardPos [2]int
	var guardDir rune
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if _, exists := directions[rune(file[r][c])]; exists {
				guardPos = [2]int{r, c}
				guardDir = rune(file[r][c])
				break
			}
		}
	}

	visited := make(map[[2]int]bool)
	visited[guardPos] = true

	for {
		dy, dx := directions[guardDir][0], directions[guardDir][1]
		nextPos := [2]int{guardPos[0] + dy, guardPos[1] + dx}

		// Check if out of bounds
		if nextPos[0] < 0 || nextPos[0] >= rows || nextPos[1] < 0 || nextPos[1] >= cols {
			break
		}

		// Check if there's an obstacle
		if file[nextPos[0]][nextPos[1]] == '#' {
			// Turn right
			guardDir = turnRight[guardDir]
		} else {
			// Move forward
			guardPos = nextPos
			visited[guardPos] = true
		}
	}

	return len(visited)
}

func day6b() int {
	file := strings.Split(readDayFile(6), "\n")
	rows := len(file)
	cols := len(file[0])

	directions := map[rune][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	// Find the guard's starting position and direction
	var guardPos [2]int
	var guardDir rune
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if _, exists := directions[rune(file[r][c])]; exists {
				guardPos = [2]int{r, c}
				guardDir = rune(file[r][c])
				break
			}
		}
	}

	loopPositions := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			obstacle := [2]int{r, c}
			if isGuardInLoop(file, guardPos, byte(guardDir), obstacle) {
				loopPositions++
			}
		}
	}

	return loopPositions
}

func isGuardInLoop(mapLines []string, guardStart [2]int, guardDir byte, obstruction [2]int) bool {
	directions := map[byte][2]int{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}
	turnRight := map[byte]byte{'^': '>', '>': 'v', 'v': '<', '<': '^'}

	rows, cols := len(mapLines), len(mapLines[0])
	guardPos := guardStart
	currentDir := guardDir

	tempMap := make([][]byte, rows)
	for i := range mapLines {
		tempMap[i] = []byte(mapLines[i])
	}
	tempMap[obstruction[0]][obstruction[1]] = '#'

	visitedStates := make(map[[3]int]bool)
	recentHistory := make([][3]int, 0)
	maxHistoryLength := 10

	steps, maxSteps := 0, rows*cols*2

	for {
		state := [3]int{guardPos[0], guardPos[1], int(currentDir)}
		if visitedStates[state] {
			for _, s := range recentHistory {
				if s == state {
					return true
				}
			}
		}

		visitedStates[state] = true
		recentHistory = append(recentHistory, state)
		if len(recentHistory) > maxHistoryLength {
			recentHistory = recentHistory[1:]
		}

		dx, dy := directions[currentDir][0], directions[currentDir][1]
		nextPos := [2]int{guardPos[0] + dx, guardPos[1] + dy}

		if nextPos[0] < 0 || nextPos[0] >= rows || nextPos[1] < 0 || nextPos[1] >= cols {
			return false
		} else if tempMap[nextPos[0]][nextPos[1]] == '#' {
			currentDir = turnRight[currentDir]
		} else {
			guardPos = nextPos
		}

		steps++
		if steps > maxSteps {
			return true
		}
	}
}

// endregion

// region Day6
func day7a() int {
	file := strings.Split(readDayFile(7), "\n")
	sum := 0

	for _, line := range file {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numberParts := strings.Fields(strings.TrimSpace(parts[1]))

		numbers := make([]int, len(numberParts))
		for i, numStr := range numberParts {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		if canAchieveTarget(target, numbers, numbers[0], 1) {
			sum += target
		}
	}

	return sum
}

func day7b() int {
	file := strings.Split(readDayFile(7), "\n")
	sum := 0

	for _, line := range file {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numberParts := strings.Fields(strings.TrimSpace(parts[1]))

		numbers := make([]int, len(numberParts))
		for i, numStr := range numberParts {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		if canAchieveTarget2(target, numbers, numbers[0], 1) {
			sum += target
		}
	}

	return sum
}

func canAchieveTarget(target int, numbers []int, currentValue, index int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	if canAchieveTarget(target, numbers, currentValue+numbers[index], index+1) {
		return true
	}

	if canAchieveTarget(target, numbers, currentValue*numbers[index], index+1) {
		return true
	}

	return false
}

func canAchieveTarget2(target int, numbers []int, currentValue, index int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	if canAchieveTarget2(target, numbers, currentValue+numbers[index], index+1) {
		return true
	}

	if canAchieveTarget2(target, numbers, currentValue*numbers[index], index+1) {
		return true
	}

	concatenated, _ := strconv.Atoi(fmt.Sprintf("%v%v", currentValue, numbers[index]))

	if canAchieveTarget2(target, numbers, concatenated, index+1) {
		return true
	}

	return false
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
