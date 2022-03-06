package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
			f, err := c.FormFile("file")
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "上传失败!",
				})
				return
			} else {

				fileExt:=strings.ToLower(path.Ext(f.Filename))
				if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
					c.JSON(200, gin.H{
						"code": 400,
						"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
					})
					return
				}
				fileName:=f.Filename
				fildDir:="./static/"

				filepath:=fmt.Sprintf("%s%s",fildDir,fileName)
				c.SaveUploadedFile(f, filepath)
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "上传成功!",
					"result":gin.H{
						"path":filepath,
					},
				})
			}
	})
	//默认端口号是8080
	r.Run(":8080")
}
