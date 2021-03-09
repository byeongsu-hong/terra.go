package market

import (
	"github.com/frostornge/terra-go"
	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	terraassets "github.com/terra-project/core/types/assets"
)

var _ ExecutorMsg = (*market)(nil)

type ExecutorMsg interface {
	BorrowStableMsg(
		acc terra.Account,
		amount cosmostypes.Int,
		to *cosmostypes.AccAddress,
	) ([]cosmostypes.Msg, error)

	RepayStableMsg(
		acc terra.Account,
		amount cosmostypes.Coin,
	) ([]cosmostypes.Msg, error)

	DepositStableMsg(
		acc terra.Account,
		amount cosmostypes.Coin,
	) ([]cosmostypes.Msg, error)

	RedeemStableMsg(
		acc terra.Account,
		amount cosmostypes.Int,
	) ([]cosmostypes.Msg, error)
}

func (m market) BorrowStableMsg(
	acc terra.Account,
	amount cosmostypes.Int,
	to *cosmostypes.AccAddress,
) ([]cosmostypes.Msg, error) {
	payload := types.Q{"borrow_amount": amount.String()}
	if to != nil {
		payload["to"] = (*to).String()
	}
	return m.MakeMessage(acc, "borrow_stable", payload, nil)
}

func (m market) RepayStableMsg(
	acc terra.Account,
	amount cosmostypes.Coin,
) ([]cosmostypes.Msg, error) {
	if amount.Denom != terraassets.MicroUSDDenom {
		return nil, errors.Errorf("invalid denom %s. currently we support uusd only.", amount.Denom)
	}
	return m.MakeMessage(acc, "repay_stable", nil, cosmostypes.Coins{amount})
}

func (m market) DepositStableMsg(
	acc terra.Account,
	amount cosmostypes.Coin,
) ([]cosmostypes.Msg, error) {
	if amount.Denom != terraassets.MicroUSDDenom {
		return nil, errors.Errorf("invalid denom %s. currently we support uusd only.", amount.Denom)
	}
	return m.MakeMessage(acc, "deposit_stable", nil, cosmostypes.Coins{amount})
}

func (m market) RedeemStableMsg(
	acc terra.Account,
	amount cosmostypes.Int,
) ([]cosmostypes.Msg, error) {
	sendMsgs, err := m.anchored.SendMsg(
		acc, m.GetAddress(), amount,
		types.Q{"redeem_stable": types.Q{}},
	)
	if err != nil {
		return nil, errors.Wrap(err, "make cw20 send message")
	}
	return sendMsgs, nil
}
