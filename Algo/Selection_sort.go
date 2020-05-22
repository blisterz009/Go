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

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var min = i
		for j := i; j < n; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
}

func main() {

	slice := generate(20)
	fmt.Println("Original slice : \n", slice)
	selectionsort(slice)
	fmt.Println("\n After sorting: \n", slice, "\n")
}
