package main

import "fmt"

type Books struct {
	// 标题
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var myBook Books
	myBook.book_id = 11000
	myBook.subject = "开发语言"
	fmt.Println("作者：", myBook.author)
	fmt.Println("主题：", myBook.subject)

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

}
