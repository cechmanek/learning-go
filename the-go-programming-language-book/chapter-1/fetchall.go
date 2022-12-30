// Fetchall fetches URLs in parallel and reports their times and sizes.
// Call this as 'go run fetchall.go http://www.example.com http://gopl.io'
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f seconds %7d  %s", seconds, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string) // makes a channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive channel ch
	}

	fmt.Printf("%.2f seconds elapsed\n", time.Since(start).Seconds())
}
