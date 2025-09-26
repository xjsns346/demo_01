// dup3, 使用os.Readfile来读取文件内容，并查找重复行。
// 原文中使用的ioutil包被废弃了，很多操作放在了os包里。
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("请在参数部分输入要读取的文件路径，以空格分隔。")
	if len(os.Args[1:]) < 1 { //通过比较长度，来判断是否输入文件路径。
		fmt.Println("未输入文件路径。")
		return
	}
	for _, file := range os.Args[1:] {
		counts := make(map[string]int) //声明counts 来存储数据。

		data, err := os.ReadFile(file) //返回的data为[]byte类型。
		if err != nil {
			fmt.Fprintf(os.Stderr, "报错 :%s", err)
		}

		for _, lines := range strings.Split(string(data), "\n") {

			counts[lines]++
		}
		for lines, n := range counts {
			if n > 1 {
				fmt.Printf("文件名 :%s\t，重复次数 :%d\t，重复内容 :%s\n", file, n, lines)
			}
		}

	}

}
