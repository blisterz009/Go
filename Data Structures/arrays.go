package main

import "fmt"

var arr [5]int

func main() {
	arr[0] = 4
	arr[1] = 5
	arr[2] = 12
	arr[3] = 122
	arr[4] = 8
	fmt.Println("This is an Array: ", arr)
	arr1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("This is also an Array: ", arr1)
}