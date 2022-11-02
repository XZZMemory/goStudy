package main

import "fmt"

// 改变地址a指向的值
func change(a *int) {
	sourceNum := *a
	*a = sourceNum + 1
}
func main() {
	var data int = 10
	var address *int = &data
	fmt.Println("原始数据：", data, " 地址是:", address)
	change(address)
	fmt.Println("改变或数据：", data, " 地址是:", address)
}
