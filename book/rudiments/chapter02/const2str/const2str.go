package main

import "fmt"

//声明芯片乐星
type ChipType int
const(
	//const declaration. It is zero-indexed.
	//const iota = 0 // Untyped int.
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Printf("%s %d",CPU, CPU)
}
