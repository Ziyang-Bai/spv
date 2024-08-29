// SC:我只是一个Go的初学者，如有错误感谢指教
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 获取URL
	fmt.Print("请输入要访问的URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)
	if url == "" {
		fmt.Println("URL不能为空！")
		return
	}

	// 获取访问次数
	var count int
	for {
		fmt.Print("请输入要访问的次数: ")
		countInput, _ := reader.ReadString('\n')
		countInput = strings.TrimSpace(countInput)
		if countInput == "" {
			fmt.Println("访问次数不能为空！")
			continue
		}

		var err error
		count, err = strconv.Atoi(countInput)
		if err != nil {
			fmt.Println("请输入有效的整数！")
			continue
		}

		break
	}

	// 获取访问间隔时间
	var interval time.Duration
	for {
		fmt.Print("请输入每次访问之间的间隔时间（秒）: ")
		intervalInput, _ := reader.ReadString('\n')
		intervalInput = strings.TrimSpace(intervalInput)
		if intervalInput == "" {
			fmt.Println("访问间隔时间不能为空！")
			continue
		}

		interval, err := strconv.Atoi(intervalInput)
		if err != nil {
			fmt.Println("请输入有效的整数！")
			continue
		}

		break
	}

	// 获取线程数
	var threads int
	for {
		fmt.Print("请输入线程数: ")
		threadsInput, _ := reader.ReadString('\n')
		threadsInput = strings.TrimSpace(threadsInput)
		if threadsInput == "" {
			fmt.Println("线程数不能为空！")
			continue
		}

		var err error
		threads, err = strconv.Atoi(threadsInput)
		if err != nil || threads <= 0 {
			fmt.Println("请输入有效的正整数！")
			continue
		}

		break
	}

	// 创建一个channel来同步goroutines
	done := make(chan struct{})
	defer close(done)

	// 同步等待组
	var wg sync.WaitGroup

	// 开始访问
	fmt.Println("开始访问...")

	// 分配任务到goroutines
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error on request %d: %v\n", i+1, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading body for request %d: %v\n", i+1, err)
				return
			}

			fmt.Printf("Request %d: %s\n", i+1, string(body))

			// 在每个请求之间等待指定的时间
			time.Sleep(time.Duration(interval) * time.Second)
		}(i)
	}

	// 等待所有goroutines完成
	wg.Wait()

	// 关闭通道，通知goroutines结束
	close(done)

	fmt.Println("访问完成.")
}
