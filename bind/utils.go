package bind

import (
	"encoding/json"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/terra-project/core/x/wasm"
)

func GenerateExecuteMsg(
	sender cosmostypes.AccAddress,
	contract cosmostypes.AccAddress,
	executeMsg interface{},
	coins cosmostypes.Coins,
) (wasm.MsgExecuteContract, error) {
	rawExecuteMsg, err := json.Marshal(executeMsg)
	if err != nil {
		return wasm.MsgExecuteContract{}, errors.Wrap(err, "marshal execute message")
	}
	return wasm.MsgExecuteContract{
		Sender:     sender,
		Contract:   contract,
		ExecuteMsg: rawExecuteMsg,
		Coins:      coins,
	}, nil
}
