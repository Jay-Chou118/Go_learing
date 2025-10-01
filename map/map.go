package main

import (
	"fmt"
	"sync"
)

type HashMap struct {
	m map[string]int
	//容量
	cap   int
	mutex sync.RWMutex
}

// 初始化 HashMap
func NewHashMap(initialCap int) *HashMap {
	if initialCap <= 0 {
		initialCap = 16 // 默认初始容量
	}
	return &HashMap{m: make(map[string]int), cap: initialCap}
}

// Put 向 HashMap 添加一个键值对
func (h *HashMap) Put(key string, value int) {

	h.mutex.Lock()
	defer h.mutex.Unlock()

	//检查是否超过容量的70%
	if len(h.m) > h.cap*7/10 {
		h.Resize(h.cap * 2)
	}
	h.m[key] = value
}

// Get 查找 HashMap 中的值
func (h *HashMap) Get(key string) (int, bool) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	value, ok := h.m[key]
	return value, ok
}

// Remove 从 HashMap 删除一个键值对
func (h *HashMap) Remove(key string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	delete(h.m, key)
}

// 扩容,当元素超过容量的70%时调用
func (h *HashMap) Resize(newCapacity int) {
	if newCapacity <= h.cap {
		return // 新容量必须大于当前容量
	}

	newMap := make(map[string]int, newCapacity)
	for k, v := range h.m {
		newMap[k] = v
	}
	h.m = newMap        // 替换为新 map
	h.cap = newCapacity // 更新容量
}

// Len 返回当前元素数量
func (h *HashMap) Len() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.m)
}

func (h *HashMap) Getcap() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.cap

}

func main() {
	h := NewHashMap(4)

	// 向 HashMap 添加数据
	// h.Put("apple", 5)
	// h.Put("banana", 3)
	// h.Put("cherry", 7)
	// h.Put("date", 2)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			h.Put(key, id)
			fmt.Printf("Goroutine %d: Put %s = %d\n", id, key, id)
		}(i)
	}
	wg.Wait()

	fmt.Println("HashMap size after adding elements:", h.Len())
	fmt.Println("Current capacity:", h.cap)

	// 演示并发读取
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if val, ok := h.Get(key); ok {
				fmt.Printf("Goroutine %d: Get %s = %d\n", id, key, val)
			} else {
				fmt.Printf("Goroutine %d: %s not found\n", id, key)
			}
		}(i)
	}
	wg.Wait()

	// 查找数据
	if value, ok := h.Get("apple"); ok {
		fmt.Println("apple:", value)
	}

	// 删除数据
	h.Remove("banana")

	// 查找删除后的数据
	if _, ok := h.Get("banana"); !ok {
		fmt.Println("banana not found")
	}
}
