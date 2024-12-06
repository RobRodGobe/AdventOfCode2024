using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.RegularExpressions;

namespace AoC_2024
{
    class Program
    {
        static void Main(string[] args)
        {
            /* Day 1 */
            /* Part a */
            Console.WriteLine(Day6a());
            /* Part b */
            Console.WriteLine(Day6b());
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
    }
}
