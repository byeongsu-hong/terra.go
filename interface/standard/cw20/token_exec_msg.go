package cw20

import (
	"encoding/json"

	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	terratypes "github.com/terra-money/core/types"
)

var _ ExecutorMsg = (*token)(nil)

type ExecutorMsg interface {
	TransferMsg(
		acc terra.Account,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)

	TransferFromMsg(
		acc terra.Account,
		owner cosmostypes.AccAddress,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)

	BurnMsg(
		acc terra.Account,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)

	BurnFromMsg(
		acc terra.Account,
		owner cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)

	SendMsg(
		acc terra.Account,
		contract cosmostypes.AccAddress,
		amount cosmostypes.Int,
		hook interface{},
	) ([]cosmostypes.Msg, error)

	SendFromMsg(
		acc terra.Account,
		owner cosmostypes.AccAddress,
		contract cosmostypes.AccAddress,
		amount cosmostypes.Int,
		hook interface{},
	) ([]cosmostypes.Msg, error)

	MintMsg(
		acc terra.Account,
		recipient cosmostypes.AccAddress,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)

	IncreaseAllowanceMsg(
		acc terra.Account,
		spender cosmostypes.AccAddress,
		amount cosmostypes.Int,
		expires interface{},
	) ([]cosmostypes.Msg, error)

	DecreaseAllowanceMsg(
		acc terra.Account,
		spender cosmostypes.AccAddress,
		amount cosmostypes.Int,
		expires interface{},
	) ([]cosmostypes.Msg, error)
}

func (t token) TransferMsg(
	acc terra.Account,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "transfer",
		types.Q{
			"recipient": recipient.String(),
			"amount":    amount.String(),
		},
		nil,
	)
}

func (t token) TransferFromMsg(
	acc terra.Account,
	owner cosmostypes.AccAddress,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "transfer_from",
		types.Q{
			"owner":     owner.String(),
			"recipient": recipient.String(),
			"amount":    amount.String(),
		},
		nil,
	)
}

func (t token) BurnMsg(
	acc terra.Account,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "burn",
		types.Q{"amount": amount.String()},
		nil,
	)
}

func (t token) BurnFromMsg(
	acc terra.Account,
	owner cosmostypes.AccAddress,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "burn_from",
		types.Q{
			"owner":  owner.String(),
			"amount": amount.String(),
		},
		nil,
	)
}

func (t token) SendMsg(
	acc terra.Account,
	contract cosmostypes.AccAddress,
	amount cosmostypes.Int,
	hook interface{},
) ([]cosmostypes.Msg, error) {
	payload := types.Q{
		"contract": contract.String(),
		"amount":   amount.String(),
	}
	if hook != nil {
		rawHook, err := json.Marshal(hook)
		if err != nil {
			return nil, errors.Wrap(err, "marshal hook message")
		}
		payload["msg"] = terratypes.Base64Bytes(rawHook)
	}
	return t.MakeMessage(acc, "send", payload, nil)
}

func (t token) SendFromMsg(
	acc terra.Account,
	owner cosmostypes.AccAddress,
	contract cosmostypes.AccAddress,
	amount cosmostypes.Int,
	hook interface{},
) ([]cosmostypes.Msg, error) {
	payload := types.Q{
		"owner":    owner.String(),
		"contract": contract.String(),
		"amount":   amount.String(),
	}
	if hook != nil {
		rawHook, err := json.Marshal(hook)
		if err != nil {
			return nil, errors.Wrap(err, "marshal hook message")
		}
		payload["msg"] = terratypes.Base64Bytes(rawHook)
	}
	return t.MakeMessage(acc, "send_from", payload, nil)
}

func (t token) MintMsg(
	acc terra.Account,
	recipient cosmostypes.AccAddress,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "mint",
		types.Q{
			"recipient": recipient.String(),
			"amount":    amount.String(),
		},
		nil,
	)
}

func (t token) IncreaseAllowanceMsg(
	acc terra.Account,
	spender cosmostypes.AccAddress,
	amount cosmostypes.Int,
	expires interface{},
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "increase_allowance",
		types.Q{
			"spender": spender.String(),
			"amount":  amount.String(),
			"expires": expires,
		},
		nil,
	)
}

func (t token) DecreaseAllowanceMsg(
	acc terra.Account,
	spender cosmostypes.AccAddress,
	amount cosmostypes.Int,
	expires interface{},
) ([]cosmostypes.Msg, error) {
	return t.MakeMessage(
		acc, "decrease_allowance",
		types.Q{
			"spender": spender.String(),
			"amount":  amount.String(),
			"expires": expires,
		},
		nil,
	)
}
