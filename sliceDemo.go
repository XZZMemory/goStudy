package main

import "fmt"

// 切片单测

func main() {
	// 声明切片，原始方法
	var mySlience []int
	if mySlience == nil {
		fmt.Println("该切片是nil ！")
	}
	// 声明切片
	var idendi []int = make([]int, 10)
	slice2 := make([]int, 10)
	slice2[9] = 1
	idendi[9] = 10
	fmt.Println("切片数据：", slice2[9])
}
