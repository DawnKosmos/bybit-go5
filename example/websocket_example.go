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

	var counter int
	err := client.TickerLinear("BTCUSDT", func(e *models.WsTickerLinear) {
		// You can Setup What to do when the Event of tickers.BTCUSDT happens, e.g. Save in a DB
		counter++
		fmt.Println(counter)
	})

	if err != nil {
		return err
	}

	for { // infite loop

	}
}
