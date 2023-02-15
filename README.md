# bybit-go5

# WIP

This Golang library implements the Bybit API v5

[Docs](https://bybit-exchange.github.io/docs/v5/intro)

REST and websocket api are implemented. I generated most structs from the documentation,
which has some type error, so some API Calls produce error! 
If you notice one Write. Please write an issue and I fix it.

===

Names of Structs and Functions and their comments are generated from the above Documentation.

Due to limitations in the GoLang Type Systems some endpoints have 3 Functions. 

1 For each possible Response struct from the Api.

E.g. GetInstrumentsInfo depending on the **Category** linear/inverse, spot or options gives different response struct





### Example
```go
package main

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5"
	"github.com/DawnKosmos/bybit-go5/models"
)

func main() {
	a := &bybit.Account{
		PublicKey: "",
		SecretKey: "",
	}
	c, _ := bybit.New(nil, bybit.URL, a, true)
	resp, err := c.GetKline(models.GetKlineRequest{
		Category: "linear",
		Symbol:   "BTCUSDT",
		Interval: "240",
		Limit:    200,
	})
	if err != nil {
		return
	}
	for _, v := range resp.ToCandle() {
		fmt.Println(v)
	}

}
```

WS Example
```go
package main

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"github.com/DawnKosmos/bybit-go5/ws"
)

func WSExample() error {
	client := ws.New(ws.Config{
		Id:            "Test Implementation",
		Ctx:           nil,
		Endpoint:      ws.LINEAR, // Use linear datapoint
		AutoReconnect: true,
		Debug:         true,
	})

	err := client.Kline("BTCUSDT", "1", func(e *models.WsKline) {
		// This Function gets triggered every time this Kline Event is received by the WS
		fmt.Println(e.Data)
	})

	if err != nil {
		return err
	}

	for { // infite loop

	}
}
```

