package quick_sort

func partition(arr []int, left, right int) int {
	swapIndex := left - 1

	for i := left; i < right; i++ {
		if arr[right] > arr[i] {
			swapIndex++
			arr[i], arr[swapIndex] = arr[swapIndex], arr[i]
		}
	}

	swapIndex++
	arr[swapIndex], arr[right] = arr[right], arr[swapIndex]

	return swapIndex
}

func Solution(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		Solution(arr, left, pivot-1)
		Solution(arr, pivot+1, right)
	}
}
