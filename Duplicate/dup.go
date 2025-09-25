//这是Go语言圣经中的dup2，将输出代码放置在了countLines函数中。

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
		countLines(os.Stdin, counts) //如果没有传递函数名，则打开标准输入，os.Stdin也是一个文件指针类型。

	} else {
		for _, file := range files {
			f, err := os.Open(file)
			//如果打开失败，则使用os.Fprintf,指定输出流为os.Stderr,打印错误信息。
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error :%v", err) //用于将错误信息输出到 标准错误流，区别于正常输出。适合日志或报错信息，保证和正常输出区分开。
				continue                                 //继续循环。
			}
			countLines(f, counts)
			f.Close() //关闭文件。

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

			fmt.Printf("%d\t%s\n", count, line)
		}
	}

}
