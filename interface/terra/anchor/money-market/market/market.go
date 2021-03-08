package market

import (
	"context"

	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/bind"
	"github.com/frostornge/terra-go/interface/terra/cw20"

	"github.com/airbloc/logger"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type Market interface {
	StableDenom() string
	Anchored() cw20.Token

	bind.BaseContract
	Executor
	ExecutorMsg
	Querier
}

type market struct {
	stableDenom string
	anchored    cw20.Token
	logger      logger.Logger

	bind.BaseContract
}

func (m market) StableDenom() string  { return m.stableDenom }
func (m market) Anchored() cw20.Token { return m.anchored }

func NewContract(ctx context.Context, addr cosmostypes.AccAddress, client terra.Client) (Market, error) {
	m := market{
		BaseContract: bind.NewBaseContract(addr, client),
		logger:       logger.New("anchor/money-market/market"),
	}

	configResp, err := m.GetConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetch market contract config")
	}

	m.stableDenom = configResp.StableDenom
	m.anchored, err = cw20.NewTokenContract(ctx, configResp.ATerraContract, client)
	if err != nil {
		return nil, errors.Wrap(err, "new cw20 token contract")
	}
	return m, nil
}
