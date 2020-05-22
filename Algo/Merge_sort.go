package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func merge(left, right []int) (res []int) {
	res = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}

	return
}
func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	mid := int(num / 2)
	var (
		left  = make([]int, mid)
		right = make([]int, num-mid)
	)
	for i := 0; i < num; i++ {
		if i < mid {
			left[i] = items[i]
		} else {
			right[i-mid] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func main() {

	slice := generate(20)
	fmt.Println("Original slice : \n", slice)
	fmt.Println("\n After sorting \n", mergeSort(slice), "\n")
}
