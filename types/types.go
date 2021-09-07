package types

import (
	"encoding/json"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type Q map[string]interface{}

type TokensHuman struct {
	Addr   cosmostypes.AccAddress `json:"addr"`
	Amount cosmostypes.Int        `json:"amount"`
}

func (t *TokensHuman) UnmarshalJSON(b []byte) error {
	var origin []string
	if err := json.Unmarshal(b, &origin); err != nil {
		return errors.Wrap(err, "unmarshal origin")
	}
	if len(origin) != 2 {
		return errors.Errorf("invalid tuple length %d", len(origin))
	}

	addr, err := cosmostypes.AccAddressFromBech32(origin[0])
	if err != nil {
		return errors.Wrap(err, "bech32 to accAddress")
	}
	amount, ok := cosmostypes.NewIntFromString(origin[1])
	if !ok {
		return errors.Errorf("invalid numeric string %s", origin[1])
	}

	t.Addr = addr
	t.Amount = amount
	return nil
}

func (t TokensHuman) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{
		t.Addr.String(),
		t.Amount.String(),
	})
}
