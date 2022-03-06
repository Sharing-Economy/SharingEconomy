package handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"ShareFish/db"
	"ShareFish/solidity"
)
//实现转账
func SendContract(c *gin.Context)  {
	faucet:= db.Faucet_huangcuihua{}
	//获取post中返回的addresso
	faucet.Address = c.PostForm("address")
	fmt.Println("address:::",faucet.Address)

	client,err := solidity.GetClient()
	if err != nil{
		respError(c,err)
		return
	}

	contract ,err:= solidity.GetFacet(client)
	if err != nil{
		respError(c,err)
		return
	}

	data,err:=solidity.SendAddresses(client,contract,common.HexToAddress(faucet.Address))
	fmt.Println("sendaddress:",data)
	if err != nil{
		respError(c,err)
		return
	}

	balance,err:=solidity.GetBalance(contract)
	fmt.Println("balance:",balance)
	if err != nil{
		respError(c,err)
		return
	}

	//实现添加记录
	err =db.Insert(faucet.Address,1)
	if err != nil{
		respError(c,err)
		return
	}
	//更新表中记录
	err =db.Update()
	if err != nil{
		respError(c,err)
		return
	}
	respOK(c,data,balance)
	client.Close()
}


