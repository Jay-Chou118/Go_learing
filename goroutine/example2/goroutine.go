package main

import (
	"fmt"
	"time"
)

// 向管道发送数据的函数
func sendData(ch chan int, name string) {
	fmt.Printf("%s: 准备发送数据\n", name)
	ch <- 1 // 发送数据
	fmt.Printf("%s: 数据发送完成\n", name)
}

// 测试无缓冲管道
func testUnbufferedChannel() {
	fmt.Println("=== 测试无缓冲管道 ===")
	ch := make(chan int) // 无缓冲管道

	// 启动goroutine发送数据
	go sendData(ch, "发送者1")

	// 等待一段时间，让发送者先执行
	time.Sleep(1 * time.Second)
	fmt.Println("主程序: 准备接收数据")

	// 接收数据
	data := <-ch
	fmt.Printf("主程序: 接收到数据 %d\n", data)

	time.Sleep(1 * time.Second) // 等待goroutine完成
	fmt.Println()
}

// 测试有缓冲管道
func testBufferedChannel() {
	fmt.Println("=== 测试有缓冲管道 ===")
	ch := make(chan int, 1) // 有缓冲管道，缓冲区大小为1

	// 启动goroutine发送数据
	go sendData(ch, "发送者2")

	// 等待一段时间，观察发送者是否能完成
	time.Sleep(1 * time.Second)
	fmt.Println("主程序: 准备接收数据")

	// 接收数据
	data := <-ch
	fmt.Printf("主程序: 接收到数据 %d\n", data)

	time.Sleep(1 * time.Second) // 等待goroutine完成
}

func main() {
	testUnbufferedChannel()
	testBufferedChannel()
}
