package main

import (
	"fmt"
	"ShareFish/cfg"
	"ShareFish/db"
	"ShareFish/handler"
)

func main() {
	//初始化数据库
	err :=db.InitDB()
	if err != nil {
		fmt.Printf("数据库服务错误:%v\n", err)
		return
	}

	// 载入配置文件
	c, err := cfg.LoadConfig("cfg.json")
	if err != nil {
		fmt.Printf("载入配置文件错误:%v\n", err)
		return
	}
	fmt.Println(*c)

	//初始化web服务器
	err = handler.Server(fmt.Sprintf("%s:%s", c.Host, c.Port), c.WebDir)
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}

}
