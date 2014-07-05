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

	var (
		startUrl   string
		homeDir    string
		depthLimit int
		buffSize   int
	)

	flag.StringVar(&startUrl, "url", "http://golang.org/", "Start url. golang.org by default")
	flag.IntVar(&depthLimit, "depth", 1, "Depth limit. 1 by default")
	flag.IntVar(&buffSize, "buffer", depthLimit, "Buffer size. Same as depthLimit by default")
	flag.StringVar(&homeDir, "dir", crawlers.HOME_DIR, "Directory to save. /home/termith/bfs/ as default")

	flag.Parse()

	urlFinder := crawlers.NewFinder(depthLimit, buffSize)

	urlReciever := make(chan string, buffSize)
	doneChannel := make(chan string)

	myQueue.Push(queue.Url{Url: startUrl, Depth: 0})

	go urlFinder.Run(myQueue, urlReciever)

	for i := buffSize; i > 0; i-- {
		go crawlers.NewSaver(homeDir).Run(urlReciever, doneChannel)
	}

	for j := buffSize; j > 0; j-- {
		<-doneChannel
	}

}
