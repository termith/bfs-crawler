package crawler

import (
	html "code.google.com/p/go.net/html"
	"io/ioutil"
	"net/http"
)

// Tools for crawler

type Crawler struct {
	isRun       bool
	visitedUrls []string
}

func NewCrawler() *Crawler {
	return &Crawler{
		isRun:       false,
		visitedUrls: make([]string, 10),
	}
}

func (c *Crawler) IsRun() bool {
	return c.isRun
}

func (c *Crawler) SavePageToDisk(url string, pathToFile string) (int, error) {
	page, getError := http.Get(url) // Do we need get page content two times - here and in FindAllUrls
	if getError != nil {
		return 1, getError
	}

	pageReader := page.Body
	file, readError := ioutil.ReadAll(pageReader)
	if readError != nil {
		return 1, readError
	}

	writeError := ioutil.WriteFile(pathToFile+url, file, 0777)
	if writeError != nil {
		return 1, writeError
	}

	return 0, nil
}

func (c *Crawler) FindAllUrls(url string) ([]string, error) {
	findedUrls := make([]string, 10)

	/*TODO: I have to draw with function because I don't understand how it works
	  And may be it can be done simple... */

	var newParser func(*html.Node)
	newParser = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					findedUrls = append(findedUrls, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			newParser(c)
		}
	}

	page, getError := http.Get(url)
	if getError != nil {
		return nil, getError
	}

	pageReader := page.Body

	pageContent, parseError := html.Parse(pageReader)
	if parseError != nil {
		return nil, parseError
	}

	newParser(pageContent)
	return findedUrls, nil

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
