package types

import (
	"encoding/json"

	"github.com/terra-project/core/x/wasm"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ cosmostypes.Msg = MsgExecuteContract{}

type MsgExecuteContract struct {
	Sender     cosmostypes.AccAddress `json:"sender"`
	Contract   cosmostypes.AccAddress `json:"contract"`
	ExecuteMsg json.RawMessage        `json:"execute_msg"`
	Coins      cosmostypes.Coins      `json:"coins"`
}

func (w MsgExecuteContract) Route() string {
	return wasm.MsgExecuteContract{}.Route()
}

func (w MsgExecuteContract) Type() string {
	return wasm.MsgExecuteContract{}.Type()
}

func (w MsgExecuteContract) ValidateBasic() error {
	return wasm.MsgExecuteContract{
		Sender:     w.Sender,
		Contract:   w.Contract,
		ExecuteMsg: []byte(w.ExecuteMsg),
		Coins:      w.Coins,
	}.ValidateBasic()
}

func (w MsgExecuteContract) GetSignBytes() []byte {
	return wasm.MsgExecuteContract{
		Sender:     w.Sender,
		Contract:   w.Contract,
		ExecuteMsg: []byte(w.ExecuteMsg),
		Coins:      w.Coins,
	}.GetSignBytes()
}

func (w MsgExecuteContract) GetSigners() []cosmostypes.AccAddress {
	return wasm.MsgExecuteContract{
		Sender:     w.Sender,
		Contract:   w.Contract,
		ExecuteMsg: []byte(w.ExecuteMsg),
		Coins:      w.Coins,
	}.GetSigners()
}
