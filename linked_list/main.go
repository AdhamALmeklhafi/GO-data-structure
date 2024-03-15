package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func main() {
	ll := LinkedList{}
	ll.insert(1)
	ll.insert(2)
	ll.insert(3)

	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
