package sort

func Popsort(arr []int) []int {
	length := len(arr)
	for i := length; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 312459869
/*
312459869
132459869
123459869
*/
