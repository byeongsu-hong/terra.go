package market

import (
	"context"

	"github.com/frostornge/terra-go"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ Executor = (*market)(nil)

type Executor interface {
	BorrowStable(
		ctx context.Context,
		acc terra.Account,
		amount cosmostypes.Int,
		to *cosmostypes.AccAddress,
	) (cosmostypes.TxResponse, error)

	RepayStable(
		ctx context.Context,
		acc terra.Account,
		amount cosmostypes.Coin,
	) (cosmostypes.TxResponse, error)

	DepositStable(
		ctx context.Context,
		acc terra.Account,
		amount cosmostypes.Coin,
	) (cosmostypes.TxResponse, error)

	RedeemStable(
		ctx context.Context,
		acc terra.Account,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)
}

func (m market) BorrowStable(
	ctx context.Context,
	acc terra.Account,
	amount cosmostypes.Int,
	to *cosmostypes.AccAddress,
) (cosmostypes.TxResponse, error) {
	msgs, err := m.BorrowStableMsg(acc, amount, to)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return m.Execute(ctx, acc, msgs, nil, nil)
}

func (m market) RepayStable(
	ctx context.Context,
	acc terra.Account,
	amount cosmostypes.Coin,
) (cosmostypes.TxResponse, error) {
	msgs, err := m.RepayStableMsg(acc, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return m.Execute(ctx, acc, msgs, nil, nil)
}

func (m market) DepositStable(
	ctx context.Context,
	acc terra.Account,
	amount cosmostypes.Coin,
) (cosmostypes.TxResponse, error) {
	msgs, err := m.DepositStableMsg(acc, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return m.Execute(ctx, acc, msgs, nil, nil)
}

func (m market) RedeemStable(
	ctx context.Context,
	acc terra.Account,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := m.RedeemStableMsg(acc, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return m.Execute(ctx, acc, msgs, nil, nil)
}
