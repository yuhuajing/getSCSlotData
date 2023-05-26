package ethclient

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// check the address whether is a smart contract address
func CheckContractAddress(addr string, client *ethclient.Client) bool {
	if !CheckAddress(addr) {
		return false
	}
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil) //nil is the latest block
	if err != nil {
		panic(err)
	}
	isContract := len(bytecode) > 0
	if isContract {
		//fmt.Println("SC address")
		return true
	}
	fmt.Println("This is normal address, but we want a smart contract address")
	return false
}
