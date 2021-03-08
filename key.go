package terra

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	terraauth "github.com/terra-project/core/x/auth"
)

type Key interface {
	AccAddress() cosmostypes.AccAddress
	ValAddress() cosmostypes.ValAddress
	PubKey() crypto.PubKey

	SignTx(msg terraauth.StdSignMsg) (terraauth.StdTx, error)
	MakeSignature(msg terraauth.StdSignMsg) (terraauth.StdSignature, error)
}

type rawKey struct {
	privKey secp256k1.PrivKeySecp256k1
}

func NewRawKey(hexed string) Key {
	var k secp256k1.PrivKeySecp256k1
	hex.Decode(k[:], []byte(hexed))
	return rawKey{privKey: k}
}

func (r rawKey) AccAddress() cosmostypes.AccAddress { return r.privKey.PubKey().Address().Bytes() }
func (r rawKey) ValAddress() cosmostypes.ValAddress { return r.privKey.PubKey().Address().Bytes() }
func (r rawKey) PubKey() crypto.PubKey              { return r.privKey.PubKey() }

func (r rawKey) SignTx(msg terraauth.StdSignMsg) (terraauth.StdTx, error) {
	sign, err := r.MakeSignature(msg)
	if err != nil {
		return terraauth.StdTx{}, errors.Wrap(err, "make signature")
	}

	signedTx := terraauth.NewStdTx(
		msg.Msgs,
		msg.Fee,
		[]terraauth.StdSignature{sign},
		msg.Memo,
	)
	return signedTx, nil
}

func (r rawKey) MakeSignature(msg terraauth.StdSignMsg) (terraauth.StdSignature, error) {
	sign, err := r.privKey.Sign(msg.Bytes())
	if err != nil {
		return terraauth.StdSignature{}, errors.Wrap(err, "sign with terra key")
	}
	return terraauth.StdSignature{
		PubKey:    r.privKey.PubKey(),
		Signature: sign,
	}, nil
}

type walletKey struct {
	name       string
	passphrase string
	info       keys.Info
	wallet     keys.Keybase
}

func NewWalletKey(name, passphrase string, wallet keys.Keybase) (Key, error) {
	info, err := wallet.Get(name)
	if err != nil {
		return nil, errors.Wrap(err, "get account info from keybase")
	}

	return walletKey{
		name:       name,
		passphrase: passphrase,
		info:       info,
		wallet:     wallet,
	}, nil
}

func (w walletKey) AccAddress() cosmostypes.AccAddress { return w.info.GetAddress().Bytes() }
func (w walletKey) ValAddress() cosmostypes.ValAddress { return w.info.GetAddress().Bytes() }
func (w walletKey) PubKey() crypto.PubKey              { return w.info.GetPubKey() }

func (w walletKey) SignTx(msg terraauth.StdSignMsg) (terraauth.StdTx, error) {
	sign, err := w.MakeSignature(msg)
	if err != nil {
		return terraauth.StdTx{}, errors.Wrap(err, "make signature")
	}

	signedTx := terraauth.NewStdTx(
		msg.Msgs,
		msg.Fee,
		[]terraauth.StdSignature{sign},
		msg.Memo,
	)
	return signedTx, nil
}

func (w walletKey) MakeSignature(msg terraauth.StdSignMsg) (terraauth.StdSignature, error) {
	return terraauth.MakeSignature(w.wallet, w.name, w.passphrase, msg)
}
