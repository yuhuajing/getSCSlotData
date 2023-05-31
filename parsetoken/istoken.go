package parsetoken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Result struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}
type ResponseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  []NFTResult `json:"result"`
}

func GetERC721FromEtherScan(addr, url, apiKey string) (*ResponseData, error) {
	//url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=%s&page=1&offset=1&apikey=%s", addr, config.ApiKEY)
	url = strings.ReplaceAll(url, "{address}", addr)
	if apiKey != "" {
		url = strings.ReplaceAll(url, "{apiKey}", apiKey)
	}
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	for err != nil {
		fmt.Println("get request failed:", err)
		time.Sleep(10 * time.Second)
		resp, err = client.Do(req)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error get:%d\n", resp.StatusCode)
		//fmt.Errorf("http code:%d", resp.StatusCode)
		return nil, fmt.Errorf("error http code:%d", resp.StatusCode)
	}

	blocks := ResponseData{}
	_ = json.NewDecoder(resp.Body).Decode(&blocks)
	return &blocks, nil
}
