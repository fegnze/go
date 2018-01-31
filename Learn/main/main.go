package main

import (
	"fmt"
	"os"
	"time"

	"../http"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go http.Fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
