package crawler

import (
	"net/http"
)

// Tools for crawler

type Crawler struct {
	isRun bool
}

func NewCrawler() {
	return &Crawler
}

func (c *Crawler) IsRun() {
	return c.isRun
}

func (c *Crawler) SavePageToDisk() int {
	return 0
}

func (c *Crawler) FindAllUrls(url) []string {
	return new([]string)
}

func (c *Crawler) CheckUrlIsVisited(url) bool {
	return true
}
