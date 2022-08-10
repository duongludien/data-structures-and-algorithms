package sorting

func SelectionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		min := arr[i]
		minKey := i
		for j := i; j < len(arr); i++ {
			if arr[j] < min {
				min = arr[j]
				minKey = j
			}

			swap(&arr[i], &arr[minKey])
		}
	}
}
