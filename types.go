package terra

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type Q map[string]interface{}

type BroadcastMode string

const (
	ModeBlock BroadcastMode = "block"
	ModeSync  BroadcastMode = "sync"
	ModeAsync BroadcastMode = "async"
)

var (
	_ json.Marshaler                = TerraAddress{}
	_ json.Unmarshaler              = &TerraAddress{}
	_ dynamodbattribute.Marshaler   = TerraAddress{}
	_ dynamodbattribute.Unmarshaler = &TerraAddress{}
)

type TerraAddress struct{ cosmostypes.AccAddress }

func (e TerraAddress) ToEthHash() ethcommon.Hash {
	hash := ethcommon.Hash{}
	copy(hash[:], e.Bytes())
	return hash
}

func (e *TerraAddress) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	val := aws.StringValue(value.S)
	addr, err := cosmostypes.AccAddressFromBech32(val)
	if err != nil {
		return errors.Wrap(err, "conv bech32 to addr")
	}

	e.AccAddress = addr
	return nil
}

func (e TerraAddress) MarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	value.S = aws.String(e.AccAddress.String())
	return nil
}

var (
	_ json.Marshaler                = TerraAddress{}
	_ json.Unmarshaler              = &TerraAddress{}
	_ dynamodbattribute.Marshaler   = TerraKey{}
	_ dynamodbattribute.Unmarshaler = &TerraKey{}
)

type TerraKey struct {
	secret []byte
	secp256k1.PrivKeySecp256k1
}

func KeyWithSecret(secret []byte, key secp256k1.PrivKeySecp256k1) (TerraKey, error) {
	return TerraKey{secret: secret, PrivKeySecp256k1: key}, nil
}

func (e *TerraKey) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	if e.secret != nil {
		encrypted := value.B
		decrypted, err := DecryptAESGCM(e.secret, encrypted[:NonceSize], encrypted[NonceSize:])
		if err != nil {
			return errors.Wrap(err, "decrypt terra key")
		}

		copy(e.PrivKeySecp256k1[:], decrypted)
	} else {
		copy(e.PrivKeySecp256k1[:], value.B)
	}
	return nil
}

func (e TerraKey) MarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	if e.secret != nil {
		encrypted, err := EncryptAESGCM(e.secret, e.PrivKeySecp256k1[:])
		if err != nil {
			return errors.Wrap(err, "encrypt terra key")
		}
		value.B = encrypted
	} else {
		value.B = e.PrivKeySecp256k1.Bytes()
	}
	return nil
}

var (
	_ dynamodbattribute.Marshaler   = HDPath{}
	_ dynamodbattribute.Unmarshaler = &HDPath{}
)

type HDPath struct{ *hd.BIP44Params }

func (e *HDPath) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	val := aws.StringValue(value.S)
	params, err := hd.NewParamsFromPath(val)
	if err != nil {
		return errors.Wrap(err, "conv path to params")
	}

	e.BIP44Params = params
	return nil
}

func (e HDPath) MarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	value.S = aws.String(e.String())
	return nil
}
