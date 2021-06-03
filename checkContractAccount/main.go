package main

import (
"fmt"
"github.com/ethereum/go-ethereum/rpc"
)

var getCode string

func main()  {
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer client.Close()
	GetCode(client, "0x72109962Ff76156F979b1bA7f9472359f0d49cC2")
}

func GetCode(client *rpc.Client, address string)  {
	err := client.Call(&getCode, "eth_getCode", address, "latest")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(getCode)
	if getCode == "0x" {
		fmt.Println(address, "是外部账户")
	} else {
		fmt.Println(address, "是合约账户")
	}
}