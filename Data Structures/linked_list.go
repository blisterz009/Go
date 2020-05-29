package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	length int
	start  *Node
}

func (l *List) insertNode(newNode *Node) {
	if l.length == 0 {
		l.start = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	l.length++
}
func (l *List) Display() {
	list := l.start
	for list != nil {
		fmt.Printf("%v ->", list.val)
		list = list.next
	}
	fmt.Println()
}

func main() {
	items := &List{}
	size := 10
	for i := 0; i < size; i++ {
		node := Node{val: rand.Intn(100) + 1}
		items.insertNode(&node)
	}
	items.Display()
}
