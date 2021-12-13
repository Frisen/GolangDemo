package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left <= right {
		index := partition(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for arr[j] >= pivot && i < j {
			j--
		}
		for arr[i] <= pivot && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[i] = arr[i], arr[left]
	return i
}
