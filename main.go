package main

import (
	"fmt"
	queue "github.com/termith/bfs-crawler/queue"
)

func main() {
	myQueue := crawler.NewQueue()

	myQueue.Push("hey!")
	myQueue.Push("My name is Dima!")

	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
}
