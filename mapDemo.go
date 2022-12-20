package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}
	var intData int64
	intData = 1752714350055518
	var stintIN string
	stintIN = ToString(intData)
	fmt.Println("var:", stintIN)

	capital, ok := countryCapitalMap["American"]
	fmt.Println("capital:", capital)
	//if capital == nil {
	//	fmt.Println("capital is nil")
	//}
	fmt.Println("ok :", ok)
	fmt.Println(countryCapitalMap["India "])

}

func ToString(o interface{}) string {
	switch v := o.(type) {
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case json.Number:
		return v.String()
	case time.Time:
		return v.String()
	case time.Duration:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
