const { parse } = require('path');

function main() {
    // Day 1 a + b
    console.log(day4a(), day4b());
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

main();