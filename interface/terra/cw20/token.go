package cw20

import (
	"context"
	"fmt"

	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/bind"

	"github.com/airbloc/logger"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type Token interface {
	Name() string
	Symbol() string

	bind.BaseContract
	Executor
	ExecutorMsg
	Querier
}

type token struct {
	name   string
	symbol string
	logger logger.Logger

	bind.BaseContract
}

func (t token) Name() string   { return t.name }
func (t token) Symbol() string { return t.symbol }

func NewTokenContract(ctx context.Context, addr cosmostypes.AccAddress, client terra.Client) (Token, error) {
	t := token{BaseContract: bind.NewBaseContract(addr, client)}

	tokenInfoResp, err := t.GetTokenInfo(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetch token info")
	}

	t.name = tokenInfoResp.Name
	t.symbol = tokenInfoResp.Symbol
	t.logger = logger.New(fmt.Sprintf("anchor/cw20/%s", tokenInfoResp.Symbol))
	return t, nil
}
