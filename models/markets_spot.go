package models

type GetInstrumentsInfoResponseSpot struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol        string `json:"symbol"`     // Symbol name
		BaseCoin      string `json:"baseCoin"`   // Base coin
		QuoteCoin     string `json:"quoteCoin"`  // Quote coin
		Innovation    string `json:"innovation"` // Whether to belong to innovation. `0`: false, `1`: true
		Status        string `json:"status"`     // Instrument status
		LotSizeFilter struct {
			BasePrecision  string `json:"basePrecision"`  // The precision of base coin
			QuotePrecision string `json:"quotePrecision"` // The precision of quote coin
			MinOrderQty    string `json:"minOrderQty"`    // Minimum order quantity
			MaxOrderQty    string `json:"maxOrderQty"`    // Maximum order quantity
			MinOrderAmt    string `json:"minOrderAmt"`    // Minimum order amount
			MaxOrderAmt    string `json:"maxOrderAmt"`    // Maximum order amount
		} `json:"lotSizeFilter"` // Size attributes
		PriceFilter struct {
			TickSize string `json:"tickSize"` // The step to increase/reduce order price
		} `json:"priceFilter"` // Price attributes
	} `json:"list"` // Object
}

type GetTickersSpotResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol        string `json:"symbol"`        // Symbol name
		Bid1Price     string `json:"bid1Price"`     // Best bid price
		Bid1Size      string `json:"bid1Size"`      // Best bid size
		Ask1Price     string `json:"ask1Price"`     // Best ask price
		Ask1Size      string `json:"ask1Size"`      // Best ask size
		LastPrice     string `json:"lastPrice"`     // Last price
		PrevPrice24h  string `json:"prevPrice24h"`  // Market price 24 hours ago
		Price24hPcnt  string `json:"price24hPcnt"`  // Percentage change of market price relative to 24h
		HighPrice24h  string `json:"highPrice24h"`  // The highest price in the last 24 hours
		LowPrice24h   string `json:"lowPrice24h"`   // The lowest price in the last 24 hours
		Turnover24h   string `json:"turnover24h"`   // Turnover for 24h
		Volume24h     string `json:"volume24h"`     // Volume for 24h
		UsdIndexPrice string `json:"usdIndexPrice"` // USD index price. It can be empty
	} `json:"list"` // Object
}
