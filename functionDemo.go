package main

import "fmt"

func main() {
	fmt.Println("结果", fi(5))

	interfac := interfaceDemo()
	// demo 接口转成结构体，同样类型才能转，否则报错
	// 1.同样类型，正常转换
	student := interfac.(*studentInfo)
	fmt.Println("接口转成结构体", student.id)
	fmt.Println("接口转成结构体", student.name)
	// 1.不同类型，转换失败
	student2 := interfac.(*studentInfo2)
	fmt.Println("接口转成结构体，不同类型报错", student2.name)
}

type studentInfo struct {
	id   string
	name string
}

type studentInfo2 struct {
	id   string
	name string
}

func fi(n int64) (result int64) {
	if n <= 0 {
		return 1
	}
	return n * fi(n-1)

}
func interfaceDemo() (req interface{}) {
	return &studentInfo{
		id:   "1701850",
		name: "肖珍珍"}

}
