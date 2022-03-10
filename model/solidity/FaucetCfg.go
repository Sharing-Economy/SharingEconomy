package solidity

//获取client
//func GetClient() (*ethclient.Client, error) {
//	client, err := ethclient.Dial("http://127.0.0.1:7545")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Close()
//	return client, nil
//}
//
////  获取合约对象
//func GetFacet(client *ethclient.Client) (*Faucet, error) {
//	contract, err := NewFaucet(common.HexToAddress("0x4E2eaD9e4269D57760Bdd2700f0659aA3D65663A"), client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("faucet:",contract)
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
//func GetPrivateKey() (*ecdsa.PrivateKey) {
//	Prikey := "a7cdc2e93dc1b17141b80238a70941c6ec91d8b64340b309827d663e62b7091e"
//	privateKey, err := crypto.HexToECDSA(Prikey)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("pk:",privateKey)
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
//		fmt.Println("gasPrice",gasPrice)
//		return gasPrice, nil
//	}
//
//}
////封装转账方法
//func SendAddresses(client *ethclient.Client,contract *Faucet ,Address common.Address) (*types.Transaction, error) {
//	opts :=Getopts()
//	opts.Value=big.NewInt(1000000000000000000)
//	res, err := contract.Withdraw(opts,Address)
//	fmt.Println("withdraw:",res)
//	opts.GasLimit=3000000
//	opts.GasPrice,err=GetgasPrice(client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//func GetBalance(contract *Faucet) (*big.Int, error) {
//
//	res, err := contract.GetBalance(nil)
//	if err != nil {
//		return big.NewInt(0), err
//	}
//	return res, nil
//}
