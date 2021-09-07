package terra

import (
	"context"
	"encoding/hex"
	"log"
	"testing"

	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/frostornge/terra-go/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	terraapp "github.com/terra-project/core/app"
	"github.com/terra-project/core/x/bank"
)

func mustParseCoins(str string) cosmostypes.Coins {
	coins, err := cosmostypes.ParseCoins(str)
	if err != nil {
		panic(err)
	}
	return coins
}

func TestEstimateFee(t *testing.T) {
	ctx := context.Background()

	client := NewClient(httpclient.New(terraapp.MakeCodec(), "https://bombay-lcd.terra.dev"))

	privKey := secp256k1.GenPrivKey()
	key := NewRawKey(hex.EncodeToString(privKey[:]))
	acc, err := NewAccount(ctx, client, key)
	assert.NoError(t, err)

	tx, err := acc.CreateTx(ctx, CreateTxOptions{
		Msgs: []cosmostypes.Msg{bank.MsgSend{
			FromAddress: acc.GetAddress(),
			ToAddress:   acc.GetAddress(),
			Amount:      mustParseCoins("2000uusd"),
		}},
		Fee:           nil,
		GasAdjustment: 0,
		GasPrices:     nil,
		Sequence:      nil,
		Memo:          "",
	})
	assert.NoError(t, err)

	moneyMarket, err := cosmostypes.AccAddressFromBech32("terra15dwd5mj8v59wpj0wvt233mf5efdff808c5tkal")
	assert.NoError(t, err)

	var configResp struct {
		Height cosmostypes.Int `json:"height"`
		Result struct {
			OwnerAddr           cosmostypes.AccAddress `json:"owner_addr"`
			AnchorTerraContract cosmostypes.AccAddress `json:"aterra_contract"`
			InterestModel       cosmostypes.AccAddress `json:"interest_model"`
			DitributionModel    cosmostypes.AccAddress `json:"ditribution_model"`
			OverseerContract    cosmostypes.AccAddress `json:"overseer_contract"`
			CollectorContract   cosmostypes.AccAddress `json:"collector_contract"`
			DistributorContract cosmostypes.AccAddress `json:"distributor_contract"`
			StableDenom         string                 `json:"stable_denom"`
			MaxBorrowFactor     cosmostypes.Dec        `json:"max_borrow_factor"`
		} `json:"result"`
	}

	assert.NoError(t, client.Contract().QueryContractStore(ctx, moneyMarket, types.Q{"config": types.Q{}}, &configResp))

	log.Println(tx)
	log.Println(configResp)
}
