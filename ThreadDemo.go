package main

import (
	"fmt"
	"time"
)

func main() {
	go run("A", 2)
	run("B", 2)

}
func run(name string, times int) {
	for i := 1; i <= times; i++ {
		time.Sleep(1000)
		fmt.Println("线程", name)
	}
}
