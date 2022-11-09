package main

import (
	"fmt"
	"unsafe"
)

type Books struct {
	// 标题
	title   string
	author  string
	subject string
	book_id int
}

type BorrowLog struct {
	book     Books
	userName string
}

func main() {

	initialOfStruct()
	//
	sizeOfEmptyStruct()

}

// 一、结构体的三种初始化方法
func initialOfStruct() {

	// 1.第一种初始化方法
	var myBook Books
	myBook.book_id = 11000
	myBook.subject = "开发语言"

	// 2.第二种初始化方法
	myBorrowLog := BorrowLog{
		myBook,
		"xiaozhenzhen001",
	}
	// 3.第三种初始化方法
	relation := &BorrowLog{
		myBook,
		"xiaozhenzhen",
	}
	fmt.Println("作者：", myBook.author)
	fmt.Println("主题：", myBook.subject)

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(myBorrowLog.userName, "借了本书：", myBorrowLog.book)
	fmt.Println(relation.userName, "借了本书：", relation.book)
}

type complexS struct {
	A struct{}
	B struct{}
}

// 二、空结构体信息
func sizeOfEmptyStruct() {
	// 1.空结构体大小是0
	var s struct{}
	fmt.Println("空结构体大小是：", unsafe.Sizeof(s))
	// 2.空结构体组成的组合类型
	complex := complexS{}
	fmt.Println("空结构体组成的组合类型大小是", unsafe.Sizeof(complex))
}
