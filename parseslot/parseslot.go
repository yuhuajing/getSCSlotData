package parseslot

import (
	"main/slotstorage"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckParameter(naddress common.Address, blockNum int64, slot, highslot, lowslot int, arrayslot string, client *ethclient.Client) [][]byte {
	resInt := make([][]byte, 0)
	if slot > 0 {
		slotInt, err := slotstorage.GetSCstorage(naddress, slot, blockNum, client)
		if err != nil {
			//log.Fatal(err)
			return nil
		}
		//fmt.Println(string(slotInt))
		resInt = append(resInt, slotInt)
		return resInt
	} else if highslot > 0 && lowslot >= 0 {
		for i := lowslot; i <= highslot; i++ {
			slotInt, err := slotstorage.GetSCstorage(naddress, i, blockNum, client)
			if err != nil {
				//log.Fatal(err)
				return nil
			}
			resInt = append(resInt, slotInt)
		}
		return resInt
	} else if arrayslot != "" {
		strArray := strings.Split(arrayslot, " ")
		for i := 0; i < len(strArray); i++ {
			num, _ := strconv.Atoi(strArray[i])
			slotInt, err := slotstorage.GetSCstorage(naddress, num, blockNum, client)
			if err != nil {
				//log.Fatal(err)
				return nil
			}
			resInt = append(resInt, slotInt)
		}
		return resInt
	} else {
		slotInt, err := slotstorage.GetSCstorage(naddress, 0, 0, client)
		if err != nil {
			//log.Fatal(err)
			return nil
		}
		resInt = append(resInt, slotInt)
	}
	return resInt
}
