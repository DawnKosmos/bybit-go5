package models

type GetInstrumentsInfoResponseLinear struct {
	Category       string `json:"category"`       // Product type
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used to pagination
	List           []struct {
		Symbol          string `json:"symbol"`          // Symbol name
		ContractType    string `json:"contractType"`    // Contract type
		Status          string `json:"status"`          // Instrument status
		BaseCoin        string `json:"baseCoin"`        // Base coin
		QuoteCoin       string `json:"quoteCoin"`       // Quote coin
		LaunchTime      string `json:"launchTime"`      // Launch timestamp (ms)
		DeliveryTime    string `json:"deliveryTime"`    // Delivery timestamp (ms). Valid for Inverse Futures
		DeliveryFeeRate string `json:"deliveryFeeRate"` // Delivery fee rate. Valid for Inverse Futures
		PriceScale      string `json:"priceScale"`      // Price scale
		LeverageFilter  struct {
			MinLeverage  string `json:"minLeverage"`  // Minimum leverage
			MaxLeverage  string `json:"maxLeverage"`  // Maximum leverage
			LeverageStep string `json:"leverageStep"` // The step to increase/reduce leverage
		} `json:"leverageFilter"` // Leverage attributes
		PriceFilter struct {
			MinPrice string `json:"minPrice"` // Minimum order price
			MaxPrice string `json:"maxPrice"` // Maximum order price
			TickSize string `json:"tickSize"` // The step to increase/reduce order price
		} `json:"priceFilter"` // Price attributes
		LotSizeFilter struct {
			MaxOrderQty         string `json:"maxOrderQty"`         // Maximum order quantity
			MinOrderQty         string `json:"minOrderQty"`         // Minimum order quantity
			QtyStep             string `json:"qtyStep"`             // The step to increase/reduce order quantity
			PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"` // Maximum order qty for PostOnly order
		} `json:"lotSizeFilter"` // Size attributes
		UnifiedMarginTrade bool   `json:"unifiedMarginTrade"` // Whether to support unified margin trade
		FundingInterval    int64  `json:"fundingInterval"`    // Funding interval (minute)
		SettleCoin         string `json:"settleCoin"`         // Settle coin
	} `json:"list"` // Object
}

type GetTickersLinearResponse struct {
	List []struct {
		Symbol                 string `json:"symbol"`                 // Symbol name
		LastPrice              string `json:"lastPrice"`              // Last price
		IndexPrice             string `json:"indexPrice"`             // Index price
		MarkPrice              string `json:"markPrice"`              // Mark price
		PrevPrice24h           string `json:"prevPrice24h"`           // Market price 24 hours ago
		Price24hPcnt           string `json:"price24hPcnt"`           // Percentage change of market price relative to 24h
		HighPrice24h           string `json:"highPrice24h"`           // The highest price in the last 24 hours
		LowPrice24h            string `json:"lowPrice24h"`            // The lowest price in the last 24 hours
		PrevPrice1h            string `json:"prevPrice1h"`            // Market price an hour ago
		OpenInterest           string `json:"openInterest"`           // Open interest size
		OpenInterestValue      string `json:"openInterestValue"`      // Open interest value
		Turnover24h            string `json:"turnover24h"`            // Turnover for 24h
		Volume24h              string `json:"volume24h"`              // Volume for 24h
		FundingRate            string `json:"fundingRate"`            // Funding rate
		NextFundingTime        string `json:"nextFundingTime"`        // Next funding time (ms)
		PredictedDeliveryPrice string `json:"predictedDeliveryPrice"` // Predicated delivery price. It has value when 30 min before delivery
		BasisRate              string `json:"basisRate"`              // Basis rate
		DeliveryFeeRate        string `json:"deliveryFeeRate"`        // Delivery fee rate
		DeliveryTime           string `json:"deliveryTime"`           // Delivery timestamp (ms)
		Ask1Size               string `json:"ask1Size"`               // Best ask size
		Bid1Price              string `json:"bid1Price"`              // Best bid price
		Ask1Price              string `json:"ask1Price"`              // Best ask price
		Bid1Size               string `json:"bid1Size"`               // Best bid size
	} `json:"list"` // Object
}
