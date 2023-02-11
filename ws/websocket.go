package ws

import "context"

/*
Implementation inspired by https://github.com/frankrap/deribit-api

	Who uses 5 Endpoints for their unified API xd
	Anyway I try my best to solve this.


Docs https://bybit-exchange.github.io/docs/v5/ws/connect
*/

// Ping intervall 20 seconds

/*
	const (
	SPOT    = "wss://stream.bybit.com/v5/public/spot"
	LINEAR  = "wss://stream.bybit.com/v5/public/linear"
	INVERSE = "wss://stream.bybit.com/v5/public/inverse"
	OPTION  = "wss://stream.bybit.com/v5/public/option"
	PRIVATE = "wss://stream.bybit.com/v5/private"
)


const (
	SPOTTEST    = "wss://stream-testnet.bybit.com/v5/public/spot"
	LINEARTEST  = "wss://stream-testnet.bybit.com/v5/public/linear"
	INVERSETEST = "wss://stream-testnet.bybit.com/v5/public/inverse"
	OPTIONTEST  = "wss://stream-testnet.bybit.com/v5/public/option"
	PRIVATETEST = "wss://stream-testnet.bybit.com/v5/private"
)




*/

type WsLink int

var MAXTRYTIMES = 10000
var READLIMIT int64 = 32768

const (
	SPOT WsLink = iota
	LINEAR
	INVERSE
	OPTION
	PRIVATE
)

func GetWsLink(id WsLink) string {
	var s string
	switch id {
	case SPOT:
		s = "wss://stream.bybit.com/v5/public/spot"
	case LINEAR:
		s = "wss://stream.bybit.com/v5/public/linear"
	case INVERSE:
		s = "wss://stream.bybit.com/v5/public/inverse"
	case OPTION:
		s = "wss://stream.bybit.com/v5/public/option"
	case PRIVATE:
		s = "wss://stream.bybit.com/v5/private"
	}
	return s
}

type Client struct {
	ctx context.Context
}

func New(c *Client) error {
	return nil
}
