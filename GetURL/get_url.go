//练习1.8函数实现，有隐患。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		//当传入的网址未添加协议名时，补上协议名。    存在隐患，当url的长度不足7时，会panic。
		if url[0:7] != "http://" {
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
	}

}
