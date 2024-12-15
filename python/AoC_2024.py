import re
import math
from typing import List, Dict, Optional, Tuple

def main():
    # Day 1
    print(day14a(), day14b())

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

# region Day11
def day11a():
    file = readDayFile(11)[0]
    stones = list(map(int, file.split()))
    rocks = {}
    for stone in stones:
        rocks[stone] = rocks.get(stone, 0) + 1

    final_rocks = blink_rocks(rocks, 25)
    return sum(final_rocks.values())

def day11b():
    file = readDayFile(11)[0]
    stones = list(map(int, file.split()))
    rocks = {}
    for stone in stones:
        rocks[stone] = rocks.get(stone, 0) + 1

    final_rocks = blink_rocks(rocks, 75)
    return sum(final_rocks.values())

def blink(rock: int) -> List[int]:
    if rock == 0:
        return [1]

    digits = math.floor(math.log10(rock)) + 1

    if digits % 2 != 0:
        return [rock * 2024]

    half_digits = digits // 2
    first = rock // (10 ** half_digits)
    second = rock % (10 ** half_digits)

    return [first, second]

def blink_rocks_iteration(rocks: Dict[int, int]) -> Dict[int, int]:
    result = {}

    for rock, count in rocks.items():
        new_rocks = blink(rock)

        for new_rock in new_rocks:
            result[new_rock] = result.get(new_rock, 0) + count

    return result

def blink_rocks(rocks: Dict[int, int], blinks: int) -> Dict[int, int]:
    current_rocks = rocks.copy()

    for _ in range(blinks):
        current_rocks = blink_rocks_iteration(current_rocks)

    return current_rocks
# endregion

# region Day12
def day12a():
    file = readDayFile(12)
    lines = [list(line.strip()) for line in file]

    return calculate_total_fencing_price(lines)

def day12b():
    file = readDayFile(12)
    lines = [list(line.strip()) for line in file]

    return calculate_total_fencing_price_with_inner_sides(lines)
   

def calculate_total_fencing_price(grid):
    n, m = len(grid), len(grid[0])
    visited = set()
    total_price = 0

    for i in range(n):
        for j in range(m):
            if (i, j) not in visited:
                area, borders = visit_region(grid, i, j, visited)
                total_price += area * len(borders)

    return total_price

def calculate_total_fencing_price_with_inner_sides(grid):
    n, m = len(grid), len(grid[0])
    visited = set()
    total_price = 0

    for i in range(n):
        for j in range(m):
            if (i, j) not in visited:
                area, borders = visit_region(grid, i, j, visited)
                total_price += area * count_sides(borders)

    return total_price

def visit_region(grid, start_i, start_j, visited):
    n, m = len(grid), len(grid[0])
    plant = grid[start_i][start_j]
    area = 0
    borders = set()

    def visit(i, j):
        nonlocal area
        if (i, j) in visited:
            return

        visited.add((i, j))
        area += 1

        dx = [-1, 1, 0, 0]
        dy = [0, 0, -1, 1]

        for k in range(4):
            i2 = i + dx[k]
            j2 = j + dy[k]

            if 0 <= i2 < n and 0 <= j2 < m and grid[i2][j2] == plant:
                visit(i2, j2)
            else:
                borders.add((i, j, i2, j2))

    visit(start_i, start_j)
    return area, borders

def count_sides(borders):
    visited = set()

    def visit_side(i, j, i2, j2):
        side = (i, j, i2, j2)
        if side in visited or side not in borders:
            return

        visited.add(side)

        if i == i2:
            visit_side(i - 1, j, i2 - 1, j2)
            visit_side(i + 1, j, i2 + 1, j2)
        else:
            visit_side(i, j - 1, i2, j2 - 1)
            visit_side(i, j + 1, i2, j2 + 1)

    num_sides = 0
    for side in borders:
        if side in visited:
            continue

        num_sides += 1
        visit_side(*side)

    return num_sides
# endregion

# region Day13
def day13a():
    file = readDayFile(13)

    return get_max_prize_for_min_tokens(file)

def day13b():
    file = readDayFile(13)
    machines = parse_claw_machine_input(file)
    for machine in machines:
        machine.prize_x += 10_000_000_000_000
        machine.prize_y += 10_000_000_000_000

    adjusted_input = []
    for machine in machines:
        adjusted_input.extend([
            f"Button A: X+{machine.ax}, Y+{machine.ay}",
            f"Button B: X+{machine.bx}, Y+{machine.by}",
            f"Prize: X={machine.prize_x}, Y={machine.prize_y}"
        ])

    return get_max_prize_for_min_tokens(adjusted_input)

class ClawMachineSettings:
    def __init__(self, ax, ay, bx, by, prize_x, prize_y):
        self.ax = ax
        self.ay = ay
        self.bx = bx
        self.by = by
        self.prize_x = prize_x
        self.prize_y = prize_y

def parse_claw_machine_input(input_data: List[str]) -> List[ClawMachineSettings]:
    machines = []
    cleaned_data = [line for line in input_data if line.strip() != ""]
    for i in range(0, len(cleaned_data), 3):
        a_move = cleaned_data[i].strip().replace("Button A: ", "").split(", ")
        b_move = cleaned_data[i + 1].strip().replace("Button B: ", "").split(", ")
        prize = cleaned_data[i + 2].strip().replace("Prize: ", "").split(", ")

        machines.append(ClawMachineSettings(
            ax=int(a_move[0].replace("X+", "")),
            ay=int(a_move[1].replace("Y+", "")),
            bx=int(b_move[0].replace("X+", "")),
            by=int(b_move[1].replace("Y+", "")),
            prize_x=int(prize[0].replace("X=", "")),
            prize_y=int(prize[1].replace("Y=", ""))
        ))

    return machines

def calculate_price(machine: ClawMachineSettings) -> Optional[int]:
    det = machine.ay * machine.bx - machine.ax * machine.by
    if det == 0:
        return None

    b = (machine.ay * machine.prize_x - machine.ax * machine.prize_y) // det
    a = (machine.prize_x - b * machine.bx) // machine.ax if machine.ax != 0 else 0

    if (machine.ax * a + machine.bx * b == machine.prize_x and
        machine.ay * a + machine.by * b == machine.prize_y and
        a >= 0 and b >= 0):
        return a * 3 + b
    return None

def get_max_prize_for_min_tokens(input_data: List[str]) -> int:
    machines = parse_claw_machine_input(input_data)
    total_tokens = 0

    for machine in machines:
        tokens = calculate_price(machine)
        if tokens is not None:
            total_tokens += tokens

    return total_tokens
# endregion

# region Day14
def day14a():
    file = readDayFile(14)
    
    return calculate_safety_factor(file)

def day14b():
    file = readDayFile(14)

    return find_robot_sequence_time(file)

class BathroomRobot:
    def __init__(self, P: Tuple[int, int], V: Tuple[int, int]):
        self.P = P
        self.V = V

    @staticmethod
    def simulate_robot(robot, mod_rows: int, mod_cols: int, ticks: int):
        row_delta = BathroomRobot.calculate_delta(robot.V[1], ticks, mod_rows)
        new_row = BathroomRobot.mod_add(robot.P[1], row_delta, mod_rows)

        col_delta = BathroomRobot.calculate_delta(robot.V[0], ticks, mod_cols)
        new_col = BathroomRobot.mod_add(robot.P[0], col_delta, mod_cols)

        return BathroomRobot((new_col, new_row), robot.V)

    @staticmethod
    def calculate_delta(velocity: int, ticks: int, mod: int) -> int:
        delta = velocity * ticks % mod
        return delta + mod if delta < 0 else delta

    @staticmethod
    def mod_add(a: int, b: int, mod: int) -> int:
        res = (a + b) % mod
        return res + mod if res < 0 else res
    
def calculate_safety_factor(file: List[str]) -> int:
    width, height = 101, 103
    duration = 100

    robots = parse_robots(file)
    final_positions = calculate_final_positions(robots, width, height, duration)

    return compute_quadrant_multiplier(final_positions, width, height)

def parse_robots(lines: List[str]) -> List[BathroomRobot]:
    return [parse_single_robot(line) for line in lines if line.strip()]

def parse_single_robot(line: str) -> BathroomRobot:
    parts = line.split()
    p = parts[0][2:].split(',')
    v = parts[1][2:].split(',')

    return BathroomRobot(
        (int(p[0]), int(p[1])), 
        (int(v[0]), int(v[1]))
    )

def calculate_final_positions(robots: List[BathroomRobot], width: int, height: int, duration: int):
    final_positions = [[0 for _ in range(height)] for _ in range(width)]

    for robot in robots:
        final_x = (robot.P[0] + robot.V[0] * duration) % width
        final_y = (robot.P[1] + robot.V[1] * duration) % height

        final_x = final_x + width if final_x < 0 else final_x
        final_y = final_y + height if final_y < 0 else final_y

        final_positions[final_x][final_y] += 1

    return final_positions

def compute_quadrant_multiplier(final_positions: List[List[int]], width: int, height: int) -> int:
    mid_x, mid_y = width // 2, height // 2

    top_left = top_right = bottom_left = bottom_right = 0

    for x in range(width):
        for y in range(height):
            if x == mid_x or y == mid_y:
                continue

            if x < mid_x and y < mid_y:
                top_left += final_positions[x][y]
            elif x >= mid_x and y < mid_y:
                top_right += final_positions[x][y]
            elif x < mid_x and y >= mid_y:
                bottom_left += final_positions[x][y]
            elif x >= mid_x and y >= mid_y:
                bottom_right += final_positions[x][y]

    return top_left * top_right * bottom_left * bottom_right

def find_robot_sequence_time(file: List[str]) -> int:
    rows, cols = 103, 101
    trunk_seq_size = 10
    max_seconds = 100000

    robots = parse_robots(file)

    for sec in range(1, max_seconds + 1):
        robots_by_col = [[] for _ in range(cols)]

        for i in range(len(robots)):
            new_robot = BathroomRobot.simulate_robot(robots[i], rows, cols, 1)
            robots_by_col[new_robot.P[0]].append(new_robot.P[1])
            robots[i] = new_robot

        for c in range(cols):
            column = robots_by_col[c]
            column.sort()

            if has_consecutive_sequence(column, trunk_seq_size):
                return sec

    return 0

def has_consecutive_sequence(sequence: List[int], required_length: int) -> bool:
    if len(sequence) < required_length:
        return False

    consecutive_count = 1
    for i in range(1, len(sequence)):
        consecutive_count = consecutive_count + 1 if sequence[i] == sequence[i-1] + 1 else 1

        if consecutive_count == required_length:
            return True

    return False
# endregion

if __name__ == "__main__":
    main()