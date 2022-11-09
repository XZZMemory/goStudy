package main

import (
	"context"
	"fmt"
	"sync"
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
func (m *Miner) Stop(ctx context.Context) error {
	// 1.try to acquire lock
	m.lk.Lock()
	// define how to handle when exceptions occur
	defer m.lk.Unlock()
	// 3. initial
	m.stopping = make(chan struct{})
	stopping := m.stopping
	close(m.stop)
	select {
	case <-stopping:
		fmt.Println("stopping")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		return fmt.Errorf("default error")
	}
}
