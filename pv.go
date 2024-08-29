package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func worker(url string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error during request:", err)
		return
	}
	defer resp.Body.Close()
	responseTime := time.Since(startTime).Seconds()
	fmt.Printf("Goroutine %d Response Time: %.2f seconds\n", getGoroutineID(), responseTime)
	fmt.Printf("Sleep", delay)
	time.Sleep(delay)
	fmt.Printf("OK!")
	
}

// A simple way to identify goroutine ID
var goroutineID int32

func getGoroutineID() int32 {
	goroutineID++
	return goroutineID
}

func main() {
	var wg sync.WaitGroup

	url := ""
	fmt.Print("请输入要访问的URL: ")
	fmt.Scanln(&url)

	var count int
	fmt.Print("请输入要访问的次数: ")
	fmt.Scanln(&count)

	var interval int
	fmt.Print("请输入每次访问之间的间隔时间（毫秒）: ")
	fmt.Scanln(&interval)

	var threads int
	fmt.Print("请输入线程数: ")
	fmt.Scanln(&threads)

	if threads <= 0 {
		fmt.Println("线程数必须为正整数！")
		return
	}

	delay := time.Duration(interval) * time.Millisecond

	for i := 0; i < threads; i++ {
		wg.Add(count)
		go worker(url, delay, &wg)
	}

	wg.Wait()
}
