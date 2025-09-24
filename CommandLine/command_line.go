package main

import (
	"fmt"
	"os"
)

func main() {
	for i, str := range os.Args {
		fmt.Println(i, str)
	}

}
