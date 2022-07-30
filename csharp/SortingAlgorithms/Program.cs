using System.Diagnostics;
using SortingAlgorithms;

// var array = new int[] {5, 8, 2, 10, 5, 12, 8, 1, 15, 4};
//
// QuickSort.SortAscending(array, 0, array.Length - 1);
//
// foreach (var item in array)
// {
//     Console.WriteLine(item);
// }

var random = new Random();
var array = Enumerable.Range(0, 1000)
    .Select(_ => random.Next(0, 1000))
    .ToArray();

var array1 = new int[array.Length];
var array2 = new int[array.Length];
array.CopyTo(array1, 0);
array.CopyTo(array2, 0);

var stopwatch1 = Stopwatch.StartNew();
QuickSort.SortAscending(array1, 0, array1.Length - 1);
stopwatch1.Stop();
Console.WriteLine($"Total time (Quick sort): {stopwatch1.ElapsedTicks}");

var stopwatch2 = Stopwatch.StartNew();
Array.Sort(array2);
stopwatch2.Stop();
Console.WriteLine($"Total time (.NET Array.Sort): {stopwatch2.ElapsedTicks}");

var correct = true;
for (var i = 0; i < array1.Length; ++i)
{
    if (array1[i] != array2[i]) correct = false;
}

Console.WriteLine(correct);
