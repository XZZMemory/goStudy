package main

import "fmt"

func main() {
	fmt.Println("结果", fi(5))

}

func fi(n int64) (result int64) {
	if n <= 0 {
		return 1
	}
	return n * fi(n-1)

}
