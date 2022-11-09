package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// 定义一个函数，入参数组、通道
func sumOfArray(array []int, c chan int) {
	sum := 0
	for _, data := range array {
		sum += data
	}
	c <- sum

}
func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	fmt.Println("切片1", array[1:len(array)/2])
	fmt.Println("切片2", array[2:len(array)/2])
	fmt.Println("切片3", array[len(array)/2:])
	go sumOfArray(array[1:len(array)/2], c)
	go sumOfArray(array[2:len(array)/2], c)
	x, y := <-c, <-c
	fmt.Println("x:", x, "y", y)
	var miner = Miner{}
	var po *Miner = &miner

	po.start(nil)
	po.Stop(nil)

	log.Println("main() 111")
	go foo2()
	log.Println("main() 222")
	<-ch1
	log.Println("main() 333")
	/**运行结果：
	2022/11/09 16:12:51 main() 111
	2022/11/09 16:12:51 main() 222
	2022/11/09 16:12:51 foo() 111
	thread.sleep()
	2022/11/09 16:12:56 foo() 222
	2022/11/09 16:12:56 foo() 333
	2022/11/09 16:12:56 main() 333

	*/

}

type Miner struct {
	lk       sync.Mutex
	stop     chan struct{}
	stopping chan struct{}
}

func (m *Miner) start(ctx context.Context) error {
	// 1.try to acquire lock
	m.lk.Lock()
	// 2.define how to handle when exceptions occur
	defer m.lk.Unlock()
	if m.stop != nil {
		fmt.Println("miner already start")
		return fmt.Errorf("miner already start")
	}

	m.stop = make(chan struct{})
	fmt.Println("miner stop")

	return nil
}

var ch1 chan struct{} = make(chan struct{})

func foo2() {

	log.Println("foo() 111")
	time.Sleep(5 * time.Second)
	log.Println("foo() 222")
	close(ch1)
	log.Println("foo() 333")

}
func (m *Miner) Stop(ctx context.Context) error {
	// 1.try to acquire lock
	m.lk.Lock()
	// define how to handle when exceptions occur
	defer m.lk.Unlock()
	// 3. initial
	m.stopping = make(chan struct{})
	stopping := m.stopping
	close(m.stop)
	//var a <- stopping
	select {
	case <-stopping:
		fmt.Println("stopping")
		return nil
	//case <-ctx.Done():
	//	return ctx.Err()
	default:
		fmt.Println("default error")
		return fmt.Errorf("default error")
	}
}
