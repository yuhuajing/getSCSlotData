package parseslot

import (
	"main/addresscheck"
	"main/etherscantx"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ParseDataToHumanReadable(resData string, client *ethclient.Client) string {
	if _res, sc := propabSc(resData); sc {
		if addresscheck.CheckContractAddress(_res, client) {
			result, _ := etherscantx.GetERC20FromEtherScan(_res)
			// if err != nil {
			// 	return err.Error()
			// }
			resultnft, _ := etherscantx.GetERC721FromEtherScan(_res)
			// if err != nil {
			// 	return err.Error()
			// }
			if len(result.Result) > 0 {
				return result.Result[0].TokenName + "(" + result.Result[0].TokenSymbol + ")"
			} else if len(resultnft.Result) > 0 {
				return resultnft.Result[0].TokenName + "(" + resultnft.Result[0].TokenSymbol + ")"
			}
		} else {
			return "Probaly sc address"
		}
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
