package crawler

import (
	html "code.google.com/p/go.net/html"
	queue "github.com/termith/bfs-crawler/queue"
	"io/ioutil"
	"net/http"
	"strings"
)

// Tools for crawler

type Crawler struct {
	isRun       bool
	visitedUrls []string
	homeDir     string
}

func NewCrawler() *Crawler {
	return &Crawler{
		isRun:       false,
		visitedUrls: make([]string, 0),
		homeDir:     "/home/ddemidov/bfs/",
	}
}

func (c *Crawler) IsRun() bool {
	return c.isRun
}

func (c *Crawler) SavePageToDisk(url string) (int, error) {
	page, getError := http.Get(url) // Do we need get page content two times - here and in FindAllUrls
	if getError != nil {
		return 1, getError
	}

	pageReader := page.Body
	file, readError := ioutil.ReadAll(pageReader)
	if readError != nil {
		return 1, readError
	}

	pathToFile := c.homeDir + strings.Replace(url, "/", "_", -1)

	writeError := ioutil.WriteFile(pathToFile, file, 0777)
	if writeError != nil {
		return 1, writeError
	}

	return 0, nil
}

func (c *Crawler) FindAllUrls(url queue.Url, urlQueue *queue.UrlQueue) error {

	/*TODO: I have to draw with function because I don't understand how it works
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

func (c *Crawler) CheckUrlIsVisited(url string) bool {
	for _, u := range c.visitedUrls {
		if url == u {
			return true
		}
	}
	return false
}

func (c *Crawler) AppendToVisitedUrls(url string) {
	c.visitedUrls = append(c.visitedUrls, url)
}
