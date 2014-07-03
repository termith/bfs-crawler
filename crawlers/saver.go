package crawlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Saver struct {
	homeDir string
}

func NewSaver(dir string) *Saver {
	return &Saver{
		homeDir: dir,
	}
}

func (s *Saver) Run(urlChannel <-chan string, doneChannel chan<- string) {
	fmt.Println("Start saving...")
	for {
		if url := <-urlChannel; url == "done" {
			doneChannel <- "done"
			return
		} else {
			err := s.SavePageToDisk(url)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (s *Saver) SavePageToDisk(url string) error {
	page, getError := http.Get(url)
	if getError != nil {
		return getError
	}

	pageReader := page.Body
	file, readError := ioutil.ReadAll(pageReader)
	if readError != nil {
		return readError
	}

	pathToFile := s.homeDir + strings.Replace(url, "/", "_", -1)

	writeError := ioutil.WriteFile(pathToFile, file, 0777)
	if writeError != nil {
		return writeError
	}

	return nil
}
