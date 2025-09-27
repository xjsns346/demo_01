//fetch函数实现，来获取URL。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		//http.Get是发出一个http请求。
		res, err := http.Get(url) //返回的res 为*http.Response类型
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error :%s", err)
			os.Exit(1) //返回错误代码 : 1
		}
		content, err := io.ReadAll(res.Body)
		defer res.Body.Close() //defer 用来注册一个函数调用，等到当前函数返回之前才执行。
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error :%s", err)
			os.Exit(1) //返回错误代码 : 1
		}
		fmt.Printf("内容为%s", string(content))
	}

}
