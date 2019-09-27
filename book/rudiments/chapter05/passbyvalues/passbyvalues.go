package main

import "fmt"

/**
函数中的参数传递效果测试
 */
type Data struct{
	complex []int
	instance InnerDate
	ptr *InnerDate
}

type InnerDate struct{
	a int
}

func passByValue(inFunc Data) Data{
	fmt.Printf("inFunc value: %+v\n", inFunc)
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}

func main() {
	in := Data{
		complex:  []int{1,2,3},
		instance: InnerDate{
			5,
		},
		ptr:      &InnerDate{a:1},
	}
	//输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	//fmt.Printf("in value: %v\n", in)//不含属性名
	//输入结构的指针地址
	fmt.Printf("in value: %p\n", &in)

	out := passByValue(in)

	//输出结构的成员情况
	fmt.Printf("in value: %+v\n", out)
	//输出结构的指针地址
	fmt.Printf("in value: %p\n", &out)
}
