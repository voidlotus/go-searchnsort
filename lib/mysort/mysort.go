package mysort

func merge(arr []int, low, mid, high int) {
	leftSize := mid - low + 1
	rightSize := high - mid

	left := make([]int, leftSize)
	right := make([]int, rightSize)

	for i := 0; i < leftSize; i++ {
		left[i] = arr[low + i]
	}
	for i := 0; i < rightSize; i++ {
		right[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, low

	for i < leftSize && j < rightSize {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}

	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int, low int, high int) {
	if low < high {
		mid := low + (high-low)/2

		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}


func swap(x *int, y *int) {
	temp := *x    /* save the value at address x */
	*x = *y    /* put y into x */
	*y = temp    /* put temp into y */
}

func partition(arr []int, low, high int) int {
	pivot, i := arr[high], (low-1)
	for j:=low; j<high; j++ {
		if arr[j] < pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[high])
	return i+1
}

func QuickSort(arr []int, low, high int){
	if low < high {
		pivotIndex := partition(arr, low, high)

		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}


