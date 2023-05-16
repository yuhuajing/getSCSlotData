package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	address                 string
	slot, highslot, lowslot int
	arrayslot               string
	client                  *ethclient.Client
	//account                 common.Address
	ethServer string
)

func init() {
	ethServer = "https://cloudflare-eth.com"
	client = getConn(ethServer)
	//account = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
}

func getConn(server string) *ethclient.Client {
	client, err := ethclient.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	return client
}

func main() {
	flag.StringVar(&ethServer, "ethServer", "", "")
	flag.StringVar(&address, "address", "", "")
	flag.IntVar(&slot, "slot", 0, "")
	flag.IntVar(&highslot, "highslot", 0, "")
	flag.IntVar(&lowslot, "lowslot", 0, "")
	flag.StringVar(&arrayslot, "arrayslot", "", "")
	flag.Parse()
	if address == "" || !checkContractAddress(address) {
		fmt.Println("--address should be provided or the address should be a smart contract address")
		return
	}
	if ethServer != "" {
		client = getConn(ethServer)
	}
	naddress := common.HexToAddress(address)
	fmt.Println(naddress)
	if slot > 0 {
		fmt.Printf("signal Slot provided, get slot %d of the address %s\n", slot, naddress.Hex())
		getSCstorage(naddress, slot)
		return
	} else if highslot > 0 && lowslot >= 0 {
		for i := lowslot; i <= highslot; i++ {
			fmt.Printf("highSlot provided, get slot %d of the address %s\n", i, naddress.Hex())
			getSCstorage(naddress, slot)
		}
		return
	} else if arrayslot != "" {
		strArray := strings.Split(arrayslot, " ")
		for i := 0; i < len(strArray); i++ {
			num, err := strconv.Atoi(strArray[i])
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			fmt.Printf("arraySlot provided, get slot %d of the address %s\n", num, naddress.Hex())
			getSCstorage(naddress, num)
		}
		return
	} else {
		fmt.Printf("nothing provided,get slot 0 of the address %s\n", naddress.Hex())
		getSCstorage(naddress, slot)
	}

}

func checkAddress(addr string) bool {
	// 16 hex 0-f
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

// check the address whether is a smart contract address
func checkContractAddress(addr string) bool {
	if !checkAddress(addr) {
		return false
	}
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil) //nil is the latest block
	if err != nil {
		panic(err)
	}
	isContract := len(bytecode) > 0
	if isContract {
		fmt.Println("SC address")
		return true
	}
	fmt.Println("Normal address")
	return false
}

func getSCstorage(address common.Address, slot int) {
	//for i := 0; i <= slot; i++ {
	t := common.BigToHash(big.NewInt(int64(slot)))
	res, _ := client.StorageAt(context.Background(), address, t, nil)
	//fmt.Println(i)
	fmt.Println(res)
	//}
}
