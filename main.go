package main

import (
	"fmt"
	queue "github.com/termith/bfs-crawler/queue"
)

func main() {
	myQueue := queue.NewQueue()

	myQueue.Push("hey!")
	myQueue.Push("My name is Dima!")

	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())

	if myQueue.Pop() == nil {
		fmt.Println("Queue is empty!")
	}

}
