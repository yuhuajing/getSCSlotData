package parseslot

import (
	"main/config"
	addresscheck "main/ethclient"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ParseDataToHumanReadable(resData string, client *ethclient.Client, chain string) string {
	if _res, sc := propabSc(resData); sc {
		return getdataBychain(_res, client, chain)
	}
	return ""
}

func propabSc(str string) (string, bool) {
	length := len(str)
	if length > 40 || length < 34 {
		return "", false
	} else if length == 40 {
		return "0x" + str, true
	} else {
		return paddingAddr(str), true
	}
}

func paddingAddr(str string) string {
	zeroString := ""
	diff := 40 - len(str)
	for i := 0; i < diff; i++ {
		zeroString += "0"
	}
	str = zeroString + str
	return "0x" + str
}

var (
	tokenurl    = config.EtherMainTokenUrl
	nftTokenurl = config.EtherMainNFTTokenUrl
	scnameurl   = config.EtherMainSCNameUrl
	apiKey      = config.EtherMainApiKEY
)

func getdataBychain(_res string, client *ethclient.Client, chain string) string {
	switch chain {
	case "bsc":
		tokenurl = config.BSCTokenUrl
		nftTokenurl = config.BSCNFTTokenUrl
		scnameurl = config.BSCSCNameUrl
		apiKey = ""
		//apiKey = config.BSCApiKey
	case "polygon":
		tokenurl = config.PolygonMainTokenUrl
		nftTokenurl = config.PolygonMainNFTTokenUrl
		scnameurl = config.PolygonMainSCNameUrl
		apiKey = config.PolygonMainApiKEY
	case "optimism":
		tokenurl = config.OptimisticMainTokenUrl
		nftTokenurl = config.OptimisticMainNFTTokenUrl
		scnameurl = config.OptimisticMainSCNameUrl
		apiKey = config.OptimisticMainApiKEY
	case "arbitrum":
		tokenurl = config.ArbitrumMainTokenUrl
		nftTokenurl = config.ArbitrumMainNFTTokenUrl
		scnameurl = config.ArbitrumMainSCNameUrl
		apiKey = config.ArbitrumMainApiKEY
	case "ethereum":
		tokenurl = config.EtherMainTokenUrl
		nftTokenurl = config.EtherMainNFTTokenUrl
		scnameurl = config.EtherMainSCNameUrl
		apiKey = config.EtherMainApiKEY
	case "goerli":
		tokenurl = config.EtherGoerliTokenUrl
		nftTokenurl = config.EtherGoerliNFTTokenUrl
		scnameurl = config.EtherGoerliSCNameUrl
		apiKey = config.EtherMainApiKEY
	}

	if addresscheck.CheckContractAddress(_res, client) {
		if tokenName := getTokenData(_res, tokenurl, nftTokenurl, apiKey); tokenName != "" {
			return "ERC20/NFT_SMART_CONTRACT_ADDRESS: " + tokenName
		} else if scName := getSCName(_res, scnameurl, apiKey); scName != "" {
			return "SMART_CONTRACT_ADDRESS: " + scName
		}
		//return
	}
	return "EOA_ADDRESS"
}
