package bdawns

import (
	"context"
	"errors"
	"github.com/DawnKosmos/bybit-go5"
	"github.com/DawnKosmos/bybit-go5/ws"
	"strings"
	"time"
)

/*
I love to use interfaces and same logic for every exchange.
dawns, are just packages that unifies the experience.

Ignore this for now :L
*/

type WsCandle struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	Volume    float64   `json:"volume"`
	StartTime time.Time `json:"startTime"`
	End       time.Time
	Ticker    string
	Finished  bool
}

type DawnBybit struct {
	b         *bybit.Client
	WsClients map[string]*ws.Stream
	testnet   bool
}

func tickerParse(in string) (category string, ticker string, err error) {
	ss := strings.Split(in, ".")
	if len(ss) != 2 {
		return "", "", errors.New("unknown ticker " + in)
	}
	switch ss[0] {
	case "s", "S":
		category = "spot"
	case "l", "L":
		category = "linear"
	case "i", "I":
		category = "inverse"
	case "o", "O":
		category = "option"
	default:
		return "", "", errors.New("unknown ticker " + in)

	}
	ticker = ss[1]
	return
}

func (d *DawnBybit) LiveKline(ticker string, resolution int64, parameters ...any) (chan WsCandle, error) {
	c, _, err := tickerParse(ticker)
	if err != nil {
		return nil, err
	}

	cl, ok := d.WsClients[c]
	if !ok {
		cl = ws.New(ws.Config{
			Ctx:           context.Background(),
			Expire:        10000,
			Endpoint:      ws.Enum(c),
			AutoReconnect: true,
			Debug:         false,
			TestNet:       d.testnet,
		})
		d.WsClients[c] = cl
	}

	var tick chan WsCandle
	return tick, nil
}

func convertResolution(res int64) string {
	if res >= 1440 {
		return "D"
	}

	return "D"
}
