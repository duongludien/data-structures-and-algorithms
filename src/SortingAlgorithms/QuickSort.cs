namespace SortingAlgorithms;

public static class QuickSort
{
    public static int FindPivot(int[] array, int left, int right)
    {
        if (left >= right)
        {
            return -1;
        }
    
        var pivotIndex = left;

        var iterator = left + 1;
        while (iterator <= right && array[iterator] == array[pivotIndex])
        {
            iterator++;
        }

        if (iterator > right)
        {
            return -1;
        }

        return array[iterator] > array[pivotIndex] ? iterator : pivotIndex;
    }

    public static int Partition(int[] array, int left, int right, int pivot)
    {
        var leftIterator = left;
        var rightIterator = right;

        while (leftIterator < rightIterator)
        {
            while (array[leftIterator] < pivot)
            {
                leftIterator++;
            }

            while (array[rightIterator] >= pivot)
            {
                rightIterator--;
            }

            if (leftIterator < rightIterator)
            {
                (array[leftIterator], array[rightIterator]) = (array[rightIterator], array[leftIterator]);
            }
        }

        return leftIterator;
    }

    public static void SortAscending(int[] array, int left, int right)
    {
        var pivotIndex = FindPivot(array, left, right);

        if (pivotIndex == -1) return;
        
        var pivot = array[pivotIndex];
        var leftIndex = Partition(array, left, right, pivot);
        
        SortAscending(array, left, leftIndex - 1);
        SortAscending(array, leftIndex, right);
    }
}