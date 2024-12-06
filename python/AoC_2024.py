import re

def main():
    # Day 1
    print(day5a(), day5b())

def readDayFile(day):
    file_path = f"../AoC_Files/{day}.txt"

    with open(file_path, "r") as file:
        file_contents = file.readlines()

    return file_contents

# region Day1
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

# endregion

# region Day2
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

# endregion

# region Day3
def day3a():
    mult = 0
    file = readDayFile(3)
    line = "".join(file)
    pattern = r"mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)"
    matches = re.findall(pattern, line)

    for match in matches:
        numbers = match.replace("mul(", "").replace(")", "").split(",")
        mult += int(numbers[0]) * int (numbers[1])

    return mult

def day3b():
    mult = 0
    file = readDayFile(3)
    line = "".join(file)
    pattern = r"mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)|do\(\)|don't\(\)"
    matches = re.findall(pattern, line)

    multiply = True

    for match in matches:
        if match == "do()":
            multiply = True
        elif match == "don't()":
            multiply = False
        
        if multiply and match.startswith("mul("):
            numbers = match.replace("mul(", "").replace(")", "").split(",")
            mult += int(numbers[0]) * int (numbers[1])

    return mult
# endregion

# region Day4
def day4a():
    file = readDayFile(4)
    word = "XMAS"
    rows = len(file)
    cols = len(file[0])
    count = 0
    word_length = len(word)

    directions = [
        (0, 1),   # Right
        (1, 0),   # Down
        (1, 1),   # Down-right
        (1, -1),  # Down-left
        (0, -1),  # Left
        (-1, 0),  # Up
        (-1, -1), # Up-left
        (-1, 1)   # Up-right
    ]

    for x in range(rows):
        for y in range(len(file[x])):
            for dx, dy in directions:
                if check_word_begin(x, y, dx, dy, word_length, rows, cols, word, file):
                    count += 1

    return count
    
def day4b():
    file = readDayFile(4)
    rows = len(file)
    count = 0

    for x in range(1, rows - 1):
        cols = len(file[x])
        for y in range(1, cols - 1):
            if is_xmas_pattern(file, x, y):
                count += 1

    return count

def check_word_begin(x, y, dx, dy, length, rows, cols, word, grid):
    for i in range(length):
        nx = x + i * dx
        ny = y + i * dy

        if nx < 0 or ny < 0 or nx >= rows or ny >= cols or grid[nx][ny] != word[i]:
            return False
    
    return True

def is_xmas_pattern(grid, x, y):
    current_row_len = len(grid[x])
    prev_row_len = len(grid[x - 1])
    next_row_len = len(grid[x + 1])

    if y - 1 < 0 or y + 1 >= current_row_len:
        return False
    if y - 1 >= prev_row_len or y + 1 >= next_row_len:
        return False

    top_left_to_bottom_right = grid[x - 1][y - 1] + grid[x][y] + grid[x + 1][y + 1]
    top_right_to_bottom_left = grid[x - 1][y + 1] + grid[x][y] + grid[x + 1][y - 1]

    return is_valid_mas_pattern(top_left_to_bottom_right) and is_valid_mas_pattern(top_right_to_bottom_left)

def is_valid_mas_pattern(pattern):
    return pattern == "MAS" or pattern == "SAM"

# endregion

# region Day5
def day5a():
    file = [line.strip() for line in readDayFile(5)]
    divider_index = file.index('')
    rules = [tuple(map(int, line.split('|'))) for line in file[:divider_index]]
    updates = [list(map(int, line.split(','))) for line in file[divider_index + 1:]]

    pages = 0
    for update in updates:
        if is_update_valid(update, rules):
            pages += get_middle_page(update)

    return pages

def day5b():
    file = [line.strip() for line in readDayFile(5)]
    divider_index = file.index('')
    rules = [tuple(map(int, line.split('|'))) for line in file[:divider_index]]
    updates = [list(map(int, line.split(','))) for line in file[divider_index + 1:]]

    pages = 0
    for update in updates:
        if not is_update_valid(update, rules):
            corrected_update = correct_update(update, rules)
            pages += get_middle_page(corrected_update)

    return pages

def is_update_valid(update, rules):
    page_positions = {page: idx for idx, page in enumerate(update)}
    for before, after in rules:
        if before in page_positions and after in page_positions:
            if page_positions[before] >= page_positions[after]:
                return False
    return True

def get_middle_page(update):
    return update[len(update) // 2]

def correct_update(update, rules):
    graph = {page: [] for page in update}
    in_degree = {page: 0 for page in update}

    for before, after in rules:
        if before in update and after in update:
            graph[before].append(after)
            in_degree[after] += 1

    queue = [page for page, degree in in_degree.items() if degree == 0]
    sorted_update = []

    while queue:
        current = queue.pop(0)
        sorted_update.append(current)

        for neighbor in graph[current]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    return sorted_update
# endregion

if __name__ == "__main__":
    main()