def main():
    # Day 1
    print(day2a(), day2b())

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

def day2a():
    file = readDayFile(2)
    safe = 0

    for line in file:
        reports = list(map(int, line.split()))
        is_ascending = True
        is_descending = True
        is_safe = True

        for j in range(1, len(reports)):
            diff = reports[j] - reports[j - 1]

            if abs(diff) > 3:
                is_ascending = False
                is_descending = False
                is_safe = False
                break

            if diff < 0:
                is_ascending = False
            if diff > 0:
                is_descending = False
            if diff == 0:
                is_ascending = False
                is_descending = False

            if not is_ascending and not is_descending:
                is_safe = False
                break

        if is_safe:
            safe += 1

    return safe

def day2b():
    file = readDayFile(2)
    safe = 0

    for line in file:
        reports = list(map(int, line.split()))

        if is_safe_report(reports, ascending=True) or is_safe_report(reports, ascending=False):
            safe += 1
            continue

        found_safe = False
        for i in range(len(reports)):
            modified_reports = reports[:i] + reports[i+1:]
            if is_safe_report(modified_reports, ascending=True) or is_safe_report(modified_reports, ascending=False):
                found_safe = True
                break
        
        if found_safe:
            safe += 1

    return safe

def is_safe_report(reports, ascending):
    for i in range(1, len(reports)):
        diff = reports[i] - reports[i - 1]
        if ascending and diff < 0:
            return False 
        if not ascending and diff > 0:
            return False 
        if abs(diff) > 3 or diff == 0:
            return False 
    return True

if __name__ == "__main__":
    main()