package main

import (
	"context"
	"fmt"
)

// 自定义类型，防止 key 冲突
type keyType string

const (
	requestIDKey keyType = "requestID"
	userIDKey    keyType = "userID"
)

func processRequest(ctx context.Context) {
	// 查找请求 ID
	requestID := ctx.Value(requestIDKey)
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("Request ID not found")
	}

	// 查找用户 ID
	userID := ctx.Value(userIDKey)
	if userID != nil {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found")
	}
}

func main() {
	// 创建父上下文
	parentCtx := context.Background()

	// 在父上下文中存储一个请求 ID
	parentCtx = context.WithValue(parentCtx, requestIDKey, "12345")

	// 创建第一个子上下文，存储用户 ID
	childCtx1 := context.WithValue(parentCtx, userIDKey, "user1")

	// 创建第二个子上下文，不存储用户 ID
	childCtx2 := context.WithValue(parentCtx, requestIDKey, "67890")

	// 传递子上下文1进行处理
	fmt.Println("Processing with childCtx1:")
	processRequest(childCtx1)

	// 传递子上下文2进行处理
	fmt.Println("\nProcessing with childCtx2:")
	processRequest(childCtx2)
}
