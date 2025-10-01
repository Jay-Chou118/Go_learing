package main

import (
	"container/list"
	"fmt"
)

// LRU 缓存结构
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
}

// 定义链表节点
type Entry struct {
	key   string
	value int
}

// 创建一个新的 LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get 从缓存中获取值
func (l *LRUCache) Get(key string) int {
	if ele, ok := l.cache[key]; ok {
		// 移动到链表头部
		l.order.MoveToFront(ele)
		return ele.Value.(*Entry).value
	}
	return -1 // 不存在
}

// Put 向缓存中添加一个值
func (l *LRUCache) Put(key string, value int) {
	if ele, ok := l.cache[key]; ok {
		// 更新值，并移动到链表头部
		ele.Value.(*Entry).value = value
		l.order.MoveToFront(ele)
	} else {
		// 插入新节点
		if l.order.Len() == l.capacity {
			// 删除最旧的元素（尾部元素）
			tail := l.order.Back()
			l.order.Remove(tail)
			delete(l.cache, tail.Value.(*Entry).key)
		}
		entry := &Entry{key, value}
		ele := l.order.PushFront(entry)
		l.cache[key] = ele
	}
}

// 打印缓存状态
func (l *LRUCache) PrintCache() {
	for e := l.order.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s: %d\n", e.Value.(*Entry).key, e.Value.(*Entry).value)
	}
}

func main() {
	cache := NewLRUCache(3)

	// 向缓存中添加数据
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// 打印缓存
	cache.PrintCache()

	// 获取数据
	fmt.Println("Get a:", cache.Get("a"))

	// 向缓存中添加数据，导致最旧的 "b" 被移除
	cache.Put("d", 4)
	cache.PrintCache()

	// 获取数据
	fmt.Println("Get b:", cache.Get("b")) // 输出 -1，因为 "b" 被移除
}
