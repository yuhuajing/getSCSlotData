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
	blockNum  int64
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
	//	fmt.Println("we have a connection")
	return client
}

func main() {
	flag.StringVar(&ethServer, "ethServer", "", "The public Ethereum server to connect to")
	flag.StringVar(&address, "address", "", "The smart contract address to get storage")
	flag.IntVar(&slot, "slot", 0, "The singal slot to get storage")
	flag.IntVar(&highslot, "highslot", 0, "The contiounus highest slot to get storage")
	flag.Int64Var(&blockNum, "blockNum", 0, "The blocknum to get storage")
	flag.IntVar(&lowslot, "lowslot", 0, "The contiounus lowest slot to get storage")
	flag.StringVar(&arrayslot, "arrayslot", "", "The specific slot to get storage`1 2 3 4 5` ")
	flag.Parse()

	if address == "" || !checkContractAddress(address) {
		fmt.Println("--address should be provided or the address should be a smart contract address")
		return
	}

	if ethServer != "" {
		client = getConn(ethServer)
	}
	naddress := common.HexToAddress(address)
	//	fmt.Println(naddress)
	if slot > 0 {
		//	fmt.Printf("signal Slot provided, get slot %d of the address %s\n", slot, naddress.Hex())
		getSCstorage(naddress, slot, blockNum)
		return
	} else if highslot > 0 && lowslot >= 0 {
		for i := lowslot; i <= highslot; i++ {
			//fmt.Printf("highSlot provided, get slot %d of the address %s\n", i, naddress.Hex())
			getSCstorage(naddress, i, blockNum)
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
			//	fmt.Printf("arraySlot provided, get slot %d of the address %s\n", num, naddress.Hex())
			getSCstorage(naddress, num, blockNum)
		}
		return
	} else {
		//	fmt.Printf("nothing provided,get slot 0 of the address %s\n", naddress.Hex())
		getSCstorage(naddress, slot, blockNum)
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
		//fmt.Println("SC address")
		return true
	}
	fmt.Println("This is normal address, but we want a smart contract address")
	return false
}

func getSCstorage(address common.Address, slot int, blockNum int64) {
	t := common.BigToHash(big.NewInt(int64(slot)))
	int256 := new(big.Int)
	if blockNum != 0 {
		//fmt.Printf("get slot %d of the address %s in the block %d\n", slot, address.Hex(), blockNum)
		blocknumBigInt := big.NewInt(int64(blockNum))
		res, _ := client.StorageAt(context.Background(), address, t, blocknumBigInt)
		//	fmt.Println(res)
		int256.SetBytes(res)
	} else {
		//fmt.Printf("get slot %d of the address %s in the latest block\n", slot, address.Hex())
		res, _ := client.StorageAt(context.Background(), address, t, nil)
		//	fmt.Println(res)
		int256.SetBytes(res)
	}
	fmt.Printf("0x%x\n", int256)
	//fmt.Printf("hexadecimal: 0x%x\n", int256)
	// fmt.Printf("uint256: %v\n", int256)
}
