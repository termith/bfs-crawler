package main

import (
	"flag"
	"fmt"
	crawler "github.com/termith/bfs-crawler/crawler"
	queue "github.com/termith/bfs-crawler/queue"
)

func saveAllPages(urlChannel, doneChannel chan string, urlReceiver *crawler.Crawler) {
	fmt.Println("Start saving...")
	for {
		if url := <-urlChannel; url == "done" {
			fmt.Println("Saving is done!")
			doneChannel <- "done"
			return
		} else {
			status, e := urlReceiver.SavePageToDisk(url)
			if status > 0 {
				panic(e)
			}
		}
	}
}

func findAllUrls(waitQueue *queue.UrlQueue, urlFinder *crawler.Crawler, urlChannel chan string, depthLimit int) {
	fmt.Println("Start searching...")
	for {
		if nextUrl, ok := waitQueue.Pop().(queue.Url); ok { //If next Url is queue.Url
			if nextUrl.Depth <= depthLimit { // We need to save pages with depth = depthLimit
				if !urlFinder.CheckUrlIsVisited(nextUrl.Url) { // If url is not visited
					urlFinder.AppendToVisitedUrls(nextUrl.Url)
					urlChannel <- nextUrl.Url                      // Send url to saver
					e := urlFinder.FindAllUrls(nextUrl, waitQueue) // Then parse page
					if e != nil {
						panic(e)
					}
				}
			} else {
				urlChannel <- "done"
				waitQueue.Clear()
				fmt.Println("Searching is done!")
				break
			}
		}
	}
}


func main() {

	const buffSize int = 5

	urlReciever := make(chan string, buffSize)
	doneChannel := make(chan string)

	urlFinder := crawler.NewCrawler()
	pageSaver := crawler.NewCrawler()

	myQueue := queue.NewQueue()

	var startUrl string
	var depthLimit int

	flag.StringVar(&startUrl, "url", "http://golang.org/", "Start url. golang.org by default")
	flag.IntVar(&depthLimit, "depth", 1, "Depth limit. 1 by default")

	flag.Parse()

	myQueue.Push(queue.Url{Url: startUrl, Depth: 0})

	go findAllUrls(myQueue, urlFinder, urlReciever, depthLimit)
	go saveAllPages(urlReciever, doneChannel, pageSaver)

	<-doneChannel

}
