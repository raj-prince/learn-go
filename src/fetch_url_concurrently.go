package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Command to execute: go build fetch_url_concurrently.go
// ./fetch_url_concurrently https://golang.org http://gopl.io https://godoc.org

// Output
// 1.02s 59766 https://golang.org
// 1.53s 17461 https://godoc.org
// 1.92s 4154 http://gopl.io
// 1.92s elapsed
func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // invokes a goroutine, main is also a goroutine(parent)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", secs, nBytes, url)
}
