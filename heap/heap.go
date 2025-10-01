package main

import (
	"container/heap"
	"fmt"
)

// 定义一个最小堆类型
type IntHeap []int

// 实现 heap.Interface 接口的 5 个方法

// Len 返回堆中元素的数量
func (h IntHeap) Len() int {
	return len(h)
}

// Less 比较堆中两个元素的大小关系，决定顺序（最大堆）
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 交换堆中两个元素的位置
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 向堆中添加一个元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int)) // 将元素追加到堆的末尾
}

// Pop 从堆中移除并返回堆顶元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// 创建一个最小堆
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	// 向堆中添加元素
	heap.Push(h, 3)
	heap.Push(h, 4)

	// 打印堆顶元素，并弹出堆顶
	fmt.Println("Heap top:", (*h)[0])

	// 弹出堆顶元素
	fmt.Println("Pop:", heap.Pop(h))
	// fmt.Println("Pop h :", h.Pop())
	fmt.Println("Heap top after Pop:", (*h)[0])

	// 打印所有堆元素
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h)) // 每次 Pop 都会打印最小的元素
		// fmt.Println("h:", h.Pop())
	}
}
