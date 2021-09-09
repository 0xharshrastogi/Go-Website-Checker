package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://google.com", "http://github.com", "http://golang.org", "http://stackoverflow.com", "http://example.com"}

	c := make(chan string)

	for _, url := range urls {
		go IsUrlActive(url, c)
	}

	for url := range c {
		go func (url string) {
				time.Sleep(3 * time.Second)
				IsUrlActive(url, c)	
		}(url)
	}
}

func IsUrlActive(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "Is Down")
	}

	fmt.Println(url, "Is Up")
	c <- url
}
