package main

import "fmt"

func main() {
	// 定义一个变量
	var a int = 10
	// 声明一个指针
	var ip *int = &a
	fmt.Println("a的值是：", a)
	fmt.Println("指针指向的地址是：", ip)
	fmt.Println("指针指向值是：", *ip)
	var ptr *int
	fmt.Println("指针指向值是：", ptr)
	if ptr == nil {
		fmt.Println("【ptr】空指针！")
	} else {
		fmt.Println("【ptr】的值是", &ptr)

	}
	fmt.Println("before... a的值是：", a)
	var pointer *int = &a
	pointerusage(pointer)
	fmt.Println("after... a的值是：", a)

}
func pointerusage(pointer *int) {
	*pointer = *pointer + 1
}
