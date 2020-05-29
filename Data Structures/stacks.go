package main

import "fmt"

func main() {
	var stack []string

	stack = append(stack, "first insert")
	stack = append(stack, "second insert ")
	stack = append(stack, "third insert ")

	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Print(stack[n], "\n")

		stack = stack[:n]
	}
}
