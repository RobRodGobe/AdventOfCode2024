using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

namespace AoC_2024
{
    class Program
    {
        static void Main(string[] args)
        {
            /* Day 1 */
            /* Part a */
            Console.WriteLine(Day2a());
            /* Part b */
            Console.WriteLine(Day2b());
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
    }
}
