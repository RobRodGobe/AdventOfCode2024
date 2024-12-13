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
            Console.WriteLine(Day13a());
            /* Part b */
            Console.WriteLine(Day13b());
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
    }
}
