package main

import (
	"fmt"
	crawler "github.com/termith/bfs-crawler/crawler"
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

	myCrawler := crawler.NewCrawler()

	testResult, err := myCrawler.FindAllUrls("http://www.yandex.ru/")
	if err != nil {
		panic(err)
	}

	fmt.Println(testResult)

}
