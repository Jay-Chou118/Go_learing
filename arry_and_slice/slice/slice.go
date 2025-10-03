package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	s1 := slice[2:4]
	s2 := s1[1:3]

	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2)

	slice[1] = 100
	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2)

	s1 = append(s1, 200)
	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 300)
	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2)

}
