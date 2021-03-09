package overseer

import (
	"context"

	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/bind"

	"github.com/airbloc/logger"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

type Overseer interface {
	bind.BaseContract
}

type overseer struct {
	logger logger.Logger

	bind.BaseContract
}

func NewContract(ctx context.Context, addr cosmostypes.AccAddress, client terra.Client) (Overseer, error) {
	o := overseer{
		BaseContract: bind.NewBaseContract(addr, client),
		logger:       logger.New("anchor/money-market/overseer"),
	}

	return o, nil
}
