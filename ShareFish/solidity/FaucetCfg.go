package solidity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)
//获取client
func GetClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	return client, nil
}

//  获取合约对象
func GetShareFish(client *ethclient.Client) (*Solidity, error) {
	contract, err := NewSolidity(common.HexToAddress("0x317c92Bb31a1c8f5678271F0A33D94Ef3d4d6bbb"), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("faucet:",contract)
	return contract, err
}

//获取Opts对象
func Getopts() *bind.TransactOpts {
	privateKey := GetPrivateKey()
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}

//获取私钥
func GetPrivateKey() (*ecdsa.PrivateKey) {
	Prikey := "9e64387a398fa1a813e2a8614cd2ebd04751755d1c2046cb0cecf0498a78591f"
	privateKey, err := crypto.HexToECDSA(Prikey)
	if err != nil {
		panic(err)
	}
	fmt.Println("pk:",privateKey)
	return privateKey
}

// 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return big.NewInt(0), err
	} else {
		fmt.Println("gasPrice",gasPrice)
		return gasPrice, nil
	}

}
//封装转账方法
func SendAddresses(client *ethclient.Client,contract *Solidity ,Address common.Address) (*types.Transaction, error) {
	opts :=Getopts() 
	opts.Value=big.NewInt(1000000000000000000)
	res, err := contract.Withdraw(opts,Address)
	fmt.Println("withdraw:",res)
	opts.GasLimit=3000000
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetBalance(contract *Solidity) (*big.Int, error) {

	res, err := contract.GetBalance(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return res, nil
}
