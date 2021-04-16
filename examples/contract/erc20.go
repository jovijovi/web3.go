package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/jovijovi/web3.go"
	"github.com/jovijovi/web3.go/types"
)

func main() {

	abiStr := `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
	var infuraURL = "https://kovan.infura.io/v3/"

	web3, err := web3.NewWeb3(infuraURL)
	if err != nil {
		panic(err)
	}

	// set default account by private key
	privateKey := "610ca682d9b48e079e9017bb000a503071a158941674d304efccc68d9b8756f9"
	kovanChainId := int64(42)
	if err := web3.Eth.SetAccount(privateKey); err != nil {
		panic(err)
	}
	web3.Eth.SetChainId(kovanChainId)
	tokenAddr := "0xa76851d55db83dff1569fbc62d2317dec84d0ac8"
	contract, err := web3.Eth.NewContract(abiStr, tokenAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Contract address: ", contract.Address())

	totalSupply, err := contract.Call("totalSupply")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total supply %v\n", totalSupply)

	data, _ := contract.EncodeABI("balanceOf", web3.Eth.Address())
	fmt.Printf("%x\n", data)

	balance, err := contract.Call("balanceOf", web3.Eth.Address())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Balance of %v is %v\n", web3.Eth.Address(), balance)

	allowance, err := contract.Call("allowance", web3.Eth.Address(), "0x0000000000000000000000000000000000000002")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Allowance is %v\n", allowance)
	approveInputData, err := contract.Methods("approve").Inputs.Pack("0x0000000000000000000000000000000000000002", web3.Utils.ToWei(0.2, "gwei"))
	if err != nil {
		panic(err)
	}
	// fmt.Println(approveInputData)

	tokenAddress := common.HexToAddress(tokenAddr)

	call := &types.CallMsg{
		From: web3.Eth.Address(),
		To:   tokenAddress,
		Data: approveInputData,
	}
	// fmt.Printf("call %v\n", call)
	gasLimit, err := web3.Eth.EstimateGas(call)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Estimate gas limit %v\n", gasLimit)

	txHash, err := web3.Eth.SyncSendRawTransaction(
		common.HexToAddress(tokenAddr),
		big.NewInt(0),
		gasLimit,
		web3.Utils.ToWei(1, "gwei"),
		approveInputData,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Send approve tx hash %v\n", txHash)
}