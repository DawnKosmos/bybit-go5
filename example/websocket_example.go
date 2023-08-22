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
		Debug:         false,
	})

	err := client.PublicTrade("BTCUSDT", func(e *models.WsTrade) {
		// You can Setup What to do when the Event of tickers.BTCUSDT happens, e.g. Save in a DB
		fmt.Println(e.Data)
	})

	if err != nil {
		return err
	}

	for { // infite loop

	}
}
