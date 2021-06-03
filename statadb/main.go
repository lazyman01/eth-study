package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func main()  {
	statadb, _ := state.New(common.Hash{},
		state.NewDatabase(rawdb.NewMemoryDatabase()), nil)// ❶

	acct1:=common.HexToAddress("0x0bB141C2F7d4d12B1D27E62F86254e6ccEd5FF9a")// ❷
	acct2:=common.HexToAddress("0x77de172A492C40217e48Ebb7EEFf9b2d7dF8151B")

	statadb.AddBalance(acct1,big.NewInt(100))
	statadb.AddBalance(acct2,big.NewInt(888))

	contract:=crypto.CreateAddress(acct1,statadb.GetNonce(acct1))//❸
	statadb.CreateAccount(contract)
	statadb.SetCode(contract,[]byte("contract code bytes"))//❹

	statadb.SetNonce(contract,1)
	statadb.SetState(contract,common.BytesToHash([]byte("owner")),common.BytesToHash(acct1.Bytes()))//❺
	statadb.SetState(contract,common.BytesToHash([]byte("name")),common.BytesToHash([]byte("ysqi")))

	statadb.SetState(contract,common.BytesToHash([]byte("online")),common.BytesToHash([]byte{1}))
	statadb.SetState(contract,common.BytesToHash([]byte("online")),common.BytesToHash([]byte{}))//❻

	statadb.Commit(true)//❼
	fmt.Println(string(statadb.Dump(true,true,true)))//❽
}

