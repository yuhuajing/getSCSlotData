package parseslot

import (
	"main/parsetoken"
)

func getSCName(_res, tokenurl, apiKey string) string {
	//fmt.Println("SCNAmetoken")
	result, _ := parsetoken.GetSCNameFromEtherScan(_res, tokenurl, apiKey)
	//result, _ := etherscantx.GetERC20FromEtherScan(_res)
	// if err != nil {
	// 	return err.Error()
	// }
	//resultnft, _ := etherscantx.GetERC721FromEtherScan(_res)
	// if err != nil {
	// 	return err.Error()
	// }
	if len(result.Result) > 0 {
		return result.Result[0].ContractName
	}
	// else {
	// 	return "normal SC addres"
	// }
	return ""
}
