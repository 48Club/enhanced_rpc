package enhanced_rpc

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestPrintEmbed(t *testing.T) {
	gasExploiter, sponsorship := GasExploiter(), Sponsorship()
	assert.Equal(t, len(gasExploiter), 5)
	assert.Equal(t, gasExploiter[0], common.HexToAddress("0x0000000000004946c0e9f43f4dee607b0ef1fa1c"))
	assert.Equal(t, len(GasWhitelist()), 182)
	assert.Equal(t, len(sponsorship), 4)
	assert.Equal(t, sponsorship[0], common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"))
}
