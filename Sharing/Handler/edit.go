package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

//func addSticker(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		respError(c,err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//   //初始化opts
//	opts :=config.Getopts()
//
//	//初始化gas
//	gasPrice,err :=config.GetgasPrice(client)
//	opts.GasLimit = 3000000
//	opts.GasPrice = gasPrice
//
//	stick :=c.PostForm("stick")
//	val,err :=contract.AddSticker(opts,stick)
//	fmt.Println(val)
//	respOK(c, val)
//
//}

//func isStickExist(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		respError(c,err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	stick :=c.PostForm("stick")
//	val,err :=contract.IsStickExist(nil,stick)
//	fmt.Println(val)
//	respOK(c, val)
//}
//func Upload(c *gin.Context) {
//	headers, err := c.FormFile("file")
//
//	if err != nil {
//		log.Printf("Error when try to get file: %v", err)
//	}
//	//headers.Size 获取文件大小
//	if headers.Size > 1024*1024*2 {
//		fmt.Println("文件太大了")
//		return
//	}
//	//headers.Header.Get("Content-Type")获取上传文件的类型
//	if headers.Header.Get("Content-Type") != "image/png" {
//		fmt.Println("只允许上传png图片")
//		return
//	}
//	c.SaveUploadedFile(headers, "./static/"+headers.Filename)
//	c.String(http.StatusOK, headers.Filename)
//}
func Upload(c *gin.Context){
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
}
