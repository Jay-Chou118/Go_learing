package main

import "fmt"

// 开启100个协程，循环打印1-1000，尾号为1的协程打印尾号1的结果
func main() {
	// ch := make(chan int)
	for i := 1; i <= 100; i++ {
		go func(id int) {
			for j := 1; j <= 10; j++ {
				num := id + (j-1)*100
				fmt.Println("协程", id, "打印:", num)
			}
		}(i)

	}

}
