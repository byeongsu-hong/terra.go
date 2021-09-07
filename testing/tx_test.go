package testing

import (
	"context"
	"testing"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/httpclient"
	anchormarket "github.com/frostornge/terra-go/interface/anchor/money-market/market"
	"github.com/frostornge/terra-go/types"
	"github.com/tj/assert"
)

func mustParseCoins(str string) cosmostypes.Coins {
	coins, err := cosmostypes.ParseCoins(str)
	if err != nil {
		panic(err)
	}
	return coins
}

func mustParseDecCoins(str string) cosmostypes.DecCoins {
	coins, err := cosmostypes.ParseDecCoins(str)
	if err != nil {
		panic(err)
	}
	return coins
}

func TestEstimateFee(t *testing.T) {
	ctx := context.Background()

	client := terra.NewClient(httpclient.New(terra.MakeCodec(), "https://bombay-lcd.terra.dev"))

	//privKey := secp256k1.GenPrivKey()
	//key := terra.NewRawKey(hex.EncodeToString(privKey[:]))
	//acc, err := terra.NewAccount(ctx, client, key)
	//assert.NoError(t, err)

	moneyMarketAddr, err := cosmostypes.AccAddressFromBech32("terra15dwd5mj8v59wpj0wvt233mf5efdff808c5tkal")
	assert.NoError(t, err)

	moneyMarket, err := anchormarket.NewContract(ctx, moneyMarketAddr, client)
	assert.NoError(t, err)

	//depositStableMsg, err := moneyMarket.DepositStableMsg(acc, mustParseCoins("2000000uusd")[0])
	//assert.NoError(t, err)
	//_ = depositStableMsg
	//
	//_, _, err = acc.CreateAndSignTx(ctx, terra.CreateTxOptions{
	//	Msgs: []cosmostypes.Msg{bank.MsgSend{
	//		FromAddress: acc.GetAddress(),
	//		ToAddress:   acc.GetAddress(),
	//		Amount:      mustParseCoins("2000uusd"),
	//	}},
	//	GasPrices:     mustParseDecCoins("0.15uluna,0.15uusd"),
	//	GasAdjustment: 1.5,
	//})
	//assert.NoError(t, err)

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
	assert.NoError(t, client.Contract().QueryContractStore(ctx, moneyMarket.GetAddress(), types.Q{"config": types.Q{}}, &configResp))
}
