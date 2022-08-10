package sorting

func BubbleSort(arr []int) {
	for i := 1; i <= len(arr); i++ {
		for j := len(arr) - 1; j >= i; j-- {
			if arr[j] < arr[j-1] {
				swap(&arr[j], &arr[j-1])
			}
		}
	}
}
