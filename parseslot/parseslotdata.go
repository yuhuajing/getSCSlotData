package parseslot

import (
	"fmt"
	"math/big"
)

func ParseslotData(slotdata [][]byte) error {
	int256 := new(big.Int)
	for _, v := range slotdata {
		int256.SetBytes(v)
		//fmt.Printf("0x%x\n", int256)
		resInt := fmt.Sprintf("0x%x", int256)
		fmt.Println(resInt)
	}
	return nil
}
