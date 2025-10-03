package main

import (
	"fmt"
	"time"
)

func A(ch chan int) {
	// 定期向channel发送数据
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(2 * time.Second) // 每2秒发送一次
	}
}

func B(ch chan string) {
	// 定期向channel发送数据
	messages := []string{"hello", "world", "golang", "channel"}
	i := 0
	for {
		ch <- messages[i%len(messages)]
		i++
		time.Sleep(3 * time.Second) // 每3秒发送一次
	}
}

func main() {
	chint := make(chan int)
	chstring := make(chan string)

	go A(chint)
	go B(chstring)

	// 循环轮询select
	for {
		select {
		case v := <-chint:
			fmt.Println("Received from chint:", v)
		case v := <-chstring:
			fmt.Println("Received from chstring:", v)
		case <-time.After(5 * time.Second):
			// 超时处理，防止长时间没有数据时程序无响应
			fmt.Println("No data received for 5 seconds")
		}
	}
}
