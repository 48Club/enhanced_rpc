package enhanced_rpc

import (
	_ "embed"
	"encoding/json"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:embed exploiter.json
var gasExploiter []byte

//go:embed whitelistContract.json
var gasWhitelist []byte

//go:embed sponsorship.json
var sponsorship []byte

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
