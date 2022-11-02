package main

import (
	"encoding/json"
	"fmt"
)

// https://old-panda.com/2019/12/11/golang-omitempty/
type address struct {
	// struct 字段：Street；json字段 street
	Street  string `json:"street"`          // 街道
	Ste     string `json:"suite,omitempty"` // 单元（可以不存在）
	City    string `json:"city"`            // 城市
	State   string `json:"state"`           // 州/省
	Zipcode string `json:"zipcode"`         // 邮编
}

func main() {
	data := `{
		"street": "200 Larkin St",
		"city": "San Francisco",
		"state": "CA",
		"zipcode": "94102"
	}`
	addr := new(address)
	json.Unmarshal([]byte(data), &addr)
	addressBytes, _ := json.MarshalIndent(addr, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}
