package main

import (
	"fmt"
	"os"
	"time"

	"gamecore/mynet/myhttp"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go myhttp.Fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("********************")
}
