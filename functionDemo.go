package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("结果", fi(5))

	//st := &studentInfo{
	//	id:   "1701850",
	//	name: "肖珍珍"}
	var name = 2
	interfac := interfaceDemo()
	// demo 接口转成结构体，同样类型才能转，否则报错
	// 1.同样类型，正常转换
	student := interfac.(*studentInfo)
	fmt.Println("接口转成结构体", student.id)
	fmt.Println("接口转成结构体", student.name)
	// 1.不同类型，转换失败
	//student2 := interfac.(*studentInfo2)
	//fmt.Println("接口转成结构体，不同类型报错", student2.name)

	fmt.Println("测试json工具", string(name))
	res, _ := Marshal(student)
	fmt.Println("测试json工具2", res)
	os.Stdout.Write(res)
	var myJson = `{"ad":{"b":1}, "c":7109625134243220263, "arr":[3,4,5],"arr_number":[1,2,3],"flag":true,"log_extra":{"origin_tpl_id":1}}`
	safeJson := NewSafeJson([]byte(myJson))
	fmt.Println(safeJson)

	fmt.Println("格式化后结果[" + strconv.FormatInt(1, 10) + "]")

}

func NewSafeJson(body []byte) *Json {
	js, err := NewJson(body)
	if err == nil {
		return js
	}

	defaultJson, _ := NewJson([]byte("{}"))
	return defaultJson
}

type Json struct {
	data interface{}
}

func NewJson(body []byte) (*Json, error) {
	j := new(Json)
	err := j.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
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
