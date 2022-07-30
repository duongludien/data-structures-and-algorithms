using SortingAlgorithms;

var array = new int[] {5, 8, 2, 10, 5, 12, 8, 1, 15, 4};

QuickSort.SortAscending(array, 0, array.Length - 1);

foreach (var item in array)
{
    Console.WriteLine(item);
}