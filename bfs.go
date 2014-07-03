package main

import (
	"flag"
	crawlers "github.com/termith/bfs-crawler/crawlers"
	queue "github.com/termith/bfs-crawler/queue"
)

func NewChannel(buffSize int) chan string {
	if buffSize < 1 {
		return make(chan string)
	} else {
		return make(chan string)
	}
}

func main() {

	myQueue := queue.NewQueue()

	var startUrl string
	var depthLimit int
	var buffSize int

	flag.StringVar(&startUrl, "url", "http://golang.org/", "Start url. golang.org by default")
	flag.IntVar(&depthLimit, "depth", 1, "Depth limit. 1 by default")
	flag.IntVar(&buffSize, "buffer", 5, "Buffer size. 5 by default")

	flag.Parse()

	urlFinder := crawlers.NewFinder(depthLimit)
	//pageSaver := crawlers.NewSaver("/home/ddemidov/bfs/")

	urlReciever := make(chan string, buffSize)
	doneChannel := make(chan string)

	myQueue.Push(queue.Url{Url: startUrl, Depth: 0})

	go urlFinder.Run(myQueue, urlReciever)
	go crawlers.NewSaver("/home/ddemidov/bfs/").Run(urlReciever, doneChannel)
	go crawlers.NewSaver("/home/ddemidov/bfs/").Run(urlReciever, doneChannel)

	<-doneChannel
	<-doneChannel

}
