package enhanced_rpc

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:embed data/exploiter.json
var gasExploiter []byte

//go:embed data/whitelistContract.json
var gasWhitelist []byte

//go:embed data/sponsorship.json
var sponsorship []byte

//go:embed data/data.csv
var ofacAddresses []byte

func GasWhitelist() (addrs []common.Address) {
	var _gasWhitelist map[common.Address]string
	unmarshal(gasWhitelist, &_gasWhitelist)

	for addr := range _gasWhitelist {
		addrs = append(addrs, addr)
	}
	return
}

func GasExploiter() (res []common.Address) {
	unmarshal(gasExploiter, &res)
	return
}

func unmarshal(data []byte, a any) {
	_ = json.Unmarshal(data, a)
}

type Sponsorship map[common.Address]*struct {
	Allow       map[string]int `json:"allow"`
	CheckAmount bool           `json:"checkAmount"`
}

func SponsorshipInit() {
	if _sponsorship != nil {
		return
	}
	unmarshal(sponsorship, &_sponsorship)
}

var (
	_sponsorship Sponsorship
	mu           sync.Mutex
)

func IsSponsorable(a *common.Address, data []byte) (b bool) {
	SponsorshipInit()

	if a == nil || len(data) <= 4 { // 格式不匹配
		return
	}
	mu.Lock()
	defer mu.Unlock()
	info, ok := _sponsorship[*a]
	if !ok { // 不在赞助列表中
		return
	}

	length, ok := info.Allow[hexutil.Encode(data[:4])]
	if !ok || length != len(data) { // 函数参数不匹配
		return
	}

	if info.CheckAmount {
		amount := new(big.Int).SetBytes(data[length-common.HashLength:])
		if big.NewInt(48e16).Cmp(amount) > 0 { // 金额小于 0.48 ether, 不赞助
			return
		}
	}

	return true
}

func GetOfacAddresses() (res []common.Address) {
	res = []common.Address{}
	reader := csv.NewReader(bytes.NewReader(ofacAddresses))

	_, err := reader.Read()
	if err != nil {
		panic(err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		if len(record) > 2 {
			if addr := common.HexToAddress(record[1]); addr != (common.Address{}) {
				res = append(res, addr)
			}
		}
	}
	return
}
