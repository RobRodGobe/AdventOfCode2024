import re
import math
from typing import List, Dict, Optional, Tuple, Set
from enum import Enum
import heapq
from collections import defaultdict, deque

def main():
    # Day 1
    print(day24a(), day24b())
    
    # Day 25
    print(day25())

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

# region Day15
def day15a():
    file = readDayFile(15)
    sections = "".join(file).split("\n\n")
    input_map = sections[0].splitlines()
    moves = [c for c in sections[1] if not c.isspace()]

    map = {}
    robot_pos = RobotVector.Zero

    for row, line in enumerate(input_map):
        for col, tile in enumerate(line):
            pos = (row, col)
            if tile in "#O":
                map[pos] = tile
            elif tile == "@":
                robot_pos = pos

    for move in moves:
        dir = {">": RobotVector.Right, "v": RobotVector.Down, "<": RobotVector.Left, "^": RobotVector.Up}[move]
        things_to_push = []
        next_pos = add_vectors(robot_pos, dir)

        while next_pos in map:
            tile = map[next_pos]
            things_to_push.append(tile)
            if tile == "#":
                break
            next_pos = add_vectors(next_pos, dir)

        if not things_to_push:
            robot_pos = add_vectors(robot_pos, dir)
        elif things_to_push[-1] == "O":
            for i in range(len(things_to_push)):
                map.pop(add_vectors(robot_pos, scale_vector(dir, 1 + i)), None)
            for i, tile in enumerate(things_to_push):
                map[add_vectors(robot_pos, scale_vector(dir, 2 + i))] = tile
            robot_pos = add_vectors(robot_pos, dir)

    total = sum(100 * pos[0] + pos[1] for pos, tile in map.items() if tile == "O")
    return total

def get_boxes_to_push(map, pos, dir):
    if pos in map:
        obstacle = map[pos]
        results = {obstacle}
        if obstacle.tile == "O":
            if dir == RobotVector.Left:
                results.update(get_boxes_to_push(map, add_vectors(obstacle.left, RobotVector.Left), dir))
            elif dir == RobotVector.Right:
                results.update(get_boxes_to_push(map, add_vectors(obstacle.right, RobotVector.Right), dir))
            else:
                results.update(get_boxes_to_push(map, add_vectors(obstacle.left, dir), dir))
                results.update(get_boxes_to_push(map, add_vectors(obstacle.right, dir), dir))
        return results
    return set()

def day15b():
    file = readDayFile(15)
    sections = "".join(file).split("\n\n")
    input_map = sections[0].splitlines()
    moves = [c for c in sections[1] if not c.isspace()]

    map = {}
    robot_pos = RobotVector.Zero

    for row, line in enumerate(input_map):
        for col, tile in enumerate(line):
            pos = (row, col * 2)
            if tile in "#O":
                right = add_vectors(pos, RobotVector.Right)
                obstacle = Obstacle(tile, pos, right)
                map[pos] = obstacle
                map[right] = obstacle
            elif tile == "@":
                robot_pos = pos

    for move in moves:
        dir = {">": RobotVector.Right, "v": RobotVector.Down, "<": RobotVector.Left, "^": RobotVector.Up}[move]
        things_to_push = get_boxes_to_push(map, add_vectors(robot_pos, dir), dir)

        if not things_to_push:
            robot_pos = add_vectors(robot_pos, dir)
        elif any(obstacle.tile == "#" for obstacle in things_to_push):
            continue
        else:
            for obstacle in things_to_push:
                del map[obstacle.left]
                del map[obstacle.right]
            for obstacle in things_to_push:
                new_obstacle = Obstacle(obstacle.tile, add_vectors(obstacle.left, dir), add_vectors(obstacle.right, dir))
                map[new_obstacle.left] = new_obstacle
                map[new_obstacle.right] = new_obstacle
            robot_pos = add_vectors(robot_pos, dir)

    coordinates = {obstacle.left for obstacle in map.values() if obstacle.tile == "O"}
    return sum(100 * coord[0] + coord[1] for coord in coordinates)

class Obstacle:
    def __init__(self, tile, left, right):
        self.tile = tile
        self.left = left
        self.right = right

class RobotVector:
    Zero = (0, 0)
    Up = (-1, 0)
    Down = (1, 0)
    Left = (0, -1)
    Right = (0, 1)

def add_vectors(a, b):
    return (a[0] + b[0], a[1] + b[1])

def scale_vector(vec, factor):
    return (vec[0] * factor, vec[1] * factor)

# endregion

# region Day16
def day16a():
    file = readDayFile(16)
    start = State(
        Position(len(file) - 2, 1), 
        Direction.East
    )

    if file[start.pos.row][start.pos.col] != 'S':
        start = State(
            Position(1, len(file[0]) - 2), 
            Direction.South
        )

    solver = solve(file, start)
    return solver.cheapest


def day16b():
    file = readDayFile(16)
    start = State(
        Position(len(file) - 2, 1), 
        Direction.East
    )

    if file[start.pos.row][start.pos.col] != 'S':
        start = State(
            Position(1, len(file[0]) - 2), 
            Direction.South
        )

    solver = solve(file, start)

    seen = set()
    queue = [solver.end]
    zero = None

    while queue:
        v = queue.pop(0)
        if v != zero:
            seen.add(v.pos)
            for parent in solver.prov[v].parents:
                queue.append(parent)

    return len(seen)

class Direction:
    East = None
    South = None
    West = None
    North = None

    def __init__(self, row: int, col: int):
        self.row = row
        self.col = col

    def turn_right(self):
        if self == Direction.East:
            return Direction.South
        if self == Direction.South:
            return Direction.West
        if self == Direction.West:
            return Direction.North
        return Direction.East

    def turn_left(self):
        if self == Direction.East:
            return Direction.North
        if self == Direction.North:
            return Direction.West
        if self == Direction.West:
            return Direction.South
        return Direction.East

Direction.East = Direction(0, 1)
Direction.South = Direction(1, 0)
Direction.West = Direction(0, -1)
Direction.North = Direction(-1, 0)

class Position:
    def __init__(self, row: int, col: int):
        self.row = row
        self.col = col

    def move(self, direction: Direction):
        return Position(self.row + direction.row, self.col + direction.col)

    def __eq__(self, other):
        if other is None:
            return False
        return self.row == other.row and self.col == other.col

    def __hash__(self):
        return hash((self.row, self.col))

class State:
    def __init__(self, pos: Position, direction: Direction):
        self.pos = pos
        self.dir = direction

    def possible(self):
        return {
            'straight': State(self.pos.move(self.dir), self.dir),
            'left': State(self.pos, self.dir.turn_left()),
            'right': State(self.pos, self.dir.turn_right())
        }

    def __eq__(self, other):
        if other is None:
            return False
        return self.pos == other.pos and self.dir == other.dir

    def __hash__(self):
        return hash((self.pos, self.dir))

class Provenance:
    def __init__(self, cost: int):
        self.cost = cost
        self.parents: List[State] = []

    def maybe_add(self, parent: Optional[State], cost: int):
        if self.cost > cost:
            self.cost = cost
            self.parents = [parent] if parent else []
        elif self.cost == cost and parent:
            self.parents.append(parent)

class Solver:
    def __init__(self, grid: List[str]):
        self.grid = grid
        self.pq: Dict[int, List[State]] = {}
        self.prov: Dict[State, Provenance] = {}
        self.visited: Dict[State, int] = {}
        self.cheapest = 0
        self.highest = 0
        self.end = None

    def add(self, v: State, prev: Optional[State], cost: int):
        if v not in self.prov:
            self.prov[v] = Provenance(cost)
        
        self.prov[v].maybe_add(prev, cost)

        existing_cost = self.visited.get(v)
        if existing_cost is None or cost < existing_cost:
            self.visited[v] = cost
            
            if cost not in self.pq:
                self.pq[cost] = []
            
            self.pq[cost].append(v)
            self.highest = max(self.highest, cost)

    def pop(self, cost: int):
        v = self.pq[cost][0]
        self.pq[cost].pop(0)
        return v

    def lookup(self, p: Position):
        return self.grid[p.row][p.col]

    def is_end(self, p: Position):
        return self.lookup(p) == 'E'

    def is_open(self, p: Position):
        return self.lookup(p) != '#'

def solve(grid: List[str], start: State) -> Solver:
    solver = Solver(grid)
    solver.add(start, None, 0)

    while True:
        while (solver.cheapest not in solver.pq or 
               len(solver.pq[solver.cheapest]) == 0):
            if solver.cheapest > solver.highest:
                raise Exception("Ran out of priority queue")
            solver.cheapest += 1

        v = solver.pop(solver.cheapest)

        if solver.is_end(v.pos):
            solver.end = v
            return solver

        possible = v.possible()
        straight, left, right = (
            possible['straight'], 
            possible['left'], 
            possible['right']
        )

        if solver.is_open(straight.pos):
            solver.add(straight, v, solver.cheapest + 1)
        if solver.is_open(left.pos):
            solver.add(left, v, solver.cheapest + 1000)
        if solver.is_open(right.pos):
            solver.add(right, v, solver.cheapest + 1000)
# endregion

# region Day17
def day17a():
    file = readDayFile(17)
    comp = init_computer(file)
    output = run(comp.A, comp.B, comp.C, comp.Program)

    return ",".join(map(str, output))

def day17b():
    file = readDayFile(17)
    comp = init_computer(file)

    queue = [{"a": 0, "n": 1}]
    while queue:
        item = queue.pop(0)
        a = item["a"]
        n = item["n"]

        if n > len(comp.Program):
            return str(a)

        for i in range(8):
            a2 = (a << 3) | i
            output = run(a2, 0, 0, comp.Program)
            target = comp.Program[-n:]

            if matches_program(output, target):
                queue.append({"a": a2, "n": n + 1})

    return 0

class SmallComputer:
    def __init__(self):
        self.A: int = 0
        self.B: int = 0
        self.C: int = 0
        self.Program: List[int] = []
        self.Out: List[int] = []

def init_computer(puzzle: List[str]) -> SmallComputer:
    res = SmallComputer()
    
    for line in puzzle:
        if "Program" in line:
            # Use regex to extract digits
            digits = [int(m) for m in re.findall(r'\d', line)]
            res.Program = digits
        elif "Register" in line:
            # Use regex to parse register and value
            match = re.search(r'Register ([A-C]): (\d+)', line)
            if match:
                register, value = match.groups()
                value = int(value)
                
                if register == "A":
                    res.A = value
                elif register == "B":
                    res.B = value
                elif register == "C":
                    res.C = value
    
    return res

def run(a: int, b: int, c: int, program: List[int]) -> List[int]:
    out = []
    pointer = 0

    while pointer < len(program):
        instruction = program[pointer]
        param = program[pointer + 1]

        # Simulate uint64 behavior with explicit masking
        a &= 0xFFFFFFFFFFFFFFFF
        b &= 0xFFFFFFFFFFFFFFFF
        c &= 0xFFFFFFFFFFFFFFFF

        combo = param
        if param == 4:
            combo = a
        elif param == 5:
            combo = b
        elif param == 6:
            combo = c

        if instruction == 0:
            a >>= combo
        elif instruction == 1:
            b ^= param
        elif instruction == 2:
            b = combo % 8
        elif instruction == 3:
            if a != 0:
                pointer = param - 2
        elif instruction == 4:
            b ^= c
        elif instruction == 5:
            out.append(combo % 8)
        elif instruction == 6:
            b = a >> combo
        elif instruction == 7:
            c = a >> combo

        # Mask to uint64
        a &= 0xFFFFFFFFFFFFFFFF
        b &= 0xFFFFFFFFFFFFFFFF
        c &= 0xFFFFFFFFFFFFFFFF
        
        pointer += 2

    return out

def matches_program(output: List[int], expected: List[int]) -> bool:
    return output == expected
# endregion

# region Day18
def day18a():
    file = readDayFile(18)

    coordinates = AllBlockedCoords(file)
    start = Coord(0, 0)
    end = Coord(70, 70)

    shortestPath, _ = ShortestPath(coordinates[:1024], start, end)

    return shortestPath

def day18b():
    file = readDayFile(18)

    coordinates = AllBlockedCoords(file)
    start = Coord(0, 0)
    end = Coord(70, 70)

    block = SearchForBlockage(coordinates, start, end)

    return f"{block.X}, {block.Y}"

class Coord:
    def __init__(self, x: int, y: int):
        self.X = x
        self.Y = y

    def __hash__(self):
        return hash((self.X, self.Y))

    def __eq__(self, other):
        return self.X == other.X and self.Y == other.Y

def blockedCoords(coords: List[Coord]) -> Dict[Coord, None]:
    return {c: None for c in coords}

def AllBlockedCoords(coordStrs: List[str]) -> List[Coord]:
    allC = []

    for c in coordStrs:
        parts = c.split(",")
        x = int(parts[0])
        y = int(parts[1])
        allC.append(Coord(x, y))

    return allC

def isValidStep(c: Coord, b: Dict[Coord, None]) -> bool:
    return c not in b and 0 <= c.X <= 70 and 0 <= c.Y <= 70

def nextCoords(c: Coord) -> List[Coord]:
    return [
        Coord(c.X + 1, c.Y), 
        Coord(c.X - 1, c.Y), 
        Coord(c.X, c.Y + 1), 
        Coord(c.X, c.Y - 1)
    ]

def ShortestPath(coords: List[Coord], start: Coord, end: Coord) -> Tuple[int, bool]:
    b = blockedCoords(coords)
    visited = {start: 0}
    q = [start]

    while q:
        node = q.pop(0)

        for c in nextCoords(node):
            if c not in visited and isValidStep(c, b):
                q.append(c)
                visited[c] = visited[node] + 1

    return visited.get(end, 0), end in visited

def SearchForBlockage(allC: List[Coord], start: Coord, end: Coord) -> Coord:
    l, r = 1024, len(allC) - 1
    m = (l + r) // 2

    while l != m and r != m:
        _, ok = ShortestPath(allC[:m], start, end)

        if ok:
            l, r, m = m, r, (m + r) // 2
        else:
            l, r, m = l, m, (l + m) // 2

    return allC[m]
# endregion

# region Day19
def day19a():
    file = readDayFile(19)
    sections = "".join(file)
    towels, patterns = parseDesignFile(sections)
    count = 0
    cache = {}
    for pattern in patterns:
        if designPossible(pattern, towels, cache):
            count += 1
    return count

def day19b():
    file = readDayFile(19)
    sections = "".join(file)
    towels, patterns = parseDesignFile(sections)
    count = 0
    cache = {}
    for pattern in patterns:
        count += waysPossible(pattern, towels, cache)
    return count

def parseDesignFile(s):
    parts = s.split("\n\n")
    t = parts[0].split(", ")
    p = parts[1].split("\n")
    return t, p

def designPossible(pattern, ts, cache):
    if pattern in cache:
        return cache[pattern]

    for t in ts:
        if t == pattern:
            return True
        elif pattern.startswith(t):
            isPoss = designPossible(pattern[len(t):], ts, cache)
            if isPoss:
                cache[pattern] = True
                return True
    
    cache[pattern] = False
    return False

def waysPossible(pattern, ts, cache):
    if pattern in cache:
        return cache[pattern]

    ways = 0
    for t in ts:
        if t == pattern:
            ways += 1
        elif pattern.startswith(t):
            ways += waysPossible(pattern[len(t):], ts, cache)
    
    cache[pattern] = ways
    return ways
# endregion

# region Day20
def day20a():
    file = readDayFile(20)
    return getCheats(file, 2)

def day20b():
    file = readDayFile(20)
    return getCheats(file, 20)

class Point:
    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y
    
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y
    
    def __hash__(self):
        return hash((self.x, self.y))

class Offset:
    def __init__(self, point: Point, distance: int):
        self.point = point
        self.distance = distance

class Shortcut:
    def __init__(self, start: Point, end: Offset):
        self.start = start
        self.end = end

def findRoute(start: Point, end: Point, walls: Dict[Point, int]) -> Dict[Point, int]:
    queue = [start]
    visited = {}

    while queue:
        current = queue.pop(0)
        visited[current] = len(visited)

        if current == end:
            return visited

        for offset in getOffsets(current, 1):
            if offset.point in visited:
                continue

            if offset.point in walls:
                continue

            queue.append(offset.point)

    raise ValueError("Cannot find route")

def findShortcuts(route: Dict[Point, int], radius: int) -> Dict[int, int]:
    shortcuts = {}
    for current, step in route.items():
        offsets = getOffsets(current, radius)
        for offset in offsets:
            if offset.point in route:
                routeStep = route[offset.point]
                saving = routeStep - step - offset.distance
                if saving > 0:
                    shortcuts[Shortcut(current, offset)] = saving

    result = {}
    for _, saving in shortcuts.items():
        result[saving] = result.get(saving, 0) + 1

    return result

def getOffsets(from_point: Point, radius: int) -> List[Offset]:
    result = []

    for y in range(-radius, radius + 1):
        for x in range(-radius, radius + 1):
            candidatePoint = Point(from_point.X + x, from_point.Y + y)
            candidate = Offset(
                candidatePoint,
                getDistance(from_point, candidatePoint)
            )

            if 0 < candidate.distance <= radius:
                result.append(candidate)

    return result

def getDistance(from_point: Point, until_point: Point) -> int:
    xDistance = abs(from_point.X - until_point.X)
    yDistance = abs(from_point.Y - until_point.Y)
    return int(xDistance + yDistance)

def getCheats(file: List[str], radius: int) -> int:
    start = None
    end = None
    walls = {}

    for y, line in enumerate(file):
        for x, r in enumerate(line):
            if r == 'S':
                start = Point(x, y)
            elif r == 'E':
                end = Point(x, y)
            elif r == '#':
                walls[Point(x, y)] = 1

    route = findRoute(start, end, walls)
    cheats = findShortcuts(route, radius)

    found = 0
    greatShortcuts = 0
    k = 0
    while found < len(cheats):
        if k in cheats:
            found += 1
            
            if k >= 100:
                greatShortcuts += cheats[k]
        k += 1

    return greatShortcuts
# endregion

# region Day21
def day21a():
    file = readDayFile(21)
    num_map = {
        "A": Coord(2, 0),
        "0": Coord(1, 0),
        "1": Coord(0, 1),
        "2": Coord(1, 1),
        "3": Coord(2, 1),
        "4": Coord(0, 2),
        "5": Coord(1, 2),
        "6": Coord(2, 2),
        "7": Coord(0, 3),
        "8": Coord(1, 3),
        "9": Coord(2, 3)
    }

    dir_map = {
        "A": Coord(2, 1),
        "^": Coord(1, 1),
        "<": Coord(0, 0),
        "v": Coord(1, 0),
        ">": Coord(2, 0)
    }

    robots = 2

    return get_sequence(file, num_map, dir_map, robots)


def day21b():
    file = readDayFile(21)
    num_map = {
        "A": Coord(2, 0),
        "0": Coord(1, 0),
        "1": Coord(0, 1),
        "2": Coord(1, 1),
        "3": Coord(2, 1),
        "4": Coord(0, 2),
        "5": Coord(1, 2),
        "6": Coord(2, 2),
        "7": Coord(0, 3),
        "8": Coord(1, 3),
        "9": Coord(2, 3)
    }

    dir_map = {
        "A": Coord(2, 1),
        "^": Coord(1, 1),
        "<": Coord(0, 0),
        "v": Coord(1, 0),
        ">": Coord(2, 0)
    }

    robots = 25

    return get_sequence(file, num_map, dir_map, robots)

def abs(n: int) -> int:
    return -n if n < 0 else n

def atoi_no_err(s: str) -> int:
    return int(s)

def get_sequence(input_lines: List[str], num_map: Dict[str, Coord], dir_map: Dict[str, Coord], robot_count: int) -> int:
    total = 0
    cache = {}

    for line in input_lines:
        chars = list(line.replace("\n", ""))
        moves = get_num_pad_sequence(chars, "A", num_map)
        length = count_sequences(moves, robot_count, 1, cache, dir_map)
        total += atoi_no_err(line[:3].lstrip('0')) * length

    return total

def get_num_pad_sequence(input_chars: List[str], start: str, num_map: Dict[str, Coord]) -> List[str]:
    curr = num_map[start]
    seq = []

    for char in input_chars:
        dest = num_map[char]
        dx, dy = dest.X - curr.X, dest.Y - curr.Y

        horiz = [">"] * abs(dx) if dx >= 0 else ["<"] * abs(dx)
        vert = ["^"] * abs(dy) if dy >= 0 else ["v"] * abs(dy)

        if curr.Y == 0 and dest.X == 0:
            seq.extend(vert)
            seq.extend(horiz)
        elif curr.X == 0 and dest.Y == 0:
            seq.extend(horiz)
            seq.extend(vert)
        elif dx < 0:
            seq.extend(horiz)
            seq.extend(vert)
        else:
            seq.extend(vert)
            seq.extend(horiz)

        curr = dest
        seq.append("A")
    
    return seq

def count_sequences(input_seq: List[str], max_robots: int, robot: int, cache: Dict[str, List[int]], dir_map: Dict[str, Coord]) -> int:
    key = "".join(input_seq)
    if key in cache and robot <= len(cache[key]) and cache[key][robot-1] != 0:
        return cache[key][robot-1]

    if key not in cache:
        cache[key] = [0] * max_robots

    seq = get_dir_pad_sequence(input_seq, "A", dir_map)
    if robot == max_robots:
        return len(seq)

    steps = split_sequence(seq)
    count = 0
    for step in steps:
        c = count_sequences(step, max_robots, robot+1, cache, dir_map)
        count += c

    cache[key][robot-1] = count
    return count

def get_dir_pad_sequence(input_chars: List[str], start: str, dir_map: Dict[str, Coord]) -> List[str]:
    curr = dir_map[start]
    seq = []

    for char in input_chars:
        dest = dir_map[char]
        dx, dy = dest.X - curr.X, dest.Y - curr.Y

        horiz = [">"] * abs(dx) if dx >= 0 else ["<"] * abs(dx)
        vert = ["^"] * abs(dy) if dy >= 0 else ["v"] * abs(dy)

        if curr.X == 0 and dest.Y == 1:
            seq.extend(horiz)
            seq.extend(vert)
        elif curr.Y == 1 and dest.X == 0:
            seq.extend(vert)
            seq.extend(horiz)
        elif dx < 0:
            seq.extend(horiz)
            seq.extend(vert)
        else:
            seq.extend(vert)
            seq.extend(horiz)

        curr = dest
        seq.append("A")
    
    return seq

def split_sequence(input_seq: List[str]) -> List[List[str]]:
    result = []
    current = []

    for char in input_seq:
        current.append(char)
        if char == "A":
            result.append(current)
            current = []
    
    return result
# endregion

# region Day22
def day22a():
    file = readDayFile(22)
    seeds = []
    for seed in file:
        seeds.append(int(seed))
    
    sum_result = 0
    for n in seeds:
        sum_result += next_n(n, 2000)
    
    return sum_result

def day22b():
    file = readDayFile(22)
    seeds = []
    for seed in file:
        seeds.append(int(seed))
    
    max_val = 0
    
    sequence_payoffs = monkey_from_seed(seeds, 2000)
    for v in sequence_payoffs.values():
        max_val = max(max_val, v)
    
    return max_val

PRUNE = 16777216

class Sequence:
    def __init__(self, first: int, seconds: int, third: int, fourth: int):
        self.first = first
        self.seconds = seconds
        self.third = third
        self.fourth = fourth
    
    def __eq__(self, other):
        if not isinstance(other, Sequence):
            return False
        return (self.first == other.first and 
                self.seconds == other.seconds and 
                self.third == other.third and 
                self.fourth == other.fourth)
    
    def __hash__(self):
        return hash((self.first, self.seconds, self.third, self.fourth))

class Price:
    def __init__(self, cost: int, change: int):
        self.cost = cost
        self.change = change

def next(n: int) -> int:
    out = prune(mix(n << 6, n))
    out = prune(mix(out >> 5, out))
    out = prune(mix(out << 11, out))
    return out

def mix(in_val: int, secret: int) -> int:
    return in_val ^ secret

def prune(n: int) -> int:
    return n % PRUNE

def next_n(seed: int, sim_steps: int) -> int:
    random = seed
    for _ in range(sim_steps):
        random = next(random)
    return random

def all_prices(seed: int, sim_steps: int) -> List[Price]:
    random = seed
    retval = [Price(seed, 0)]
    
    for _ in range(sim_steps):
        random = next(random)
        cost = random % 10
        new_price = Price(cost, cost - retval[-1].cost)
        retval.append(new_price)
    
    return retval

def monkey_from_seed(seeds: List[int], sim_steps: int) -> Dict[Sequence, int]:
    retval = {}
    
    for seed in seeds:
        prices = all_prices(seed, sim_steps)
        seen = set()
        
        for i in range(4, sim_steps + 1):
            seq = Sequence(
                prices[i-3].change, 
                prices[i-2].change, 
                prices[i-1].change, 
                prices[i].change
            )
            
            if seq not in seen:
                seen.add(seq)
                retval[seq] = retval.get(seq, 0) + prices[i].cost
    
    return retval
# endregion

# region Day23
def day23a():
    file = [line.strip() for line in readDayFile(23)]

    np = NetworkProcessor()
    np.process_links(file)
    np.find_networks()
    return np.count_networks()


def day23b():
    file = readDayFile(23)

    np = NetworkProcessor()
    np.process_links(file)
    np.find_networks()
    return np.find_biggest_network()


class Network:
    def __init__(self, output: str, connections: Set[str]):
        self.output = output
        self.connections = connections

    def add(self, candidate: str) -> "Network":
        new_connections = self.connections | {candidate}
        new_output = ",".join(sorted(new_connections))
        return Network(new_output, new_connections)


class Computer:
    def __init__(self, name):
        self.name = name
        self.links = set()


class NetworkProcessor:
    def __init__(self):
        self.networks: Set[str] = set()
        self.comps: Dict[str, Computer] = {}
        self.simple_comps: Dict[str, List[str]] = defaultdict(list)

    def process_links(self, link_strs: List[str]):
        for link_str in link_strs:
            first, second = link_str.strip().split("-")
            self.add_or_update_computer(first, second)
            self.add_or_update_computer(second, first)
        self.populate_simple_comps()

    def add_or_update_computer(self, comp_name: str, linked_comp_name: str):
        if comp_name not in self.comps:
            self.comps[comp_name] = Computer(comp_name)
        self.comps[comp_name].links.add(linked_comp_name)

    def populate_simple_comps(self):
        for comp in self.comps.values():
            self.simple_comps[comp.name] = list(comp.links)

    def find_networks(self):
        for name, links in self.simple_comps.items():
            for i in range(len(links) - 1):
                for j in range(i + 1, len(links)):
                    i_name = links[i]
                    j_name = links[j]
                    if j_name in self.comps[i_name].links:
                        names = ",".join(sorted([name, i_name, j_name]))
                        self.networks.add(names)

    def count_networks(self) -> int:
        return sum(1 for network in self.networks 
                if network[0] == 't' or 
                    (len(network) > 3 and network[3] == 't') or 
                    (len(network) > 6 and network[6] == 't'))

    def find_biggest_network(self) -> str:
        network_cache = set()
        retval = ""

        for comp in self.comps.values():
            longest_found = self.bfs(comp, network_cache)
            if len(longest_found) > len(retval):
                retval = longest_found

        return retval

    def bfs(self, comp: Computer, network_cache: Set[str]) -> str:
        start = Network(comp.name, {comp.name})
        network_cache.add(start.output)
        queue = deque([start])

        n = start
        while queue:
            n = queue.popleft()
            for candidate in comp.links:
                if candidate in n.connections:
                    continue
                if self.is_connected(n, candidate):
                    next_network = n.add(candidate)
                    if next_network.output not in network_cache:
                        queue.append(next_network)
                        network_cache.add(next_network.output)

        return n.output

    def is_connected(self, n: Network, candidate: str) -> bool:
        return all(candidate in self.comps[existing].links for existing in n.connections)
# endregion

# region Day24
def day24a():
    file = readDayFile(24)
    wires, gates = parse_puzzle_input(file)

    while gates:
        for wire_name, gate in list(gates.items()):
            if can_eval_gate(gate, wires):
                wires[wire_name] = evaluate_gate(gate, wires)
                del gates[wire_name]

    z_wires = sorted(w for w in wires if w.startswith("z"))

    result = 0
    for wire in reversed(z_wires):
        result = (result << 1) | wires[wire]

    return str(result)

def day24b():
    file = readDayFile(24)
    _, gates = parse_puzzle_input(file)

    swapped = []
    carry = None

    gate_strings = [
        f"{gate['inputs'][0]} {['AND', 'OR', 'XOR'][gate['operation']]} {gate['inputs'][1]} -> {wire_name}"
        for wire_name, gate in gates.items()
    ]

    for i in range(45):
        n = f"{i:02}"
        m1 = find(f"x{n}", f"y{n}", "XOR", gate_strings)
        n1 = find(f"x{n}", f"y{n}", "AND", gate_strings)

        r1, z1, c1 = None, None, None

        if carry:
            r1 = find(carry, m1, "AND", gate_strings)
            if not r1:
                m1, n1 = n1, m1
                swapped.extend([m1, n1])
                r1 = find(carry, m1, "AND", gate_strings)

            z1 = find(carry, m1, "XOR", gate_strings)

            if m1.startswith("z"):
                m1, z1 = z1, m1
                swapped.extend([m1, z1])
            if n1.startswith("z"):
                n1, z1 = z1, n1
                swapped.extend([n1, z1])
            if r1.startswith("z"):
                r1, z1 = z1, r1
                swapped.extend([r1, z1])

            c1 = find(r1, n1, "OR", gate_strings)

        if c1 and c1.startswith("z") and c1 != "z45":
            c1, z1 = z1, c1
            swapped.extend([c1, z1])

        carry = n1 if carry is None else c1

    swapped.sort()
    return ",".join(swapped)

def parse_puzzle_input(input: List[str]) -> Tuple[Dict[str, int], Dict[str, Dict]]:
    parts = "\n".join(line.strip() for line in input).split("\n\n")
    wires = {line.split(": ")[0]: int(line.split(": ")[1]) for line in parts[0].split("\n")}

    gates = {}
    for line in parts[1].split("\n"):
        if not line:
            continue
        left, output = line.split(" -> ")
        inputs = left.split(" ")

        if len(inputs) == 3:
            operation = ["AND", "OR", "XOR"].index(inputs[1])
            gates[output] = {"operation": operation, "inputs": [inputs[0], inputs[2]], "output": output}

    return wires, gates

def evaluate_gate(gate: Dict, wires: Dict[str, int]) -> int:
    in1 = wires[gate["inputs"][0]]
    in2 = wires[gate["inputs"][1]]

    if gate["operation"] == 0:
        return in1 & in2
    elif gate["operation"] == 1:
        return in1 | in2
    elif gate["operation"] == 2:
        return in1 ^ in2
    return 0

def can_eval_gate(gate: Dict, wires: Dict[str, int]) -> bool:
    return all(input_name in wires for input_name in gate["inputs"])

def find(a: str, b: str, operator: str, gates: List[str]) -> str:
    for gate in gates:
        if gate.startswith(f"{a} {operator} {b}") or gate.startswith(f"{b} {operator} {a}"):
            return gate.split(" -> ")[-1]
    return ""
# endregion

# region Day25
def day25():
    file = readDayFile(25)
    locks = []
    keys = []

    for i in range(0, len(file), 8):
        if i + 7 > len(file):
            break

        heights = [0] * 5
        is_lock = False

        for row in range(7):
            for col, char in enumerate(file[i + row].strip()):
                if char == "#":
                    heights[col] += 1

            if row == 0 and file[i][0] == "#":
                is_lock = True

        for j in range(len(heights)):
            heights[j] -= 1

        if is_lock:
            locks.append(heights)
        else:
            keys.append(heights)

    matches = 0

    for lock in locks:
        for key in keys:
            if check_match(lock, key):
                matches += 1

    return matches

def check_match(lock, key):
    for i in range(5):
        if lock[i] + key[i] > 5:
            return False
    return True
# endregion

if __name__ == "__main__":
    main()