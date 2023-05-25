package slotstorage

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetSCstorage(address common.Address, slot int, blockNum int64, client *ethclient.Client) ([]byte, error) {
	t := common.BigToHash(big.NewInt(int64(slot)))
	//res :=
	if blockNum != 0 {
		blocknumBigInt := big.NewInt(int64(blockNum))
		res, err := client.StorageAt(context.Background(), address, t, blocknumBigInt)
		if err != nil {
			return nil, fmt.Errorf("client.getStorage error: %v", err)
		}
		return res, nil
	} else {
		res, err := client.StorageAt(context.Background(), address, t, nil)
		if err != nil {
			return nil, fmt.Errorf("client.getStorage error: %v", err)
		}
		return res, nil
		//int256.SetBytes(res)
	}
}
