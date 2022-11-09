package main

import (
	"log"
	"time"
)

var ch chan struct{} = make(chan struct{})

func foo() {
	ch <- struct{}{}
	log.Println("foo() 111")
	time.Sleep(5 * time.Second)
	log.Println("foo() 222")
	close(ch)
	log.Println("foo() 333")
}

func main() {

	log.Println("main() 111")
	go foo()
	log.Println("main() 222")
	<-ch
	log.Println("main() 333")
}

/** 运行结果：
2022/11/09 16:10:49 main() 111
2022/11/09 16:10:49 main() 222
2022/11/09 16:10:49 foo() 111
2022/11/09 16:10:49 main() 333
*/
