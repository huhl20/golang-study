package main

import (
	"fmt"
)

type Invoker interface{
	Call(interface{})
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}){
	f(p)
}

func main() {
	var invoker Invoker
	//将匿名函数转为FuncCaller类型
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function",v)
	})

	invoker.Call("hello")
}
