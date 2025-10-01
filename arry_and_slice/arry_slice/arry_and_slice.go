package main

import "fmt"

func main() {

	a := make([]int, 0, 2)
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("i=%d len=%d cap=%d addr=%p\n", i, len(a), cap(a), &a[0])
	}

	base := []int{1, 2, 3, 4}
	s1 := base[:2]  // [1 2]
	s2 := base[1:3] // [2 3]
	fmt.Println("before:", s1, s2, base)
	s1[1] = 99
	fmt.Println("after :", s1, s2, base)
}
