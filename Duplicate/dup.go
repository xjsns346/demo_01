// 这是Go语言圣经中的练习1.4，增加了打印文件名的功能。
// 27行 ，避免了传入多个文件时，出现重复打印。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] //提取命令行中的参数
	if len(files) == 0 {
		fmt.Println("请输入内容，按下'ctrl + z + 回车'退出。")
		countLines(os.Stdin, counts) //如果没有传递函数名，则打开标准输入，os.Stdin也是一个文件指针类型。需要按下ctrl + z + 回车才可以退出os.Stdin

	} else {
		for _, file := range files {
			f, err := os.Open(file)
			//如果打开失败，则使用os.Fprintf,指定输出流为os.Stderr,打印错误信息。
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error :%v", err) //用于将错误信息输出到 标准错误流，区别于正常输出。适合日志或报错信息，保证和正常输出区分开。
				continue                                 //继续循环。
			}
			countLines(f, counts)
			counts = make(map[string]int) //清空counts中的键值对，从而避免读取文件内容重复。
			f.Close()                     //关闭文件。

		}
	}

}

// 声明了一个countLines函数，参数为一个文件指针，一个map类型的键值对集合，输出了count:重复行出现的次数，line:重复行的内容。
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, count := range counts {
		if count > 1 {
			//调用f.Name()方法来获取文件名。
			fmt.Printf("文件名：%s\t重复次数：%d\t重复内容：%s\n", f.Name(), count, line)
		}
	}

}
