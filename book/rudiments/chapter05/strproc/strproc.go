package main

import "strings"

/**
字符串的链式处理
 */
func StringProcess(list []string, chain []func(string)string){
	for index,str := range list{

		for _,proc := range chain{
			str = proc(str)
		}
		list[index] = str
	}
}

func removePrefix(str string) string{
	return strings.TrimPrefix(str,"go")
}

func main() {
	list := []string{
		"go sacnner",
		"go phaser",
		"go compiler",
	}

	chain := []func(string)string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	//切片传递引用
	StringProcess(list, chain)

	for _,str := range list{
		println(str)
	}
}
