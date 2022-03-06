package config
//
//import (
//	"context"
//	"crypto/ecdsa"
//	"fmt"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"log"
//	"math/big"
//
//	"Sharing/Agreement"
//)
//
////获取client
//func GetClient() (*ethclient.Client, error) {
//	client, err := ethclient.Dial("http://127.0.0.1:7545")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Close()
//	return client, nil
//}
//
//// 获取合约对象
//func GetAddress(client *ethclient.Client) (*Agreement.User, error) {
//	contract, err := Agreement.NewUser(common.HexToAddress("0x9d63a7EaeFefBc117b22Aa863Df21112B4e80119"), client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("faucet:", contract)
//	return contract, err
//}
//
////获取Opts对象
//func Getopts() *bind.TransactOpts {
//	privateKey := GetPrivateKey()
//	opts := bind.NewKeyedTransactor(privateKey)
//	fmt.Println("opts:", opts)
//	return opts
//}
//
////获取私钥
//func GetPrivateKey() *ecdsa.PrivateKey {
//	Prikey := "3b80f1a89549326722a715bd3af3c2240d97c05b22e1f1416c9c51fa78d32b1a"
//	privateKey, err := crypto.HexToECDSA(Prikey)
//	if err != nil {
//		panic(err)
//	}
//	return privateKey
//}
//
//// 获取 gasPrice
//func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//
//	if err != nil {
//		return big.NewInt(0), err
//	} else {
//		fmt.Println("gasPrice", gasPrice)
//		return gasPrice, nil
//	}
//
//}
