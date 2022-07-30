using System.Collections.Generic;
using SortingAlgorithms;
using Xunit;

namespace UnitTests;

public class QuickSortTests
{
    [Fact]
    public void FindPivot_WhenAllItemsAreEqual_ShouldReturnNegativeIndex()
    {
        var array = new[] {6, 6, 6, 6};
        var pivotIndex = QuickSort.FindPivot(array, 0, 3);
        Assert.Equal(-1, pivotIndex);
    }

    [Fact]
    public void FindPivot_WhenLeftGreaterThanOrEqualToRight_ShouldReturnNegativeIndex()
    {
        var array = new[] {3, 5, 9, 8};
        var pivotIndex = QuickSort.FindPivot(array, 2, 1);
        Assert.Equal(-1, pivotIndex);
    }
    
    [Theory]
    [InlineData(new[] {3, 5, 9, 8}, 0, 3, 1)]
    [InlineData(new[] {5, 3, 9, 8}, 0, 3, 0)]
    [InlineData(new[] {5, 5, 3, 8}, 0, 3, 0)]
    [InlineData(new[] {5, 5, 5, 9}, 0, 3, 3)]
    public void FindPivot_ShouldReturnValidPivotIndex(int[] array, int left, int right, int expectedPivotIndex)
    {
        var pivotIndex = QuickSort.FindPivot(array, left, right);
        Assert.Equal(expectedPivotIndex, pivotIndex);
    }

    [Fact]
    public void Partition_ShouldReturnCorrectLeftIndex()
    {
        var array = new[] {5, 4, 2, 1, 5};
        var leftIndex = QuickSort.Partition(array, 0, 4, 5);
        Assert.Equal(3, leftIndex);

        var leftPart = new[] {1, 4, 2};
        var rightPart = new[] {5, 5};

        VerifyPartition(array, 0, leftIndex - 1, leftPart);
        VerifyPartition(array, leftIndex, array.Length - 1, rightPart);
    }

    private static void VerifyPartition(IReadOnlyList<int> array, int left, int right, IReadOnlyList<int> expectedArray)
    {
        var iterator = left;
        var expectedArrayIndex = 0;

        while (iterator <= right)
        {
            Assert.Equal(expectedArray[expectedArrayIndex], array[iterator]);
            expectedArrayIndex++;
            iterator++;
        }
    } 
}