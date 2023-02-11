package ws_models

/*
orderbook.{depth}.{symbol} e.g., orderbook.1.BTCUSDT
publicTrade.{symbol}
tickers.{symbol}
kline.{interval}.{symbol} e.g., kline.30.BTCUSDT
liquidation.{symbol} e.g., liquidation.BTCUSDT
kline_lt.{interval}.{symbol} e.g., kline_lt.30.BTC3SUSDT
tickers_lt.{symbol} e.g.,tickers_lt.BTC3SUSDT
lt.{symbol} e.g.,lt.BTC3SUSDT


*/

type Pong struct {
	ReqId  string   `json:"req_id"`
	Op     string   `json:"op"`
	Args   []string `json:"args"`
	ConnId string   `json:"conn_id"`
}
