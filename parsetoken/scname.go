package parsetoken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type scResult struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	OptimizationUsed     string `json:"OptimizationUsed"`
	Runs                 string `json:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
}
type scResponseData struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []scResult `json:"result"`
}

func GetSCNameFromEtherScan(addr, url, apiKey string) (*scResponseData, error) {
	//url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=%s&page=1&offset=1&apikey=%s", addr, config.ApiKEY)
	url = strings.ReplaceAll(url, "{address}", addr)
	if apiKey != "" {
		url = strings.ReplaceAll(url, "{apiKey}", apiKey)
	}
	//fmt.Println(url)
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

	blocks := scResponseData{}
	_ = json.NewDecoder(resp.Body).Decode(&blocks)
	return &blocks, nil
}
