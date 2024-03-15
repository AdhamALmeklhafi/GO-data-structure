package main

import "fmt"

type Stacks struct {
	items []int
}

func (s *Stacks) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stacks) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stacks{}
	fmt.Println(myStack)
	myStack.Push(0)
	myStack.Push(200)
	myStack.Push(100)
	fmt.Println(myStack)

	myStack.Pop()
	myStack.Pop()
	// the last item is 0 in the output
	fmt.Println(myStack)
}
