function main() {
    // Day 1 a + b
    console.log(day1a(), day1b());
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

main();