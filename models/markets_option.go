package models

type GetInstrumentsInfoResponseOption struct {
	Category       string `json:"category"`       // Product type
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used to pagination
	List           []struct {
		Symbol          string `json:"symbol"`          // Symbol name
		OptionsType     string `json:"optionsType"`     // Option type. Call, Put
		Status          string `json:"status"`          // Instrument status
		BaseCoin        string `json:"baseCoin"`        // Base coin
		QuoteCoin       string `json:"quoteCoin"`       // Quote coin
		SettleCoin      bool   `json:"settleCoin"`      // Settle coin
		LaunchTime      string `json:"launchTime"`      // Launch timestamp (ms)
		DeliveryTime    string `json:"deliveryTime"`    // Delivery timestamp (ms)
		DeliveryFeeRate string `json:"deliveryFeeRate"` // Delivery fee rate
		PriceFilter     struct {
			MinPrice string `json:"minPrice"` // Minimum order price
			MaxPrice string `json:"maxPrice"` // Maximum order price
			TickSize string `json:"tickSize"` // The step to increase/reduce order price
		} `json:"priceFilter"` // Price attributes
		LotSizeFilter struct {
			MaxOrderQty string `json:"maxOrderQty"` // Maximum order quantity
			MinOrderQty string `json:"minOrderQty"` // Minimum order quantity
			QtyStep     string `json:"qtyStep"`     // The step to increase/reduce order quantity
		} `json:"lotSizeFilter"` // Size attributes
	} `json:"list"` // Object
}

type GetTickersOptionResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol                 string `json:"symbol"`                 // Symbol name
		Bid1Price              string `json:"bid1Price"`              // Best bid price
		Bid1Size               string `json:"bid1Size"`               // Best bid size
		Bid1Iv                 string `json:"bid1Iv"`                 // Best bid iv
		Ask1Price              string `json:"ask1Price"`              // Best ask price
		Ask1Size               string `json:"ask1Size"`               // Best ask size
		Ask1Iv                 string `json:"ask1Iv"`                 // Best ask iv
		LastPrice              string `json:"lastPrice"`              // Last price
		HighPrice24h           string `json:"highPrice24h"`           // The highest price in the last 24 hours
		LowPrice24h            string `json:"lowPrice24h"`            // The lowest price in the last 24 hours
		MarkPrice              string `json:"markPrice"`              // Mark price
		IndexPrice             string `json:"indexPrice"`             // Index price
		MarkIv                 string `json:"markIv"`                 // Mark price iv
		UnderlyingPrice        string `json:"underlyingPrice"`        // Underlying price
		OpenInterest           string `json:"openInterest"`           // Open interest size
		Turnover24h            string `json:"turnover24h"`            // Turnover for 24h
		Volume24h              string `json:"volume24h"`              // Volume for 24h
		TotalVolume            string `json:"totalVolume"`            // Total volume
		TotalTurnover          string `json:"totalTurnover"`          // Total turnover
		Delta                  string `json:"delta"`                  // Delta
		Gamma                  string `json:"gamma"`                  // Gamma
		Vega                   string `json:"vega"`                   // Vega
		Theta                  string `json:"theta"`                  // Theta
		PredictedDeliveryPrice string `json:"predictedDeliveryPrice"` // Predicated delivery price. It has value when 30 min before delivery
		Change24h              string `json:"change24h"`              // The change in the last 24 hous
	} `json:"list"` // Object
}
