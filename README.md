# getSCSlotData


Compile for mac
```golang
go build -o ./getSlotDataForMac
```
Compile for Linux
```golang
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./getSlotDataForLinux 
```

```bash
Usage of ./getSlotDataForMac:
  --address string
        The smart contract address.
  --arrayslot "1 2 3 4 5"
        The slot array to get storage like "1 2 3 4 5"
  --blockNum int
        The blocknum, default: the latest block.
  --chain string
        The blockchain.
  --highslot int
        The highest slot.
  --lowslot int
        The lowest slot, default: 0.
  --slot int
        The singal slot, default: 0.
  --human 
        Print the SC name in a readable way.
```

1. chain configure

Now support ethereum、bsc、polygon、optimisim、arbitrum
```golang
switch xx{
    case "bsc":
			Get data from BSC
	case "polygon":
			Get data from Polygon
	case "optimism":
			Get data from Optimism
	case "arbitrum":
			Get data from Arbitrum
      case "ethereum":
			Get data from Ethereum Mainnet
      case "goerli":
			Get data from Ethereum Testnet Goerli
}
```
# Get smart contract slot data from Ethereum

1. Get single slot data
mac

./getSlotDataForMac --chain="ethereum" --address="0x0000a26b00c1F0DF003000390027140000fAa719" --highslot=10 --blockNum=17378079 --human

Example
```bash
0xc669b5f25f03be2ac0323037cb57f49eb543657a : SMART_CONTRACT_ADDRESS: GnosisSafeProxy
0x0
0x361ba2d30db481dd0e2840508d259084e07f20ea : EOA_ADDRESS
0x0
0x0
0x0
0x0
0x0
0x0
0x0
0x0
```

2. Get continuous slot data, from 0 to the highslot data. The default lowslot is 0.

 ./getSlotDataForMac --chain="ethereum" --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --highslot=1 --blockNum=17269518 --human

 Example
```bash
0x577261707065642045746865720000000000000000000000000000000000001a
0x5745544800000000000000000000000000000000000000000000000000000008
``` 

3. Get continuous slot data, from lowslot to the highslot data.

./getSlotDataForMac --chain="ethereum" --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --lowslot=1 --highslot=2 --blockNum=17269518 --human

 Example
```bash
0x5745544800000000000000000000000000000000000000000000000000000008
0x12
``` 

4. get specific slot data by providing string slot.

 ./getSlotDataForMac --chain="ethereum" --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --arrayslot="0 1 2" --blockNum=17269518 --human

 Example
```bash
0x577261707065642045746865720000000000000000000000000000000000001a
0x5745544800000000000000000000000000000000000000000000000000000008
0x12
``` 

# Get smart contract slot data from BSC

./getSlotDataForMac --chain="bsc" --address="0x245E0BA4562E96CC3873b7C1594954AaD877a4dE" --lowslot=0 --highslot=10 --blockNum=28312630 --human

 Example
```bash
0x55a3b37957bfbd3345bed9968e7e8dd56d67066 : EOA_ADDRESS
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
0x10000000000
0x10000000000000000
0x100000000000000000000000000000000
0x0
0x0
0x55a3b37957bfbd3345bed9968e7e8dd56d67066 : EOA_ADDRESS
0x4946c0e9f43f4dee607b0ef1fa1c
0x0
0x0
```

# Get smart contract slot data from polygon

./getSlotDataForMac --chain="polygon" --address="0x53fe4D1aB48363f1bFf8BFB5d96148E738f4ed6B" --lowslot=0 --highslot=10 --blockNum=42846967 --human

 Example
```bash
0x5069204e465420322e3000000000000000000000000000000000000000000014
0x5049000000000000000000000000000000000000000000000000000000000004
0x0
0x0
0x0
0x0
0x0
0x0
0x2c86
0x0
0x142b27b469bb2bb1257d8e077efd6d8b666cf1e : EOA_ADDRESS
```

# Get smart contract slot data from optimism

./getSlotDataForMac --chain="optimism" --address="0x86Bb63148d17d445Ed5398ef26Aa05Bf76dD5b59" --highslot=10 --human

Q: 无法指定blockNum

 Example
```bash
0x9cff694c1689f5c4bc4482b75b013e9634319b7b
0x0
0x0
0x0
0x0
0x1
0x6c
0x0
0x0
0x0
```

# Get smart contract slot data from arbitrum

./getSlotDataForMac --chain="arbitrum" --address="0x15b2fb8f08E4Ac1Ce019EADAe02eE92AeDF06851" --highslot=10 --blockNum=91881435 --human

 Example
```bash
0x3c27899d2f495928d00000
0x0
0x0
0x233001101d35d7a5cab7d51bcfbab3679ae4dc1501
0x0
0x0
0x0
0x0
0x0
0x0
0x0
```