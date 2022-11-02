package main

import "fmt"

type Phone interface {
	call()
}

// 定义手机结构
type IPhone struct {
}

// 为手机结构添加方法
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 定义nokia手机结构
type NokiaPhone struct {
	birthday string
}

// 为nokia手机添加call() 方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!", nokiaPhone.birthday, "是我的生日")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
