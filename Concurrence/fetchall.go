//使用goroutine,并发获取多个URL,并输出耗时，字节数。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	//每一次循环都创建了一个goroutine，执行fetch函数
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a goroutine
	}
	//打印每一个goroutinr中ch的内容。
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	//输出多少秒过去了。
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// 实现发送请求，获取HTTP请求体，统计多少字节，以及耗费的时间。
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("打开%s的时候报错:%v", url, err)
		return //退出函数fetch
	}
	nBytes, err := io.Copy(io.Discard, resp.Body) //使用io.discard，将响应体的内容丢弃，返回复制的字节数以及err。
	defer resp.Body.Close()                       //关闭响应体。
	if err != nil {
		ch <- fmt.Sprintf("拷贝HTTP响应体的时候报错:%v", err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("耗时: %g s,字节数: %7d,URL: %s", seconds, nBytes, url)

}
