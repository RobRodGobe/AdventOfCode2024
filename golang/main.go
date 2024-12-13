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
	fmt.Println(day13a())
	fmt.Println(day13b())
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

func readDayFile(day int32) string {
	filePath := fmt.Sprintf("../AoC_Files/%d.txt", day)

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(content)
	return fileContents
}
