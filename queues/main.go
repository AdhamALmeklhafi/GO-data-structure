package main

import "fmt"

type NodeQueues struct {
	items []int
}

func (qu *NodeQueues) Enqueue(i int) {
	qu.items = append(qu.items, i)
}

func (qu *NodeQueues) Dequeue() int {
	toRemove := qu.items[0]
	qu.items = qu.items[1:]
	return toRemove
}

func main() {
	myQueue := NodeQueues{}
	fmt.Println(myQueue.items)
	myQueue.Enqueue(0)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue.items)

	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue.items)
}
