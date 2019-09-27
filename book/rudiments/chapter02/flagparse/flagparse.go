package main

import (
	"flag"
	"fmt"
)

func main() {
	var mode = flag.String("mode", "", "process mode")

	//解析命令行参数
	fmt.Println("------")
	flag.Parse()

	//输出命令行参数
	fmt.Println("======")
	fmt.Println(*mode)
}
