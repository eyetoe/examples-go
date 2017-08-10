package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(tasksCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for {
		url, ok := <-tasksCh
		if !ok {
			return
		}
		//conn, err := net.Dial("tcp", "google.com:80")
		conn, err := net.Dial("tcp", url+":80")
		if err != nil {
			fmt.Println("  Problem connecting to:", url)
			fmt.Println("  error was:", err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		elapsed := time.Since(start)

		if err != nil {
			fmt.Println(url, "Status:", status)
		} else {
			fmt.Println(url, "		", elapsed)
		}

	}
}

func pool(wg *sync.WaitGroup, workers int) {
	tasksCh := make(chan string)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	all := []string{
		"google.com",
		"youtube.com",
		"amazon.com",
		"wikipedia.org",
		"google.co.in",
		"github.com",
		"xonk.org",
		"openwager.com",
	}

	for i := 0; i < len(all); i++ {
		tasksCh <- all[i]
	}

	close(tasksCh)
}

func main() {
	var wg sync.WaitGroup

	nWorkers := 10

	wg.Add(nWorkers)
	go pool(&wg, nWorkers)
	wg.Wait()
}
