package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("A")
		ch <- 1
	}()

	go func() {
		fmt.Println("B")
		ch <- 2
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
