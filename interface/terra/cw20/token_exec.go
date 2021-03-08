package cw20

import (
	"context"

	"github.com/frostornge/terra-go"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ Executor = (*token)(nil)

type Executor interface {
	Transfer(
		ctx context.Context,
		acc terra.Account,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)

	TransferFrom(
		ctx context.Context,
		acc terra.Account,
		owner cosmostypes.AccAddress,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)

	Burn(
		ctx context.Context,
		acc terra.Account,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)

	BurnFrom(
		ctx context.Context,
		acc terra.Account,
		owner cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)

	Send(
		ctx context.Context,
		acc terra.Account,
		contract cosmostypes.AccAddress,
		amount cosmostypes.Int,
		hook interface{},
	) (cosmostypes.TxResponse, error)

	SendFrom(
		ctx context.Context,
		acc terra.Account,
		owner cosmostypes.AccAddress,
		contract cosmostypes.AccAddress,
		amount cosmostypes.Int,
		hook interface{},
	) (cosmostypes.TxResponse, error)

	Mint(
		ctx context.Context,
		acc terra.Account,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) (cosmostypes.TxResponse, error)

	IncreaseAllowance(
		ctx context.Context,
		acc terra.Account,
		spender cosmostypes.AccAddress,
		amount cosmostypes.Int,
		expires interface{},
	) (cosmostypes.TxResponse, error)

	DecreaseAllowance(
		ctx context.Context,
		acc terra.Account,
		spender cosmostypes.AccAddress,
		amount cosmostypes.Int,
		expires interface{},
	) (cosmostypes.TxResponse, error)
}

func (t token) Transfer(
	ctx context.Context,
	acc terra.Account,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := t.TransferMsg(acc, recipient, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) TransferFrom(
	ctx context.Context,
	acc terra.Account,
	owner cosmostypes.AccAddress,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := t.TransferFromMsg(acc, owner, recipient, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) Burn(
	ctx context.Context,
	acc terra.Account,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := t.BurnMsg(acc, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) BurnFrom(
	ctx context.Context,
	acc terra.Account,
	owner cosmostypes.AccAddress,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := t.BurnFromMsg(acc, owner, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) Send(
	ctx context.Context,
	acc terra.Account,
	contract cosmostypes.AccAddress,
	amount cosmostypes.Int,
	hook interface{},
) (cosmostypes.TxResponse, error) {
	msgs, err := t.SendMsg(acc, contract, amount, hook)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) SendFrom(
	ctx context.Context,
	acc terra.Account,
	owner cosmostypes.AccAddress,
	contract cosmostypes.AccAddress,
	amount cosmostypes.Int,
	hook interface{},
) (cosmostypes.TxResponse, error) {
	msgs, err := t.SendFromMsg(acc, owner, contract, amount, hook)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) Mint(
	ctx context.Context,
	acc terra.Account,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) (cosmostypes.TxResponse, error) {
	msgs, err := t.MintMsg(acc, recipient, amount)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) IncreaseAllowance(
	ctx context.Context,
	acc terra.Account,
	spender cosmostypes.AccAddress,
	amount cosmostypes.Int,
	expires interface{},
) (cosmostypes.TxResponse, error) {
	msgs, err := t.IncreaseAllowanceMsg(acc, spender, amount, expires)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}

func (t token) DecreaseAllowance(
	ctx context.Context,
	acc terra.Account,
	spender cosmostypes.AccAddress,
	amount cosmostypes.Int,
	expires interface{},
) (cosmostypes.TxResponse, error) {
	msgs, err := t.DecreaseAllowanceMsg(acc, spender, amount, expires)
	if err != nil {
		return cosmostypes.TxResponse{}, err
	}
	return t.Execute(ctx, acc, msgs, nil, nil)
}
