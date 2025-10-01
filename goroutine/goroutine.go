package main

import (
	"fmt"
	"time"
)

func goroutine() {
	fmt.Println("This is a goroutine example.")
	// ch <- struct{}{}

}
func main() {
	go goroutine()
	time.Sleep(2 * time.Second)
	fmt.Println("Main function")
}
