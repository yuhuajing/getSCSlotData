package main

import (
	"flag"
	"fmt"
	"log"
	"main/addresscheck"
	"main/parseslot"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ethServer  = "https://cloudflare-eth.com"
	goerliNet  = "https://goerli.infura.io/v3/d7b27eea18a54fb389c2562ba19f8e36"
	bscRpc     = "https://bsc-mainnet.nodereal.io/v1/64a9df0874fb4a93b9d0a3849de012d3"
	polygonRpc = "https://polygon-mainnet.nodereal.io/v1/f510fc4d083b49d1ab383d25246cc7de"
	opRpc      = "https://opt-mainnet.nodereal.io/v1/1fd7be3e976444759d636dd367aae9ac"
	arbitrum   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/arbitrum-nitro/"
	avalanch   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/avalanche-c/ext/bc/C/avax"
	salana     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/solana/"
	near       = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/near/"
	fantom     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/fantom/"
)

var (
	address                 string
	slot, highslot, lowslot int
	arrayslot               string
	chain                   string
	blockNum                int64
	client                  *ethclient.Client
)

func init() {
	server := "https://cloudflare-eth.com"
	client = getConn(server)
	//account = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
}

func getConn(server string) *ethclient.Client {
	client, err := ethclient.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("we have a connection")
	return client
}

func main() {
	flag.StringVar(&chain, "chain", "", "The public Ethereum server to connect to")
	flag.StringVar(&address, "address", "", "The smart contract address to get storage")
	flag.IntVar(&slot, "slot", 0, "The singal slot to get storage")
	flag.IntVar(&highslot, "highslot", 0, "The contiounus highest slot to get storage")
	flag.Int64Var(&blockNum, "blockNum", 0, "The blocknum to get storage")
	flag.IntVar(&lowslot, "lowslot", 0, "The contiounus lowest slot to get storage")
	flag.StringVar(&arrayslot, "arrayslot", "", "The specific slot to get storage like `1 2 3 4 5` ")
	flag.Parse()

	if chain != "" {
		switch chain {
		case "bsc":
			client = getConn(bscRpc)
		case "polygon":
			client = getConn(polygonRpc)
		case "optimism":
			client = getConn(opRpc)
		case "arbitrum":
			client = getConn(arbitrum)
		case "ethereum":
			client = getConn(ethServer)
		case "goerli":
			client = getConn(goerliNet)
			// case "avalanch":
			// 	client = getConn(avalanch)
			// case "solana":
			// 	client = getConn(salana)
			// case "near":
			// 	client = getConn(near)
			// case "fantom":
			// 	client = getConn(fantom)
		}
	}

	if address == "" || !addresscheck.CheckContractAddress(address, client) {
		fmt.Println("--address should be provided or the address should be a smart contract address")
		return
	}

	if ethServer != "" {
		client = getConn(ethServer)
	}
	naddress := common.HexToAddress(address)
	slotRes := parseslot.CheckParameter(naddress, blockNum, slot, highslot, lowslot, arrayslot, client)
	if slotRes != nil {
		err := parseslot.ParseslotData(slotRes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
