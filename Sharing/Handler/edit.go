package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

func addSticker(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}
  //初始化opts
	opts :=config.Getopts()

	//初始化gas
	gasPrice,err :=config.GetgasPrice(client)
	opts.GasLimit = 3000000
	opts.GasPrice = gasPrice

	stick :=c.PostForm("stick")
	val,err :=contract.AddSticker(opts,stick)
	fmt.Println(val)
	respOK(c, val)

}

func isStickExist(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	stick :=c.PostForm("stick")
	val,err :=contract.IsStickExist(nil,stick)
	fmt.Println(val)
	respOK(c, val)
}

func upload(c *gin.Context) {
	f, err := c.FormFile("files")
	fmt.Println(f)
	if err != nil {
		fmt.Println(err)
		respError(c,err)
		return
	} else {
		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			respOK(c, "OK！")
			return
		}
		fileName:=f.Filename
		fildDir:="./Static/"

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