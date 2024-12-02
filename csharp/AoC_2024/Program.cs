using System;
using System.IO;
using System.Linq;

namespace AoC_2024
{
    class Program
    {
        static void Main(string[] args)
        {
            /* Day 1 */
            /* Part a */
            Console.WriteLine(Day1a());
            /* Part b */
            Console.WriteLine(Day1b());
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
    }
}
