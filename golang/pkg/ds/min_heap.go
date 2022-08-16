package ds

import "errors"

func FindMin(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("the heap is nil or empty")
	}

	return arr[0], nil
}

func Heapify(arr []int) ([]int, error) {
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("the given array is nil or empty")
	}

	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		arr = pushDown(arr, i, len(arr)-1)
	}

	return arr, nil
}

func pushDown(arr []int, first int, last int) []int {
	i := first

	for i <= lastInnerNodeIndex(last) {
		if leftChild(i) == last {
			// i has only one child and it is the last item of the array
			if arr[i] > arr[leftChild(i)] {
				swap(&arr[i], &arr[leftChild(i)])
			} else {
				// Do nothing
			}
			i = leftChild(i)
		} else {
			// i has two children
			if arr[i] > arr[leftChild(i)] && arr[leftChild(i)] <= arr[rightChild(i)] {
				swap(&arr[i], &arr[leftChild(i)])
				i = leftChild(i)
			} else if arr[i] > arr[rightChild(i)] && arr[rightChild(i)] < arr[leftChild(i)] {
				swap(&arr[i], &arr[rightChild(i)])
				i = rightChild(i)
			} else {
				i = last
			}
		}
	}

	return arr
}

func swap(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
}

func lastInnerNodeIndex(last int) int {
	return (last - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func isLeaf(arr []int, index int) bool {
	return !hasLeftChild(arr, index) && !hasRightChild(arr, index)
}

func hasLeftChild(arr []int, index int) bool {
	return isValidArrayIndex(arr, 2*index+1)
}

func hasRightChild(arr []int, index int) bool {
	return isValidArrayIndex(arr, 2*index+2)
}

func isValidArrayIndex(arr []int, index int) bool {
	return 0 <= index && index < len(arr)
}
