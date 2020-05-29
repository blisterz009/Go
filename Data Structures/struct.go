package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "Pintu",
		lastName:  "Kumar",
		age:       23,
	}

	fmt.Println("Person: ", p1)
}
