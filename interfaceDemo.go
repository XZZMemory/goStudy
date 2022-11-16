package main

import "fmt"

type Phone interface {
	call()
}

// 定义手机结构
type IPhone struct {
	price int
}

// 为手机结构添加方法
func (iphone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!", iphone.price)
}

// 定义nokia手机结构
type NokiaPhone struct {
	birthday string
}
type MyStruct struct {
	NokiaPhone
}

// 为nokia手机添加call() 方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!", nokiaPhone.birthday, "是我的生日")
}

func main() {
	var phone Phone
	//
	phone = new(NokiaPhone)
	phone.call()

	var sss MyStruct = MyStruct{}

	fmt.Println(sss.birthday)
	fmt.Println(sss.call)
	phone = IPhone{price: 1}
	//phone.price=
	phone.call()

}
