package main

import (
	"cmp"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(day24a())
	fmt.Println(day24b())

	fmt.Println(day25())
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

// region Day7
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

// region Day8
type coordinate struct {
	x, y int
}

func day8a() int {
	file := strings.Split(readDayFile(8), "\n")

	var matrix [][]rune
	for _, line := range file {
		matrix = append(matrix, []rune(line))
	}

	antennaMap := getAntennaMap(matrix)

	var allAntinodes [][2]int

	for _, coords := range antennaMap {
		antinodes := getAntinodes(coords, matrix)
		allAntinodes = append(allAntinodes, antinodes...)
	}

	uniqueAntinodes := getUniqueAntinodes(allAntinodes)

	return len(uniqueAntinodes)
}

func day8b() int {
	file := strings.Split(readDayFile(8), "\n")

	var matrix [][]rune
	for _, line := range file {
		matrix = append(matrix, []rune(line))
	}

	antennaMap := getAntennaMap(matrix)

	antinodeMatrix := make([][]bool, len(file))
	for i := range file {
		antinodeMatrix[i] = make([]bool, len(file[0]))
	}

	for _, coords := range antennaMap {
		convertedCoords := make([]coordinate, len(coords))
		for i, coord := range coords {
			convertedCoords[i] = coordinate{x: coord[0], y: coord[1]}
		}
		processAntinodeLines(convertedCoords, matrix, antinodeMatrix)
	}

	return getUniqueAntinodesCount(antinodeMatrix)
}

func withinBoundaries(value, min, max int) bool {
	return value >= min && value < max
}

func getAntennaMap(matrix [][]rune) map[rune][][2]int {
	antennaMap := make(map[rune][][2]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			cell := matrix[i][j]
			if cell == '.' {
				continue
			}

			antennaMap[cell] = append(antennaMap[cell], [2]int{i, j})
		}
	}

	return antennaMap
}

func getAntinodes(coords [][2]int, matrix [][]rune) [][2]int {
	var antinodes [][2]int

	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords); j++ {
			if i == j {
				continue
			}

			ax, ay := coords[i][0], coords[i][1]
			bx, by := coords[j][0], coords[j][1]

			cx := 2*bx - ax
			cy := 2*by - ay

			if withinBoundaries(cx, 0, len(matrix)) && withinBoundaries(cy, 0, len(matrix[0])) {
				antinodes = append(antinodes, [2]int{cx, cy})
			}
		}
	}

	return antinodes
}

func getUniqueAntinodes(antinodes [][2]int) [][2]int {
	uniqueSet := make(map[[2]int]bool)

	for _, antinode := range antinodes {
		uniqueSet[antinode] = true
	}

	var uniqueAntinodes [][2]int
	for antinode := range uniqueSet {
		uniqueAntinodes = append(uniqueAntinodes, antinode)
	}

	return uniqueAntinodes
}

func processAntinodeLines(coords []coordinate, matrix [][]rune, antinodeMatrix [][]bool) {
	for i, c1 := range coords {
		for j, c2 := range coords {
			if i != j {
				x1, y1 := c1.x, c1.y
				x2, y2 := c2.x, c2.y

				for x := 0; x < len(matrix); x++ {
					for y := 0; y < len(matrix[0]); y++ {
						if !antinodeMatrix[x][y] {
							lineResult := (y1-y2)*x + (x2-x1)*y + (x1*y2 - x2*y1)
							if lineResult == 0 {
								antinodeMatrix[x][y] = true
							}
						}
					}
				}
			}
		}
	}
}

func getUniqueAntinodesCount(antinodeMatrix [][]bool) int {
	count := 0
	for _, row := range antinodeMatrix {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

// endregion

// region Day9
type DiskFile struct {
	ID       int
	Length   int
	StartIdx int
}

func day9a() int {
	file := readDayFile(9)
	diskMap := parseDiskMap(file)
	diskMap = compactDisk(diskMap)
	return calculateChecksum(diskMap)
}

func day9b() int {
	file := readDayFile(9)
	diskMap := parseDiskMap(file)
	diskMap = compactDiskByFile(diskMap)
	return calculateChecksum(diskMap)
}

func parseDiskMap(line string) []string {
	nums := []string{}
	index := 0

	for i := 0; i < len(line); i++ {
		count, _ := strconv.Atoi(string(line[i]))
		if i%2 == 0 {
			for j := 0; j < count; j++ {
				nums = append(nums, strconv.Itoa(index))
			}
			index++
		} else {
			for j := 0; j < count; j++ {
				nums = append(nums, ".")
			}
		}
	}

	return nums
}

func compactDisk(diskMap []string) []string {
	L, R := 0, len(diskMap)-1

	for L <= R {
		if diskMap[L] == "." && diskMap[R] != "." {
			diskMap[L], diskMap[R] = diskMap[R], diskMap[L]
			R--
			L++
		} else if diskMap[R] == "." {
			R--
		} else {
			L++
		}
	}

	return diskMap
}

func calculateChecksum(diskMap []string) int {
	checksum := 0
	for i, block := range diskMap {
		if block != "." {
			blockInt, _ := strconv.Atoi(block)
			checksum += i * blockInt
		}
	}
	return checksum
}

func analyzeDisk(diskMap []string) ([]DiskFile, map[int]int) {
	files := []DiskFile{}
	spaces := make(map[int]int)
	spaceStartIdx := -1

	for i, item := range diskMap {
		if item == "." {
			if spaceStartIdx == -1 {
				spaceStartIdx = i
			}
			spaces[spaceStartIdx]++
		} else {
			if spaceStartIdx != -1 {
				spaceStartIdx = -1
			}

			fileID, _ := strconv.Atoi(item)
			for len(files) <= fileID {
				files = append(files, DiskFile{ID: len(files), Length: 0, StartIdx: i})
			}

			files[fileID].Length++
		}
	}

	return files, spaces
}

func getFirstAvailableSpaceIdx(spaces map[int]int, fileLength int) int {
	spaceIndices := make([]int, 0, len(spaces))
	for idx := range spaces {
		spaceIndices = append(spaceIndices, idx)
	}

	sortInts(spaceIndices)

	for _, idx := range spaceIndices {
		if spaces[idx] >= fileLength {
			return idx
		}
	}
	return -1
}

func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func updateSpaces(spaces map[int]int, spaceIdx int, fileLength int) {
	if spaces[spaceIdx] == fileLength {
		delete(spaces, spaceIdx)
	} else {
		remainingSpace := spaces[spaceIdx] - fileLength
		delete(spaces, spaceIdx)
		spaces[spaceIdx+fileLength] = remainingSpace
	}
}

func moveFile(diskMap []string, file DiskFile, targetIdx int) {
	fileSegments := make([]string, file.Length)
	copy(fileSegments, diskMap[file.StartIdx:file.StartIdx+file.Length])

	for i := 0; i < file.Length; i++ {
		diskMap[file.StartIdx+i] = "."
	}

	for i, segment := range fileSegments {
		diskMap[targetIdx+i] = segment
	}
}

func compactDiskByFile(diskMap []string) []string {
	files, spaces := analyzeDisk(diskMap)

	// Sort files in descending order of ID
	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].ID < files[j].ID {
				files[i], files[j] = files[j], files[i]
			}
		}
	}

	for _, file := range files {
		targetIdx := getFirstAvailableSpaceIdx(spaces, file.Length)
		if targetIdx != -1 && targetIdx < file.StartIdx {
			moveFile(diskMap, file, targetIdx)
			updateSpaces(spaces, targetIdx, file.Length)
		}
	}

	return diskMap
}

// endregion

// region Day10
func day10a() int {
	file := strings.Split(readDayFile(10), "\n")

	return solveTopographicMap(file)
}

func day10b() int {
	file := strings.Split(readDayFile(10), "\n")

	return solveTopographicMapTrailRatings(file)
}

func solveTopographicMap(input []string) int {
	rows := len(input)
	cols := len(input[0])
	mapGrid := make([][]int, rows)

	for r := range input {
		mapGrid[r] = make([]int, cols)
		for c := range input[r] {
			mapGrid[r][c] = int(input[r][c] - '0')
		}
	}

	totalTrailheadScore := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mapGrid[r][c] == 0 {
				trailheadScore := findTrailheadScore(r, c, mapGrid, rows, cols)
				totalTrailheadScore += trailheadScore
			}
		}
	}

	return totalTrailheadScore
}

func findTrailheadScore(startRow, startCol int, mapGrid [][]int, rows, cols int) int {
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}

	ninePositions := make(map[string]bool)

	var dfs func(int, int, int) bool
	dfs = func(row, col, expectedHeight int) bool {
		if row < 0 || row >= rows || col < 0 || col >= cols ||
			visited[row][col] || mapGrid[row][col] != expectedHeight {
			return false
		}

		visited[row][col] = true

		if expectedHeight == 9 {
			ninePositions[fmt.Sprintf("%d:%d", row, col)] = true
		}

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			if dfs(row+dir[0], col+dir[1], expectedHeight+1) {
				return true
			}
		}

		return false
	}

	dfs(startRow, startCol, 0)
	return len(ninePositions)
}

func solveTopographicMapTrailRatings(input []string) int {
	mapGrid, rows, cols := parseTopographicMap(input)
	scoresMap := make(map[string]map[string]int)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mapGrid[r][c] == 0 {
				startHike(r, c, mapGrid, rows, cols, scoresMap)
			}
		}
	}

	return getScoresSum(scoresMap)
}

func startHike(startR, startC int, mapGrid [][]int, rows, cols int, scoresMap map[string]map[string]int) {
	type Route struct {
		r, c, height int
		initialCell  string
	}

	routes := []Route{{
		r:           startR,
		c:           startC,
		height:      0,
		initialCell: serializeCoordinates(startR, startC),
	}}

	for len(routes) > 0 {
		route := routes[0]
		routes = routes[1:]

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			newR := route.r + dir[0]
			newC := route.c + dir[1]
			newHeight := route.height + 1

			if newR < 0 || newR >= rows || newC < 0 || newC >= cols {
				continue
			}

			newCell := mapGrid[newR][newC]

			if newCell != newHeight {
				continue
			}

			if newCell == 9 {
				if _, exists := scoresMap[route.initialCell]; !exists {
					scoresMap[route.initialCell] = make(map[string]int)
				}

				endKey := serializeCoordinates(newR, newC)
				scoresMap[route.initialCell][endKey]++
			} else {
				routes = append(routes, Route{
					r:           newR,
					c:           newC,
					height:      newHeight,
					initialCell: route.initialCell,
				})
			}
		}
	}
}

func serializeCoordinates(r, c int) string {
	return fmt.Sprintf("%d:%d", r, c)
}

func getScoresSum(scoresMap map[string]map[string]int) int {
	sum := 0
	for _, scores := range scoresMap {
		for _, score := range scores {
			sum += score
		}
	}
	return sum
}

func parseTopographicMap(input []string) ([][]int, int, int) {
	rows := len(input)
	cols := len(input[0])
	mapGrid := make([][]int, rows)

	for r := range input {
		mapGrid[r] = make([]int, cols)
		for c := range input[r] {
			mapGrid[r][c] = int(input[r][c] - '0')
		}
	}

	return mapGrid, rows, cols
}

// endregion

// region Day11
func day11a() int64 {
	file := readDayFile(11)
	stones := strings.Split(file, " ")
	rocks := make(map[int64]int64)

	for _, stoneStr := range stones {
		stone, _ := strconv.ParseInt(stoneStr, 10, 64)
		rocks[stone]++
	}

	finalRocks := blinkRocks(rocks, 25)
	var total int64
	for _, count := range finalRocks {
		total += count
	}

	return total
}

func day11b() int64 {
	file := readDayFile(11)
	stones := strings.Split(file, " ")
	rocks := make(map[int64]int64)

	for _, stoneStr := range stones {
		stone, _ := strconv.ParseInt(stoneStr, 10, 64)
		rocks[stone]++
	}

	finalRocks := blinkRocks(rocks, 75)
	var total int64
	for _, count := range finalRocks {
		total += count
	}

	return total
}

func blink(rock int64) []int64 {
	if rock == 0 {
		return []int64{1}
	}

	digits := int64(math.Floor(math.Log10(float64(rock)))) + 1

	if digits%2 != 0 {
		return []int64{rock * 2024}
	}

	halfDigits := digits / 2
	first := rock / int64(math.Pow(10, float64(halfDigits)))
	second := rock % int64(math.Pow(10, float64(halfDigits)))

	return []int64{first, second}
}

func blinkRocksIteration(rocks map[int64]int64) map[int64]int64 {
	result := make(map[int64]int64)

	for rock, count := range rocks {
		newRocks := blink(rock)

		for _, newRock := range newRocks {
			result[newRock] += count
		}
	}

	return result
}

func blinkRocks(rocks map[int64]int64, blinks int) map[int64]int64 {
	currentRocks := make(map[int64]int64)
	for k, v := range rocks {
		currentRocks[k] = v
	}

	for i := 0; i < blinks; i++ {
		currentRocks = blinkRocksIteration(currentRocks)
	}

	return currentRocks
}

// endregion

// region Day12
func day12a() int {
	file := strings.Split(readDayFile(12), "\n")

	return calculateTotalFencingPrice(file)
}

func day12b() int {
	file := strings.Split(readDayFile(12), "\n")

	return calculateTotalFencingPriceWithInnerSides(file)
}

func calculateTotalFencingPrice(grid []string) int {
	n, m := len(grid), len(grid[0])
	visited := make(map[string]bool)
	totalPrice := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			if !visited[key] {
				area, borders := visitRegion(grid, i, j, visited)
				totalPrice += area * len(borders)
			}
		}
	}

	return totalPrice
}

func calculateTotalFencingPriceWithInnerSides(grid []string) int {
	n, m := len(grid), len(grid[0])
	visited := make(map[string]bool)
	totalPrice := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			if !visited[key] {
				area, borders := visitRegion(grid, i, j, visited)
				totalPrice += area * countSides(borders)
			}
		}
	}

	return totalPrice
}

func visitRegion(grid []string, startI, startJ int, visited map[string]bool) (int, map[string]bool) {
	n, m := len(grid), len(grid[0])
	plant := grid[startI][startJ]
	area := 0
	borders := make(map[string]bool)

	var visit func(int, int)
	visit = func(i, j int) {
		key := fmt.Sprintf("%d,%d", i, j)
		if visited[key] {
			return
		}

		visited[key] = true
		area++

		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}

		for k := 0; k < 4; k++ {
			i2 := i + dx[k]
			j2 := j + dy[k]

			if i2 >= 0 && i2 < n && j2 >= 0 && j2 < m && grid[i2][j2] == plant {
				visit(i2, j2)
			} else {
				borderKey := fmt.Sprintf("%d,%d,%d,%d", i, j, i2, j2)
				borders[borderKey] = true
			}
		}
	}

	visit(startI, startJ)
	return area, borders
}

func countSides(borders map[string]bool) int {
	visited := make(map[string]bool)

	var visitSide func(int, int, int, int)
	visitSide = func(i, j, i2, j2 int) {
		side := fmt.Sprintf("%d,%d,%d,%d", i, j, i2, j2)
		if visited[side] || !borders[side] {
			return
		}

		visited[side] = true

		if i == i2 {
			visitSide(i-1, j, i2-1, j2)
			visitSide(i+1, j, i2+1, j2)
		} else {
			visitSide(i, j-1, i2, j2-1)
			visitSide(i, j+1, i2, j2+1)
		}
	}

	numSides := 0
	for side := range borders {
		if visited[side] {
			continue
		}

		numSides++
		parts := strings.Split(side, ",")
		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])
		i2, _ := strconv.Atoi(parts[2])
		j2, _ := strconv.Atoi(parts[3])
		visitSide(i, j, i2, j2)
	}

	return numSides
}

// endregion

// region Day13
func day13a() int64 {
	file := strings.Split(readDayFile(13), "\n")

	return GetMaxPrizeForMinTokens(file)
}

func day13b() int64 {
	file := strings.Split(readDayFile(13), "\n")

	machines := ParseClawMachineInput(file)

	for i := range machines {
		machines[i].PrizeX += 10_000_000_000_000
		machines[i].PrizeY += 10_000_000_000_000
	}

	var adjustedInput []string
	for _, machine := range machines {
		adjustedInput = append(adjustedInput,
			fmt.Sprintf("Button A: X+%d, Y+%d", machine.AX, machine.AY),
			fmt.Sprintf("Button B: X+%d, Y+%d", machine.BX, machine.BY),
			fmt.Sprintf("Prize: X=%d, Y=%d", machine.PrizeX, machine.PrizeY),
		)
	}

	return GetMaxPrizeForMinTokens(adjustedInput)
}

type ClawMachineSettings struct {
	AX, AY, BX, BY int
	PrizeX, PrizeY int64
}

func ParseClawMachineInput(input []string) []ClawMachineSettings {
	var machines []ClawMachineSettings
	var cleanedData []string

	for _, line := range input {
		if strings.TrimSpace(line) != "" {
			cleanedData = append(cleanedData, line)
		}
	}

	for i := 0; i < len(cleanedData); i += 3 {
		aMove := strings.Split(strings.ReplaceAll(cleanedData[i], "Button A: ", ""), ", ")
		bMove := strings.Split(strings.ReplaceAll(cleanedData[i+1], "Button B: ", ""), ", ")
		prize := strings.Split(strings.ReplaceAll(cleanedData[i+2], "Prize: ", ""), ", ")

		ax, _ := strconv.Atoi(strings.ReplaceAll(aMove[0], "X+", ""))
		ay, _ := strconv.Atoi(strings.ReplaceAll(aMove[1], "Y+", ""))
		bx, _ := strconv.Atoi(strings.ReplaceAll(bMove[0], "X+", ""))
		by, _ := strconv.Atoi(strings.ReplaceAll(bMove[1], "Y+", ""))
		prizeX, _ := strconv.ParseInt(strings.ReplaceAll(prize[0], "X=", ""), 10, 64)
		prizeY, _ := strconv.ParseInt(strings.ReplaceAll(prize[1], "Y=", ""), 10, 64)

		machines = append(machines, ClawMachineSettings{AX: ax, AY: ay, BX: bx, BY: by, PrizeX: prizeX, PrizeY: prizeY})
	}

	return machines
}

func CalculatePrice(machine ClawMachineSettings) *int64 {
	det := int64(machine.AY*machine.BX - machine.AX*machine.BY)
	if det == 0 {
		return nil
	}

	b := (int64(machine.AY)*machine.PrizeX - int64(machine.AX)*machine.PrizeY) / det
	a := int64(0)
	if machine.AX != 0 {
		a = (machine.PrizeX - b*int64(machine.BX)) / int64(machine.AX)
	}

	if int64(machine.AX)*a+int64(machine.BX)*b == machine.PrizeX &&
		int64(machine.AY)*a+int64(machine.BY)*b == machine.PrizeY &&
		a >= 0 && b >= 0 {
		result := a*3 + b
		return &result
	}

	return nil
}

func GetMaxPrizeForMinTokens(input []string) int64 {
	machines := ParseClawMachineInput(input)
	var totalTokens int64

	for _, machine := range machines {
		tokens := CalculatePrice(machine)
		if tokens != nil {
			totalTokens += *tokens
		}
	}

	return totalTokens
}

// endregion

// region Day14
func day14a() int {
	file := strings.Split(readDayFile(14), "\n")

	return calculateSafetyFactor(file)
}

func day14b() int {
	file := strings.Split(readDayFile(14), "\n")

	return findRobotSequenceTime(file)
}

type BathroomRobot struct {
	P struct{ X, Y int }
	V struct{ X, Y int }
}

func simulateRobot(robot BathroomRobot, modRows, modCols, ticks int) BathroomRobot {
	rowDelta := calculateDelta(robot.V.Y, ticks, modRows)
	newRow := modAdd(robot.P.Y, rowDelta, modRows)

	colDelta := calculateDelta(robot.V.X, ticks, modCols)
	newCol := modAdd(robot.P.X, colDelta, modCols)

	return BathroomRobot{
		P: struct{ X, Y int }{X: newCol, Y: newRow},
		V: robot.V,
	}
}

func calculateDelta(velocity, ticks, mod int) int {
	delta := velocity * ticks % mod
	if delta < 0 {
		delta += mod
	}
	return delta
}

func modAdd(a, b, mod int) int {
	res := (a + b) % mod
	if res < 0 {
		res += mod
	}
	return res
}

func calculateSafetyFactor(file []string) int {
	width, height := 101, 103
	duration := 100

	robots := parseRobots(file)
	finalPositions := calculateFinalPositions(robots, width, height, duration)

	return computeQuadrantMultiplier(finalPositions, width, height)
}

func parseRobots(lines []string) []BathroomRobot {
	var robots []BathroomRobot
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			robots = append(robots, parseSingleRobot(line))
		}
	}
	return robots
}

func parseSingleRobot(line string) BathroomRobot {
	parts := strings.Split(line, " ")
	p := strings.Split(parts[0][2:], ",")
	v := strings.Split(parts[1][2:], ",")

	px, _ := strconv.Atoi(p[0])
	py, _ := strconv.Atoi(p[1])
	vx, _ := strconv.Atoi(v[0])
	vy, _ := strconv.Atoi(v[1])

	return BathroomRobot{
		P: struct{ X, Y int }{X: px, Y: py},
		V: struct{ X, Y int }{X: vx, Y: vy},
	}
}

func calculateFinalPositions(robots []BathroomRobot, width, height, duration int) [][]int {
	finalPositions := make([][]int, width)
	for i := range finalPositions {
		finalPositions[i] = make([]int, height)
	}

	for _, robot := range robots {
		finalX := (robot.P.X + robot.V.X*duration) % width
		finalY := (robot.P.Y + robot.V.Y*duration) % height

		if finalX < 0 {
			finalX += width
		}
		if finalY < 0 {
			finalY += height
		}

		finalPositions[finalX][finalY]++
	}

	return finalPositions
}

func computeQuadrantMultiplier(finalPositions [][]int, width, height int) int {
	midX, midY := width/2, height/2

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x == midX || y == midY {
				continue
			}

			switch {
			case x < midX && y < midY:
				topLeft += finalPositions[x][y]
			case x >= midX && y < midY:
				topRight += finalPositions[x][y]
			case x < midX && y >= midY:
				bottomLeft += finalPositions[x][y]
			case x >= midX && y >= midY:
				bottomRight += finalPositions[x][y]
			}
		}
	}

	return topLeft * topRight * bottomLeft * bottomRight
}

func findRobotSequenceTime(file []string) int {
	rows, cols := 103, 101
	trunkSeqSize := 10
	maxSeconds := 100000

	robots := parseRobots(file)

	for sec := 1; sec <= maxSeconds; sec++ {
		robotsByCol := make([][]int, cols)
		for i := range robotsByCol {
			robotsByCol[i] = []int{}
		}

		for i := range robots {
			newRobot := simulateRobot(robots[i], rows, cols, 1)
			robotsByCol[newRobot.P.X] = append(robotsByCol[newRobot.P.X], newRobot.P.Y)
			robots[i] = newRobot
		}

		for c := 0; c < cols; c++ {
			column := robotsByCol[c]
			sort.Ints(column)

			if hasConsecutiveSequence(column, trunkSeqSize) {
				return sec
			}
		}
	}

	return 0
}

func hasConsecutiveSequence(sequence []int, requiredLength int) bool {
	if len(sequence) < requiredLength {
		return false
	}

	consecutiveCount := 1
	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1]+1 {
			consecutiveCount++
		} else {
			consecutiveCount = 1
		}

		if consecutiveCount == requiredLength {
			return true
		}
	}

	return false
}

// endregion

// region Day15
func day15a() int {
	file := readDayFile(15)
	sections := strings.Split(file, "\n\n")
	inputMap := strings.Split(sections[0], "\n")
	moves := strings.ReplaceAll(sections[1], " ", "")

	mapTiles := make(map[Vector]rune)
	robotPos := Zero

	for row, line := range inputMap {
		for col, tile := range line {
			pos := Vector{row, col}
			if tile == '#' || tile == 'O' {
				mapTiles[pos] = tile
			} else if tile == '@' {
				robotPos = pos
			}
		}
	}

	directions := map[rune]Vector{
		'>': Right,
		'v': Down,
		'<': Left,
		'^': Up,
	}

	for _, move := range moves {
		dir := directions[rune(move)]
		var thingsToPush []rune
		next := add(robotPos, dir)

		for tile, ok := mapTiles[next]; ok; tile, ok = mapTiles[next] {
			thingsToPush = append(thingsToPush, tile)
			if tile == '#' {
				break
			}
			next = add(next, dir)
		}

		if len(thingsToPush) == 0 {
			robotPos = add(robotPos, dir)
		} else if thingsToPush[len(thingsToPush)-1] == 'O' {
			for i := range thingsToPush {
				delete(mapTiles, add(robotPos, scale(dir, 1+i)))
			}
			for i, tile := range thingsToPush {
				mapTiles[add(robotPos, scale(dir, 2+i))] = tile
			}
			robotPos = add(robotPos, dir)
		}
	}

	total := 0
	for pos, tile := range mapTiles {
		if tile == 'O' {
			total += 100*pos.Row + pos.Col
		}
	}
	return total
}

func day15b() int {
	file := readDayFile(15)
	sections := strings.Split(file, "\n\n")
	inputMap := strings.Split(sections[0], "\n")
	moves := strings.ReplaceAll(sections[1], " ", "")

	mapTiles := make(map[Vector]*Obstacle)
	robotPos := Zero

	for row, line := range inputMap {
		for col, tile := range line {
			pos := Vector{row, col * 2}
			if tile == '#' || tile == 'O' {
				right := add(pos, Right)
				obstacle := &Obstacle{Tile: tile, Left: pos, Right: right}
				mapTiles[pos] = obstacle
				mapTiles[right] = obstacle
			} else if tile == '@' {
				robotPos = pos
			}
		}
	}

	directions := map[rune]Vector{
		'>': Right,
		'v': Down,
		'<': Left,
		'^': Up,
	}

	for _, move := range moves {
		dir := directions[rune(move)]
		thingsToPush := getBoxesToPush(mapTiles, add(robotPos, dir), dir)

		if len(thingsToPush) == 0 {
			robotPos = add(robotPos, dir)
		} else {
			hasWall := false
			for _, obstacle := range thingsToPush {
				if obstacle.Tile == '#' {
					hasWall = true
					break
				}
			}
			if hasWall {
				continue
			}

			for _, obstacle := range thingsToPush {
				delete(mapTiles, obstacle.Left)
				delete(mapTiles, obstacle.Right)
			}
			for _, obstacle := range thingsToPush {
				newObstacle := &Obstacle{
					Tile:  obstacle.Tile,
					Left:  add(obstacle.Left, dir),
					Right: add(obstacle.Right, dir),
				}
				mapTiles[newObstacle.Left] = newObstacle
				mapTiles[newObstacle.Right] = newObstacle
			}
			robotPos = add(robotPos, dir)
		}
	}

	coordinates := make(map[Vector]bool)
	for _, obstacle := range mapTiles {
		if obstacle.Tile == 'O' {
			coordinates[obstacle.Left] = true
		}
	}

	total := 0
	for coord := range coordinates {
		total += 100*coord.Row + coord.Col
	}
	return total
}

type Obstacle struct {
	Tile  rune
	Left  Vector
	Right Vector
}

func getBoxesToPush(mapTiles map[Vector]*Obstacle, pos Vector, dir Vector) map[Vector]*Obstacle {
	results := make(map[Vector]*Obstacle)
	if obstacle, exists := mapTiles[pos]; exists {
		results[pos] = obstacle
		if obstacle.Tile == 'O' {
			if dir == Left {
				for k, v := range getBoxesToPush(mapTiles, add(obstacle.Left, Left), dir) {
					results[k] = v
				}
			} else if dir == Right {
				for k, v := range getBoxesToPush(mapTiles, add(obstacle.Right, Right), dir) {
					results[k] = v
				}
			} else {
				for k, v := range getBoxesToPush(mapTiles, add(obstacle.Left, dir), dir) {
					results[k] = v
				}
				for k, v := range getBoxesToPush(mapTiles, add(obstacle.Right, dir), dir) {
					results[k] = v
				}
			}
		}
	}
	return results
}

type Vector struct {
	Row, Col int
}

var (
	Zero  = Vector{0, 0}
	Up    = Vector{-1, 0}
	Down  = Vector{1, 0}
	Left  = Vector{0, -1}
	Right = Vector{0, 1}
)

func add(a, b Vector) Vector {
	return Vector{a.Row + b.Row, a.Col + b.Col}
}

func scale(v Vector, factor int) Vector {
	return Vector{v.Row * factor, v.Col * factor}
}

// endregion

// region Day16
func day16a() int {
	file := strings.Split(readDayFile(16), "\n")
	start := state{pos: position{row: len(file) - 2, col: 1}, dir: east}
	if file[start.pos.row][start.pos.col] != 'S' {
		start = state{pos: position{row: 1, col: len(file[0]) - 2}, dir: south}
	}
	s := solve(file, start)
	return s.cheapest
}

func day16b() int {
	file := strings.Split(readDayFile(16), "\n")
	start := state{pos: position{row: len(file) - 2, col: 1}, dir: east}
	if file[start.pos.row][start.pos.col] != 'S' {
		start = state{pos: position{row: 1, col: len(file[0]) - 2}, dir: south}
	}
	s := solve(file, start)

	seen := make(map[position]bool)
	q := []state{s.end}
	var zero state
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v != zero {
			seen[v.pos] = true
			q = append(q, s.prov[v].parents...)
		}
	}
	return len(seen)
}

const (
	dot  = '.'
	end  = 'E'
	wall = '#'
)

type direction struct{ row, col int }
type position struct{ row, col int }

var (
	east       = direction{col: 1}
	south      = direction{row: 1}
	west       = direction{col: -1}
	north      = direction{row: -1}
	directions = []direction{east, south, west, north}
)

func (d direction) turnright() direction {
	switch d {
	case east:
		return south
	case south:
		return west
	case west:
		return north
	case north:
		return east
	default:
		log.Fatalf("unknown direction %v", d)
		return d
	}
}

func (d direction) turnleft() direction {
	switch d {
	case east:
		return north
	case north:
		return west
	case west:
		return south
	case south:
		return east
	default:
		log.Fatalf("unknown direction %v", d)
		return d
	}
}

func (p position) move(dir direction) position {
	return position{row: p.row + dir.row, col: p.col + dir.col}
}

type state struct {
	pos position
	dir direction
}

func (s state) possible() (straight, left, right state) {
	straight = state{pos: s.pos.move(s.dir), dir: s.dir}
	left = state{pos: s.pos, dir: s.dir.turnleft()}
	right = state{pos: s.pos, dir: s.dir.turnright()}
	return
}

type provenance struct {
	cost    int
	parents []state
}

func (p *provenance) maybeAdd(parent state, cost int) {
	if p.cost > cost {
		p.cost = cost
		p.parents = []state{parent}
	} else if p.cost == cost {
		p.parents = append(p.parents, parent)
	}
}

type solver struct {
	grid     []string
	pq       map[int][]state
	cheapest int
	highest  int
	end      state
	visited  map[state]int
	prov     map[state]*provenance
}

func (s *solver) add(v, prev state, cost int) {
	if cost < s.cheapest {
		log.Fatalf("Trying to add %v at cost %d but cheapest is %d", v, cost, s.cheapest)
	}
	p := s.prov[v]
	if p == nil {
		p = &provenance{cost: cost}
		s.prov[v] = p
	}
	p.maybeAdd(prev, cost)
	if c, ok := s.visited[v]; !ok || cost < c {
		s.visited[v] = cost
		s.pq[cost] = append(s.pq[cost], v)
		if cost > s.highest {
			s.highest = cost
		}
	}
}

func (s *solver) pop(cost int) state {
	v := s.pq[cost][0]
	s.pq[cost] = s.pq[cost][1:]
	return v
}

func (s *solver) lookup(p position) byte { return s.grid[p.row][p.col] }

func (s *solver) isend(p position) bool { return s.lookup(p) == end }

func (s *solver) isopen(p position) bool { return s.lookup(p) != wall }

func solve(grid []string, start state) *solver {
	s := &solver{grid: grid, pq: map[int][]state{}, visited: map[state]int{}, prov: map[state]*provenance{}}
	s.add(start, state{}, 0)
	for {
		for len(s.pq[s.cheapest]) == 0 {
			if s.cheapest > s.highest {
				log.Fatalf("Ran out of priority queue: %d > %d", s.cheapest, s.highest)
			}
			s.cheapest++
		}
		v := s.pop(s.cheapest)
		if s.isend(v.pos) {
			s.end = v
			return s
		}
		straight, left, right := v.possible()
		if s.isopen(straight.pos) {
			s.add(straight, v, s.cheapest+1)
		}
		if s.isopen(left.pos) {
			s.add(left, v, s.cheapest+1000)
		}
		if s.isopen(right.pos) {
			s.add(right, v, s.cheapest+1000)
		}
	}
}

// endregion

// region Day17
func day17a() string {
	file := strings.Split(readDayFile(17), "\n")
	var res string
	comp := initComputer(file)
	out := Run(comp.A, comp.B, comp.C, comp.Program)
	for i, num := range out {
		if i > 0 {
			res += ","
		}

		res += strconv.FormatUint(num, 10)
	}

	return res
}

func day17b() uint64 {
	file := strings.Split(readDayFile(17), "\n")
	comp := initComputer(file)

	type QueueItem struct {
		a uint64
		n int
	}

	queue := list.New()
	queue.PushBack(QueueItem{a: 0, n: 1})

	for queue.Len() > 0 {
		item := queue.Remove(queue.Front()).(QueueItem)
		a, n := item.a, item.n

		if n > len(comp.Program) {
			return a
		}

		for i := uint64(0); i < 8; i++ {
			a2 := (a << 3) | i
			out := Run(a2, 0, 0, comp.Program)
			target := comp.Program[len(comp.Program)-n:]

			if matchesProgram(out, target) {
				queue.PushBack(QueueItem{a: a2, n: n + 1})
			}
		}
	}

	return 0
}

type SmallComputer struct {
	A, B, C uint64
	Program []uint64
	Out     []uint64
}

func initComputer(puzzle []string) *SmallComputer {
	var res = SmallComputer{
		A:       0,
		B:       0,
		C:       0,
		Program: nil,
		Out:     nil,
	}
	reR := regexp.MustCompile(`Register ([A|B|C]): (\d+)`)
	reP := regexp.MustCompile(`\d`)
	for _, line := range puzzle {
		if strings.Contains(line, "Program") {
			match := reP.FindAllStringSubmatch(line, -1)
			for m := 0; m < len(match); m++ {
				instruction, _ := strconv.ParseUint(match[m][0], 10, 64)
				res.Program = append(res.Program, instruction)

			}
		} else if strings.Contains(line, "Register") {
			match := reR.FindAllStringSubmatch(line, -1)

			register, _ := strconv.Atoi(match[0][2])
			if match[0][1] == "A" {
				res.A = uint64(register)
			} else if match[0][1] == "B" {
				res.B = uint64(register)
			} else if match[0][1] == "C" {
				res.C = uint64(register)
			}
		}
	}
	return &res
}

func Run(a, b, c uint64, program []uint64) []uint64 {
	var instruction uint64
	var param uint64
	out := []uint64{}

	for pointer := uint64(0); pointer < uint64(len(program)); pointer += 2 {
		instruction, param = program[pointer], program[pointer+1]

		combo := param
		switch param {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}
		switch instruction {
		case 0:
			a >>= combo
		case 1:
			b ^= param
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				pointer = param - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}
	return out
}

func matchesProgram(output []uint64, expected []uint64) bool {
	if len(output) != len(expected) {
		return false
	}
	for i := range output {
		if output[i] != expected[i] {
			return false
		}
	}
	return true
}

// endregion

// region Day18
func day18a() int {
	file := readDayFile(18)

	coordinates := AllBlockedCoords(file)
	start := Coord{0, 0}
	end := Coord{70, 70}

	shortestPath, _ := ShortestPath(coordinates[:1024], start, end)

	return shortestPath
}

func day18b() string {
	file := readDayFile(18)

	coordinates := AllBlockedCoords(file)
	start := Coord{0, 0}
	end := Coord{70, 70}

	block := SearchForBlockage(coordinates, start, end)

	return fmt.Sprintf("%d, %d", block.X, block.Y)
}

type Coord struct {
	X, Y int
}

type BlockedCoords map[Coord]struct{}

func blockedCoords(coords []Coord) BlockedCoords {
	bc := make(BlockedCoords)
	for _, c := range coords {
		bc[c] = struct{}{}
	}
	return bc
}

func AllBlockedCoords(s string) []Coord {
	coordStrs := strings.Split(s, "\n")
	allC := make([]Coord, len(coordStrs))

	for i, c := range coordStrs {
		parts := strings.Split(c, ",")
		x, errX := strconv.ParseInt(parts[0], 10, 0)
		y, errY := strconv.ParseInt(parts[1], 10, 0)
		if errX != nil || errY != nil {
			log.Fatal("Couldn't parse int.")
		}
		allC[i] = Coord{int(x), int(y)}
	}

	return allC
}

func (c Coord) isValidStep(b BlockedCoords) bool {
	_, blocked := b[c]
	return !(blocked || c.X < 0 || c.Y < 0 || c.X > 70 || c.Y > 70)
}

func (c Coord) nextCoords() []Coord {
	return []Coord{Coord{c.X + 1, c.Y}, Coord{c.X - 1, c.Y}, Coord{c.X, c.Y + 1}, Coord{c.X, c.Y - 1}}
}

func ShortestPath(coords []Coord, start, end Coord) (int, bool) {
	b := blockedCoords(coords)
	visited := map[Coord]int{start: 0}
	q := []Coord{start}
	var node Coord

	for len(q) > 0 {
		node, q = q[0], q[1:]

		for _, c := range node.nextCoords() {
			_, alreadyReached := visited[c]
			if c.isValidStep(b) && !alreadyReached {
				q = append(q, c)
				visited[c] = visited[node] + 1
			}
		}
	}

	distance, validPath := visited[end]
	return distance, validPath
}

func SearchForBlockage(allC []Coord, start, end Coord) Coord {
	l := 1024
	r := len(allC) - 1
	m := (l + r) / 2

	for l != m && r != m {
		_, ok := ShortestPath(allC[:m], start, end)

		if ok {
			l, r, m = m, r, (m+r)/2
		} else {
			l, r, m = l, m, (l+m)/2
		}
	}
	return allC[m]
}

// endregion

// region Day19
func day19a() int {
	file := readDayFile(19)
	towels, patterns := parseDesignFile(file)
	count := 0
	cache := make(map[string]bool)
	for _, pattern := range patterns {
		if designPossible(pattern, towels, cache) {
			count++
		}
	}
	return count
}

func day19b() int {
	file := readDayFile(19)
	towels, patterns := parseDesignFile(file)
	count := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		count += waysPossible(pattern, towels, cache)
	}
	return count
}

func parseDesignFile(s string) ([]string, []string) {
	parts := strings.Split(s, "\n\n")
	t := strings.Split(parts[0], ", ")
	p := strings.Split(parts[1], "\n")
	return t, p
}

func designPossible(pattern string, ts []string, cache map[string]bool) bool {
	b, ok := cache[pattern]
	if ok {
		return b
	}

	for _, t := range ts {
		if t == pattern {
			return true
		} else if strings.HasPrefix(pattern, t) {
			isPoss := designPossible(strings.TrimPrefix(pattern, t), ts, cache)
			if isPoss {
				cache[pattern] = true
				return true
			}
		}
	}
	cache[pattern] = false
	return false
}

func waysPossible(pattern string, ts []string, cache map[string]int) (ways int) {

	w, ok := cache[pattern]
	if ok {
		return w
	}

	for _, t := range ts {
		if t == pattern {
			ways++
		} else if strings.HasPrefix(pattern, t) {
			ways += waysPossible(strings.TrimPrefix(pattern, t), ts, cache)
		}
	}
	cache[pattern] = ways
	return
}

// endregion

// region Day20
func day20a() int {
	file := strings.Split(readDayFile(20), "\n")

	return getCheats(file, 2)
}

func day20b() int {
	file := strings.Split(readDayFile(20), "\n")

	return getCheats(file, 20)
}

type shortcut struct {
	start point
	end   offset
}

type offset struct {
	point    point
	distance int
}

type point struct {
	x int
	y int
}

func findRoute(start, end point, walls map[point]int) map[point]int {
	queue := []point{start}
	visited := make(map[point]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current] = len(visited)

		if current == end {
			return visited
		}

		for _, offset := range getOffsets(current, 1) {
			if _, found := visited[offset.point]; found {
				continue
			}

			if _, found := walls[offset.point]; found {
				continue
			}

			queue = append(queue, offset.point)
		}
	}

	panic("Cannot find route")
}

func findShortcuts(route map[point]int, radius int) map[int]int {
	shortcuts := make(map[shortcut]int)
	for current, step := range route {

		offsets := getOffsets(current, radius)
		for _, offset := range offsets {
			routeStep, inRoute := route[offset.point]
			if inRoute {
				saving := routeStep - step - offset.distance
				if saving > 0 {
					shortcuts[shortcut{current, offset}] = saving
				}
			}
		}
	}

	result := make(map[int]int)
	for _, saving := range shortcuts {
		result[saving]++
	}

	return result
}

func getOffsets(from point, radius int) []offset {
	result := []offset{}

	for y := radius * -1; y <= radius; y++ {
		for x := radius * -1; x <= radius; x++ {
			candidatePoint := point{from.x + x, from.y + y}
			candidate := offset{
				candidatePoint,
				getDistance(from, candidatePoint),
			}

			if candidate.distance > 0 && candidate.distance <= radius {
				result = append(result, candidate)
			}
		}
	}

	return result
}

func getDistance(from, until point) int {
	xDistance := math.Abs(float64(from.x - until.x))
	yDistance := math.Abs(float64(from.y - until.y))
	return int(xDistance + yDistance)
}

func getCheats(file []string, radius int) int {
	var start point
	var end point
	var walls map[point]int = make(map[point]int)

	for y, line := range file {
		for x, r := range line {
			switch r {
			case 'S':
				start = point{x, y}
			case 'E':
				end = point{x, y}
			case '#':
				walls[point{x, y}]++
			}
		}
	}

	route := findRoute(start, end, walls)
	cheats := findShortcuts(route, radius)

	var found int
	var greatShortcuts int
	for k := 0; found < len(cheats); k++ {
		if v, ok := cheats[k]; ok {
			found++

			if k >= 100 {
				greatShortcuts += v
			}
		}
	}

	return greatShortcuts
}

// endregion

// region Day21
func day21a() int {
	file := strings.Split(readDayFile(21), "\n")
	numMap := map[string]Coord{
		"A": {2, 0},
		"0": {1, 0},
		"1": {0, 1},
		"2": {1, 1},
		"3": {2, 1},
		"4": {0, 2},
		"5": {1, 2},
		"6": {2, 2},
		"7": {0, 3},
		"8": {1, 3},
		"9": {2, 3},
	}

	dirMap := map[string]Coord{
		"A": {2, 1},
		"^": {1, 1},
		"<": {0, 0},
		"v": {1, 0},
		">": {2, 0},
	}

	robots := 2

	return getSequence(file, numMap, dirMap, robots)
}

func day21b() int {
	file := strings.Split(readDayFile(21), "\n")
	numMap := map[string]Coord{
		"A": {2, 0},
		"0": {1, 0},
		"1": {0, 1},
		"2": {1, 1},
		"3": {2, 1},
		"4": {0, 2},
		"5": {1, 2},
		"6": {2, 2},
		"7": {0, 3},
		"8": {1, 3},
		"9": {2, 3},
	}

	dirMap := map[string]Coord{
		"A": {2, 1},
		"^": {1, 1},
		"<": {0, 0},
		"v": {1, 0},
		">": {2, 0},
	}

	robots := 25

	return getSequence(file, numMap, dirMap, robots)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func AtoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getSequence(input []string, numMap, dirMap map[string]Coord, robotCount int) int {
	total := 0
	cache := make(map[string][]int)

	for _, line := range input {
		chars := strings.Split(line, "")
		moves := getNumPadSequence(chars, "A", numMap)
		length := countSequences(moves, robotCount, 1, cache, dirMap)
		total += AtoiNoErr(strings.TrimLeft(line[:3], "0")) * length
	}

	return total
}

func getNumPadSequence(input []string, start string, numMap map[string]Coord) []string {
	curr := numMap[start]
	seq := []string{}

	for _, char := range input {
		dest := numMap[char]
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := []string{}, []string{}

		for i := 0; i < Abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		for i := 0; i < Abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		if curr.Y == 0 && dest.X == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if curr.X == 0 && dest.Y == 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func countSequences(input []string, maxRobots, robot int, cache map[string][]int, dirMap map[string]Coord) int {
	key := strings.Join(input, "")
	if val, ok := cache[key]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	if _, ok := cache[key]; !ok {
		cache[key] = make([]int, maxRobots)
	}

	seq := getDirPadSequence(input, "A", dirMap)
	if robot == maxRobots {
		return len(seq)
	}

	steps := splitSequence(seq)
	count := 0
	for _, step := range steps {
		c := countSequences(step, maxRobots, robot+1, cache, dirMap)
		count += c
	}

	cache[key][robot-1] = count
	return count
}

func getDirPadSequence(input []string, start string, dirMap map[string]Coord) []string {
	curr := dirMap[start]
	seq := []string{}

	for _, char := range input {
		dest := dirMap[char]
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := []string{}, []string{}

		for i := 0; i < Abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		for i := 0; i < Abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		if curr.X == 0 && dest.Y == 1 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if curr.Y == 1 && dest.X == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func splitSequence(input []string) [][]string {
	var result [][]string
	var current []string

	for _, char := range input {
		current = append(current, char)
		if char == "A" {
			result = append(result, current)
			current = []string{}
		}
	}
	return result
}

// endregion

// region Day22
func day22a() int {
	file := strings.Split(readDayFile(22), "\n")

	seeds := make([]int, len(file))
	for i, seed := range file {
		s, err := strconv.ParseInt(seed, 10, 0)
		if err != nil {
			log.Fatal("Couldn't parse int:", seed)
		}
		seeds[i] = int(s)
	}

	sum := 0
	for _, n := range seeds {
		sum += NextN(n, 2000)
	}
	return sum
}

func day22b() int {
	file := strings.Split(readDayFile(22), "\n")
	seeds := make([]int, len(file))
	for i, seed := range file {
		s, err := strconv.ParseInt(seed, 10, 0)
		if err != nil {
			log.Fatal("Couldn't parse int:", seed)
		}
		seeds[i] = int(s)
	}

	max := 0

	sequencePayoffs := MonkeyFromSeed(seeds, 2000)
	for _, v := range sequencePayoffs {
		if v > max {
			max = v
		}
	}
	return max
}

const pRUNE = 16777216

func next(n int) (out int) {
	out = prune(mix(n<<6, n))
	out = prune(mix(out>>5, out))
	out = prune(mix(out<<11, out))
	return
}

func mix(in, secret int) int {
	return in ^ secret
}

func prune(n int) int {
	return n % pRUNE
}

func NextN(seed, simSteps int) (random int) {
	random = seed
	for i := 0; i < simSteps; i++ {
		random = next(random)
	}
	return
}

func allPrices(seed, simSteps int) []price {
	random := seed
	retval := make([]price, simSteps+1)
	retval[0] = price{seed, 0}
	for i := 0; i < simSteps; i++ {
		random = next(random)
		cost := random % 10
		newPrice := price{cost, cost - retval[i].cost}
		retval[i+1] = newPrice
	}

	return retval
}

func MonkeyFromSeed(seeds []int, simSteps int) Monkey {
	retval := make(map[sequence]int)
	for _, seed := range seeds {
		prices := allPrices(seed, simSteps)
		seen := make(map[sequence]struct{})
		for i := 4; i <= simSteps; i++ {
			seq := sequence{prices[i-3].change, prices[i-2].change, prices[i-1].change, prices[i].change}
			if _, ok := seen[seq]; !ok {
				seen[seq] = struct{}{}
				retval[seq] += prices[i].cost
			}
		}
	}
	return retval
}

type sequence struct {
	first, seconds, third, fourth int
}

type price struct {
	cost, change int
}

type Monkey map[sequence]int

// endregion

// region Day23
func day23a() int {
	file := strings.Split(readDayFile(23), "\n")

	np := newNetworkProcessor()
	np.processLinks(file)
	np.findNetworks()
	return np.countNetworks()
}

func day23b() string {
	file := strings.Split(readDayFile(23), "\n")

	np := newNetworkProcessor()
	np.processLinks(file)
	np.findNetworks()
	return np.findBiggestNetwork()
}

type link struct {
	first, second string
}

type network struct {
	output      string
	connections map[string]struct{}
}

type computer struct {
	name  string
	links map[string]struct{}
}

type networkProcessor struct {
	networks    map[string]struct{}
	comps       map[string]computer
	simpleComps map[string][]string
}

func newNetworkProcessor() *networkProcessor {
	return &networkProcessor{
		networks:    make(map[string]struct{}),
		comps:       make(map[string]computer),
		simpleComps: make(map[string][]string),
	}
}

func (np *networkProcessor) processLinks(linkStrs []string) {
	for _, linkStr := range linkStrs {
		computers := strings.Split(linkStr, "-")
		l := link{computers[0], computers[1]}

		np.addOrUpdateComputer(l.first, l.second)
		np.addOrUpdateComputer(l.second, l.first)
	}

	np.populateSimpleComps()
}

func (np *networkProcessor) addOrUpdateComputer(compName, linkedCompName string) {
	comp, exists := np.comps[compName]
	if !exists {
		comp = computer{compName, make(map[string]struct{})}
	}
	comp.links[linkedCompName] = struct{}{}
	np.comps[compName] = comp
}

func (np *networkProcessor) populateSimpleComps() {
	for _, v := range np.comps {
		links := make([]string, 0, len(v.links))
		for li := range v.links {
			links = append(links, li)
		}
		np.simpleComps[v.name] = links
	}
}

func (np *networkProcessor) findNetworks() {
	for name, com := range np.simpleComps {
		linkCount := len(com)
		for i := 0; i < linkCount-1; i++ {
			for j := i + 1; j < linkCount; j++ {
				iName, jName := com[i], com[j]
				iCom := np.comps[iName]
				if _, iContainsJ := iCom.links[jName]; iContainsJ {
					names := []string{name, iName, jName}
					sort.Strings(names)
					np.networks[strings.Join(names, ",")] = struct{}{}
				}
			}
		}
	}
}

func (np *networkProcessor) countNetworks() (count int) {
	for n := range np.networks {
		if n[0] == 't' || n[3] == 't' || n[6] == 't' {
			count++
		}
	}
	return
}

func (np *networkProcessor) findBiggestNetwork() string {
	networkCache := make(map[string]struct{})
	var retval string

	for _, v := range np.comps {
		longestFound := np.bfs(v, networkCache)
		if len(longestFound) > len(retval) {
			retval = longestFound
		}
	}

	return retval
}

func (np *networkProcessor) bfs(c computer, networkCache map[string]struct{}) string {
	start := network{c.name, map[string]struct{}{c.name: {}}}
	networkCache[start.output] = struct{}{}
	queue := []network{start}

	var n network
	for len(queue) > 0 {
		n, queue = queue[0], queue[1:]
		for candidate := range c.links {
			if _, ok := n.connections[candidate]; ok {
				continue
			}
			if np.isConnected(n, candidate) {
				next := n.add(candidate)
				if _, ok := networkCache[next.output]; !ok {
					queue = append(queue, next)
					networkCache[next.output] = struct{}{}
				}
			}
		}
	}

	return n.output
}

func (np *networkProcessor) isConnected(n network, candidate string) bool {
	for existing := range n.connections {
		if _, ok := np.comps[existing].links[candidate]; !ok {
			return false
		}
	}
	return true
}

func (n network) add(c string) network {
	newConnections := make(map[string]struct{})
	for k, v := range n.connections {
		newConnections[k] = v
	}
	newConnections[c] = struct{}{}

	newOutput := make([]string, 0, len(newConnections))
	for k := range newConnections {
		newOutput = append(newOutput, k)
	}
	sort.Strings(newOutput)

	return network{
		output:      strings.Join(newOutput, ","),
		connections: newConnections,
	}
}

// endregion

// region Day24
func day24a() string {
	file := readDayFile(24)
	wires, gates := parsePuzzleInput(file)

	for len(gates) > 0 {
		for wireName, gate := range gates {
			if canEvalGate(gate, wires) {
				wires[wireName] = evaluateGate(gate, wires)
				delete(gates, wireName)
			}
		}
	}

	zWires := []string{}
	for wire := range wires {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	sort.Strings(zWires)

	result := 0
	for i := len(zWires) - 1; i >= 0; i-- {
		result = (result << 1) | wires[zWires[i]]
	}

	return strconv.Itoa(result)
}

func day24b() string {
	file := readDayFile(24)
	_, gates := parsePuzzleInput(file)

	var swapped []string
	var carry string

	gateStrings := []string{}
	for wireName, gate := range gates {
		gateStr := fmt.Sprintf("%s %s %s -> %s",
			gate.inputs[0],
			[]string{"AND", "OR", "XOR"}[gate.operation],
			gate.inputs[1],
			wireName)
		gateStrings = append(gateStrings, gateStr)
	}

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		m1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "XOR", gateStrings)
		n1 = find(fmt.Sprintf("x%s", n), fmt.Sprintf("y%s", n), "AND", gateStrings)

		if carry != "" {
			r1 = find(carry, m1, "AND", gateStrings)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(carry, m1, "AND", gateStrings)
			}

			z1 = find(carry, m1, "XOR", gateStrings)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gateStrings)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if carry == "" {
			carry = n1
		} else {
			carry = c1
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

type GateInfo struct {
	operation int
	inputs    []string
	output    string
}

func parsePuzzleInput(input string) (map[string]int, map[string]GateInfo) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	wires := make(map[string]int)
	gates := make(map[string]GateInfo)

	for _, line := range strings.Split(strings.TrimSpace(parts[0]), "\n") {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = AtoiNoErr(parts[1])
	}

	for _, line := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		inputs := strings.Split(parts[0], " ")

		var operation int
		var ins []string

		if len(inputs) == 3 {
			switch inputs[1] {
			case "AND":
				operation = 0
			case "OR":
				operation = 1
			case "XOR":
				operation = 2
			}
			ins = []string{inputs[0], inputs[2]}
		}

		gates[parts[1]] = GateInfo{
			operation: operation,
			inputs:    ins,
			output:    parts[1],
		}
	}

	return wires, gates
}

func evaluateGate(gate GateInfo, wires map[string]int) int {
	in1 := wires[gate.inputs[0]]
	in2 := wires[gate.inputs[1]]

	switch gate.operation {
	case 0:
		return in1 & in2
	case 1:
		return in1 | in2
	case 2:
		return in1 ^ in2
	}
	return 0
}

func canEvalGate(gate GateInfo, wires map[string]int) bool {
	_, hasIn1 := wires[gate.inputs[0]]
	_, hasIn2 := wires[gate.inputs[1]]
	return hasIn1 && hasIn2
}

func find(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			parts := strings.Split(gate, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}

// endregion

// region Day25
func day25() int {
	file := strings.Split(readDayFile(25), "\n")

	var locks, keys [][]int

	for i := 0; i < len(file); i += 8 {
		if i+7 > len(file) {
			break
		}

		heights := make([]int, 5)
		isLock := false

		for row := 0; row < 7; row++ {
			for col, char := range file[i+row] {
				if char == '#' {
					heights[col]++
				}
			}

			if row == 0 && file[i][0] == '#' {
				isLock = true
			}
		}

		for i := range heights {
			heights[i]--
		}

		if isLock {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	matches := 0

	for _, lock := range locks {
		for _, key := range keys {
			if checkMatch(lock, key) {
				matches++
			}
		}
	}

	return matches
}

func checkMatch(lock, key []int) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
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
