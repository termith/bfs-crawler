package main

import (
	"flag"
	"fmt"
	crawler "github.com/termith/bfs-crawler/crawler"
	queue "github.com/termith/bfs-crawler/queue"
)

func main() {
	myQueue := queue.NewQueue()
	myCrawler := crawler.NewCrawler()

	var startUrl string
	var depthLimit int

	flag.StringVar(&startUrl, "url", "http://golang.org/", "Start url. golang.org by default")
	flag.IntVar(&depthLimit, "depth", 1, "Depth limit. 1 by default")

	flag.Parse()

	myQueue.Push(queue.Url{Url: startUrl, Depth: 0})

	for {
		if nextUrl, ok := myQueue.Pop().(queue.Url); ok { //If next Url is queue.Url
			if nextUrl.Depth <= depthLimit { // We need to save pages with depth = depthLimit
				if !myCrawler.CheckUrlIsVisited(nextUrl.Url) { // If url is not visited
					if status, e := myCrawler.SavePageToDisk(nextUrl.Url); status != 0 { // Try to save page on disk
						panic(e)
					}
					e := myCrawler.FindAllUrls(nextUrl, myQueue) // Then parse page
					if e != nil {
						panic(e)
					}
				}
			} else {
				myQueue.Clear()
				fmt.Println("We are done!")
				break
			}
		}
	}
}
