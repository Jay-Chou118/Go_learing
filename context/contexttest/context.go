package main

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context) {
	// 获取上下文中的请求 ID
	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Processing request with ID:", requestID)
	} else {
		fmt.Println("No request ID found")
	}
}

func main() {
	// 创建一个包含请求 ID 的上下文
	ctx := context.WithValue(context.Background(), "requestID", "12345")

	// 传递上下文给其他函数
	processRequest(ctx)
}
