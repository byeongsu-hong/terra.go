package bind

import (
	"encoding/json"

	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

func GenerateExecuteMsg(
	sender cosmostypes.AccAddress,
	contract cosmostypes.AccAddress,
	executeMsg interface{},
	coins cosmostypes.Coins,
) (types.MsgExecuteContract, error) {
	rawExecuteMsg, err := json.Marshal(executeMsg)
	if err != nil {
		return types.MsgExecuteContract{}, errors.Wrap(err, "marshal execute message")
	}
	return types.MsgExecuteContract{
		Sender:     sender,
		Contract:   contract,
		ExecuteMsg: rawExecuteMsg,
		Coins:      coins,
	}, nil
}
