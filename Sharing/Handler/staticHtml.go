package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户信息
func UserProfile(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		fmt.Println(err)
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	userName,people,integral,email,sign,goodsNum,balance,err:=config.GetUserMethod(contract,loginUser)
	fmt.Println("res",userName)
	fmt.Println("peo",people,integral)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/profile.html", gin.H{"userName": userName,"address": people,"integral": integral,"email":email,"sign":sign,"goodsNum":goodsNum,"balance":balance})
}

//导航栏渲染
func Navigate(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		fmt.Println(err)
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	userName,people,_,_,_,_,_,err:=config.GetUserMethod(contract,loginUser)
	fmt.Println("res",userName)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/index.html", gin.H{"userName": userName,"address": people})
}