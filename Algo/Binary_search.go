package main

import "fmt"

func binarySearch(item int, arr []int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < item {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != item {
		return false
	}

	return true
}

func main() {
	items := []int{3, 5, 8, 12, 79, 342, 542, 722}
	fmt.Println(binarySearch(79, items))
}
