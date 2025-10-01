package main

import "fmt"

// RingBuffer 队列结构体
type RingBuffer[T any] struct {
	buf  []T // 底层数组，用来存储数据
	head int // 队头索引
	tail int // 队尾索引
	size int // 当前队列中元素的个数
	cap  int // 队列的容量
}

// 创建一个新的 RingBuffer
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, capacity),
		cap: capacity,
	}
}

// Enqueue 向队列中添加一个元素
func (rb *RingBuffer[T]) Enqueue(value T) {
	if rb.size == rb.cap {
		// 队列已满，覆盖最旧的元素
		rb.head = (rb.head + 1) % rb.cap
	} else {
		rb.size++
	}
	rb.buf[rb.tail] = value
	rb.tail = (rb.tail + 1) % rb.cap
}

// Dequeue 从队列中移除并返回一个元素
func (rb *RingBuffer[T]) Dequeue() (T, bool) {
	if rb.size == 0 {
		var zero T
		return zero, false // 空队列，返回默认值
	}
	value := rb.buf[rb.head]
	rb.head = (rb.head + 1) % rb.cap
	rb.size--
	return value, true
}

// Len 返回队列当前的长度
func (rb *RingBuffer[T]) Len() int {
	return rb.size
}

// Cap 返回队列的容量
func (rb *RingBuffer[T]) Cap() int {
	return rb.cap
}

// 显示当前队列的内容（仅用于调试）
func (rb *RingBuffer[T]) Display() {
	fmt.Printf("RingBuffer: head=%d tail=%d size=%d cap=%d\n", rb.head, rb.tail, rb.size, rb.cap)
	fmt.Println("Elements:", rb.buf)
}

func main() {
	rb := NewRingBuffer[int](5) // 创建一个容量为 5 的环形队列

	// 入队操作
	for i := 0; i < 6; i++ { // 插入 6 个元素，观察覆盖行为
		rb.Enqueue(i)
		rb.Display()
	}

	// 出队操作
	for i := 0; i < 3; i++ {
		value, ok := rb.Dequeue()
		if ok {
			fmt.Printf("Dequeued: %d\n", value)
		} else {
			fmt.Println("Queue is empty!")
		}
		rb.Display()
	}
}
