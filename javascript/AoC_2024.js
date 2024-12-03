function main() {
    // Day 1 a + b
    console.log(day2a(), day2b());
}

function readDayFile(day){
    const fs = require('fs');
    
    const filePath = `../AoC_Files/${day}.txt`
    const fileContents = fs.readFileSync(filePath, 'utf-8');

    return fileContents
}

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

main();