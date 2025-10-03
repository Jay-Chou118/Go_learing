package main

import (
	"fmt"
	"sync"
	"time"
)

func A(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("A:", id)
}

func B(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 2
	fmt.Println("B: 已发送数据")
}

func main() {
	fmt.Println("Hello, World!")
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) // 我们有两个goroutine需要等待
	go B(ch, &wg)
	time.Sleep(1 * time.Second) // 确保B先运行
	go A(<-ch, &wg)

	wg.Wait() // 等待所有goroutine完成
}
