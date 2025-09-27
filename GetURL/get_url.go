//练习1.9函数实现

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//使用strings.HasPrefix()函数，不会越界（自动处理长度问题）并考虑了 HTTPS 的情况。
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		//http.Get是发出一个http请求。
		res, err := http.Get(url) //返回的res 为*http.Response类型
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error :%s", err)
			os.Exit(1) //返回错误代码 : 1
		}
		//调用io.copy函数，将HTTP响应体拷贝到标准输出。
		n, err := io.Copy(os.Stdout, res.Body)
		defer res.Body.Close() //defer 用来注册一个函数调用，等到当前函数返回之前才执行。
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error :%s", err)
			os.Exit(1) //返回错误代码 : 1
		}
		fmt.Printf("%d的字符被复制。", n)
		//返回HTTP协议的状态码。
		fmt.Printf("状态码为: %v", res.Status)
	}

}
