package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)

	original := []int{1, 2, 3} // 长度3，容量3（无预留空间）
	appendWithResize(original)

	fmt.Println(original)      // 输出 [1 2 3] → 原切片完全不变
	fmt.Println(len(original)) // 输出 3 → 原切片长度不变（关键！）
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	          i++
	  }
	*/

	for i := range s {
		s[i] += 1
	}
}

// 向切片追加元素（触发扩容）
func appendWithResize(s []int) {
	s = append(s, 4) // 容量不足，触发扩容（创建新数组）
	s[0] = 100
	fmt.Println(s)      // 输出 [1 2 3] → 原切片完全不变
	fmt.Println(len(s)) // 输出 3 → 原切片长度不变（关键！）
}
