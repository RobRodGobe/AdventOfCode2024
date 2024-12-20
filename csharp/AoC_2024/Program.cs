using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text;
using System.Text.RegularExpressions;

namespace AoC_2024
{
    class Program
    {
        static void Main(string[] args)
        {
            /* Day 1 */
            /* Part a */
            Console.WriteLine(Day21a());
            /* Part b */
            Console.WriteLine(Day21b());
        }

        static string[] ReadDayFile(int day)
        {
            string filePath = $"../../AoC_Files/{day}.txt";
            string[] fileContents = File.ReadAllLines(filePath);
            return fileContents;
        }

        #region Day1

        static int Day1a()
        {            
            // Read file
            string[] file = ReadDayFile(1);
            int[] list1 = new int[file.Length];
            int[] list2 = new int[file.Length];

            for (int i = 0; i < file.Length; i++)
            {
                string[] pairs = file[i].Split("   ");
                list1[i] = Int32.Parse(pairs[0]);
                list2[i] = Int32.Parse(pairs[1]);
            }

            Array.Sort(list1);
            Array.Sort(list2);

            int diff = 0;

            for (int i = 0; i < list1.Length; i++)
            {
                diff += Math.Abs(list1[i] - list2[i]);
            }

            return diff;
        }

        static int Day1b()
        {            
            // Read file
            string[] file = ReadDayFile(1);
            int[] list1 = new int[file.Length];
            int[] list2 = new int[file.Length];

            for (int i = 0; i < file.Length; i++)
            {
                string[] pairs = file[i].Split("   ");
                list1[i] = Int32.Parse(pairs[0]);
                list2[i] = Int32.Parse(pairs[1]);
            }

            int similar = 0;

            for (int i = 0; i < list1.Length; i++)
            {
                similar += list1[i] * list2.Where(l => l == list1[i]).Count();
            }
            
            return similar;
        }

        #endregion

        #region Day2
        static int Day2a()
        {
            string[] file = ReadDayFile(2);
            int safe = 0;

            for (int i = 0; i < file.Length; i++)
            {
                List<int> reports = file[i].Split(" ").Select(int.Parse).ToList();
                bool isAscending = true;
                bool isDescending = true;
                bool isSafe = true;

                for (int j = 1; j < reports.Count(); j++)
                {
                    int diff = reports[j] - reports[j - 1];

                    if (Math.Abs(diff) > 3)
                    {
                        isAscending = false;
                        isDescending = false;
                        isSafe = false;
                        break;
                    }

                    if (diff < 0) isAscending = false;
                    if (diff > 0) isDescending = false;
                    if (diff == 0) 
                    {
                        isAscending = false;
                        isDescending = false;
                    }

                    if (!isAscending && !isDescending)
                    {
                        isSafe = false;
                        break;
                    }
                }

                if (isSafe)
                {
                    safe++;
                }
            }

            return safe;
        }

        static int Day2b()
        {
        string[] file = ReadDayFile(2); // Read the input file
        int safe = 0;

            for (int i = 0; i < file.Length; i++)
            {
                List<int> reports = file[i].Split(" ").Select(int.Parse).ToList(); 

                if (IsSafeReport(reports, true) || IsSafeReport(reports, false))
                {
                    safe++;
                    continue;
                }

                bool isSafe = false;
                for (int j = 0; j < reports.Count(); j++)
                {
                    List<int> modifiedReports = reports.Where((_, idx) => idx != j).ToList();
                    if (IsSafeReport(modifiedReports, true) || IsSafeReport(modifiedReports, false))
                    {
                        isSafe = true;
                        break;
                    }
                }

                if (isSafe)
                {
                    safe++;
                }
            }

            return safe;
        }

        static bool IsSafeReport(List<int> reports, bool ascending)
        {
            for (int i = 1; i < reports.Count; i++)
            {
                int diff = reports[i] - reports[i - 1];
                if (ascending && diff < 0) return false;
                if (!ascending && diff > 0) return false;
                if (Math.Abs(diff) > 3 || diff == 0) return false;
            }
            return true;
        }

        #endregion

        #region Day3
        static int Day3a()
        {
            int mult = 0;
            string[] file = ReadDayFile(3);
            string line = string.Join("", file);

            string pattern = @"mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)";
            RegexOptions options = RegexOptions.Multiline;
            MatchCollection matches = Regex.Matches(line, pattern, options);

            for (int i = 0; i < matches.Count; i++)
            {
                string[] numbers = matches[i].Value.Replace("mul(", "").Replace(")","").Split(",");
                mult += Int32.Parse(numbers[0]) * Int32.Parse(numbers[1]);
            }

            return mult;
        }

        static int Day3b()
        {
            int mult = 0;
            string[] file = ReadDayFile(3);
            string line = string.Join("", file);

            string pattern = @"mul\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)|do\(\)|don't\(\)";
            RegexOptions options = RegexOptions.Multiline;
            MatchCollection matches = Regex.Matches(line, pattern, options);

            bool multiply = true;

            for (int i = 0; i < matches.Count; i++)
            {
                if (matches[i].Value == "do()") multiply = true;
                else if (matches[i].Value == "don't()") multiply = false;
                
                if (multiply && !matches[i].Value.Contains("do"))
                {
                    string[] numbers = matches[i].Value.Replace("mul(", "").Replace(")","").Split(",");
                    mult += Int32.Parse(numbers[0]) * Int32.Parse(numbers[1]);
                }
            }

            return mult;
        }
        #endregion

        #region Day4
        static int Day4a() 
        {
            string[] file = ReadDayFile(4);
            string word = "XMAS";
            int rows = file.Length;
            int cols = file[0].Length;
            int count = 0;
            int wordLength = word.Length;

            int[,] directions = {
                { 0, 1 },   // Right
                { 1, 0 },   // Down
                { 1, 1 },   // Down-right
                { 1, -1 },  // Down-left
                { 0, -1 },  // Left
                { -1, 0 },  // Up
                { -1, -1 }, // Up-left
                { -1, 1 }   // Up-right
            };

            for (int x = 0; x < rows; x++)
            {
                for (int y = 0; y < cols; y++)
                {
                    for (int d = 0; d < directions.GetLength(0); d++)
                    {
                        int dx = directions[d, 0];
                        int dy = directions[d, 1];
                        if (CheckWordBegin(x, y, dx, dy, wordLength, rows, cols, word, file))
                        {
                            count++;
                        }
                    }
                }
            }

            return count;
        }

        static int Day4b() 
        {
            string[] file = ReadDayFile(4);
            int rows = file.Length;
            int cols = file[0].Length;
            int count = 0;

            for (int x = 1; x < rows - 1; x++)
            {
                for (int y = 1; y < cols - 1; y++)
                {
                    if (IsXMasPattern(file, x, y))
                    {
                        count++;
                    }
                }
            }

            return count;
        }

        static bool CheckWordBegin(int x, int y, int dx, int dy, int length, int rows, int cols, string word, string[] grid)
        {
            for (int i = 0; i < length; i++)
            {
                int nx = x + i * dx;
                int ny = y + i * dy;

                if (nx < 0 || ny < 0 || nx >= cols || ny >= cols || grid[nx][ny] != word[i])
                {
                    return false;
                }
            }

            return true;
        }

        static bool IsXMasPattern(string[] grid, int x, int y)
        {
            string topLeftToBottomRight = $"{grid[x - 1][y - 1]}{grid[x][y]}{grid[x + 1][y + 1]}";
            string topRightToBottomLeft = $"{grid[x - 1][y + 1]}{grid[x][y]}{grid[x + 1][y - 1]}";

            return (IsValidMasPattern(topLeftToBottomRight) && IsValidMasPattern(topRightToBottomLeft));
        }

        static bool IsValidMasPattern(string pattern)
        {
            return pattern == "MAS" || pattern == "SAM";
        }

        #endregion

        #region Day5
        static int Day5a()
        {
            string[] file = ReadDayFile(5);
            List<(int Before, int After)> rules = new List<(int Before, int After)>();
            int dividerIndex = Array.IndexOf(file, "");
            int pages = 0;

            for (int i = 0; i < dividerIndex; i++)
            {
                var parts = file[i].Split('|').Select(int.Parse).ToArray();
                rules.Add((parts[0], parts[1]));
            }

            List<List<int>> updates = new List<List<int>>();

            for (int i = dividerIndex + 1; i < file.Length; i++)
            {
                updates.Add(file[i].Split(',').Select(int.Parse).ToList());
            }

            for (int i = 0; i < updates.Count(); i++)
            {
                if (IsUpdateValid(updates[i], rules))
                {                    
                    pages += GetMiddlePage(updates[i]);
                }
            }

            return pages;
        }

        static int Day5b()
        {
            string[] file = ReadDayFile(5);
            List<(int Before, int After)> rules = new List<(int Before, int After)>();
            int dividerIndex = Array.IndexOf(file, "");
            int pages = 0;

            for (int i = 0; i < dividerIndex; i++)
            {
                var parts = file[i].Split('|').Select(int.Parse).ToArray();
                rules.Add((parts[0], parts[1]));
            }

            List<List<int>> updates = new List<List<int>>();

            for (int i = dividerIndex + 1; i < file.Length; i++)
            {
                updates.Add(file[i].Split(',').Select(int.Parse).ToList());
            }

            for (int i = 0; i < updates.Count(); i++)
            {
                if (!IsUpdateValid(updates[i], rules))
                {
                    List<int> correctedUpdate = CorrectUpdate(updates[i], rules);
                    pages += GetMiddlePage(correctedUpdate);
                }
            }

            return pages;
        }

        private static bool IsUpdateValid(List<int> update, List<(int Before, int After)> rules)
        {
            var pagePositions = update.Select((page, index) => (Page: page, Position: index))
                                    .ToDictionary(x => x.Page, x => x.Position);

            for (int i = 0; i < rules.Count(); i++)
            {
                if (pagePositions.ContainsKey(rules[i].Before) && pagePositions.ContainsKey(rules[i].After))
                {
                    if (pagePositions[rules[i].Before] >= pagePositions[rules[i].After])
                    {
                        return false;
                    }
                }
            }

            return true;
        }

        private static int GetMiddlePage(List<int> update)
        {
            int midIndex = update.Count / 2;
            return update[midIndex];
        }

        private static List<int> CorrectUpdate(List<int> update, List<(int Before, int After)> rules)
        {
            // Create a dependency graph
            var graph = new Dictionary<int, List<int>>();
            var inDegree = new Dictionary<int, int>();

            for (int i = 0; i < update.Count(); i++)
            {
                graph[update[i]] = new List<int>();
                inDegree[update[i]] = 0;
            }

            for (int i = 0; i < rules.Count(); i++)
            {
                if (update.Contains(rules[i].Before) && update.Contains(rules[i].After))
                {
                    graph[rules[i].Before].Add(rules[i].After);
                    inDegree[rules[i].After]++;
                }
            }

            // Perform topological sort
            var sorted = new List<int>();
            var queue = new Queue<int>(inDegree.Where(kvp => kvp.Value == 0).Select(kvp => kvp.Key));

            while (queue.Count > 0)
            {
                int current = queue.Dequeue();
                sorted.Add(current);

                for (int i = 0; i < graph[current].Count(); i++)
                {
                    inDegree[graph[current][i]]--;
                    if (inDegree[graph[current][i]] == 0)
                    {
                        queue.Enqueue(graph[current][i]);
                    }
                }
            }

            return sorted;
        }

        #endregion

        #region Day6
        static int Day6a()
        {
            string[] file = ReadDayFile(6);
            int rows = file.Length;
            int cols = file[0].Length;

            Dictionary<char, (int, int)> directions = new Dictionary<char, (int, int)>
            {
                { '^', (-1, 0) },
                { '>', (0, 1) },
                { 'v', (1, 0) },
                { '<', (0, -1) }
            };

            Dictionary<char, char> turnRight = new Dictionary<char, char>
            {
                { '^', '>' },
                { '>', 'v' },
                { 'v', '<' },
                { '<', '^' }
            };

            (int Row, int Col) guardPos = (0, 0);
            char guardDir = ' ';
            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    if (directions.ContainsKey(file[r][c]))
                    {
                        guardPos = (r, c);
                        guardDir = file[r][c];
                        break;
                    }
                }
            }

            HashSet<(int, int)> visited = new HashSet<(int, int)> { guardPos };

            while (true)
            {
                (int dy, int dx) = directions[guardDir];
                var nextPos = (Row: guardPos.Row + dy, Col: guardPos.Col + dx);

                if (nextPos.Row < 0 || nextPos.Row >= rows || nextPos.Col < 0 || nextPos.Col >= cols)
                    break;

                if (file[nextPos.Row][nextPos.Col] == '#')
                {
                    guardDir = turnRight[guardDir];
                }
                else
                {
                    guardPos = nextPos;
                    visited.Add(guardPos);
                }
            }

            return visited.Count;
        }

        static int Day6b()
        {
            string[] file = ReadDayFile(6);
            int rows = file.Length;
            int cols = file[0].Length;

            Dictionary<char, (int, int)> directions = new Dictionary<char, (int, int)>
            {
                {'^', (-1, 0)},
                {'>', (0, 1)},
                {'v', (1, 0)},
                {'<', (0, -1)}
            };

            (int x, int y) guardPos = (0, 0);
            char guardDir = ' ';
            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    if (directions.ContainsKey(file[r][c]))
                    {
                        guardPos = (r, c);
                        guardDir = file[r][c];
                        break;
                    }
                }
            }

            int loopPositions = 0;

            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    if (IsGuardInLoop(file, guardPos, guardDir, (r, c)))
                        loopPositions++;
                }
            }

            return loopPositions;
        }

        static bool IsGuardInLoop(string[] mapLines, (int Row, int Col) guardStart, char guardDir, (int Row, int Col) obstruction)
        {
            Dictionary<char, (int, int)> directions = new()
            {
                { '^', (-1, 0) },
                { '>', (0, 1) },
                { 'v', (1, 0) },
                { '<', (0, -1) }
            };

            Dictionary<char, char> turnRight = new()
            {
                { '^', '>' },
                { '>', 'v' },
                { 'v', '<' },
                { '<', '^' }
            };

            int rows = mapLines.Length;
            int cols = mapLines[0].Length;

            // Add the obstruction
            char[][] tempMap = mapLines.Select(row => row.ToCharArray()).ToArray();
            tempMap[obstruction.Row][obstruction.Col] = '#';

            (int Row, int Col) guardPos = guardStart;
            char currentDir = guardDir;

            HashSet<(int, int, char)> visitedStates = new();
            Queue<(int, int, char)> recentHistory = new();

            int steps = 0, maxSteps = rows * cols * 2;

            while (true)
            {
                var state = (guardPos.Row, guardPos.Col, currentDir);

                if (visitedStates.Contains(state))
                {
                    if (recentHistory.Contains(state))
                        return true;
                }

                visitedStates.Add(state);
                recentHistory.Enqueue(state);
                if (recentHistory.Count > 10)
                    recentHistory.Dequeue();

                int dx = directions[currentDir].Item1;
                int dy = directions[currentDir].Item2;
                var nextPos = (Row: guardPos.Row + dx, Col: guardPos.Col + dy);

                if (nextPos.Row < 0 || nextPos.Row >= rows || nextPos.Col < 0 || nextPos.Col >= cols)
                {
                    return false;
                }
                else if (tempMap[nextPos.Row][nextPos.Col] == '#')
                {
                    currentDir = turnRight[currentDir];
                }
                else
                {
                    guardPos = nextPos;
                }

                steps++;
                if (steps > maxSteps)
                    return true; // Assume infinite loop if guard doesn't leave in a reasonable number of steps
            }
        }

        #endregion
    
        #region Day7
        static long Day7a() {
            string[] file = ReadDayFile(7);
            long sum = 0;

            for (int i = 0; i < file.Length; i++)
            {                
                string[] nums = file[i].Split(":");
                long total = long.Parse(nums[0]);
                long[] factors = nums[1].Trim().Split(" ").Select(long.Parse).ToArray();                
                if (CanCalibrate(total, factors, factors[0], 1))
                    sum += total;
            }

            return sum;
        }

        static long Day7b() {
            string[] file = ReadDayFile(7);
            long sum = 0;

            for (int i = 0; i < file.Length; i++)
            {                
                string[] nums = file[i].Split(":");
                long total = long.Parse(nums[0]);
                long[] factors = nums[1].Trim().Split(" ").Select(long.Parse).ToArray();                
                if (CanCalibrate2(total, factors, factors[0], 1))
                    sum += total;
            }

            return sum;
        }

        static bool CanCalibrate(long target, long[] numbers, long current, int i)
        {
            if (i == numbers.Count())
                return current == target;

            if (CanCalibrate(target, numbers, current + numbers[i], i + 1))
                return true;

            if (CanCalibrate(target, numbers, current * numbers[i], i + 1))
                return true;

            return false;
        }

        static bool CanCalibrate2(long target, long[] numbers, long current, int i)
        {
            if (i == numbers.Count())
                return current == target;

            if (CanCalibrate2(target, numbers, current + numbers[i], i + 1))
                return true;

            if (CanCalibrate2(target, numbers, current * numbers[i], i + 1))
                return true;

            if (CanCalibrate2(target, numbers, long.Parse($"{current}{numbers[i]}"), i + 1))
                return true;

            return false;
        }
        #endregion
    
        #region Day8
        static int Day8a()
        {
            string[] file = ReadDayFile(8);
            List<List<char>> matrix = file.Select(line => line.ToCharArray().ToList()).ToList();

            Dictionary<char, List<(int x, int y)>> antennaMap = GetAntennaMap(matrix);

            List<(int, int)> allAntinodes = new List<(int, int)>();

            foreach (var coords in antennaMap.Values)
            {
                List<(int x, int y)> antinodes = GetAntinodes(coords, matrix);
                allAntinodes.AddRange(antinodes);
            }

            List<(int x, int y)> uniqueAntinodes = GetUniqueAntinodes(allAntinodes);

            return uniqueAntinodes.Count();
        }

        static int Day8b()
        {
            string[] file = ReadDayFile(8);
            List<List<char>> matrix = file.Select(line => line.ToCharArray().ToList()).ToList();

            Dictionary<char, List<(int x, int y)>> antennaMap = GetAntennaMap(matrix);

            bool[,] antinodeMatrix = new bool[matrix.Count, matrix[0].Count()];

            foreach (var coords in antennaMap.Values)
            {
                ProcessAntinodeLines(coords, matrix, antinodeMatrix);
            }

            return GetUniqueAntinodesCount(antinodeMatrix);
        }

        static Dictionary<char, List<(int x, int y)>> GetAntennaMap(List<List<char>> matrix)
        {
            var map = new Dictionary<char, List<(int x, int y)>>();

            for (int i = 0; i < matrix.Count(); i++)
            {
                for (int j = 0; j < matrix[i].Count(); j++)
                {
                    char cell = matrix[i][j];
                    if (cell != '.')
                    {
                        if (!map.ContainsKey(cell))
                            map[cell] = new List<(int x, int y)>();

                        map[cell].Add((i, j));
                    }
                }
            }

            return map;
        }

        static List<(int x, int y)> GetAntinodes(List<(int x, int y)> coords, List<List<char>> matrix)
        {
            List<(int x, int y)> antinodes = new List<(int x, int y)>();

            for (int i = 0; i < coords.Count(); i++)
            {
                for (int j = 0; j < coords.Count(); j++)
                {
                    if (i != j)
                    {
                        var (ax, ay) = coords[i];
                        var (bx, by) = coords[j];

                        int cx = 2 * bx - ax;
                        int cy = 2 * by - ay;

                        if (WithinBoundaries(cx, 0, matrix.Count) && WithinBoundaries(cy, 0, matrix[0].Count))
                        {
                            antinodes.Add((cx, cy));
                        }
                    }
                }
            }

            return antinodes;
        }

        static bool WithinBoundaries(int value, int min, int max)
        {
            return value >= min && value < max;
        }

        static List<(int x, int y)> GetUniqueAntinodes(List<(int x, int y)> antinodes)
        {
            HashSet<(int, int)> uniqueSet = new HashSet<(int, int)>(antinodes);
            return uniqueSet.ToList();
        }

        static void ProcessAntinodeLines(List<(int x, int y)> coords, List<List<char>> matrix, bool[,] antinodeMatrix)
        {
            for (int i = 0; i < coords.Count(); i++)
            {
                for (int j = 0; j < coords.Count(); j++)
                {
                    if (i != j)
                    {
                        var (x1, y1) = coords[i];
                        var (x2, y2) = coords[j];

                        for (int x = 0; x < matrix.Count(); x++)
                        {
                            for (int y = 0; y < matrix[0].Count(); y++)
                            {
                                if (!antinodeMatrix[x, y])
                                {
                                    int lineResult = (y1 - y2) * x + (x2 - x1) * y + (x1 * y2 - x2 * y1);

                                    if (lineResult == 0)
                                    {
                                        antinodeMatrix[x, y] = true;
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }

        static int GetUniqueAntinodesCount(bool[,] antinodeMatrix)
        {
            int count = 0;

            for (int x = 0; x < antinodeMatrix.GetLength(0); x++)
            {
                for (int y = 0; y < antinodeMatrix.GetLength(1); y++)
                {
                    if (antinodeMatrix[x, y])
                    {
                        count++;
                    }
                }
            }

            return count;
        }
        #endregion
    
        #region Day9
        static long Day9a()
        {
            string[] file = ReadDayFile(9);
            string line = file[0];

            List<string> diskMap = ParseDiskMap(line);
            
            diskMap = CompactDisk(diskMap);

            return CalculateChecksum(diskMap);
        }

        static long Day9b()
        {
            string[] file = ReadDayFile(9);
            string line = file[0];

            List<string> diskMap = ParseDiskMap(line);

            diskMap = CompactDiskByFile(diskMap);

            return CalculateChecksum(diskMap);
        }

        static List<string> ParseDiskMap(string line)
        {
            List<string> nums = new List<string>();
            int index = 0;

            for (int i = 0; i < line.Length; i++)
            {
                int count = line[i] - '0';
                if (i % 2 == 0)
                {
                    for (int j = 0; j < count; j++)
                    {
                        nums.Add(index.ToString());
                    }
                    index++;
                }
                else
                {
                    for (int j = 0; j < count; j++)
                    {
                        nums.Add(".");
                    }
                }
            }

            return nums;
        }

        static List<string> CompactDisk(List<string> diskMap)
        {
            int L = 0;
            int R = diskMap.Count() - 1;

            while (L <= R)
            {
                if (diskMap[L] == "." && diskMap[R] != ".")
                {
                    string temp = diskMap[L];
                    diskMap[L] = diskMap[R];
                    diskMap[R] = temp;
                    R--;
                    L++;
                }
                else if (diskMap[R] == ".")
                    R--;
                else
                    L++;
            }

            return diskMap;
        }

        static long CalculateChecksum(List<string> diskMap)
        {
            long checksum = 0;

            for (int i = 0; i < diskMap.Count(); i++)
            {
                if (diskMap[i] != ".")
                {
                    checksum += i * (long.Parse(diskMap[i]));
                }
            }

            return checksum;
        }

        class DiskFile
        {
            public int Id { get; set; }
            public int Length { get; set; }
            public int StartIdx { get; set; }

            public DiskFile(int id, int length, int startIdx)
            {
                Id = id;
                Length = length;
                StartIdx = startIdx;
            }
        }

        static (List<DiskFile> files, Dictionary<int, int> spaces) AnalyzeDisk(List<string> diskMap)
        {
            var files = new List<DiskFile>();
            var spaces = new Dictionary<int, int>();
            int spaceStartIdx = -1;

            for (int i = 0; i < diskMap.Count; i++)
            {
                if (diskMap[i] == ".")
                {
                    if (spaceStartIdx == -1)
                        spaceStartIdx = i;

                    if (!spaces.ContainsKey(spaceStartIdx))
                        spaces[spaceStartIdx] = 0;

                    spaces[spaceStartIdx]++;
                }
                else
                {
                    if (spaceStartIdx != -1)
                        spaceStartIdx = -1;

                    int fileId = int.Parse(diskMap[i]);
                    if (files.Count <= fileId)
                        files.Add(new DiskFile(fileId, 0, i));

                    files[fileId].Length++;
                }
            }

            return (files, spaces);
        }

        static int GetFirstAvailableSpaceIdx(Dictionary<int, int> spaces, int fileLength)
        {
            foreach (var space in spaces)
            {
                if (space.Value >= fileLength)
                    return space.Key;
            }
            return -1;
        }

        static void UpdateSpaces(Dictionary<int, int> spaces, int spaceIdx, int fileLength)
        {
            if (spaces[spaceIdx] == fileLength)
            {
                spaces.Remove(spaceIdx);
            }
            else
            {
                int remainingSpace = spaces[spaceIdx] - fileLength;
                spaces.Remove(spaceIdx);
                spaces[spaceIdx + fileLength] = remainingSpace;
            }
        }

        static void MoveFile(List<string> diskMap, DiskFile file, int targetIdx)
        {
            for (int i = 0; i < file.Length; i++)
            {
                diskMap[targetIdx + i] = diskMap[file.StartIdx + i];
                diskMap[file.StartIdx + i] = ".";
            }
        }

        static List<string> CompactDiskByFile(List<string> diskMap)
        {
            var (files, spaces) = AnalyzeDisk(diskMap);

            files.Sort((a, b) => b.Id.CompareTo(a.Id));

            foreach (var file in files)
            {
                int targetIdx = GetFirstAvailableSpaceIdx(spaces, file.Length);
                if (targetIdx != -1 && targetIdx < file.StartIdx)
                {
                    MoveFile(diskMap, file, targetIdx);
                    UpdateSpaces(spaces, targetIdx, file.Length);
                }
            }

            return diskMap;
        }

        #endregion

        #region Day10
        static int Day10a()
        {
            string[] file = ReadDayFile(10);

            return SolveTopographicMap(file);
        }

        static int Day10b()
        {
            string[] file = ReadDayFile(10);

            return SolveTopographicMapTrailRatings(file);
        }

        static int SolveTopographicMap(string[] input)
        {
            int rows = input.Length;
            int cols = input[0].Length;
            int[,] map = new int[rows, cols];

            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    map[r, c] = input[r][c] - '0';
                }
            }

            int totalTrailheadScore = 0;

            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    if (map[r, c] == 0)
                    {
                        int trailheadScore = FindTrailheadScore(r, c, map, rows, cols);
                        totalTrailheadScore += trailheadScore;
                    }
                }
            }

            return totalTrailheadScore;
        }

        static int FindTrailheadScore(int startRow, int startCol, int[,] map, int rows, int cols)
        {
            bool[,] visited = new bool[rows, cols];
            HashSet<(int, int)> ninePositions = new HashSet<(int, int)>();

            DFS(startRow, startCol, 0, ninePositions, map, rows, cols, visited);

            return ninePositions.Count;
        }

        static bool DFS(int row, int col, int expectedHeight, HashSet<(int, int)> ninePositions, int[,] map, int rows, int cols, bool[,] visited)
        {
            if (row < 0 || row >= rows || col < 0 || col >= cols || 
                visited[row, col] || map[row, col] != expectedHeight)
            {
                return false;
            }

            visited[row, col] = true;

            if (expectedHeight == 9)
            {
                ninePositions.Add((row, col));
            }

            bool found = DFS(row - 1, col, expectedHeight + 1, ninePositions, map, rows, cols, visited) ||
                        DFS(row + 1, col, expectedHeight + 1, ninePositions, map, rows, cols, visited) ||
                        DFS(row, col - 1, expectedHeight + 1, ninePositions, map, rows, cols, visited) ||
                        DFS(row, col + 1, expectedHeight + 1, ninePositions, map, rows, cols, visited);

            return found;
        }

        static int SolveTopographicMapTrailRatings(string[] input)
        {
            var (map, rows, cols) = ParseTopographicMap(input);
            var scoresMap = new Dictionary<string, Dictionary<string, int>>();

            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    if (map[r, c] == 0)
                    {
                        StartHike(r, c, map, rows, cols, scoresMap);
                    }
                }
            }

            return GetScoresSum(scoresMap);
        }

        static void StartHike(int startR, int startC, int[,] map, int rows, int cols, 
                            Dictionary<string, Dictionary<string, int>> scoresMap)
        {
            var routes = new Queue<(int r, int c, int height, string initialCell)>();
            routes.Enqueue((startR, startC, 0, SerializeCoordinates(startR, startC)));

            while (routes.Count > 0)
            {
                var route = routes.Dequeue();

                foreach (var (dr, dc) in GetValidDirections())
                {
                    int newR = route.r + dr;
                    int newC = route.c + dc;
                    int newHeight = route.height + 1;

                    if (newR < 0 || newR >= rows || 
                        newC < 0 || newC >= cols)
                        continue;

                    int newCell = map[newR, newC];
                    
                    if (newCell != newHeight) continue;

                    if (newCell == 9)
                    {
                        if (!scoresMap.ContainsKey(route.initialCell))
                            scoresMap[route.initialCell] = new Dictionary<string, int>();

                        string endKey = SerializeCoordinates(newR, newC);
                        if (!scoresMap[route.initialCell].ContainsKey(endKey))
                            scoresMap[route.initialCell][endKey] = 0;

                        scoresMap[route.initialCell][endKey]++;
                    }
                    else
                    {
                        routes.Enqueue((newR, newC, newHeight, route.initialCell));
                    }
                }
            }
        }

        static IEnumerable<(int, int)> GetValidDirections()
        {
            return new[] { (-1, 0), (1, 0), (0, -1), (0, 1) };
        }

        static string SerializeCoordinates(int r, int c)
        {
            return $"{r}:{c}";
        }

        static int GetScoresSum(Dictionary<string, Dictionary<string, int>> scoresMap)
        {
            int sum = 0;
            foreach (var entry in scoresMap.Values)
            {
                sum += entry.Values.Sum();
            }
            return sum;
        }

        static (int[,] map, int rows, int cols) ParseTopographicMap(string[] input)
        {
            int rows = input.Length;
            int cols = input[0].Length;
            int[,] map = new int[rows, cols];

            for (int r = 0; r < rows; r++)
            {
                for (int c = 0; c < cols; c++)
                {
                    map[r, c] = input[r][c] - '0';
                }
            }

            return (map, rows, cols);
        }
        #endregion
    
        #region Day11
        static long Day11a()
        {
            string file = ReadDayFile(11)[0];
            string[] stones = file.Split(" ");
            Dictionary<long, long> rocks = stones.Select(long.Parse)
                                                .GroupBy(x => x)
                                                .ToDictionary(g => g.Key, g => (long)g.Count());
            
            Dictionary<long, long> finalRocks = BlinkRocks(rocks, 25);
            return finalRocks.Values.Sum();
        }

        static long Day11b()
        {
            string file = ReadDayFile(11)[0];
            string[] stones = file.Split(" ");
            Dictionary<long, long> rocks = stones.Select(long.Parse)
                                                .GroupBy(x => x)
                                                .ToDictionary(g => g.Key, g => (long)g.Count());
            
            Dictionary<long, long> finalRocks = BlinkRocks(rocks, 75);
            return finalRocks.Values.Sum();            
        }

        static List<long> Blink(long rock)
        {
            if (rock == 0) return [1];

            long digits = (long)Math.Floor(Math.Log10(rock)) + 1;

            if (digits % 2 != 0) return [rock * 2024];

            long halfDigits = digits / 2;
            long first = rock / (long)Math.Pow(10, halfDigits);
            long second = rock % (long)Math.Pow(10, halfDigits);

            return [first, second];
        }

        static Dictionary<long, long> BlinkRocksIteration(Dictionary<long, long> rocks)
        {
            Dictionary<long, long> result = new Dictionary<long, long>();

            foreach(var (rock, count) in rocks)
            {
                List<long> newRocks = Blink(rock);

                for (int i = 0; i < newRocks.Count(); i++)
                {
                    result[newRocks[i]] = result.GetValueOrDefault(newRocks[i]) + count;
                }
            }

            return result;
        }

        static Dictionary<long, long> BlinkRocks(Dictionary<long, long> rocks, int blinks)
        {
            Dictionary<long, long> currentRocks = new Dictionary<long, long>(rocks);

            for (int i = 0; i < blinks; i++)
            {
                currentRocks = BlinkRocksIteration(currentRocks);
            }

            return currentRocks;
        }
        #endregion
    
        #region Day12
        static int Day12a()
        {
            string[] file = ReadDayFile(12);
            return CalculateTotalFencingPrice(file);
        }

        static int Day12b()
        {
            string[] file = ReadDayFile(12);
            return CalculateTotalFencingPriceWithInnerSides(file);
        }

        static int CalculateTotalFencingPrice(string[] grid)
        {
            int n = grid.Length;
            int m = grid[0].Length;
            HashSet<(int, int)> visited = new HashSet<(int, int)>();
            int totalPrice = 0;

            for (int i = 0; i < n; i++)
            {
                for (int j = 0; j < m; j++)
                {
                    if (!visited.Contains((i, j)))
                    {
                        var (area, borders) = VisitRegion(grid, i, j, visited);
                        totalPrice += area * borders.Count;
                    }
                }
            }

            return totalPrice;
        }

        static int CalculateTotalFencingPriceWithInnerSides(string[] grid)
        {
            int n = grid.Length;
            int m = grid[0].Length;
            HashSet<(int, int)> visited = new HashSet<(int, int)>();
            int totalPrice = 0;

            for (int i = 0; i < n; i++)
            {
                for (int j = 0; j < m; j++)
                {
                    if (!visited.Contains((i, j)))
                    {
                        var (area, borders) = VisitRegion(grid, i, j, visited);
                        totalPrice += area * CountSides(borders);
                    }
                }
            }

            return totalPrice;
        }

        static (int area, HashSet<(int, int, int, int)> borders) VisitRegion(string[] grid, int startI, int startJ, HashSet<(int, int)> visited)
        {
            int n = grid.Length;
            int m = grid[0].Length;
            char plant = grid[startI][startJ];
            int area = 0;
            HashSet<(int, int, int, int)> borders = new HashSet<(int, int, int, int)>();

            void Visit(int i, int j)
            {
                if (visited.Contains((i, j)))
                    return;

                visited.Add((i, j));
                area++;

                int[] dx = { -1, 1, 0, 0 };
                int[] dy = { 0, 0, -1, 1 };

                for (int k = 0; k < 4; k++)
                {
                    int i2 = i + dx[k];
                    int j2 = j + dy[k];

                    if (i2 >= 0 && i2 < n && j2 >= 0 && j2 < m && grid[i2][j2] == plant)
                    {
                        Visit(i2, j2);
                    }
                    else
                    {
                        borders.Add((i, j, i2, j2));
                    }
                }
            }

            Visit(startI, startJ);
            return (area, borders);
        }

        static int CountSides(HashSet<(int, int, int, int)> borders)
        {
            HashSet<(int, int, int, int)> visited = new HashSet<(int, int, int, int)>();

            void VisitSide(int i, int j, int i2, int j2)
            {
                var side = (i, j, i2, j2);
                if (visited.Contains(side) || !borders.Contains(side))
                    return;

                visited.Add(side);

                if (i == i2)
                {
                    VisitSide(i - 1, j, i2 - 1, j2);
                    VisitSide(i + 1, j, i2 + 1, j2);
                }
                else
                {
                    VisitSide(i, j - 1, i2, j2 - 1);
                    VisitSide(i, j + 1, i2, j2 + 1);
                }
            }

            int numSides = 0;
            foreach (var side in borders)
            {
                if (visited.Contains(side))
                    continue;

                numSides++;
                VisitSide(side.Item1, side.Item2, side.Item3, side.Item4);
            }

            return numSides;
        }
        #endregion
    
        #region Day13
        static long Day13a()
        {
            string[] file = ReadDayFile(13);
            return GetMaxPrizeForMinTokens(file);
        }

        static long Day13b()
        {
            string[] file = ReadDayFile(13);
            List<ClawMachineSettings> machines = ParseClawMachineInput(file);

            for (int i = 0; i < machines.Count(); i++)
            {
                machines[i].PrizeX += 10_000_000_000_000;
                machines[i].PrizeY += 10_000_000_000_000;
            }

            List<string> adjustedInput = new List<string>();
            for (int i = 0; i < machines.Count(); i++)
            {
                adjustedInput.Add($"Button A: X+{machines[i].AX}, Y+{machines[i].AY}");
                adjustedInput.Add($"Button B: X+{machines[i].BX}, Y+{machines[i].BY}");
                adjustedInput.Add($"Prize: X={machines[i].PrizeX}, Y={machines[i].PrizeY}");
            }

            return GetMaxPrizeForMinTokens(adjustedInput.ToArray());
        }        

        static List<ClawMachineSettings> ParseClawMachineInput(string[] input)
        {
            var machines = new List<ClawMachineSettings>();
            var cleanedData = input.Where(line => !string.IsNullOrWhiteSpace(line)).ToArray();

            for (int i = 0; i < cleanedData.Length; i += 3)
            {
                var aMove = cleanedData[i].Replace("Button A: ", "").Split(", ");
                var bMove = cleanedData[i + 1].Replace("Button B: ", "").Split(", ");
                var prize = cleanedData[i + 2].Replace("Prize: ", "").Split(", ");

                machines.Add(new ClawMachineSettings(
                    int.Parse(aMove[0].Replace("X+", "")),
                    int.Parse(aMove[1].Replace("Y+", "")),
                    int.Parse(bMove[0].Replace("X+", "")),
                    int.Parse(bMove[1].Replace("Y+", "")),
                    long.Parse(prize[0].Replace("X=", "")),
                    long.Parse(prize[1].Replace("Y=", ""))
                ));
            }

            return machines;
        }

        static long? CalculatePrice(ClawMachineSettings machine)
        {
            long det = machine.AY * machine.BX - machine.AX * machine.BY;
            if (det == 0) return null;

            long b = (machine.AY * machine.PrizeX - machine.AX * machine.PrizeY) / det;
            long a = machine.AX != 0 ? (machine.PrizeX - b * machine.BX) / machine.AX : 0;

            if (machine.AX * a + machine.BX * b == machine.PrizeX &&
                machine.AY * a + machine.BY * b == machine.PrizeY &&
                a >= 0 && b >= 0)
            {
                return a * 3 + b;
            }

            return null;
        }

        static long GetMaxPrizeForMinTokens(string[] input)
        {
            var machines = ParseClawMachineInput(input);
            long totalTokens = 0;

            foreach (var machine in machines)
            {
                var tokens = CalculatePrice(machine);
                if (tokens.HasValue)
                {
                    totalTokens += tokens.Value;
                }
            }

            return totalTokens;
        }

        public class ClawMachineSettings
        {
            public int AX { get; set; }
            public int AY { get; set; }
            public int BX { get; set; }
            public int BY { get; set; }
            public long PrizeX { get; set; }
            public long PrizeY { get; set; }

            public ClawMachineSettings(int ax, int ay, int bx, int by, long prizeX, long prizeY)
            {
                AX = ax;
                AY = ay;
                BX = bx;
                BY = by;
                PrizeX = prizeX;
                PrizeY = prizeY;
            }
        }

        #endregion
    
        #region Day14
        record struct BathroomRobot((int X, int Y) P, (int X, int Y) V);

        static int Day14a()
        {
            string[] file = ReadDayFile(14);
            return CalculateSafetyFactor(file);
        }

        static int Day14b()
        {
            string[] file = ReadDayFile(14);
            return FindRobotSequenceTime(file);
        }

        static int CalculateSafetyFactor(string[] file)
        {
            const int width = 101;
            const int height = 103;
            const int duration = 100;

            var robots = ParseRobots(file);
            var finalPositions = CalculateFinalPositions(robots, width, height, duration);

            return ComputeQuadrantMultiplier(finalPositions, width, height);
        }

        static BathroomRobot[] ParseRobots(string[] lines)
        {
            return lines
                .Where(line => !string.IsNullOrWhiteSpace(line))
                .Select(ParseSingleRobot)
                .ToArray();
        }

        static BathroomRobot ParseSingleRobot(string line)
        {
            var parts = line.Split(' ');
            var p = parts[0].Substring(2).Split(',');
            var v = parts[1].Substring(2).Split(',');

            return new BathroomRobot(
                (int.Parse(p[0]), int.Parse(p[1])), 
                (int.Parse(v[0]), int.Parse(v[1]))
            );
        }

        static int[,] CalculateFinalPositions(BathroomRobot[] robots, int width, int height, int duration)
        {
            int[,] finalPositions = new int[width, height];

            foreach (var robot in robots)
            {
                int finalX = (robot.P.X + robot.V.X * duration) % width;
                int finalY = (robot.P.Y + robot.V.Y * duration) % height;

                finalX = finalX < 0 ? finalX + width : finalX;
                finalY = finalY < 0 ? finalY + height : finalY;

                finalPositions[finalX, finalY]++;
            }

            return finalPositions;
        }

        static int ComputeQuadrantMultiplier(int[,] finalPositions, int width, int height)
        {
            int midX = width / 2;
            int midY = height / 2;

            int topLeft = 0, topRight = 0, bottomLeft = 0, bottomRight = 0;

            for (int x = 0; x < width; x++)
            {
                for (int y = 0; y < height; y++)
                {
                    if (x == midX || y == midY) continue;

                    if (x < midX && y < midY) topLeft += finalPositions[x, y];
                    else if (x >= midX && y < midY) topRight += finalPositions[x, y];
                    else if (x < midX && y >= midY) bottomLeft += finalPositions[x, y];
                    else if (x >= midX && y >= midY) bottomRight += finalPositions[x, y];
                }
            }

            return topLeft * topRight * bottomLeft * bottomRight;
        }

        static int FindRobotSequenceTime(string[] file)
        {
            const int rows = 103;
            const int cols = 101;
            const int trunkSeqSize = 10;
            const int maxSeconds = 100000;

            var robots = ParseRobots(file);

            for (int sec = 1; sec <= maxSeconds; sec++)
            {
                var robotsByCol = new List<int>[cols];
                for (int i = 0; i < cols; i++)
                {
                    robotsByCol[i] = new List<int>();
                }

                for (int i = 0; i < robots.Length; i++)
                {
                    var newRobot = SimulateRobot(robots[i], rows, cols, 1);
                    robotsByCol[newRobot.P.X].Add(newRobot.P.Y);
                    robots[i] = newRobot;
                }

                for (int c = 0; c < cols; c++)
                {
                    var column = robotsByCol[c];
                    column.Sort();

                    if (HasConsecutiveSequence(column, trunkSeqSize))
                    {
                        return sec;
                    }
                }
            }

            return 0;
        }

        static bool HasConsecutiveSequence(List<int> sequence, int requiredLength)
        {
            if (sequence.Count < requiredLength) return false;

            int consecutiveCount = 1;
            for (int i = 1; i < sequence.Count; i++)
            {
                consecutiveCount = sequence[i] == sequence[i - 1] + 1 
                    ? consecutiveCount + 1 
                    : 1;

                if (consecutiveCount == requiredLength) return true;
            }

            return false;
        }

        static BathroomRobot SimulateRobot(BathroomRobot robot, int modRows, int modCols, int ticks)
        {
            int rowDelta = CalculateDelta(robot.V.Y, ticks, modRows);
            int newRow = ModAdd(robot.P.Y, rowDelta, modRows);

            int colDelta = CalculateDelta(robot.V.X, ticks, modCols);
            int newCol = ModAdd(robot.P.X, colDelta, modCols);

            return new BathroomRobot((newCol, newRow), robot.V);
        }

        static int CalculateDelta(int velocity, int ticks, int mod)
        {
            int delta = velocity * ticks % mod;
            if (delta < 0)
            {
                delta += mod;
            }
            return delta;
        }

        static int ModAdd(int a, int b, int mod)
        {
            int res = (a + b) % mod;
            if (res < 0)
            {
                res += mod;
            }
            return res;
        }
        #endregion
    
        #region Day15
        static long Day15a()
        {
            string[] file = ReadDayFile(15);
            var sections = String.Join("\n", file).Split("\n\n");
            var inputMap = sections[0].Split('\n');
            Dictionary<RobotVector, char> map = new Dictionary<RobotVector, char>();
            RobotVector robotPos = RobotVector.Zero;

            for (int row = 0; row < inputMap.Length; row++)
            {
                for (int col = 0; col < inputMap[row].Length; col++)
                {
                    RobotVector pos = new RobotVector(row, col);
                    char tile = inputMap[row][col];
                    if (tile == '#' || tile == 'O')
                    {
                        map[pos] = tile;
                    }
                    else if (tile == '@')
                    {
                        robotPos = pos;
                    }
                }
            }

            char[] moves = sections[1].Where(c => !char.IsWhiteSpace(c)).ToArray();

            foreach (char move in moves)
            {
                RobotVector dir = move switch
                {
                    '>' => RobotVector.Right,
                    'v' => RobotVector.Down,
                    '<' => RobotVector.Left,
                    '^' => RobotVector.Up,
                    _ => throw new Exception("Invalid direction")
                };
                List<char> thingsToPush = new();
                RobotVector next = robotPos + dir;
                while (true)
                {
                    if (map.TryGetValue(next, out var tile))
                    {
                        thingsToPush.Add(tile);
                        if (tile == '#')
                        {
                            break;
                        }
                        else
                        {
                            next += dir;
                        }
                    }
                    else
                    {
                        break;
                    }
                }
                if (thingsToPush.Count == 0)
                {
                    robotPos += dir;
                }
                else if (thingsToPush.Last() == 'O')
                {
                    for (int i = 0; i < thingsToPush.Count; i++)
                    {
                        map.Remove(robotPos + dir.Scale(1 + i));
                    }
                    for (int i = 0; i < thingsToPush.Count; i++)
                    {
                        map[robotPos + dir.Scale(2 + i)] = thingsToPush[i];
                    }
                    robotPos += dir;
                }
            }

            long total = 0;
            foreach (var (Position, Tile) in map)
            {
                if (Tile == 'O')
                {
                    var coordinate = 100 * Position.Row + Position.Col;
                    total += coordinate;
                }
            }
            return total;
        }

        static long Day15b()
        {
            string[] file = ReadDayFile(15);
            var sections = String.Join("\n", file).Split("\n\n");
            var inputMap = sections[0].Split('\n');
            Dictionary<RobotVector, Obstacle> map = new();
            RobotVector robotPos = RobotVector.Zero;
            for (int row = 0; row < inputMap.Length; row++)
            {
                for (int col = 0; col < inputMap[row].Length; col++)
                {
                    RobotVector pos = new RobotVector(row, col * 2);
                    char tile = inputMap[row][col];
                    if (tile == '#' || tile == 'O')
                    {
                        RobotVector right = pos + RobotVector.Right;
                        Obstacle obstacle = new(tile, pos, right);
                        map[pos] = obstacle;
                        map[right] = obstacle;
                    }
                    else if (tile == '@')
                    {
                        robotPos = pos;
                    }
                }
            }

            char[] moves = sections[1].Where(c => !char.IsWhiteSpace(c)).ToArray();

            foreach (char move in moves)
            {
                RobotVector dir = move switch
                {
                    '>' => RobotVector.Right,
                    'v' => RobotVector.Down,
                    '<' => RobotVector.Left,
                    '^' => RobotVector.Up,
                    _ => throw new Exception("Invalid direction")
                };
                
                HashSet<Obstacle> GetBoxesToPush(RobotVector pos)
                {
                    if (map.TryGetValue(pos, out var obstacle))
                    {
                        HashSet<Obstacle> results = [obstacle];
                        if (obstacle.Tile == 'O')
                        {
                            if (dir == RobotVector.Left)
                            {
                                results.UnionWith(GetBoxesToPush(obstacle.Left + RobotVector.Left));
                            }
                            else if (dir == RobotVector.Right)
                            {
                                results.UnionWith(GetBoxesToPush(obstacle.Right + RobotVector.Right));
                            }
                            else
                            {
                                results.UnionWith(GetBoxesToPush(obstacle.Left + dir));
                                results.UnionWith(GetBoxesToPush(obstacle.Right + dir));
                            }
                        }
                        return results;
                    }
                    else
                    {
                        return new HashSet<Obstacle>();
                    }
                }
                HashSet<Obstacle> thingsToPush = GetBoxesToPush(robotPos + dir);
                if (thingsToPush.Count == 0)
                {
                    robotPos += dir;
                }
                else if (thingsToPush.Any(obstacle => obstacle.Tile == '#'))
                {
                    continue;
                }
                else
                {
                    foreach (var obstacle in thingsToPush)
                    {
                        map.Remove(obstacle.Left);
                        map.Remove(obstacle.Right);
                    }
                    foreach (var obstacle in thingsToPush)
                    {
                        Obstacle newObstacle = new(obstacle.Tile, obstacle.Left + dir, obstacle.Right + dir);
                        map[newObstacle.Left] = newObstacle;
                        map[newObstacle.Right] = newObstacle;
                    }
                    robotPos += dir;
                }
            }

            HashSet<RobotVector> coordinates = new();
            foreach (var kvp in map)
            {
                if (kvp.Value.Tile == 'O')
                {
                    coordinates.Add(kvp.Value.Left);
                }
            }
            var total = coordinates.Sum(c => 100L * c.Row + c.Col);
            return total;
        }

        record class Obstacle(char Tile, RobotVector Left, RobotVector Right);

        public readonly record struct RobotVector(int Row, int Col)
        {
            public static readonly RobotVector Zero = new RobotVector(0, 0);
            public static readonly RobotVector Up = new RobotVector(-1, 0);
            public static readonly RobotVector Down = new RobotVector(+1, 0);
            public static readonly RobotVector Left = new RobotVector(0, -1);
            public static readonly RobotVector Right = new RobotVector(0, +1);

            public static RobotVector operator +(RobotVector left, RobotVector right)
            {
                return new RobotVector(left.Row + right.Row, left.Col + right.Col);
            }
            public static RobotVector operator -(RobotVector left, RobotVector right)
            {
                return new RobotVector(left.Row - right.Row, left.Col - right.Col);
            }
            public static RobotVector operator -(RobotVector vec)
            {
                return new RobotVector(-vec.Row, -vec.Col);
            }

            public readonly RobotVector Scale(int factor)
            {
                return new RobotVector(Row * factor, Col * factor);
            }
        }

        #endregion
    
        #region Day16
        public static int Day16a()
        {
            string[] file = ReadDayFile(16);
            State start = new State(new Position(file.Length - 2, 1), East);
            if (file[start.Pos.Row][start.Pos.Col] != 'S')
            {
                start = new State(new Position(1, file[0].Length - 2), South);
            }
            Solver solver = Solve(file, start);
            return solver.Cheapest;
        }

        public static int Day16b()
        {
            string[] file = ReadDayFile(16);
            State start = new State(new Position(file.Length - 2, 1), East);
            if (file[start.Pos.Row][start.Pos.Col] != 'S')
            {
                start = new State(new Position(1, file[0].Length - 2), South);
            }
            Solver solver = Solve(file, start);

            HashSet<Position> seen = new HashSet<Position>();
            Queue<State> queue = new Queue<State>();
            queue.Enqueue(solver.End);
            State zero = null;

            while (queue.Count > 0)
            {
                State v = queue.Dequeue();
                if (v != zero)
                {
                    seen.Add(v.Pos);
                    foreach (State parent in solver.Prov[v].Parents)
                    {
                        queue.Enqueue(parent);
                    }
                }
            }
            return seen.Count;
        }

        static readonly Direction East = new Direction(0, 1);
        static readonly Direction South = new Direction(1, 0);
        static readonly Direction West = new Direction(0, -1);
        static readonly Direction North = new Direction(-1, 0);

        static Solver Solve(string[] grid, State start)
        {
            Solver solver = new Solver(grid);
            solver.Add(start, null, 0);

            while (true)
            {
                while (!solver.PQ.ContainsKey(solver.Cheapest) || solver.PQ[solver.Cheapest].Count == 0)
                {
                    if (solver.Cheapest > solver.Highest)
                    {
                        throw new Exception("Ran out of priority queue");
                    }
                    solver.Cheapest++;
                }

                State v = solver.Pop(solver.Cheapest);

                if (solver.IsEnd(v.Pos))
                {
                    solver.End = v;
                    return solver;
                }

                (State straight, State left, State right) = v.Possible();

                if (solver.IsOpen(straight.Pos))
                {
                    solver.Add(straight, v, solver.Cheapest + 1);
                }
                if (solver.IsOpen(left.Pos))
                {
                    solver.Add(left, v, solver.Cheapest + 1000);
                }
                if (solver.IsOpen(right.Pos))
                {
                    solver.Add(right, v, solver.Cheapest + 1000);
                }
            }
        }

        class Direction
        {
            public int Row { get; }
            public int Col { get; }

            public Direction(int row, int col)
            {
                Row = row;
                Col = col;
            }

            public Direction TurnRight()
            {
                if (this == East) return South;
                if (this == South) return West;
                if (this == West) return North;
                return East;
            }

            public Direction TurnLeft()
            {
                if (this == East) return North;
                if (this == North) return West;
                if (this == West) return South;
                return East;
            }
        }

        class Position
        {
            public int Row { get; }
            public int Col { get; }

            public Position(int row, int col)
            {
                Row = row;
                Col = col;
            }

            public Position Move(Direction dir) => new Position(Row + dir.Row, Col + dir.Col);

            public override bool Equals(object obj)
            {
                if (obj is Position other)
                {
                    return Row == other.Row && Col == other.Col;
                }
                return false;
            }

            public override int GetHashCode() => HashCode.Combine(Row, Col);
        }

        class State
        {
            public Position Pos { get; }
            public Direction Dir { get; }

            public State(Position pos, Direction dir)
            {
                Pos = pos;
                Dir = dir;
            }

            public (State straight, State left, State right) Possible()
            {
                return (
                    new State(Pos.Move(Dir), Dir),
                    new State(Pos, Dir.TurnLeft()),
                    new State(Pos, Dir.TurnRight())
                );
            }

            public override bool Equals(object obj)
            {
                if (obj is State other)
                {
                    return Pos.Equals(other.Pos) && Dir.Equals(other.Dir);
                }
                return false;
            }

            public override int GetHashCode() => HashCode.Combine(Pos, Dir);
        }

        class Provenance
        {
            public int Cost { get; set; }
            public List<State> Parents { get; }

            public Provenance(int cost)
            {
                Cost = cost;
                Parents = new List<State>();
            }

            public void MaybeAdd(State parent, int cost)
            {
                if (Cost > cost)
                {
                    Cost = cost;
                    Parents.Clear();
                    Parents.Add(parent);
                }
                else if (Cost == cost)
                {
                    Parents.Add(parent);
                }
            }
        }

        class Solver
        {
            public string[] Grid { get; }
            public Dictionary<int, List<State>> PQ { get; }
            public Dictionary<State, Provenance> Prov { get; }
            public Dictionary<State, int> Visited { get; }
            public int Cheapest { get; set; }
            public int Highest { get; set; }
            public State End { get; set; }

            public Solver(string[] grid)
            {
                Grid = grid;
                PQ = new Dictionary<int, List<State>>();
                Prov = new Dictionary<State, Provenance>();
                Visited = new Dictionary<State, int>();
                Cheapest = 0;
                Highest = 0;
            }

            public void Add(State v, State prev, int cost)
            {
                if (!Prov.ContainsKey(v))
                {
                    Prov[v] = new Provenance(cost);
                }
                Prov[v].MaybeAdd(prev, cost);

                if (!Visited.TryGetValue(v, out int existingCost) || cost < existingCost)
                {
                    Visited[v] = cost;
                    if (!PQ.ContainsKey(cost))
                    {
                        PQ[cost] = new List<State>();
                    }
                    PQ[cost].Add(v);
                    Highest = Math.Max(Highest, cost);
                }
            }

            public State Pop(int cost)
            {
                State v = PQ[cost][0];
                PQ[cost].RemoveAt(0);
                return v;
            }

            public char Lookup(Position p) => Grid[p.Row][p.Col];

            public bool IsEnd(Position p) => Lookup(p) == 'E';

            public bool IsOpen(Position p) => Lookup(p) != '#';
        }
        #endregion

        #region Day17
        static string Day17a()
        {
            string[] file = ReadDayFile(17);
            SmallComputer comp = InitComputer(file);
            List<ulong> output = Run(comp.A, comp.B, comp.C, comp.Program);

            string result = string.Join(",", output);
            return result;
        }

        static ulong Day17b()
        {
            string[] file = ReadDayFile(17);
            SmallComputer comp = InitComputer(file);

            Queue<(ulong a, int n)> queue = new Queue<(ulong, int)>();
            queue.Enqueue((0, 1));

            while (queue.Count > 0)
            {
                var (a, n) = queue.Dequeue();

                if (n > comp.Program.Count)
                    return a;

                for (ulong i = 0; i < 8; i++)
                {
                    ulong a2 = (a << 3) | i;
                    List<ulong> output = Run(a2, 0, 0, comp.Program);
                    List<ulong> target = comp.Program.GetRange(comp.Program.Count - n, n);

                    if (MatchesProgram(output, target))
                    {
                        queue.Enqueue((a2, n + 1));
                    }
                }
            }

            return 0;
        }

        class SmallComputer
        {
            public ulong A { get; set; }
            public ulong B { get; set; }
            public ulong C { get; set; }
            public List<ulong> Program { get; set; } = new List<ulong>();
            public List<ulong> Out { get; set; } = new List<ulong>();
        }

        static SmallComputer InitComputer(string[] puzzle)
        {
            SmallComputer res = new SmallComputer();
            Regex reR = new Regex(@"Register ([A|B|C]): (\d+)");
            Regex reP = new Regex(@"\d");

            foreach (string line in puzzle)
            {
                if (line.Contains("Program"))
                {
                    MatchCollection matches = reP.Matches(line);
                    foreach (Match match in matches)
                    {
                        res.Program.Add(ulong.Parse(match.Value));
                    }
                }
                else if (line.Contains("Register"))
                {
                    Match match = reR.Match(line);
                    int registerValue = int.Parse(match.Groups[2].Value);
                    if (match.Groups[1].Value == "A") res.A = (ulong)registerValue;
                    if (match.Groups[1].Value == "B") res.B = (ulong)registerValue;
                    if (match.Groups[1].Value == "C") res.C = (ulong)registerValue;
                }
            }

            return res;
        }

        static List<ulong> Run(ulong a, ulong b, ulong c, List<ulong> program)
        {
            List<ulong> output = new List<ulong>();
            ulong instruction, param;

            for (ulong pointer = 0; pointer < (ulong)program.Count; pointer += 2)
            {
                instruction = program[(int)pointer];
                param = program[(int)pointer + 1];

                ulong combo = param;
                if (param == 4) combo = a;
                else if (param == 5) combo = b;
                else if (param == 6) combo = c;

                switch (instruction)
                {
                    case 0:
                        a >>= (int)combo;
                        break;
                    case 1:
                        b ^= param;
                        break;
                    case 2:
                        b = combo % 8;
                        break;
                    case 3:
                        if (a != 0)
                            pointer = param - 2;
                        break;
                    case 4:
                        b ^= c;
                        break;
                    case 5:
                        output.Add(combo % 8);
                        break;
                    case 6:
                        b = a >> (int)combo;
                        break;
                    case 7:
                        c = a >> (int)combo;
                        break;
                }
            }

            return output;
        }

        static bool MatchesProgram(List<ulong> output, List<ulong> expected)
        {
            if (output.Count != expected.Count)
                return false;

            for (int i = 0; i < output.Count; i++)
            {
                if (output[i] != expected[i])
                    return false;
            }

            return true;
        }

        #endregion    
    
        #region Day18
        static int Day18a()
        {
            string[] file = ReadDayFile(18);

            var coordinates = AllBlockedCoords(file);
            var start = new Coord(0, 0);
            var end = new Coord(70, 70);

            var (shortestPath, _) = ShortestPath(coordinates.Take(1024).ToList(), start, end);

            return shortestPath;
        }

        static string Day18b()
        {
            string[] file = ReadDayFile(18);

            var coordinates = AllBlockedCoords(file);
            var start = new Coord(0, 0);
            var end = new Coord(70, 70);

            var block = SearchForBlockage(coordinates, start, end);

            return $"{block.X}, {block.Y}";
        }

        struct Coord
        {
            public int X, Y;

            public Coord(int x, int y)
            {
                X = x;
                Y = y;
            }

            public bool IsValidStep(HashSet<Coord> blocked)
            {
                return !blocked.Contains(this) && X >= 0 && Y >= 0 && X <= 70 && Y <= 70;
            }

            public IEnumerable<Coord> NextCoords()
            {
                return new List<Coord>
                {
                    new Coord(X + 1, Y),
                    new Coord(X - 1, Y),
                    new Coord(X, Y + 1),
                    new Coord(X, Y - 1)
                };
            }
        }

        static HashSet<Coord> BlockedCoords(IEnumerable<Coord> coords)
        {
            return new HashSet<Coord>(coords);
        }

        static List<Coord> AllBlockedCoords(string[] lines)
        {
            var allCoords = new List<Coord>();

            foreach (var line in lines)
            {
                var parts = line.Split(',');
                if (int.TryParse(parts[0], out int x) && int.TryParse(parts[1], out int y))
                {
                    allCoords.Add(new Coord(x, y));
                }
                else
                {
                    throw new Exception("Couldn't parse int.");
                }
            }

            return allCoords;
        }

        static (int, bool) ShortestPath(List<Coord> coords, Coord start, Coord end)
        {
            var blocked = BlockedCoords(coords);
            var visited = new Dictionary<Coord, int> { [start] = 0 };
            var queue = new Queue<Coord>();
            queue.Enqueue(start);

            while (queue.Count > 0)
            {
                var node = queue.Dequeue();

                foreach (var c in node.NextCoords())
                {
                    if (c.IsValidStep(blocked) && !visited.ContainsKey(c))
                    {
                        queue.Enqueue(c);
                        visited[c] = visited[node] + 1;
                    }
                }
            }

            return visited.TryGetValue(end, out int distance) ? (distance, true) : (0, false);
        }

        static Coord SearchForBlockage(List<Coord> allCoords, Coord start, Coord end)
        {
            int l = 1024;
            int r = allCoords.Count - 1;
            int m = (l + r) / 2;

            while (l != m && r != m)
            {
                var (_, ok) = ShortestPath(allCoords.Take(m).ToList(), start, end);

                if (ok)
                {
                    l = m;
                    m = (m + r) / 2;
                }
                else
                {
                    r = m;
                    m = (l + m) / 2;
                }
            }

            return allCoords[m];
        }
        #endregion
    
        #region Day19
        public static int Day19a()
        {
            string[] file = ReadDayFile(19);
            string rawFile = String.Join("\n", file);
            var (towels, patterns) = ParseDesignFile(rawFile);
            int count = 0;
            var cache = new Dictionary<string, bool>();

            foreach (var pattern in patterns)
            {
                if (DesignPossible(pattern, towels, cache))
                {
                    count++;
                }
            }
            return count;
        }

        public static long Day19b()
        {
            string[] file = ReadDayFile(19);
            string rawFile = String.Join("\n", file);
            var (towels, patterns) = ParseDesignFile(rawFile);
            long count = 0;
            var cache = new Dictionary<string, long>();

            foreach (var pattern in patterns)
            {
                count += WaysPossible(pattern, towels, cache);
            }
            return count;
        }

        private static (List<string>, List<string>) ParseDesignFile(string input)
        {
            var parts = input.Split(new[] { "\n\n" }, StringSplitOptions.None);
            var towels = parts[0].Split(", ").ToList();
            var patterns = parts[1].Split(new[] { "\n" }, StringSplitOptions.RemoveEmptyEntries).ToList();
            return (towels, patterns);
        }

        private static bool DesignPossible(string pattern, List<string> towels, Dictionary<string, bool> cache)
        {
            if (cache.TryGetValue(pattern, out bool result))
            {
                return result;
            }

            foreach (var towel in towels)
            {
                if (towel == pattern)
                {
                    return true;
                }
                else if (pattern.StartsWith(towel))
                {
                    var isPossible = DesignPossible(pattern.Substring(towel.Length), towels, cache);
                    if (isPossible)
                    {
                        cache[pattern] = true;
                        return true;
                    }
                }
            }

            cache[pattern] = false;
            return false;
        }

        private static long WaysPossible(string pattern, List<string> towels, Dictionary<string, long> cache)
        {
            if (cache.TryGetValue(pattern, out long ways))
            {
                return ways;
            }

            ways = 0;

            foreach (var towel in towels)
            {
                if (towel == pattern)
                {
                    ways++;
                }
                else if (pattern.StartsWith(towel))
                {
                    ways += WaysPossible(pattern.Substring(towel.Length), towels, cache);
                }
            }

            cache[pattern] = ways;
            return ways;
        }
        #endregion
    
        #region Day20
        static int Day20a()
        {
            string[] file = ReadDayFile(20);
            return GetCheats(file, 2);
        }

        static int Day20b()
        {
            string[] file = ReadDayFile(20);
            return GetCheats(file, 20);
        }

        record Point(int X, int Y)
        {
            public override int GetHashCode() => HashCode.Combine(X, Y);
        }

        record Offset(Point Point, int Distance);

        static Dictionary<Point, int> FindRoute(Point start, Point end, HashSet<Point> walls)
        {
            var queue = new Queue<Point>();
            queue.Enqueue(start);

            var visited = new Dictionary<Point, int> { [start] = 0 };

            while (queue.Count > 0)
            {
                var current = queue.Dequeue();
                if (current == end) return visited;

                foreach (var offset in GetOffsets(current, 1))
                {
                    if (!visited.ContainsKey(offset.Point) && !walls.Contains(offset.Point))
                    {
                        queue.Enqueue(offset.Point);
                        visited[offset.Point] = visited[current] + 1;
                    }
                }
            }

            throw new InvalidOperationException("Cannot find route");
        }

        static Dictionary<int, int> FindShortcuts(Dictionary<Point, int> route, int radius)
        {
            var shortcuts = new Dictionary<(Point, Point), int>();

            foreach (var (current, step) in route)
            {
                foreach (var offset in GetOffsets(current, radius))
                {
                    if (route.TryGetValue(offset.Point, out var routeStep))
                    {
                        var saving = routeStep - step - offset.Distance;
                        if (saving > 0)
                        {
                            shortcuts[(current, offset.Point)] = saving;
                        }
                    }
                }
            }

            var result = new Dictionary<int, int>();
            foreach (var saving in shortcuts.Values)
            {
                if (result.ContainsKey(saving))
                    result[saving]++;
                else
                    result[saving] = 1;
            }

            return result;
        }

        static IEnumerable<Offset> GetOffsets(Point from, int radius)
        {
            for (int y = -radius; y <= radius; y++)
            {
                for (int x = -radius; x <= radius; x++)
                {
                    var candidatePoint = new Point(from.X + x, from.Y + y);
                    var distance = GetDistance(from, candidatePoint);

                    if (distance > 0 && distance <= radius)
                        yield return new Offset(candidatePoint, distance);
                }
            }
        }

        static int GetDistance(Point from, Point until)
        {
            return Math.Abs(from.X - until.X) + Math.Abs(from.Y - until.Y);
        }

        static int GetCheats(string[] file, int radius)
        {
            Point start = null, end = null;
            var walls = new HashSet<Point>();

            for (int y = 0; y < file.Length; y++)
            {
                var line = file[y];
                for (int x = 0; x < line.Length; x++)
                {
                    var point = new Point(x, y);
                    switch (line[x])
                    {
                        case 'S': start = point; break;
                        case 'E': end = point; break;
                        case '#': walls.Add(point); break;
                    }
                }
            }

            var route = FindRoute(start, end, walls);
            var cheats = FindShortcuts(route, radius);

            int found = 0, greatShortcuts = 0;
            for (int k = 0; found < cheats.Count; k++)
            {
                if (cheats.TryGetValue(k, out var count))
                {
                    found++;
                    if (k >= 100) greatShortcuts += count;
                }
            }

            return greatShortcuts;
        }
        #endregion
    
        #region Day21
        static int Day21a()
        {
            string[] file = ReadDayFile(21);

            return file.Count();
        }

        static int Day21b()
        {
            string[] file = ReadDayFile(21);

            return file.Count();
        }
        #endregion
    }
}
