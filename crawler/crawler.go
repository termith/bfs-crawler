package crawler

import (
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
	page, getError := http.Get(url)
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

func (c *Crawler) FindAllUrls(url string) *[]string {
	return new([]string)
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
