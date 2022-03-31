package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strings"
)
//实现转账
//func SendContract(c *gin.Context)  {
//	faucet:= db.Faucet_huangcuihua{}
//	//获取post中返回的address
//	faucet.Address = c.PostForm("address")
//	fmt.Println("address:::",faucet.Address)
//
//	client,err := solidity.GetClient()
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	contract ,err:= solidity.GetFacet(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	data,err:=solidity.SendAddresses(client,contract,common.HexToAddress(faucet.Address))
//	fmt.Println("sendaddress:",data)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	balance,err:=solidity.GetBalance(contract)
//	fmt.Println("balance:",balance)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	//实现添加记录
//	err =db.Insert(faucet.Address,1)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//	//更新表中记录
//	err =db.Update()
//	if err != nil{
//		respError(c,err)
//		return
//	}
//	respOK(c,data,balance)
//	client.Close()
//}
func Upload(c *gin.Context){
	f, err := c.FormFile("imgname")
	//file, handler, err := c.FormFile("imgname")
	fmt.Println(err)
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
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
		fmt.Println("filename",fileName)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result":gin.H{
				"path":filepath,
			},
		})
	}
}





