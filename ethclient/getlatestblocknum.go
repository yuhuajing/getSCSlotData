package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockNum(client *ethclient.Client) int64 {
	blockNum, _ := client.BlockNumber(context.Background())
	return int64(blockNum)
}
