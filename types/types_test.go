package types

import (
	"encoding/json"
	"fmt"
	"testing"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTokensHuman(t *testing.T) {
	Convey("init test", t, func() {
		testAddr, err := cosmostypes.AccAddressFromBech32("terra12avq876h9mn3wehchcezaafd4kdyjzer4u79kc")
		So(err, ShouldBeNil)
		testAmount := cosmostypes.NewIntWithDecimal(100, 6)

		rawData := fmt.Sprintf(
			`["%s","%s"]`,
			testAddr.String(), testAmount.String(),
		)
		Convey("#UnmarshalJSON", func() {
			var format TokensHuman
			So(json.Unmarshal([]byte(rawData), &format), ShouldBeNil)
			So(format.Addr[:], ShouldResemble, testAddr[:])
			So(format.Amount.String(), ShouldEqual, testAmount.String())
		})
		Convey("#MarshalJSON", func() {
			data := TokensHuman{
				Addr:   testAddr,
				Amount: testAmount,
			}
			marshaled, err := json.Marshal(data)
			So(err, ShouldBeNil)
			So(string(marshaled), ShouldEqual, fmt.Sprintf(
				`["%s","%s"]`,
				testAddr.String(), testAmount.String(),
			))
		})
		Convey("In Array", func() {
			data := []TokensHuman{{Addr: testAddr, Amount: testAmount}}
			marshaled, err := json.Marshal(data)
			So(err, ShouldBeNil)
			So(string(marshaled), ShouldEqual, fmt.Sprintf(`[%s]`, rawData))
		})
		Convey("In Struct", func() {
			data := struct{ Data TokensHuman }{
				Data: TokensHuman{Addr: testAddr, Amount: testAmount},
			}
			marshaled, err := json.Marshal(data)
			So(err, ShouldBeNil)
			So(string(marshaled), ShouldEqual, fmt.Sprintf(`{"Data":%s}`, rawData))
		})
	})
}
