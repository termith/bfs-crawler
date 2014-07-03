package crawlers

import (
	html "code.google.com/p/go.net/html"
	"fmt"
	queue "github.com/termith/bfs-crawler/queue"
	"net/http"
	"strings"
)

// Tools for crawler

type Finder struct {
	visitedUrls []string
	depthLimit  int
}

func NewFinder(maxDepth int) *Finder {
	return &Finder{
		visitedUrls: make([]string, 0),
		depthLimit:  maxDepth,
	}
}

func (f *Finder) Run(waitQueue *queue.UrlQueue, urlChannel chan<- string) {
	fmt.Println("Start searching...")
	for {
		if nextUrl, ok := waitQueue.Pop().(queue.Url); ok { //If next Url is queue.Url
			if nextUrl.Depth <= f.depthLimit { // We need to save pages with depth = depthLimit
				if !f.CheckUrlIsVisited(nextUrl.Url) { // If url is not visited
					f.AppendToVisitedUrls(nextUrl.Url)
					urlChannel <- nextUrl.Url              // Send url to saver
					e := f.FindAllUrls(nextUrl, waitQueue) // Then parse page
					if e != nil {
						panic(e)
					}
				}
			} else {
				urlChannel <- "done"
				urlChannel <- "done"
				waitQueue.Clear()
				fmt.Println("Searching is done!")
				break
			}
		}
	}
}

func (f *Finder) FindAllUrls(url queue.Url, urlQueue *queue.UrlQueue) error {

	/*TODO: I have to draw this function because I don't understand how it works
	  And may be it can be done simple... */

	var newParser func(*html.Node)
	newParser = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.Contains(a.Val, "http") {
						urlQueue.Push(queue.Url{Url: a.Val, Depth: url.Depth + 1})
						break
					}

				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			newParser(c)
		}
	}

	page, getError := http.Get(url.Url)
	if getError != nil {
		return getError
	}

	pageReader := page.Body

	pageContent, parseError := html.Parse(pageReader)
	if parseError != nil {
		return parseError
	}

	newParser(pageContent)
	return nil

}

func (f *Finder) CheckUrlIsVisited(url string) bool {
	for _, u := range f.visitedUrls {
		if url == u {
			return true
		}
	}
	return false
}

func (f *Finder) AppendToVisitedUrls(url string) {
	f.visitedUrls = append(f.visitedUrls, url)
}
