// echo 输出其命令行参数
package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strconv"
)

func main() {
	var s string
	var sep = "89"
	for i := 1; i < 10; i++ {
		s += sep + strconv.Itoa(i) + " "

	}
	fmt.Println(&sep)
	fmt.Println("Hello, world")
	var myArray = [10]int{101, 10, 11}
	//myArray[0]=10
	fmt.Println(myArray[0])
	fmt.Println(myArray[9])
	/*	// 出现次数大于1 的行
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			//counts[len(input.Text())]++*/
	//}
}
