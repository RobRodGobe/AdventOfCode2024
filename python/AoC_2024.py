def main():
    # Day 1
    print(day1a(), day1b())

def readDayFile(day):
    file_path = f"../AoC_Files/{day}.txt"

    with open(file_path, "r") as file:
        file_contents = file.readlines()

    return file_contents

def day1a():
    file = readDayFile(1)
    list1 = []
    list2 = []

    for line in file:
        pairs = line.strip().split()
        list1.append(int(pairs[0]))
        list2.append(int(pairs[1]))

    list1.sort()
    list2.sort()

    diff = 0

    for i, line in enumerate(list1):
        diff += abs(list1[i] - list2[i])
    
    return diff      

def day1b():
    file = readDayFile(1)
    list1 = []
    list2 = []

    for line in file:
        pairs = line.strip().split()
        list1.append(int(pairs[0]))
        list2.append(int(pairs[1]))

    similar = 0

    for i, line in enumerate(list1):
        similar += list1[i] * list2.count(list1[i])

    return similar

if __name__ == "__main__":
    main()