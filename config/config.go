package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthServer  = "https://cloudflare-eth.com"
	GoerliNet  = "https://goerli.infura.io/v3/d7b27eea18a54fb389c2562ba19f8e36"
	BscRpc     = "https://bsc-mainnet.nodereal.io/v1/64a9df0874fb4a93b9d0a3849de012d3"
	PolygonRpc = "https://polygon-mainnet.nodereal.io/v1/f510fc4d083b49d1ab383d25246cc7de"
	OpRpc      = "https://opt-mainnet.nodereal.io/v1/1fd7be3e976444759d636dd367aae9ac"
	Arbitrum   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/arbitrum-nitro/"
	Avalanch   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/avalanche-c/ext/bc/C/avax"
	Salana     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/solana/"
	Near       = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/near/"
	Fantom     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/fantom/"
	// trunk-ignore(gitleaks/generic-api-key)
	ApiKEY = "KXTV96AQDTB5F1M6CKPSJH69ZJ7EF5R713"
)

func GetConn(server string) *ethclient.Client {
	client, err := ethclient.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("we have a connection")
	return client
}
