package ws

import (
	"context"
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"testing"
)

var ever = true

func TestPing(t *testing.T) {
	ws := New(Config{
		Id:            "Test",
		Ctx:           context.Background(),
		Endpoint:      LINEAR,
		A:             nil,
		AutoReconnect: false,
		Debug:         true,
		TestNet:       false,
	})

	err := ws.Orderbook("BTCUSDT", 1, func(e *models.WsOrderbook) {
		fmt.Println(e.Ts, e.Data)
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = ws.PublicTrade("BTCUSDT", func(e *models.WsTrade) {
		fmt.Println(e.Ts, e.Data)
	})

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	for ever {

	}
}
