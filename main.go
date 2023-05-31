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
	server                  string
)

func main() {
	flag.StringVar(&chain, "chain", "", "Blockchain name,for example: bsc,polygon,optimism,arbitrum,ethereum,goerli")
	flag.StringVar(&server, "server", "", "Blockchain RPC server")
	flag.StringVar(&address, "address", "", "Smart contract address")
	flag.Int64Var(&blockNum, "blockNum", 0, "The blocknum to get storage, default:lastest block")
	flag.IntVar(&slot, "slot", 0, "The target slot")
	flag.IntVar(&highslot, "highslot", 0, "The highest slot")
	flag.IntVar(&lowslot, "lowslot", 0, "The lowest slot, default: 0")
	flag.StringVar(&arrayslot, "arrayslot", "", "The slot array like `1 2 3 4 5` ")
	human := flag.Bool("human", false, "Human readable data")
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
		default:
			fmt.Println("UNSUPPORTED_BLOCKCHAIN")
			return
		}
	} else {
		client = config.GetConn(config.EthServer)
	}

	if server != "" {
		client = config.GetConn(server)
	}

	if client == nil {
		return
	}

	if address == "" {
		fmt.Println("ADDRESS_IS_NEEDED")
		return
	} else if !addresscheck.CheckContractAddress(address, client) {
		fmt.Println("NOT_SMART_CONTRACT")
		return
	}

	latestblockNum := addresscheck.GetLatestBlockNum(client)
	if latestblockNum > 0 && blockNum > latestblockNum {
		fmt.Printf("INVALID_BLOCK_NUMBER %d WITH_LATEST_BLOCK_NUMBER: %d", blockNum, latestblockNum)
		return
	} else if blockNum == 0 {
		blockNum = latestblockNum
	}

	naddress := common.HexToAddress(address)
	slotRes := parseslot.CheckParameter(naddress, blockNum, slot, highslot, lowslot, arrayslot, client)
	if slotRes != nil {
		resString := parseslot.ParseslotData(slotRes)
		for _, res := range *resString {

			if *human {
				tokenname := parseslot.ParseDataToHumanReadable(res, client, chain)
				if tokenname != "" {
					fmt.Println("0x" + res + " : " + tokenname)
				} else {
					fmt.Println("0x" + res)
				}
			} else {
				fmt.Println("0x" + res)
			}
		}
	} else {
		return
	}
}
