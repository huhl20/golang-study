package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	//1.创建白色底图
	const size = 300

	pic := image.NewGray(image.Rect(0,0,size,size))

	for x :=0; x < size; x++ {
		for y:=0; y< size; y++ {
			//255代表白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//2.绘制正弦函数轨迹
	for x:=0; x <size; x++  {
		s := float64(x)*2*math.Pi/size
		y := size/2-math.Sin(s)*size/2
		//0代表黑色
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//3.保存到文件
	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(file, pic)
	}

	png.Encode(file, pic)

	file.Close()
}
