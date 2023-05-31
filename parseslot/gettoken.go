package parseslot

import (
	"main/parsetoken"
)

func getTokenData(_res, tokenurl, nftTokenurl, apiKey string) string {
	result, _ := parsetoken.GetERC20FromEtherScan(_res, tokenurl, apiKey)
	resultnft, _ := parsetoken.GetERC721FromEtherScan(_res, nftTokenurl, apiKey)
	//result, _ := etherscantx.GetERC20FromEtherScan(_res)
	// if err != nil {
	// 	return err.Error()
	// }
	//resultnft, _ := etherscantx.GetERC721FromEtherScan(_res)
	// if err != nil {
	// 	return err.Error()
	// }
	if len(result.Result) > 0 {
		return result.Result[0].TokenName + "(" + result.Result[0].TokenSymbol + ")"
	} else if len(resultnft.Result) > 0 {
		return resultnft.Result[0].TokenName + "(" + resultnft.Result[0].TokenSymbol + ")"
	}
	// else {
	// 	return "normal SC addres"
	// }
	return ""
}
