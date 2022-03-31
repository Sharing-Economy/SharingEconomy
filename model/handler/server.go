package handler

import (
	"github.com/gin-gonic/gin"
)

func Server(addr,webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()

	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器
		r.Static("/", webDir)
	}
	//设置路由组
	router := r.Group("")
	{
		router.POST("/index", Upload)
	}

	r.Run(addr)
	return err
}
// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{},data2 interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
		"data2":data2,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
