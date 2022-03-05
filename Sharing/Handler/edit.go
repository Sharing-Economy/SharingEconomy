package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"Sharing/Config"
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
