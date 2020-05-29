package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["abc"] = 1
	m["def"] = 2
	m["ghi"] = 3
	for key, value := range m {
		fmt.Printf("%v : %d \n", key, value)
	}
}
