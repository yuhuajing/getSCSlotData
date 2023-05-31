package config

import (
	"fmt"

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
	//Ethereum Mainnet
	// trunk-ignore(gitleaks/generic-api-key)
	EtherMainApiKEY      = "KXTV96AQDTB5F1M6CKPSJH69ZJ7EF5R713"
	EtherMainTokenUrl    = "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1&apikey={apiKey}"
	EtherMainNFTTokenUrl = "https://api.etherscan.io/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1&apikey={apiKey}"
	EtherMainSCNameUrl   = "https://api.etherscan.io/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"

	//Ethereum Testnet Goerli
	EtherGoerliTokenUrl    = "https://api-goerli.etherscan.io/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1&apikey={apiKey}"
	EtherGoerliNFTTokenUrl = "https://api-goerli.etherscan.io/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1&apikey={apiKey}"
	EtherGoerliSCNameUrl   = "https://api-goerli.etherscan.io/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"
	//BSC Mainnet
	// trunk-ignore(gitleaks/generic-api-key)
	BSCApiKey      = "HK7M7ZSD5MXW3U8X6IZ6732HV8KZVX449S"
	BSCTokenUrl    = "https://api.bscscan.com/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1"
	BSCNFTTokenUrl = "https://api.bscscan.com/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1"
	BSCSCNameUrl   = "https://api.bscscan.com/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"
	// polygon mainnet
	// trunk-ignore(gitleaks/generic-api-key)
	PolygonMainApiKEY      = "7PMJA7NTUDF5BCEND1XA44MH7NC7MBCA6N"
	PolygonMainTokenUrl    = "https://api.polygonscan.com/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	PolygonMainNFTTokenUrl = "https://api.polygonscan.com/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	PolygonMainSCNameUrl   = "https://api.polygonscan.com/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"
	// optimistic mainnet
	// trunk-ignore(gitleaks/generic-api-key)
	OptimisticMainApiKEY      = "1GMDJZ4CYZQIX4W6U5MEMDZAGQT64NVUWG"
	OptimisticMainTokenUrl    = "https://api-optimistic.etherscan.io/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	OptimisticMainNFTTokenUrl = "https://api-optimistic.etherscan.io/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	OptimisticMainSCNameUrl   = "https://api-optimistic.etherscan.io/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"
	//Arbitrum mainnet
	//	https://api.arbiscan.io/api?module=account&action=tokentx&contractaddress=0xda10009cbd5d07dd0cecc66161fc93d7c9000da1&page=1&offset=1&sort=asc&apikey=
	ArbitrumMainApiKEY      = "EHJDJWN9N6QI684FB9VH2QHDBWSSYNAVCW"
	ArbitrumMainTokenUrl    = "https://api.arbiscan.io/api?module=account&action=tokentx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	ArbitrumMainNFTTokenUrl = "https://api.arbiscan.io/api?module=account&action=tokennfttx&contractaddress={address}&page=1&offset=1&sort=asc&apikey={apiKey}"
	ArbitrumMainSCNameUrl   = "https://api.arbiscan.io/api?module=contract&action=getsourcecode&address={address}&apikey={apiKey}"
)

func GetConn(server string) *ethclient.Client {
	client, err := ethclient.Dial(server)
	if err != nil {
		fmt.Printf("The server is not supported any more with err %s, replace a new server please.%s", err, server)
		//log.Fatal(err)
		return nil
	}
	return client
}
