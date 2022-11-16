package main

import "fmt"

type iphone struct {
	operatorSystem string //操作系统
	birthday       string // 出厂日期

}

func main() {
	var countryCapital1 map[string]string
	countryCapital1 = make(map[string]string)
	countryCapital1["河南"] = "郑州"

	var countryCapital2 = new(map[string]string)
	//(*countryCapital2)["河南"] = "郑州"

	var name = 1
	fmt.Println(name, countryCapital2)
}
