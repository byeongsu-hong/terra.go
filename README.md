# terra.go

## Prerequisites

* `>= go 1.15`

## Install

``` bash
$ go get github.com/frostornge/terra-go
```

## Packages

* bind
  * contract binding helper [ref](./interface/anchor/money-market/market)
* httpclient
  * http middleware to process codec encoded respones
* interface
  * main contract binding (WIP)
* service
  * LCD biding
* types

## How to use

``` golang
import (
  "log"
  
  terra "github.com/frostornge/terra-go"
  "github.com/frostornge/terra-go/httpclient"
  "github.com/frostornge/terra-go/types"
  
  "github.com/cosmos/cosmos-sdk/crypto/keys"
  cosmostypes "github.com/cosmos/cosmos-sdk/types"
  terraapp "github.com/terra-project/core/app"
  terratypes "github.com/terra-project/core/types"
)

func must(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  lcdClient := terra.NewClient(
    httpclient.New(
      terraapp.MakeCodec(), 
      "https://tequila-lcd.terra.dev",
    )
  )
 
  var (
    acc terra.Account
    key terra.Key
    err error
  )
 
  // with hexed ECDSA private key
  {
    key = terra.NewRawKey("{private_key}")
    acc, err = terra.NewAccount(context.Background(), lcdClient, key)
    must(err)
  }
  
  // with KeyBase
  {
    keyBase := keys.NewInMemory()
    _, err = keyBase.CreateAccount(
      "example", "{mnemonic}", 
      keys.DefaultBIP39Passphrase, "{passphrase}", 
      terratypes.FullFundraiserPath, keys.Secp256k1,
    )
    must(err)
    
    key, err = terra.NewWalletKey("example", "{passphrase}", keyBase)
    must(err)
    
    acc, err = terra.NewAccount(context.Background(), lcdClient, key)
    must(err)
  }


  // example#1 - bank send
  {
    var (
      to     cosmostypes.AccAddress
      amount cosmostypes.Int
    )
  
    tx, msg, err := acc.CreateAndSignTx(context.Background(), terra.CreateTxOptions{
      Msgs: []cosmostypes.Msg{
        bank.MsgSend{
          FromAddress: acc.GetAddress(),
          ToAddress: to,
          Amount: cosmostypes.Coins{{
            Denom: "uusd",
            Amount: amount,
          }},
        }
      },
    })
    must(err)
    
    txResp, err := lcdClient.Transaction().BroadcastTx(ctx, tx, types.ModeBlock)
    must(err)
    
    log.Println(txResp.TxHash)
  }
  
}

```
