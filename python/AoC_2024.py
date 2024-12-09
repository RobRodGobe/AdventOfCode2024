import re

def main():
    # Day 1
    print(day10a(), day10b())

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

# region Day6
def day6a():
    file = readDayFile(6)
    grid = [list(line.strip()) for line in file]
    rows, cols = len(grid), len(grid[0])

    directions = {"^": (-1, 0), ">": (0, 1), "v": (1, 0), "<": (0, -1)}
    turn_right = {"^": ">", ">": "v", "v": "<", "<": "^"}

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] in directions:
                guard_pos = (r, c)
                guard_dir = grid[r][c]
                break

    visited = set()
    visited.add(guard_pos)

    while True:
        dy, dx = directions[guard_dir]
        next_pos = (guard_pos[0] + dy, guard_pos[1] + dx)

        if not (0 <= next_pos[0] < rows and 0 <= next_pos[1] < cols):
            break

        if grid[next_pos[0]][next_pos[1]] == "#":
            guard_dir = turn_right[guard_dir]
        else:
            guard_pos = next_pos
            visited.add(guard_pos)

    return len(visited)

def day6b():
    file = readDayFile(6)
    grid = [list(line.strip()) for line in file]
    rows, cols = len(grid), len(grid[0])

    directions = {"^": (-1, 0), ">": (0, 1), "v": (1, 0), "<": (0, -1)}

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] in directions:
                guard_pos = (r, c)
                guard_dir = grid[r][c]
                break

    loop_positions = 0

    for r in range(rows):
        for c in range(cols):
            if is_guard_in_loop(grid, guard_pos, guard_dir, (r, c)):
                loop_positions += 1

    return loop_positions

def is_guard_in_loop(map_lines, guard_start, guard_dir, obstruction):
    directions = {'^': (-1, 0), '>': (0, 1), 'v': (1, 0), '<': (0, -1)}
    turn_right = {'^': '>', '>': 'v', 'v': '<', '<': '^'}

    rows, cols = len(map_lines), len(map_lines[0])
    guard_pos = guard_start
    current_dir = guard_dir

    temp_map = [list(row) for row in map_lines]
    temp_map[obstruction[0]][obstruction[1]] = '#'

    visited_states = set()
    recent_history = []
    max_history_length = 10

    steps = 0
    max_steps = rows * cols * 2

    while True:
        state = (guard_pos, current_dir)
        if state in visited_states:
            if state in recent_history:
                return True

        visited_states.add(state)
        recent_history.append(state)
        if len(recent_history) > max_history_length:
            recent_history.pop(0)

        dx, dy = directions[current_dir]
        next_pos = (guard_pos[0] + dx, guard_pos[1] + dy)
        
        if (next_pos[0] < 0 or next_pos[0] >= rows or next_pos[1] < 0 or next_pos[1] >= cols):
            return False
        elif temp_map[next_pos[0]][next_pos[1]] == '#':
            current_dir = turn_right[current_dir]
        else:
            guard_pos = next_pos

        steps += 1
        if steps > max_steps:
            return True
# endregion

# region Day7
def day7a():
    file = readDayFile(7)
    sum = 0

    for line in file:
        nums = line.split(":")
        total = nums[0]
        factors = list(map(int, nums[1].strip().split()))
        if can_calibrate(total, factors, factors[0], 1):
            sum += int(total.strip())

    return sum

def day7b():
    file = readDayFile(7)
    sum = 0

    for line in file:
        nums = line.split(":")
        total = nums[0]
        factors = list(map(int, nums[1].strip().split()))
        if can_calibrate_2(total, factors, factors[0], 1):
            sum += int(total.strip())

    return sum

def can_calibrate(target, numbers, current, i):
    if i == len(numbers):
        return int(current) == int(target)
    
    if can_calibrate(target, numbers, current + numbers[i], i + 1):
        return True
    
    if can_calibrate(target, numbers, current * numbers[i], i + 1):
        return True
    
    return False

def can_calibrate_2(target, numbers, current, i):
    if i == len(numbers):
        return int(current) == int(target)
    
    if can_calibrate_2(target, numbers, current + numbers[i], i + 1):
        return True
    
    if can_calibrate_2(target, numbers, current * numbers[i], i + 1):
        return True
    
    if can_calibrate_2(target, numbers, int(f"{current}{numbers[i]}"), i + 1):
        return True
    
    return False
# endregion

# region Day8
def day8a():
    file = readDayFile(8)
    matrix = [list(line.strip()) for line in file]

    antenna_map = get_antenna_map(matrix)

    all_antinodes = []

    for coords in antenna_map.values():
        antinodes = get_antinodes(coords, matrix)
        all_antinodes.extend(antinodes)

    unique_antinodes = get_unique_antinodes(all_antinodes)

    return len(unique_antinodes)

def day8b():
    file = readDayFile(8)
    matrix = [list(line.strip()) for line in file]

    antenna_map = get_antenna_map(matrix)
    antinode_matrix = [[False] * len(matrix[0]) for _ in range(len(matrix))]

    for coords in antenna_map.values():
        process_antinode_lines(coords, matrix, antinode_matrix)

    return get_unique_antinodes_count(antinode_matrix)

def get_antenna_map(matrix):
    antenna_map = {}
    for i, row in enumerate(matrix):
        for j, cell in enumerate(row):
            if cell != '.':
                if cell not in antenna_map:
                    antenna_map[cell] = []
                antenna_map[cell].append((i, j))
    return antenna_map

def get_antinodes(coords, matrix):
    antinodes = []

    for i, (ax, ay) in enumerate(coords):
        for j, (bx, by) in enumerate(coords):
            if i != j:
                cx, cy = 2 * bx - ax, 2 * by - ay

                if within_boundaries(cx, 0, len(matrix)) and within_boundaries(cy, 0, len(matrix[0])):
                    antinodes.append((cx, cy))

    return antinodes

def within_boundaries(value, min_value, max_value):
    return min_value <= value < max_value

def get_unique_antinodes(antinodes):
    return list(set(antinodes))

def process_antinode_lines(coords, matrix, antinode_matrix):
    for i, (x1, y1) in enumerate(coords):
        for j, (x2, y2) in enumerate(coords):
            if i != j:
                for x in range(len(matrix)):
                    for y in range(len(matrix[0])):
                        if not antinode_matrix[x][y]:
                            line_result = (y1 - y2) * x + (x2 - x1) * y + (x1 * y2 - x2 * y1)
                            if line_result == 0:
                                antinode_matrix[x][y] = True

def get_unique_antinodes_count(antinode_matrix):
    count = 0
    for row in antinode_matrix:
        count += sum(row)
    return count

# endregion

# region Day9
def day9a():
    file = readDayFile(9)
    line = file[0].strip()
    diskMap = parse_disk_map(line)
    diskMap = compact_disk(diskMap)
    return calculate_checksum(diskMap)

def day9b():
    file = readDayFile(9)
    line = file[0].strip()
    diskMap = parse_disk_map(line)
    diskMap = compact_disk_by_file(diskMap)
    return calculate_checksum(diskMap)

def parse_disk_map(line):
    nums = []
    index = 0

    for i in range(len(line)):
        count = int(line[i])
        if i % 2 == 0:
            for _ in range(count):
                nums.append(str(index))
            index += 1
        else:
            for _ in range(count):
                nums.append(".")
    
    return nums

def compact_disk(diskMap):
    L = 0
    R = len(diskMap) - 1

    while L <= R:
        if diskMap[L] == "." and diskMap[R] != ".":
            diskMap[L], diskMap[R] = diskMap[R], diskMap[L]
            R -= 1
            L += 1
        elif diskMap[R] == ".":
            R -= 1
        else:
            L += 1

    return diskMap

def calculate_checksum(diskMap):
    return sum(i * int(block) for i, block in enumerate(diskMap) if block != ".")

class DiskFile:
    def __init__(self, id, length, startIdx):
        self.id = id
        self.length = length
        self.startIdx = startIdx

def analyze_disk(diskMap):
    files = []
    spaces = {}
    spaceStartIdx = -1

    for i, item in enumerate(diskMap):
        if item == ".":
            if spaceStartIdx == -1:
                spaceStartIdx = i
            spaces[spaceStartIdx] = spaces.get(spaceStartIdx, 0) + 1
        else:
            if spaceStartIdx != -1:
                spaceStartIdx = -1

            fileId = int(item)
            while len(files) <= fileId:
                files.append(DiskFile(len(files), 0, i))

            files[fileId].length += 1

    return files, spaces

def get_first_available_space_idx(spaces, fileLength):
    sorted_indices = sorted(spaces.keys())
    for idx in sorted_indices:
        if spaces[idx] >= fileLength:
            return idx
    return -1

def move_file(diskMap, file, targetIdx):
    file_segments = diskMap[file.startIdx:file.startIdx + file.length]
    
    for i in range(file.length):
        diskMap[file.startIdx + i] = "."
    
    for i, segment in enumerate(file_segments):
        diskMap[targetIdx + i] = segment

def update_spaces(spaces, spaceIdx, fileLength):
    if spaces[spaceIdx] == fileLength:
        del spaces[spaceIdx]
    else:
        remainingSpace = spaces[spaceIdx] - fileLength
        del spaces[spaceIdx]
        spaces[spaceIdx + fileLength] = remainingSpace

def compact_disk_by_file(diskMap):
    files, spaces = analyze_disk(diskMap)
     
    files.sort(key=lambda x: x.id, reverse=True)
    
    for file in files:
        targetIdx = get_first_available_space_idx(spaces, file.length)
        
        if targetIdx != -1 and targetIdx < file.startIdx:
            move_file(diskMap, file, targetIdx)
            update_spaces(spaces, targetIdx, file.length)
            
    return diskMap

# endregion

# region Day10
def day10a():
    file = readDayFile(10)
    trail = [list(line.strip()) for line in file]

    return solve_topographic_map(trail)

def day10b():
    file = readDayFile(10)
    trail = [list(line.strip()) for line in file]

    return solve_topographic_map_trail_ratings(trail)

def solve_topographic_map(input_data):
    rows = len(input_data)
    cols = len(input_data[0])
    map_grid = [[int(c) for c in row] for row in input_data]

    total_trailhead_score = 0

    for r in range(rows):
        for c in range(cols):
            if map_grid[r][c] == 0:
                trailhead_score = find_trailhead_score(r, c, map_grid, rows, cols)
                total_trailhead_score += trailhead_score

    return total_trailhead_score

def find_trailhead_score(start_row, start_col, map_grid, rows, cols):
    visited = [[False] * cols for _ in range(rows)]
    nine_positions = set()

    def dfs(row, col, expected_height):
        if (row < 0 or row >= rows or col < 0 or col >= cols or 
            visited[row][col] or map_grid[row][col] != expected_height):
            return False

        visited[row][col] = True

        if expected_height == 9:
            nine_positions.add((row, col))

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        return any(
            dfs(row + dr, col + dc, expected_height + 1)
            for dr, dc in directions
        )

    dfs(start_row, start_col, 0)
    return len(nine_positions)

def solve_topographic_map_trail_ratings(input_data):
    map_grid, rows, cols = parse_topographic_map(input_data)
    scores_map = {}

    for r in range(rows):
        for c in range(cols):
            if map_grid[r][c] == 0:
                start_hike(r, c, map_grid, rows, cols, scores_map)

    return get_scores_sum(scores_map)

def start_hike(start_r, start_c, map_grid, rows, cols, scores_map):
    routes = [(start_r, start_c, 0, f"{start_r}:{start_c}")]

    while routes:
        route = routes.pop(0)
        r, c, height, initial_cell = route

        for dr, dc in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            new_r = r + dr
            new_c = c + dc
            new_height = height + 1

            if (new_r < 0 or new_r >= rows or 
                new_c < 0 or new_c >= cols):
                continue

            new_cell = map_grid[new_r][new_c]
            
            if new_cell != new_height:
                continue

            if new_cell == 9:
                if initial_cell not in scores_map:
                    scores_map[initial_cell] = {}
                
                end_key = f"{new_r}:{new_c}"
                scores_map[initial_cell][end_key] = scores_map[initial_cell].get(end_key, 0) + 1
            else:
                routes.append((new_r, new_c, new_height, initial_cell))

def get_scores_sum(scores_map):
    return sum(sum(scores.values()) for scores in scores_map.values())

def parse_topographic_map(input_data):
    rows = len(input_data)
    cols = len(input_data[0])
    map_grid = [[int(c) for c in row] for row in input_data]

    return map_grid, rows, cols
# endregion

if __name__ == "__main__":
    main()