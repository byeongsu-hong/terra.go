package terra

import (
	"encoding/hex"

	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	cosmosclienttx "github.com/cosmos/cosmos-sdk/client/tx"
	cosmoscrypto "github.com/cosmos/cosmos-sdk/crypto"
	cosmoscryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

type Key interface {
	AccAddress() cosmostypes.AccAddress
	ValAddress() cosmostypes.ValAddress
	PubKey() cosmoscryptotypes.PubKey

	SignTx(cosmosclienttx.Factory, cosmosclient.TxBuilder, bool) error
}

const (
	rawKeyPassphrase = "raw_passphrase"
	rawKeyUID        = "raw_uid"
)

type rawKey struct {
	keyring keyring.Keyring
}

func NewRawKey(hexed string) Key {
	var k secp256k1.PrivKey
	if _, err := hex.Decode(k.Key[:], []byte(hexed)); err != nil {
		panic(err)
	}

	armoredPrivKey := cosmoscrypto.EncryptArmorPrivKey(cosmoscryptotypes.PrivKey(&k), rawKeyPassphrase, string(cosmoscryptohd.Secp256k1.Name()))
	keys := keyring.NewInMemory()
	if err := keys.ImportPrivKey(rawKeyUID, armoredPrivKey, rawKeyPassphrase); err != nil {
		panic(err)
	}
	return rawKey{keyring: keys}
}

func (r rawKey) fetchKey() keyring.Info {
	info, err := r.keyring.Key(rawKeyUID)
	if err != nil {
		panic(err)
	}
	return info
}
func (r rawKey) AccAddress() cosmostypes.AccAddress {
	return r.fetchKey().GetPubKey().Address().Bytes()
}
func (r rawKey) ValAddress() cosmostypes.ValAddress {
	return r.fetchKey().GetPubKey().Address().Bytes()
}
func (r rawKey) PubKey() cosmoscryptotypes.PubKey { return r.fetchKey().GetPubKey() }

func (r rawKey) SignTx(factory cosmosclienttx.Factory, builder cosmosclient.TxBuilder, overwriteSig bool) error {
	return cosmosclienttx.Sign(factory.WithKeybase(r.keyring), rawKeyUID, builder, overwriteSig)
}
