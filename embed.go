package enhanced_rpc

import (
	_ "embed"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

//go:embed exploiter.json
var gasExploiter []byte

//go:embed whitelistContract.json
var gasWhitelist []byte

//go:embed sponsorship.json
var sponsorship []byte

func GasWhitelist() (addrs []common.Address) {
	var _gasWhitelist map[common.Address]string
	if json.Unmarshal(gasWhitelist, &_gasWhitelist) != nil {
		return
	}
	for addr := range _gasWhitelist {
		addrs = append(addrs, addr)
	}
	return
}

func GasExploiter() []common.Address {
	return unmarshal(gasExploiter)
}

func Sponsorship() []common.Address {
	return unmarshal(sponsorship)
}

func unmarshal(data []byte) (addrs []common.Address) {
	_ = json.Unmarshal(data, &addrs)
	return
}
