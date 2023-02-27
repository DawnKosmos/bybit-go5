package ws

/*
Part of Implementation was inspired by https://github.com/frankrap/deribit-api

	Bybit uses 5 Websocket Endpoints for their "unified" API
	I choose to implement only 1 endpoint per Stream Class
	So if you want to use all 5 endpoint, you have to create 5 Stream Classes with different Endpoint.

	Docs https://bybit-exchange.github.io/docs/v5/ws/connect
*/

var MAXTRYTIMES = 10000
var READLIMIT int64 = 32768

type WsLink int

const (
	SPOT WsLink = iota
	LINEAR
	INVERSE
	OPTION
	PRIVATE
)

// Hardcoded to Get The Right Link
func GetWsLink(id WsLink, testnet bool) string {
	var s string
	if testnet {
		switch id {
		case SPOT:
			s = "wss://stream-testnet.bybit.com/v5/public/spot"
		case LINEAR:
			s = "wss://stream-testnet.bybit.com/v5/public/linear"
		case INVERSE:
			s = "wss://stream-testnet.bybit.com/v5/public/inverse"
		case OPTION:
			s = "wss://stream-testnet.bybit.com/v5/public/option"
		case PRIVATE:
			s = "wss://stream-testnet.bybit.com/v5/private"
		}
	} else {
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
	}
	return s
}

func Enum(endpoint string) WsLink {
	switch endpoint {
	case "spot":
		return SPOT
	case "linear":
		return LINEAR
	case "inverse":
		return INVERSE
	case "option":
		return OPTION
	case "private":
		return PRIVATE
	}
	return LINEAR
}
