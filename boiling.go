package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const boiling = 212.0

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	var f = boiling
	//var c = (f - 32) * 5 / 9
	fmt.Printf("%g or %g \n", f, fToc(f))
	//fmt.Println(c)
	//var g, h, j int
	var b = "hello world"
	fmt.Println(b)

}

// 函数:温度转换
func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
