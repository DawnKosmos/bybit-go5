package models

type Subscription struct {
	Topic string `json:"topic"`          // Topic name
	Type  string `json:"type,omitempty"` // Data type. snapshot,delta
	Cs    int64  `json:"cs,omitempty"`   // Cross sequence
	Ts    int64  `json:"ts,omitempty"`   // The timestamp (ms) that the system generates the data
}

type WsOrderbook struct {
	Subscription
	Data struct {
		Symbol string      `json:"s"` // Symbol name
		Bid    [][2]string `json:"b"` // Bids. The element is sorted by price in descending order
		Ask    [][2]string `json:"a"` // Asks. The element is sorted by price in ascending order
	} `json:"data"` // Object
}

type WsTrade struct {
	Subscription
	Data []struct {
		Timestamp    int64  `json:"T"`           // The timestamp (ms) that the order is filled
		Symbol       string `json:"s"`           // Symbol name
		Side         string `json:"S"`           // Side. Buy,Sell
		Volume       string `json:"v"`           // Trade size
		Price        string `json:"p"`           // Trade price
		Direction    string `json:"L,omitempty"` // Direction of price change. Unique field for future
		Id           string `json:"i"`           // Trade ID
		IsBlockTrade bool   `json:"BT"`          // Whether it is a block trade order or not
	} `json:"data"` // Object
}

type WsTickerLinear struct {
	Subscription
	Data struct {
		Symbol            string `json:"symbol"`            // Symbol name
		TickDirection     string `json:"tickDirection"`     // Tick direction
		Price24hPcnt      string `json:"price24hPcnt"`      // Percentage change of market price in the last 24 hours
		LastPrice         string `json:"lastPrice"`         // Last price
		PrevPrice24h      string `json:"prevPrice24h"`      // Market price 24 hours ago
		HighPrice24h      string `json:"highPrice24h"`      // The highest price in the last 24 hours
		LowPrice24h       string `json:"lowPrice24h"`       // The lowest price in the last 24 hours
		PrevPrice1h       string `json:"prevPrice1h"`       // Market price an hour ago
		MarkPrice         string `json:"markPrice"`         // Mark price
		IndexPrice        string `json:"indexPrice"`        // Index price
		OpenInterest      string `json:"openInterest"`      // Open interest size
		OpenInterestValue string `json:"openInterestValue"` // Open interest value
		Turnover24h       string `json:"turnover24h"`       // Turnover for 24h
		Volume24h         string `json:"volume24h"`         // Volume for 24h
		NextFundingTime   string `json:"nextFundingTime"`   // Next funding timestamp (ms)
		FundingRate       string `json:"fundingRate"`       // Funding rate
		Bid1Price         string `json:"bid1Price"`         // Best bid price
		Bid1Size          string `json:"bid1Size"`          // Best bid size
		Ask1Price         string `json:"ask1Price"`         // Best ask price
		Ask1Size          string `json:"ask1Size"`          // Best ask size
	}
}

type WsTickerOption struct {
	Subscription
	Data struct {
		Symbol            string `json:"symbol"`            // Symbol name
		TickDirection     string `json:"tickDirection"`     // Tick direction
		Price24hPcnt      string `json:"price24hPcnt"`      // Percentage change of market price in the last 24 hours
		LastPrice         string `json:"lastPrice"`         // Last price
		PrevPrice24h      string `json:"prevPrice24h"`      // Market price 24 hours ago
		HighPrice24h      string `json:"highPrice24h"`      // The highest price in the last 24 hours
		LowPrice24h       string `json:"lowPrice24h"`       // The lowest price in the last 24 hours
		PrevPrice1h       string `json:"prevPrice1h"`       // Market price an hour ago
		MarkPrice         string `json:"markPrice"`         // Mark price
		IndexPrice        string `json:"indexPrice"`        // Index price
		OpenInterest      string `json:"openInterest"`      // Open interest size
		OpenInterestValue string `json:"openInterestValue"` // Open interest value
		Turnover24h       string `json:"turnover24h"`       // Turnover for 24h
		Volume24h         string `json:"volume24h"`         // Volume for 24h
		NextFundingTime   string `json:"nextFundingTime"`   // Next funding timestamp (ms)
		FundingRate       string `json:"fundingRate"`       // Funding rate
		Bid1Price         string `json:"bid1Price"`         // Best bid price
		Bid1Size          string `json:"bid1Size"`          // Best bid size
		Ask1Price         string `json:"ask1Price"`         // Best ask price
		Ask1Size          string `json:"ask1Size"`          // Best ask size
	}
}

type WsTickerSpot struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Type  string `json:"type"`
	Cs    int64  `json:"cs"`
	Data  struct {
		Symbol        string `json:"symbol"`
		LastPrice     string `json:"lastPrice"`
		HighPrice24H  string `json:"highPrice24h"`
		LowPrice24H   string `json:"lowPrice24h"`
		PrevPrice24H  string `json:"prevPrice24h"`
		Volume24H     string `json:"volume24h"`
		Turnover24H   string `json:"turnover24h"`
		Price24HPcnt  string `json:"price24hPcnt"`
		UsdIndexPrice string `json:"usdIndexPrice"`
	} `json:"data"`
}

type WsKline struct {
	Topic string `json:"topic"` // Topic name
	Type  string `json:"type"`  // Data type. snapshot
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Data  []struct {
		Start     int64  `json:"start"`     // The start timestamp (ms)
		End       int64  `json:"end"`       // The end timestamp (ms). It is current timestamp if it does not reach to the end time of candle
		Interval  string `json:"interval"`  // Kline interval
		Open      string `json:"open"`      // Open price
		Close     string `json:"close"`     // Close price
		High      string `json:"high"`      // Highest price
		Low       string `json:"low"`       // Lowest price
		Volume    string `json:"volume"`    // Trade volume
		Turnover  string `json:"turnover"`  // Turnover
		Confirm   bool   `json:"confirm"`   // Weather the tick is ended or not
		Timestamp int64  `json:"timestamp"` // The timestamp (ms) of the last matched order in the candle
	} `json:"data"` // Object
}

type WsLiquidation struct {
	Topic string `json:"topic"` // Topic name
	Type  string `json:"type"`  // Data type. snapshot
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Data  []struct {
		UpdateTime int64  `json:"updateTime"` // The updated timestamp (ms)
		Symbol     string `json:"symbol"`     // Symbol name
		Side       string `json:"side"`       // Order side. Buy,Sell
		Size       string `json:"size"`       // Executed size
		Price      string `json:"price"`      // Executed price
	} `json:"data"` // Object
}

type WsLTKline struct {
	Topic string `json:"topic"` // Topic name
	Type  string `json:"type"`  // Data type. snapshot
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Data  []struct {
		Start     int64  `json:"start"`     // The start timestamp (ms)
		End       int64  `json:"end"`       // The end timestamp (ms). It is current timestamp if it does not reach to the end time of candle
		Interval  string `json:"interval"`  // Kline interval
		Open      string `json:"open"`      // Open price
		Close     string `json:"close"`     // Close price
		High      string `json:"high"`      // Highest price
		Low       string `json:"low"`       // Lowest price
		Confirm   bool   `json:"confirm"`   // Weather the tick is ended or not
		Timestamp int64  `json:"timestamp"` // The timestamp (ms) of the last matched order in the candle
	} `json:"data"` // Object
}

type WsLTTicker struct {
	Topic string `json:"topic"` // Topic name
	Type  string `json:"type"`  // Data type. snapshot
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Data  []struct {
		Symbol       string `json:"symbol"`       // Symbol name
		Price24hPcnt string `json:"price24hPcnt"` // Market price change percentage in the past 24 hours
		LastPrice    string `json:"lastPrice"`    // The last price
		PrevPrice24h string `json:"prevPrice24h"` // Market price 24 hours ago
		HighPrice24h string `json:"highPrice24h"` // Highest price in the past 24 hours
		LowPrice24h  string `json:"lowPrice24h"`  // Lowest price in the past 24 hours
	} `json:"data"` // Object
}

type WsLTNav struct {
	Topic string `json:"topic"` // Topic name
	Type  string `json:"type"`  // Data type. snapshot
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Data  []struct {
		Time           int64  `json:"time"`           // The generated timestamp of nav
		Symbol         string `json:"symbol"`         // Symbol name
		Nav            string `json:"nav"`            // Net asset value
		BasketPosition string `json:"basketPosition"` // Total position value = basket value * total circulation
		Leverage       string `json:"leverage"`       // Leverage
		BasketLoan     string `json:"basketLoan"`     // Basket loan
		Circulation    string `json:"circulation"`    // Circulation
		Basket         string `json:"basket"`         // Basket
	} `json:"data"` // Object
}
