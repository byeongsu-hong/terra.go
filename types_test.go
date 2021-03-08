package terra

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
)

func TestTerraKey(t *testing.T) {
	Convey("init test", t, func() {
		encryptionKey := sha3.Sum256([]byte("test"))

		Convey("#marshal & #unmarshal", func() {
			key, err := KeyWithSecret(encryptionKey[:], secp256k1.GenPrivKey())
			So(err, ShouldBeNil)

			marshaled, err := dynamodbattribute.Marshal(key)
			So(err, ShouldBeNil)

			newKey, err := KeyWithSecret(encryptionKey[:], secp256k1.PrivKeySecp256k1{})
			So(err, ShouldBeNil)
			So(dynamodbattribute.Unmarshal(marshaled, &newKey), ShouldBeNil)

			So(
				key.PrivKeySecp256k1[:],
				ShouldResemble,
				newKey.PrivKeySecp256k1[:],
			)
			So(
				key.PrivKeySecp256k1.PubKey().Bytes(),
				ShouldResemble,
				newKey.PrivKeySecp256k1.PubKey().Bytes(),
			)
		})
	})
}
