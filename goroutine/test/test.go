package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Hello, World!", <-ch)
	ch <- 4 // 会报错，阻塞

}
