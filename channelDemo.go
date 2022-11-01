package main

import "fmt"

// 定义一个函数，入参数组、通道
func sumOfArray(array []int, c chan int) {
	sum := 0
	for _, data := range array {
		sum += data
	}
	c <- sum

}
func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	fmt.Println("切片1", array[1:len(array)/2])
	fmt.Println("切片2", array[2:len(array)/2])
	fmt.Println("切片3", array[len(array)/2:])
	go sumOfArray(array[1:len(array)/2], c)
	go sumOfArray(array[2:len(array)/2], c)
	x, y := <-c, <-c
	fmt.Println("x:", x, "y", y)

}
