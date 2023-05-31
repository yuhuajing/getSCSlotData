package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockNum(client *ethclient.Client) int64 {
	blockNum, err := client.BlockNumber(context.Background())
	if err == nil {
		return int64(blockNum)
	}
	return 0
}
