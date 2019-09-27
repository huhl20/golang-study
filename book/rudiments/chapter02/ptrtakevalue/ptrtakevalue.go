package main

import "fmt"

func main() {
	var house= "Malibu Point 10880, 90265"

	ptr := &house
	//打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	//打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	value := *ptr
	//取值后的值
	fmt.Printf("value type: %T\n", value)
	//指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	/*
	ptr type: *string
	address: 0xc0000421f0
	value type: string
	value: Malibu Point 10880, 90265
	 */
}
