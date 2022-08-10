package sorting

func InsertionSort(arr []int) {

    for i := 1; i < len(arr); i++ {
        j := i
        for j > 0 && arr[j] < arr[j-1] {
            swap(&arr[j], &arr[j-1])
            j--
        }
    }
}
