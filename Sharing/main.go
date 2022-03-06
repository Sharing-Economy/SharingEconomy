package main

import (
	"Sharing/handler"
	"fmt"
)

func main() {
	// 初始化web服务器
	err := handler.Start(fmt.Sprintf("%s:%s", "127.0.0.1", "8080"),"static")
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}

}
