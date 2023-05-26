package parseslot

import (
	"fmt"
	"math/big"
)

func ParseslotData(slotdata [][]byte) *[]string {
	int256 := new(big.Int)
	resString := make([]string, 0)
	for _, v := range slotdata {
		int256.SetBytes(v)
		resInt := fmt.Sprintf("%x", int256)
		resString = append(resString, resInt)
		//fmt.Println(resInt)
	}
	return &resString
}
