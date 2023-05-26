package main

import (
	"flag"
	"fmt"
	"main/config"
	addresscheck "main/ethclient"
	"main/parseslot"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	address                 string
	slot, highslot, lowslot int
	arrayslot               string
	chain                   string
	blockNum                int64
	client                  *ethclient.Client
)

func main() {
	flag.StringVar(&chain, "chain", "", "The public Ethereum server to connect to")
	flag.StringVar(&address, "address", "", "The smart contract address to get storage")
	flag.IntVar(&slot, "slot", 0, "The singal slot to get storage")
	flag.IntVar(&highslot, "highslot", 0, "The contiounus highest slot to get storage")
	flag.Int64Var(&blockNum, "blockNum", -1, "The blocknum to get storage")
	flag.IntVar(&lowslot, "lowslot", 0, "The contiounus lowest slot to get storage")
	flag.StringVar(&arrayslot, "arrayslot", "", "The specific slot to get storage like `1 2 3 4 5` ")
	human := flag.Bool("h", false, "Parse the slot data to a human readable data structure")
	flag.Parse()

	if chain != "" {
		switch chain {
		case "bsc":
			client = config.GetConn(config.BscRpc)
		case "polygon":
			client = config.GetConn(config.PolygonRpc)
		case "optimism":
			client = config.GetConn(config.OpRpc)
		case "arbitrum":
			client = config.GetConn(config.Arbitrum)
		case "ethereum":
			client = config.GetConn(config.EthServer)
		case "goerli":
			client = config.GetConn(config.GoerliNet)
			// case "avalanch":
			// 	client = getConn(avalanch)
			// case "solana":
			// 	client = getConn(salana)
			// case "near":
			// 	client = getConn(near)
			// case "fantom":
			// 	client = getConn(fantom)
		}
	} else {
		client = config.GetConn(config.EthServer)
	}

	if address == "" || !addresscheck.CheckContractAddress(address, client) {
		fmt.Println("--address should be provided or the address should be a smart contract address")
		return
	}

	if blockNum < 0 {
		blockNum = addresscheck.GetLatestBlockNum(client)
	}

	naddress := common.HexToAddress(address)
	slotRes := parseslot.CheckParameter(naddress, blockNum, slot, highslot, lowslot, arrayslot, client)
	if slotRes != nil {
		resString := parseslot.ParseslotData(slotRes)
		for _, res := range *resString {
			fmt.Println("0x" + res)
			if *human {
				tokenname := parseslot.ParseDataToHumanReadable(res, client, chain)
				if tokenname != "" {
					fmt.Println(tokenname)
				}
			}

		}
	}
}
