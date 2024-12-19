const { parse } = require('path');

function main() {
    // Day 1 a + b
    console.log(day16a(), day16b());
}

function readDayFile(day){
    const fs = require('fs');
    
    const filePath = `../AoC_Files/${day}.txt`
    const fileContents = fs.readFileSync(filePath, 'utf-8');

    return fileContents
}

// #region Day1
function day1a(){
    const file = readDayFile(1);
    const list1 = []
    const list2 = []

    file.split('\n').forEach(line => {
        const pairs = line.trim().split("   ");
        list1.push(pairs[0]);
        list2.push(pairs[1]);
    });

    list1.sort();
    list2.sort();

    let diff = 0;

    for (let i = 0; i < list1.length; i++) {
        diff += Math.abs(list1[i] - list2[i]);
    }

    return diff;
}

function day1b(){
    const file = readDayFile(1);
    const list1 = []
    const list2 = []

    file.split('\n').forEach(line => {
        const pairs = line.trim().split("   ");
        list1.push(pairs[0]);
        list2.push(pairs[1]);
    });

    let similar = 0;

    for (let i = 0; i < list1.length; i++) {
        similar += list1[i] * list2.filter(a => a == list1[i]).length;
    }

    return similar;
}

// #endregion

// #region Day2
function day2a() {
    const file = readDayFile(2).split("\n");
    let safe = 0;

    for (let i = 0; i < file.length; i++) {
        let reports = file[i].split(" ").map(Number);
        let isAscending = true;
        let isDescending = true;
        let isSafe = true;

        for (let j = 1; j < reports.length; j++) {
            let diff = reports[j] - reports[j - 1];

            if (Math.abs(diff) > 3) {
                isAscending = false;
                isDescending = false;
                isSafe = false;
                break;
            }

            if (diff < 0) isAscending = false;
            if (diff > 0) isDescending = false;
            if (diff === 0) 
            {
                isAscending = false;
                isDescending = false;
            }
            
            if (!isAscending && !isDescending) {
                isSafe = false;
                break;
            }
        }

        if (isSafe) {
            safe++;
        }
    }

    return safe;
}

function day2b() {
    const file = readDayFile(2).split("\n");
    let safe = 0;

    for (let i = 0; i < file.length; i++) {
        let reports = file[i].split(" ").map(Number);

        if (isSafeReport(reports, true) || isSafeReport(reports, false)) {
            safe++;
            continue;
        }

        let foundSafe = false;
        for (let j = 0; j < reports.length; j++) {
            let modifiedReports = [...reports.slice(0, j), ...reports.slice(j + 1)];
            if (isSafeReport(modifiedReports, true) || isSafeReport(modifiedReports, false)) {
                foundSafe = true;
                break;
            }
        }

        if (foundSafe) {
            safe++;
        }
    }

    return safe;
}

function isSafeReport(reports, ascending) {
    for (let i = 1; i < reports.length; i++) {
        let diff = reports[i] - reports[i - 1];
        if (ascending && diff < 0) return false;  
        if (!ascending && diff > 0) return false; 
        if (Math.abs(diff) > 3 || diff === 0) return false; 
    }
    return true;
}

// #endregion

// #region Day3
function day3a() {
    let mult = 0;
    const file = readDayFile(3);

    const pattern = /mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)/gm;
    const matches = file.match(pattern);

    for (let i = 0; i < matches.length; i++) {
        const numbers = matches[i].replace("mul(", "").replace(")", "").split(",");
        mult += parseInt(numbers[0]) * parseInt(numbers[1]);
    }

    return mult;
}

function day3b() {
    let mult = 0;
    const file = readDayFile(3);

    const pattern = /mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)|do\(\)|don't\(\)/gm;
    const matches = file.match(pattern);

    let multiply = true;

    for (let i = 0; i < matches.length; i++) {
        if (matches[i] === "do()") {
            multiply = true;
        }
        else if (matches[i] === "don't()") {
            multiply = false;
        }

        if (multiply && !matches[i].includes("do")) {
            const numbers = matches[i].replace("mul(", "").replace(")", "").split(",");
            mult += parseInt(numbers[0]) * parseInt(numbers[1]);
        }
    }
    
    return mult;
}
// #endregion

// #region Day4
function day4a() {
    const file = readDayFile(4).split("\n");
    const word = "XMAS";
    const rows = file.length;
    const cols = file[0].length;
    let count = 0;
    const wordLength = word.length;

    const directions = [
        [0, 1],   // Right
        [1, 0],   // Down
        [1, 1],   // Down-right
        [1, -1],  // Down-left
        [0, -1],  // Left
        [-1, 0],  // Up
        [-1, -1], // Up-left
        [-1, 1]   // Up-right
    ];

    for (let x = 0; x < rows; x++) {
        for (let y = 0; y < cols; y++) {
            for (let [dx, dy] of directions) {
                if (checkWordBegin(x, y, dx, dy, wordLength, rows, cols, word, file)) {
                    count++;
                }
            }
        }
    }

    return count;
}

function day4b() {
    const file = readDayFile(4).split("\n");
    const rows = file.length;
    const cols = file[0].length;
    let count = 0;

    for (let x = 1; x < rows - 1; x++) {
        for (let y = 1; y < cols - 1; y++) {
            if (isXmasPattern(file, x, y)) {
                count++;
            }
        }
    }

    return count;
}

function checkWordBegin(x, y, dx, dy, length, rows, cols, word, grid) {
    for (let i = 0; i < length; i++) {
        const nx = x + i * dx;
            const ny = y + i * dy;

            if (nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] !== word[i]) {
                return false;
            }
    }

    return true;
}

function isXmasPattern(grid, x, y) {
    const topLeftToBottomRight = `${grid[x - 1][y - 1]}${grid[x][y]}${grid[x + 1][y + 1]}`;
    const topRightToBottomLeft = `${grid[x - 1][y + 1]}${grid[x][y]}${grid[x + 1][y - 1]}`;

    return (isValidMasPattern(topLeftToBottomRight) && isValidMasPattern(topRightToBottomLeft));
}

function isValidMasPattern(pattern) {
    return pattern === "MAS" || pattern === "SAM";
}
// #endregion

// #region Day5
function day5a() {
    const file = readDayFile(5).split('\n');
    const dividerIndex = file.indexOf('');
    const rules = file.slice(0, dividerIndex).map(line => {
        const [before, after] = line.split('|').map(Number);
        return { before, after };
    });
    const updates = file.slice(dividerIndex + 1).map(line => line.split(',').map(Number));

    let pages = 0;
    for (const update of updates) {
        if (isUpdateValid(update, rules)) {
            pages += getMiddlePage(update);
        }
    }

    return pages;
}

function day5b() {
    const file = readDayFile(5).split('\n');
    const dividerIndex = file.indexOf('');
    const rules = file.slice(0, dividerIndex).map(line => {
        const [before, after] = line.split('|').map(Number);
        return { before, after };
    });
    const updates = file.slice(dividerIndex + 1).map(line => line.split(',').map(Number));

    let pages = 0;
    for (const update of updates) {
        if (!isUpdateValid(update, rules)) {
            const correctedUpdate = correctUpdate(update, rules);
            pages += getMiddlePage(correctedUpdate);
        }
    }

    return pages;
}

function isUpdateValid(update, rules) {
    const pagePositions = new Map(update.map((page, index) => [page, index]));
    for (const { before, after } of rules) {
        if (pagePositions.has(before) && pagePositions.has(after)) {
            if (pagePositions.get(before) >= pagePositions.get(after)) {
                return false;
            }
        }
    }
    return true;
}

function getMiddlePage(update) {
    const midIndex = Math.floor(update.length / 2);
    return update[midIndex];
}

function correctUpdate(update, rules) {
    const graph = new Map(update.map(page => [page, []]));
    const inDegree = new Map(update.map(page => [page, 0]));

    for (const { before, after } of rules) {
        if (update.includes(before) && update.includes(after)) {
            graph.get(before).push(after);
            inDegree.set(after, inDegree.get(after) + 1);
        }
    }

    const queue = [...inDegree.entries()].filter(([_, degree]) => degree === 0).map(([page]) => page);
    const sorted = [];

    while (queue.length > 0) {
        const current = queue.shift();
        sorted.push(current);

        for (const neighbor of graph.get(current) || []) {
            inDegree.set(neighbor, inDegree.get(neighbor) - 1);
            if (inDegree.get(neighbor) === 0) {
                queue.push(neighbor);
            }
        }
    }

    return sorted;
}
// #endregion

// #region Day6
function day6a() {
    const file = readDayFile(6).split("\n").map(line => line.trim());;
    const rows = file.length;
    const cols = file[0].length;

    const directions = {
        "^": [-1, 0],
        ">": [0, 1],
        "v": [1, 0],
        "<": [0, -1]
    };

    const turnRight = {
        "^": ">",
        ">": "v",
        "v": "<",
        "<": "^"
    };
    
    let guardPos = [0, 0];
    let guardDir = "";

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (directions[file[r][c]]) {
                guardPos = [r, c];
                guardDir = file[r][c];
                break;
            }
        }
    }

    const visited = new Set();
    visited.add(`${guardPos[0]},${guardPos[1]}`);

    while (true) {
        const [dy, dx] = directions[guardDir];
        const nextPos = [guardPos[0] + dy, guardPos[1] + dx];

        if (nextPos[0] < 0 || nextPos[0] >= rows || nextPos[1] < 0 || nextPos[1] >= cols) {
            break;
        }

        if (file[nextPos[0]][nextPos[1]] === "#") {
            guardDir = turnRight[guardDir];
        }
        else {
            guardPos = nextPos;
            visited.add(`${guardPos[0]},${guardPos[1]}`);
        }
    }

    return visited.size;
}

function day6b() {
    const file = readDayFile(6).split("\n").map(line => line.trim());;
    const rows = file.length;
    const cols = file[0].length;

    let guardPos = [0, 0];
    let guardDir = "";

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if ("^>v<".includes(file[r][c])) {
                guardPos = [r, c];
                guardDir = file[r][c];
            }
        }
    }

    let loopPositions = 0;

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (isGuardInLoop(file, guardPos, guardDir,[r, c]))
                loopPositions++;
        }
    }

    return loopPositions;
}

function isGuardInLoop(mapLines, guardStart, guardDir, obstruction) {
    const directions = { '^': [-1, 0], '>': [0, 1], 'v': [1, 0], '<': [0, -1] };
    const turnRight = { '^': '>', '>': 'v', 'v': '<', '<': '^' };

    const rows = mapLines.length;
    const cols = mapLines[0].length;

    // Add obstruction
    const tempMap = mapLines.map(row => row.split(''));
    tempMap[obstruction[0]][obstruction[1]] = '#';

    let guardPos = guardStart;
    let currentDir = guardDir;

    const visitedStates = new Set();
    const recentHistory = [];
    const maxHistoryLength = 10;

    let steps = 0;
    const maxSteps = rows * cols * 2;

    while (true) {
        const state = `${guardPos[0]},${guardPos[1]},${currentDir}`;
        if (visitedStates.has(state)) {
            if (recentHistory.includes(state)) {
                return true;
            }
        }

        visitedStates.add(state);
        recentHistory.push(state);
        if (recentHistory.length > maxHistoryLength) {
            recentHistory.shift();
        }

        const [dx, dy] = directions[currentDir];
        const nextPos = [guardPos[0] + dx, guardPos[1] + dy];

        if (nextPos[0] < 0 || nextPos[0] >= rows || nextPos[1] < 0 || nextPos[1] >= cols) {
            return false;
        }
        else if (tempMap[nextPos[0]][nextPos[1]] === '#') {
            currentDir = turnRight[currentDir];
        } else {
            guardPos = nextPos;
        }

        steps++;
        if (steps > maxSteps) {
            return true;
        }
    }
}
// #endregion

// #region Day7
function day7a() {
    const file = readDayFile(7).split("\n");
    let sum = 0;

    for (let i = 0; i < file.length; i++) {
        const nums = file[i].split(":");
        const total = nums[0];
        const factors = nums[1].trim().split(" ").map(Number);
        if (canCalibrate(total, factors, factors[0], 1))
            sum += parseInt(total);
    }
    
    return sum;
}

function day7b() {
    const file = readDayFile(7).split("\n");
    let sum = 0;

    for (let i = 0; i < file.length; i++) {
        const nums = file[i].split(":");
        const total = nums[0];
        const factors = nums[1].trim().split(" ").map(Number);
        if (canCalibrate2(total, factors, factors[0], 1))
            sum += parseInt(total);
    }
    
    return sum;
}

function canCalibrate(target, numbers, current, i) {
    if (i === numbers.length) {
        return parseInt(current) === parseInt(target);
    }

    if (canCalibrate(target, numbers, current + numbers[i], i + 1)) {
        return true;
    }

    if (canCalibrate(target, numbers, current * numbers[i], i + 1)) {
        return true;
    }
    
    return false;
}

function canCalibrate2(target, numbers, current, i) {
    if (i === numbers.length) {
        return parseInt(current) === parseInt(target);
    }

    if (canCalibrate2(target, numbers, current + numbers[i], i + 1)) {
        return true;
    }

    if (canCalibrate2(target, numbers, current * numbers[i], i + 1)) {
        return true;
    }

    if (canCalibrate2(target, numbers, parseInt(`${current}${numbers[i]}`), i + 1)) {
        return true;
    }
    
    return false;
}
// #endregion

// #region Day8
function day8a() {
    const file = readDayFile(8).split("\n");
    const matrix = file.map(line => line.split(''));

    const antennaMap = getAntennaMap(matrix);

    const allAntinodes = [];

    Object.values(antennaMap).forEach(coords => {
        const antinodes = getAntinodes(coords, matrix);
        allAntinodes.push(...antinodes);
    });

    const uniqueAntinodes = getUniqueAntinodes(allAntinodes);

    return uniqueAntinodes.length;
}

function day8b() {
    const file = readDayFile(8).split("\n");
    const matrix = file.map(line => line.split(''));

    const antennaMap = getAntennaMap(matrix);

    const antinodeMatrix = Array.from({ length: matrix.length }, () => Array(matrix[0].length).fill(false));

    for (const coords of Object.values(antennaMap)) {
        processAntinodeLines(coords, matrix, antinodeMatrix);
    }

    return getUniqueAntinodesCount(antinodeMatrix);
}

function getAntennaMap(matrix) {
    const map = {};
    for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
            const cell = matrix[i][j];
            if (cell !== '.') {
                if (!map[cell]) map[cell] = [];
                map[cell].push({ x: i, y: j });
            }
        }
    }
    return map;
}

function getAntinodes(coords, matrix) {
    const antinodes = [];
    for (let i = 0; i < coords.length; i++) {
        for (let j = 0; j < coords.length; j++) {
            if (i !== j) {
                const { x: ax, y: ay } = coords[i];
                const { x: bx, y: by } = coords[j];

                const cx = 2 * bx - ax;
                const cy = 2 * by - ay;

                if (withinBoundaries(cx, 0, matrix.length) && withinBoundaries(cy, 0, matrix[0].length)) {
                    antinodes.push({ x: cx, y: cy });
                }
            }
        }
    }
    return antinodes;
}

function withinBoundaries(value, min, max) {
    return value >= min && value < max;
}

function getUniqueAntinodes(antinodes) {
    const uniqueSet = new Set(antinodes.map(({ x, y }) => `${x}:${y}`));
    return Array.from(uniqueSet).map(key => {
        const [x, y] = key.split(':').map(Number);
        return { x, y };
    });
}

function processAntinodeLines(coords, matrix, antinodeMatrix) {
    for (let i = 0; i < coords.length; i++) {
        for (let j = 0; j < coords.length; j++) {
            if (i !== j) {
                const { x: x1, y: y1 } = coords[i];
                const { x: x2, y: y2 } = coords[j];

                for (let x = 0; x < matrix.length; x++) {
                    for (let y = 0; y < matrix[0].length; y++) {
                        if (!antinodeMatrix[x][y]) {
                            const lineResult = (y1 - y2) * x + (x2 - x1) * y + (x1 * y2 - x2 * y1);
                            if (lineResult === 0) {
                                antinodeMatrix[x][y] = true;
                            }
                        }
                    }
                }
            }
        }
    }
}

function getUniqueAntinodesCount(antinodeMatrix) {
    let count = 0;
    for (const row of antinodeMatrix) {
        for (const cell of row) {
            if (cell) count++;
        }
    }
    return count;
}
// #endregion

// #region Day9
function day9a() {
    const file = readDayFile(9);

    let diskMap = parseDiskMap(file);
    diskMap = compactDisk(diskMap);

    return calculateChecksum(diskMap);
}

function day9b() {
    const file = readDayFile(9);

    let diskMap = parseDiskMap(file);
    diskMap = compactDiskByFile(diskMap);

    return calculateChecksum(diskMap);
}

function parseDiskMap(line) {
    const nums = [];
    let index = 0;

    for (let i = 0; i < line.length; i++) {
        const count = parseInt(line[i], 10);
        if (i % 2 === 0) {
            for (let j = 0; j < count; j++) {
                nums.push(index.toString());
            }
            index++;
        } else {
            for (let j = 0; j < count; j++) {
                nums.push(".");
            }
        }
    }

    return nums;
}

function compactDisk(diskMap) {
    let L = 0;
    let R = diskMap.length - 1;

    while (L <= R) {
        if (diskMap[L] === "." && diskMap[R] !== ".") {
            [diskMap[L], diskMap[R]] = [diskMap[R], diskMap[L]];
            R--;
            L++;
        } else if (diskMap[R] === ".") {
            R--;
        } else {
            L++;
        }
    }

    return diskMap;
}

function calculateChecksum(diskMap) {
    let checksum = 0;

    for (let i = 0; i < diskMap.length; i++) {
        if (diskMap[i] !== ".") {
            checksum += i * parseInt(diskMap[i], 10);
        }
    }

    return checksum;
}

class DiskFile {
    constructor(id, length, startIdx) {
        this.id = id;
        this.length = length;
        this.startIdx = startIdx;
    }
}

function analyzeDisk(diskMap) {
    const files = [];
    const spaces = {};
    let spaceStartIdx = -1;

    diskMap.forEach((item, i) => {
        if (item === ".") {
            if (spaceStartIdx === -1) spaceStartIdx = i;
            spaces[spaceStartIdx] = (spaces[spaceStartIdx] || 0) + 1;
        } else {
            if (spaceStartIdx !== -1) spaceStartIdx = -1;

            const fileId = parseInt(item, 10);
            if (files.length <= fileId) files[fileId] = new DiskFile(fileId, 0, i);

            files[fileId].length++;
        }
    });

    return { files, spaces };
}

function getFirstAvailableSpaceIdx(spaces, fileLength) {
    for (const [idx, count] of Object.entries(spaces)) {
        if (count >= fileLength) return parseInt(idx, 10);
    }
    return -1;
}

function updateSpaces(spaces, spaceIdx, fileLength) {
    if (spaces[spaceIdx] === fileLength) {
        delete spaces[spaceIdx];
    } else {
        const remainingSpace = spaces[spaceIdx] - fileLength;
        delete spaces[spaceIdx];
        spaces[spaceIdx + fileLength] = remainingSpace;
    }
}

function moveFile(diskMap, file, targetIdx) {
    for (let i = 0; i < file.length; i++) {
        diskMap[targetIdx + i] = diskMap[file.startIdx + i];
        diskMap[file.startIdx + i] = ".";
    }
}

function compactDiskByFile(diskMap) {
    const { files, spaces } = analyzeDisk(diskMap);

    files.sort((a, b) => b.id - a.id);

    files.forEach((file) => {
        const targetIdx = getFirstAvailableSpaceIdx(spaces, file.length);
        if (targetIdx !== -1 && targetIdx < file.startIdx) {
            moveFile(diskMap, file, targetIdx);
            updateSpaces(spaces, targetIdx, file.length);
        }
    });

    return diskMap;
}
// #endregion

// #region Day10
function day10a() {
    const file = readDayFile(10).split("\n");

    return solveTopographicMap(file);
}

function day10b() {
    const file = readDayFile(10).split("\n");

    return solveTopographicMapTrailRatings(file);
}

function solveTopographicMap(input) {
    const rows = input.length;
    const cols = input[0].length;
    const map = Array.from({ length: rows }, (_, r) => 
        input[r].split('').map(c => parseInt(c))
    );

    let totalTrailheadScore = 0;

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (map[r][c] === 0) {
                const trailheadScore = findTrailheadScore(r, c, map, rows, cols);
                totalTrailheadScore += trailheadScore;
            }
        }
    }

    return totalTrailheadScore;
}

function findTrailheadScore(startRow, startCol, map, rows, cols) {
    const visited = Array.from({ length: rows }, () => 
        Array(cols).fill(false)
    );
    const ninePositions = new Set();

    dfs(startRow, startCol, 0, ninePositions, map, rows, cols, visited);

    return ninePositions.size;
}

function dfs(row, col, expectedHeight, ninePositions, map, rows, cols, visited) {
    if (row < 0 || row >= rows || col < 0 || col >= cols || 
        visited[row][col] || map[row][col] !== expectedHeight) {
        return false;
    }

    visited[row][col] = true;

    if (expectedHeight === 9) {
        ninePositions.add(`${row},${col}`);
    }

    const directions = [
        [-1, 0], [1, 0], [0, -1], [0, 1]
    ];

    return directions.some(([dr, dc]) => 
        dfs(row + dr, col + dc, expectedHeight + 1, ninePositions, map, rows, cols, visited)
    );
}

function solveTopographicMapTrailRatings(input) {
    const { map, rows, cols } = parseTopographicMap(input);
    const scoresMap = new Map();

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (map[r][c] === 0) {
                startHike(r, c, map, rows, cols, scoresMap);
            }
        }
    }

    return getScoresSum(scoresMap);
}

function startHike(startR, startC, map, rows, cols, scoresMap) {
    const routes = [{ r: startR, c: startC, height: 0, initialCell: serializeCoordinates(startR, startC) }];

    while (routes.length > 0) {
        const route = routes.shift();

        for (const [dr, dc] of [[-1, 0], [1, 0], [0, -1], [0, 1]]) {
            const newR = route.r + dr;
            const newC = route.c + dc;
            const newHeight = route.height + 1;

            if (newR < 0 || newR >= rows || newC < 0 || newC >= cols) continue;

            const newCell = map[newR][newC];
            
            if (newCell !== newHeight) continue;

            if (newCell === 9) {
                if (!scoresMap.has(route.initialCell)) {
                    scoresMap.set(route.initialCell, new Map());
                }
                const endKey = serializeCoordinates(newR, newC);
                const currentScore = (scoresMap.get(route.initialCell).get(endKey) || 0) + 1;
                scoresMap.get(route.initialCell).set(endKey, currentScore);
            } else {
                routes.push({
                    r: newR,
                    c: newC,
                    height: newHeight,
                    initialCell: route.initialCell
                });
            }
        }
    }
}

function serializeCoordinates(r, c) {
    return `${r}:${c}`;
}

function getScoresSum(scoresMap) {
    let sum = 0;
    for (const scores of scoresMap.values()) {
        for (const score of scores.values()) {
            sum += score;
        }
    }
    return sum;
}

function parseTopographicMap(input) {
    const rows = input.length;
    const cols = input[0].length;
    const map = Array.from({ length: rows }, (_, r) => 
        input[r].split('').map(c => parseInt(c))
    );

    return { map, rows, cols };
}

// #endregion

// #region Day11
function day11a() {
    const file = readDayFile(11);
    const stones = file.split(" ").map(Number);
    const rocks = stones.reduce((map, stone) => {
        map.set(stone, (map.get(stone) || 0) + 1);
        return map;
    }, new Map());

    const finalRocks = blinkRocks(rocks, 25);
    return Array.from(finalRocks.values()).reduce((a, b) => a + b, 0);
}

function day11b() {
    const file = readDayFile(11);
    const stones = file.split(" ").map(Number);
    const rocks = stones.reduce((map, stone) => {
        map.set(stone, (map.get(stone) || 0) + 1);
        return map;
    }, new Map());

    const finalRocks = blinkRocks(rocks, 75);
    return Array.from(finalRocks.values()).reduce((a, b) => a + b, 0);
}

function blink(rock) {
    if (rock === 0) return [1];

    const digits = Math.floor(Math.log10(rock)) + 1;

    if (digits % 2 !== 0) return [rock * 2024];

    const halfDigits = Math.floor(digits / 2);
    const first = Math.floor(rock / Math.pow(10, halfDigits));
    const second = rock % Math.pow(10, halfDigits);

    return [first, second];
}

function blinkRocksIteration(rocks) {
    const result = new Map();

    for (const [rock, count] of rocks.entries()) {
        const newRocks = blink(rock);

        for (const newRock of newRocks) {
            result.set(newRock, (result.get(newRock) || 0) + count);
        }
    }

    return result;
}

function blinkRocks(rocks, blinks) {
    let currentRocks = new Map(rocks);

    for (let i = 0; i < blinks; i++) {
        currentRocks = blinkRocksIteration(currentRocks);
    }

    return currentRocks;
}
// #endregion

// #region Day12
function day12a() {
    const file = readDayFile(12).split("\n");
    return calculateTotalFencingPrice(file);
}

function day12b() {
    const file = readDayFile(12).split("\n");
    return calculateTotalFencingPriceWithInnerSides(file);
}

function calculateTotalFencingPrice(grid) {
    const n = grid.length;
    const m = grid[0].length;
    const visited = new Set();
    let totalPrice = 0;

    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            if (!visited.has(`${i},${j}`)) {
                const [area, borders] = visitRegion(grid, i, j, visited);
                totalPrice += area * borders.size;
            }
        }
    }

    return totalPrice;
}

function calculateTotalFencingPriceWithInnerSides(grid) {
    const n = grid.length;
    const m = grid[0].length;
    const visited = new Set();
    let totalPrice = 0;

    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            if (!visited.has(`${i},${j}`)) {
                const [area, borders] = visitRegion(grid, i, j, visited);
                totalPrice += area * countSides(borders);
            }
        }
    }

    return totalPrice;
}

function visitRegion(grid, startI, startJ, visited) {
    const n = grid.length;
    const m = grid[0].length;
    const plant = grid[startI][startJ];
    let area = 0;
    const borders = new Set();

    function visit(i, j) {
        if (visited.has(`${i},${j}`)) return;

        visited.add(`${i},${j}`);
        area++;

        const dx = [-1, 1, 0, 0];
        const dy = [0, 0, -1, 1];

        for (let k = 0; k < 4; k++) {
            const i2 = i + dx[k];
            const j2 = j + dy[k];

            if (i2 >= 0 && i2 < n && j2 >= 0 && j2 < m && grid[i2][j2] === plant) {
                visit(i2, j2);
            } else {
                borders.add(`${i},${j},${i2},${j2}`);
            }
        }
    }

    visit(startI, startJ);
    return [area, borders];
}

function countSides(borders) {
    const visited = new Set();

    function visitSide(i, j, i2, j2) {
        const side = `${i},${j},${i2},${j2}`;
        if (visited.has(side) || !borders.has(side)) return;

        visited.add(side);

        if (i === i2) {
            visitSide(i - 1, j, i2 - 1, j2);
            visitSide(i + 1, j, i2 + 1, j2);
        } else {
            visitSide(i, j - 1, i2, j2 - 1);
            visitSide(i, j + 1, i2, j2 + 1);
        }
    }

    let numSides = 0;
    for (const side of borders) {
        const [i, j, i2, j2] = side.split(',').map(Number);
        if (visited.has(side)) continue;

        numSides++;
        visitSide(i, j, i2, j2);
    }

    return numSides;
}
// #endregion

// #region Day13
function day13a() {
    const file = readDayFile(13).split("\n");
    return getMaxPrizeForMinTokens(file);
}

function day13b() {
    const file = readDayFile(13).split("\n");
    const machines = parseClawMachineInput(file);

    machines.forEach(machine => {
        machine.prizeX += 10_000_000_000_000;
        machine.prizeY += 10_000_000_000_000;
    });

    const adjustedInput = [];
    machines.forEach(machine => {
        adjustedInput.push(`Button A: X+${machine.ax}, Y+${machine.ay}`);
        adjustedInput.push(`Button B: X+${machine.bx}, Y+${machine.by}`);
        adjustedInput.push(`Prize: X=${machine.prizeX}, Y=${machine.prizeY}`);
    });

    return getMaxPrizeForMinTokens(adjustedInput);
}

class ClawMachineSettings {
    constructor(ax, ay, bx, by, prizeX, prizeY) {
        this.ax = ax;
        this.ay = ay;
        this.bx = bx;
        this.by = by;
        this.prizeX = prizeX;
        this.prizeY = prizeY;
    }
}

function parseClawMachineInput(inputData) {
    const machines = [];
    const cleanedData = inputData.filter(line => line.trim() !== "");

    for (let i = 0; i < cleanedData.length; i += 3) {
        const aMove = cleanedData[i].replace("Button A: ", "").split(", ");
        const bMove = cleanedData[i + 1].replace("Button B: ", "").split(", ");
        const prize = cleanedData[i + 2].replace("Prize: ", "").split(", ");

        machines.push(new ClawMachineSettings(
            parseInt(aMove[0].replace("X+", "")),
            parseInt(aMove[1].replace("Y+", "")),
            parseInt(bMove[0].replace("X+", "")),
            parseInt(bMove[1].replace("Y+", "")),
            parseInt(prize[0].replace("X=", "")),
            parseInt(prize[1].replace("Y=", ""))
        ));
    }

    return machines;
}

function calculatePrice(machine) {
    const det = machine.ay * machine.bx - machine.ax * machine.by;
    if (det === 0) return null;

    const b = Math.floor((machine.ay * machine.prizeX - machine.ax * machine.prizeY) / det);
    const a = machine.ax !== 0 ? Math.floor((machine.prizeX - b * machine.bx) / machine.ax) : 0;

    if (machine.ax * a + machine.bx * b === machine.prizeX &&
        machine.ay * a + machine.by * b === machine.prizeY &&
        a >= 0 && b >= 0) {
        return a * 3 + b;
    }

    return null;
}

function getMaxPrizeForMinTokens(inputData) {
    const machines = parseClawMachineInput(inputData);
    let totalTokens = 0;

    machines.forEach(machine => {
        const tokens = calculatePrice(machine);
        if (tokens !== null) {
            totalTokens += tokens;
        }
    });

    return totalTokens;
}
// #endregion

// #region Day14
function day14a() {
    const file = readDayFile(14).split("\n");

    return calculateSafetyFactor(file);
}

function day14b() {
    const file = readDayFile(14).split("\n");

    return findRobotSequenceTime(file);
}

class BathroomRobot {
    constructor(P, V) {
        this.P = P;
        this.V = V;
    }

    static simulateRobot(robot, modRows, modCols, ticks) {
        const rowDelta = BathroomRobot.calculateDelta(robot.V.Y, ticks, modRows);
        const newRow = BathroomRobot.modAdd(robot.P.Y, rowDelta, modRows);

        const colDelta = BathroomRobot.calculateDelta(robot.V.X, ticks, modCols);
        const newCol = BathroomRobot.modAdd(robot.P.X, colDelta, modCols);

        return new BathroomRobot({X: newCol, Y: newRow}, robot.V);
    }

    static calculateDelta(velocity, ticks, mod) {
        let delta = velocity * ticks % mod;
        return delta < 0 ? delta + mod : delta;
    }

    static modAdd(a, b, mod) {
        let res = (a + b) % mod;
        return res < 0 ? res + mod : res;
    }
}

function calculateSafetyFactor(file) {
    const width = 101;
    const height = 103;
    const duration = 100;

    const robots = parseRobots(file);
    const finalPositions = calculateFinalPositions(robots, width, height, duration);

    return computeQuadrantMultiplier(finalPositions, width, height);
}

function parseRobots(lines) {
    return lines
        .filter(line => line.trim() !== '')
        .map(parseSingleRobot);
}

function parseSingleRobot(line) {
    const parts = line.split(' ');
    const p = parts[0].substring(2).split(',');
    const v = parts[1].substring(2).split(',');

    return new BathroomRobot(
        {X: parseInt(p[0]), Y: parseInt(p[1])},
        {X: parseInt(v[0]), Y: parseInt(v[1])}
    );
}

function calculateFinalPositions(robots, width, height, duration) {
    const finalPositions = Array.from({length: width}, () => 
        Array(height).fill(0)
    );

    for (const robot of robots) {
        let finalX = (robot.P.X + robot.V.X * duration) % width;
        let finalY = (robot.P.Y + robot.V.Y * duration) % height;

        finalX = finalX < 0 ? finalX + width : finalX;
        finalY = finalY < 0 ? finalY + height : finalY;

        finalPositions[finalX][finalY]++;
    }

    return finalPositions;
}

function computeQuadrantMultiplier(finalPositions, width, height) {
    const midX = Math.floor(width / 2);
    const midY = Math.floor(height / 2);

    let topLeft = 0, topRight = 0, bottomLeft = 0, bottomRight = 0;

    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            if (x === midX || y === midY) continue;

            if (x < midX && y < midY) topLeft += finalPositions[x][y];
            else if (x >= midX && y < midY) topRight += finalPositions[x][y];
            else if (x < midX && y >= midY) bottomLeft += finalPositions[x][y];
            else if (x >= midX && y >= midY) bottomRight += finalPositions[x][y];
        }
    }

    return topLeft * topRight * bottomLeft * bottomRight;
}

function findRobotSequenceTime(file) {
    const rows = 103;
    const cols = 101;
    const trunkSeqSize = 10;
    const maxSeconds = 100000;

    let robots = parseRobots(file);

    for (let sec = 1; sec <= maxSeconds; sec++) {
        const robotsByCol = Array.from({length: cols}, () => []);

        for (let i = 0; i < robots.length; i++) {
            const newRobot = BathroomRobot.simulateRobot(robots[i], rows, cols, 1);
            robotsByCol[newRobot.P.X].push(newRobot.P.Y);
            robots[i] = newRobot;
        }

        for (let c = 0; c < cols; c++) {
            const column = robotsByCol[c];
            column.sort((a, b) => a - b);

            if (hasConsecutiveSequence(column, trunkSeqSize)) {
                return sec;
            }
        }
    }

    return 0;
}

function hasConsecutiveSequence(sequence, requiredLength) {
    if (sequence.length < requiredLength) return false;

    let consecutiveCount = 1;
    for (let i = 1; i < sequence.length; i++) {
        consecutiveCount = sequence[i] === sequence[i-1] + 1 
            ? consecutiveCount + 1 
            : 1;

        if (consecutiveCount === requiredLength) return true;
    }

    return false;
}
// #endregion

// #region Day15
function day15a() {
    const file = readDayFile(15);
    const { inputMap, moves } = parseInput(file);
    const map = new Map();
    let robotPos = RobotVector.Zero;

    inputMap.forEach((line, row) => {
        for (let col = 0; col < line.length; col++) {
            const tile = line[col];
            const pos = new RobotVector(row, col);
            if (tile === "#" || tile === "O") map.set(`${pos.row},${pos.col}`, tile);
            else if (tile === "@") robotPos = pos;
        }
    });

    for (const move of moves) {
        const dir = { ">": RobotVector.Right, "v": RobotVector.Down, "<": RobotVector.Left, "^": RobotVector.Up }[move];
        const thingsToPush = [];
        let next = robotPos.add(dir);

        while (true) {
            const tile = map.get(`${next.row},${next.col}`);
            if (tile) {
                thingsToPush.push(tile);
                if (tile === "#") break;
                next = next.add(dir);
            } else break;
        }

        if (thingsToPush.length === 0) {
            robotPos = robotPos.add(dir);
        } else if (thingsToPush[thingsToPush.length - 1] === "O") {
            for (let i = 0; i < thingsToPush.length; i++) {
                const pos = robotPos.add(dir.scale(1 + i));
                map.delete(`${pos.row},${pos.col}`);
            }
            for (let i = 0; i < thingsToPush.length; i++) {
                const pos = robotPos.add(dir.scale(2 + i));
                map.set(`${pos.row},${pos.col}`, thingsToPush[i]);
            }
            robotPos = robotPos.add(dir);
        }
    }

    let total = 0;
    for (const [key, value] of map) {
        if (value === "O") {
            const [row, col] = key.split(",").map(Number);
            total += 100 * row + col;
        }
    }
    return total;
}

function day15b() {
    const file = readDayFile(15);
    const { inputMap, moves } = parseInput(file);
    const map = new Map();
    let robotPos = RobotVector.Zero;

    inputMap.forEach((line, row) => {
        for (let col = 0; col < line.length; col++) {
            const tile = line[col];
            const pos = new RobotVector(row, col * 2);
            if (tile === "#" || tile === "O") {
                const right = pos.add(RobotVector.Right);
                const obstacle = new Obstacle(tile, pos, right);
                map.set(`${pos.row},${pos.col}`, obstacle);
                map.set(`${right.row},${right.col}`, obstacle);
            } else if (tile === "@") {
                robotPos = pos;
            }
        }
    });

    function getBoxesToPush(pos, dir) {
        const results = new Set();
        const obstacle = map.get(`${pos.row},${pos.col}`);
        if (obstacle) {
            results.add(obstacle);
            if (obstacle.tile === "O") {
                if (dir === RobotVector.Left) {
                    getBoxesToPush(obstacle.left.add(RobotVector.Left), dir).forEach(o => results.add(o));
                } else if (dir === RobotVector.Right) {
                    getBoxesToPush(obstacle.right.add(RobotVector.Right), dir).forEach(o => results.add(o));
                } else {
                    getBoxesToPush(obstacle.left.add(dir), dir).forEach(o => results.add(o));
                    getBoxesToPush(obstacle.right.add(dir), dir).forEach(o => results.add(o));
                }
            }
        }
        return results;
    }

    for (const move of moves) {
        const dir = { ">": RobotVector.Right, "v": RobotVector.Down, "<": RobotVector.Left, "^": RobotVector.Up }[move];
        const thingsToPush = getBoxesToPush(robotPos.add(dir), dir);

        if (thingsToPush.size === 0) {
            robotPos = robotPos.add(dir);
        } else if ([...thingsToPush].some(obstacle => obstacle.tile === "#")) {
            continue;
        } else {
            for (const obstacle of thingsToPush) {
                map.delete(`${obstacle.left.row},${obstacle.left.col}`);
                map.delete(`${obstacle.right.row},${obstacle.right.col}`);
            }
            for (const obstacle of thingsToPush) {
                const newObstacle = new Obstacle(obstacle.tile, obstacle.left.add(dir), obstacle.right.add(dir));
                map.set(`${newObstacle.left.row},${newObstacle.left.col}`, newObstacle);
                map.set(`${newObstacle.right.row},${newObstacle.right.col}`, newObstacle);
            }
            robotPos = robotPos.add(dir);
        }
    }

    const coordinates = new Set();
    for (const [key, obstacle] of map.entries()) {
        if (obstacle.tile === "O") {
            coordinates.add(obstacle.left);
        }
    }

    return [...coordinates].reduce((total, coord) => total + 100 * coord.row + coord.col, 0);
}

class Obstacle {
    constructor(tile, left, right) {
        this.tile = tile;
        this.left = left;
        this.right = right;
    }
}

class RobotVector {
    constructor(row, col) {
        this.row = row;
        this.col = col;
    }

    static Zero = new RobotVector(0, 0);
    static Up = new RobotVector(-1, 0);
    static Down = new RobotVector(1, 0);
    static Left = new RobotVector(0, -1);
    static Right = new RobotVector(0, 1);

    add(vector) {
        return new RobotVector(this.row + vector.row, this.col + vector.col);
    }

    scale(factor) {
        return new RobotVector(this.row * factor, this.col * factor);
    }
}

function parseInput(file) {
    const sections = file.split("\n\n");
    const inputMap = sections[0].split("\n");
    const moves = sections[1].replace(/\s/g, "").split("");
    return { inputMap, moves };
}
// #endregion

// #region Day16
function day16a() {
    const file = readDayFile(16).split("\n");
    let start = new State(
        new Position(file.length - 2, 1), 
        Direction.East
    );

    if (file[start.pos.row][start.pos.col] !== 'S') {
        start = new State(
            new Position(1, file[0].length - 2), 
            Direction.South
        );
    }

    const solver = solve(file, start);
    return solver.cheapest;
}

function day16b() {
    const file = readDayFile(16).split("\n");
    let start = new State(
        new Position(file.length - 2, 1), 
        Direction.East
    );

    if (file[start.pos.row][start.pos.col] !== 'S') {
        start = new State(
            new Position(1, file[0].length - 2), 
            Direction.South
        );
    }

    const solver = solve(file, start);

    const seen = new Set();
    const queue = [solver.end];
    const zero = null;

    while (queue.length > 0) {
        const v = queue.shift();
        if (v !== zero) {
            seen.add(v.pos.hashCode());
            for (const parent of solver.prov.get(v.hashCode()).parents) {
                queue.push(parent);
            }
        }
    }

    return seen.size;
}

class Direction {
    static East = null;
    static South = null;
    static West = null;
    static North = null;

    constructor(row, col) {
        this.row = row;
        this.col = col;
    }

    turnRight() {
        if (this === Direction.East) return Direction.South;
        if (this === Direction.South) return Direction.West;
        if (this === Direction.West) return Direction.North;
        return Direction.East;
    }

    turnLeft() {
        if (this === Direction.East) return Direction.North;
        if (this === Direction.North) return Direction.West;
        if (this === Direction.West) return Direction.South;
        return Direction.East;
    }

    equals(other) {
        if (!other) return false;
        return this.row === other.row && this.col === other.col;
    }

    hashCode() {
        return `${this.row},${this.col}`;
    }
}

Direction.East = new Direction(0, 1);
Direction.South = new Direction(1, 0);
Direction.West = new Direction(0, -1);
Direction.North = new Direction(-1, 0);

class Position {
    constructor(row, col) {
        this.row = row;
        this.col = col;
    }

    move(direction) {
        return new Position(this.row + direction.row, this.col + direction.col);
    }

    equals(other) {
        if (!other) return false;
        return this.row === other.row && this.col === other.col;
    }

    hashCode() {
        return `${this.row},${this.col}`;
    }
}

class State {
    constructor(pos, direction) {
        this.pos = pos;
        this.dir = direction;
    }

    possible() {
        return {
            'straight': new State(this.pos.move(this.dir), this.dir),
            'left': new State(this.pos, this.dir.turnLeft()),
            'right': new State(this.pos, this.dir.turnRight())
        };
    }

    equals(other) {
        if (!other) return false;
        return this.pos.equals(other.pos) && this.dir.equals(other.dir);
    }

    hashCode() {
        return `${this.pos.hashCode()},${this.dir.hashCode()}`;
    }
}

class Provenance {
    constructor(cost) {
        this.cost = cost;
        this.parents = [];
    }

    maybeAdd(parent, cost) {
        if (this.cost > cost) {
            this.cost = cost;
            this.parents = parent ? [parent] : [];
        } else if (this.cost === cost && parent) {
            this.parents.push(parent);
        }
    }
}

class Solver {
    constructor(grid) {
        this.grid = grid;
        this.pq = {};
        this.prov = new Map();
        this.visited = new Map();
        this.cheapest = 0;
        this.highest = 0;
        this.end = null;
    }

    add(v, prev, cost) {
        if (!this.prov.has(v.hashCode())) {
            this.prov.set(v.hashCode(), new Provenance(cost));
        }
        
        this.prov.get(v.hashCode()).maybeAdd(prev, cost);

        const existingCost = this.visited.get(v.hashCode());
        if (existingCost === undefined || cost < existingCost) {
            this.visited.set(v.hashCode(), cost);
            
            if (!this.pq[cost]) {
                this.pq[cost] = [];
            }
            
            this.pq[cost].push(v);
            this.highest = Math.max(this.highest, cost);
        }
    }

    pop(cost) {
        const v = this.pq[cost][0];
        this.pq[cost].shift();
        return v;
    }

    lookup(p) {
        return this.grid[p.row][p.col];
    }

    isEnd(p) {
        return this.lookup(p) === 'E';
    }

    isOpen(p) {
        return this.lookup(p) !== '#';
    }
}

function solve(grid, start) {
    const solver = new Solver(grid);
    solver.add(start, null, 0);

    while (true) {
        while (!solver.pq[solver.cheapest] || 
               solver.pq[solver.cheapest].length === 0) {
            if (solver.cheapest > solver.highest) {
                throw new Error("Ran out of priority queue");
            }
            solver.cheapest++;
        }

        const v = solver.pop(solver.cheapest);

        if (solver.isEnd(v.pos)) {
            solver.end = v;
            return solver;
        }

        const possible = v.possible();
        const straight = possible['straight'];
        const left = possible['left'];
        const right = possible['right'];

        if (solver.isOpen(straight.pos)) {
            solver.add(straight, v, solver.cheapest + 1);
        }
        if (solver.isOpen(left.pos)) {
            solver.add(left, v, solver.cheapest + 1000);
        }
        if (solver.isOpen(right.pos)) {
            solver.add(right, v, solver.cheapest + 1000);
        }
    }
}
// #endregion

main();