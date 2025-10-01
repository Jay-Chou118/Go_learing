package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个链表
	l := list.New()

	// 向链表中添加元素
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(0) // 将元素插入到链表头部

	// 输出链表中的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 删除一个元素（从链表中移除）
	l.Remove(l.Front()) // 删除链表的第一个元素

	// 输出链表中的元素
	fmt.Println("After removing front element:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
