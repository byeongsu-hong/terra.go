package types

import (
	"os"
	"testing"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	terratypes "github.com/terra-project/core/types"
)

func TestMain(m *testing.M) {
	// use terra types
	config := cosmostypes.GetConfig()
	config.SetBech32PrefixForAccount(terratypes.Bech32PrefixAccAddr, terratypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(terratypes.Bech32PrefixValAddr, terratypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(terratypes.Bech32PrefixConsAddr, terratypes.Bech32PrefixConsPub)
	config.SetCoinType(terratypes.CoinType)
	config.SetFullFundraiserPath(terratypes.FullFundraiserPath)
	config.Seal()

	code := m.Run()
	os.Exit(code)
}
