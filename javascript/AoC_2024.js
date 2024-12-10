const { parse } = require('path');

function main() {
    // Day 1 a + b
    console.log(day9a(), day9b());
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

main();