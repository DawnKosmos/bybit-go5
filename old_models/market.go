package models

import (
	"strconv"
	"time"
)

type GetKlineRequest struct {
	Category string `url:"category"`
	Symbol   string `url:"symbol"`
	Interval string `url:"interval"`
	Start    int64  `url:"start,omitempty"` // [optional] The start timestamp (ms)
	End      int64  `url:"end,omitempty"`   // [optional] The end timestamp (ms)
	Limit    int    `url:"limit,omitempty"` // [optional] Default is 200
}

type GetKlineResponse struct {
	Category string      `json:"category"`
	Symbol   string      `json:"symbol"`
	List     [][7]string `json:"list"`
}

type Candle struct {
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Close     float64   `json:"close"`
	Low       float64   `json:"low"`
	Volume    float64   `json:"volume"`
	StartTime time.Time `json:"startTime"`
}

func (r *GetKlineResponse) ToCandle() []Candle {
	ch := make([]Candle, 0, len(r.List))
	for _, v := range r.List {
		var c Candle
		c.Open, _ = strconv.ParseFloat(v[1], 64)
		c.High, _ = strconv.ParseFloat(v[2], 64)
		c.Low, _ = strconv.ParseFloat(v[3], 64)
		c.Close, _ = strconv.ParseFloat(v[4], 64)
		c.Volume, _ = strconv.ParseFloat(v[5], 64)
		t, _ := strconv.ParseInt(v[0], 10, 64)
		c.StartTime = time.Unix(t/1000, 0)

		ch = append(ch, c)
	}
	return ch
}

type GetMarkPricelineResponse struct {
	Category string      `json:"category"`
	Symbol   string      `json:"symbol"`
	List     [][5]string `json:"list"`
}

type GetInstrumentsInfoRequest struct {
	Category string `url:"category"`           // Product type. spot,linear,inverse,option
	Symbol   string `url:"symbol,omitempty"`   // [optional]Symbol name
	BaseCoin string `url:"baseCoin,omitempty"` // [optional]Base coin. linear,inverse,option only
	Limit    int64  `url:"limit,omitempty"`    // [optional]Limit for data size per page. [1, 1000]. Default: 500
	Cursor   string `url:"cursor,omitempty"`   // [optional]Cursor. Used for pagination
}

type GetOrderbookRequest struct {
	Category string `url:"category"`        // Product type. spot, linear, inverse, option
	Symbol   string `url:"symbol"`          // Symbol name
	Limit    int64  `url:"limit,omitempty"` // [optional]Limit size for each bid and ask spot: [1, 50]. Default: 1. linear&inverse: [1, 200]. Default: 25. option: [1, 25]. Default: 1.
}

type GetOrderbookResponse struct {
	S  string      `json:"s"`  // Symbol name
	B  [][2]string `json:"b"`  // Bid, buyer. Sort by price desc
	A  [][2]string `json:"a"`  // Ask, seller. Order by price asc
	Ts int64       `json:"ts"` // The timestamp (ms) that the system generates the data
	U  int64       `json:"u"`  // Update ID, is always in sequence For future, it is corresponding to u in the wss 200-level orderbook For spot, it is corresponding to u in the wss 50-level orderbook
}

type GetTickersRequest struct {
	Category string `url:"category"`           // Product type. spot,linear,inverse,option
	Symbol   string `url:"symbol,omitempty"`   // [optional]Symbol name
	BaseCoin string `url:"baseCoin,omitempty"` // [optional]Base coin. For option only
	ExpDate  string `url:"expDate,omitempty"`  // [optional]Expiry date. e.g., 25DEC22. For option only
}

type GetFundingRateHistoryRequest struct {
	Category  string `url:"category"`            // Product type. linear,inverse
	Symbol    string `url:"symbol"`              // Symbol name
	StartTime int64  `url:"startTime,omitempty"` // [optional]The start timestamp (ms)
	EndTime   int64  `url:"endTime,omitempty"`   // [optional]The end timestamp (ms)
	Limit     int64  `url:"limit,omitempty"`     // [optional]Limit for data size per page. [1, 200]. Default: 200
}

type GetFundingRateHistoryResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol               string `json:"symbol"`               // Symbol name
		FundingRate          string `json:"fundingRate"`          // Funding rate
		FundingRateTimestamp string `json:"fundingRateTimestamp"` // Funding rate timestamp (ms)
	} `json:"list"` // Object
}

type GetPublicTradingHistoryRequest struct {
	Category   string `url:"category"`             // Product type. spot,linear,inverse,option
	Symbol     string `url:"symbol"`               // Symbol name
	BaseCoin   string `url:"baseCoin,omitempty"`   // [optional]Base coin. For option only. If not passed, return BTC data by default
	OptionType string `url:"optionType,omitempty"` // [optional]Option type. Send or Put. For option only
	Limit      int64  `url:"limit,omitempty"`      // [optional]Limit for data size per page. spot: [1,60], default: 60. others: [1,1000], default: 500
}

type GetPublicTradingHistoryResponse struct {
	Category string `json:"category"` // Products category
	List     []struct {
		ExecId       string `json:"execId"`       // Execution ID
		Symbol       string `json:"symbol"`       // Symbol name
		Price        string `json:"price"`        // Trade price
		Size         string `json:"size"`         // Trade size
		Side         string `json:"side"`         // Buy, Sell
		Time         string `json:"time"`         // Trade time (ms)
		IsBlockTrade bool   `json:"isBlockTrade"` // Whether the trade is block trade
	} `json:"list"` // Object
}

type GetOpenInterestRequest struct {
	Category     string `url:"category"`            // Product type. linear,inverse
	Symbol       string `url:"symbol"`              // Symbol name
	IntervalTime string `url:"intervalTime"`        // Interval. 5min,15min,30min,1h,4h,1d
	StartTime    int64  `url:"startTime,omitempty"` // [optional]The start timestamp (ms)
	EndTime      int64  `url:"endTime,omitempty"`   // [optional]The end timestamp (ms)
	Limit        int64  `url:"limit,omitempty"`     // [optional]Limit for data size per page. [1, 200]. Default: 50
	Cursor       string `url:"cursor,omitempty"`    // [optional]Cursor. Used to paginate
}

type GetOpenInterestResponse struct {
	Symbol   string `json:"symbol"`
	Category string `json:"category"`
	List     []struct {
		OpenInterest string `json:"openInterest"`
		Timestamp    string `json:"timestamp"`
	} `json:"list"`
	NextPageCursor string `json:"nextPageCursor"`
}

type GetHistoricalVolatilityRequest struct {
	Category  string `url:"category"`            // Product type. option
	BaseCoin  string `url:"baseCoin,omitempty"`  // [optional]Base coin. Default: return BTC data
	Period    int64  `url:"period,omitempty"`    // [optional]Period
	StartTime int64  `url:"startTime,omitempty"` // [optional]The start timestamp (ms)
	EndTime   int64  `url:"endTime,omitempty"`   // [optional]The end timestamp (ms)
}

type GetHistoricalVolatilityResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Period int64  `json:"period"` // Period
		Value  string `json:"value"`  // Volatility
		Time   string `json:"time"`   // Timestamp (ms)
	} `json:"list"` // Object
}

type GetInsuranceRequest struct {
	Coin string `url:"coin,omitempty"` // [optional]coin. Default: return all insurance coins
}

type GetInsuranceResponse struct {
	UpdatedTime string `json:"updatedTime"` // Data updated time (ms)
	List        []struct {
		Coin    string `json:"coin"`    // Coin
		Balance string `json:"balance"` // Balance
		Value   string `json:"value"`   // USD value
	} `json:"list"` // Object
}

type GetRiskLimitRequest struct {
	Category string `url:"category,omitempty"` // [optional]Product type. linear,inverse
	Symbol   string `url:"symbol,omitempty"`   // [optional]Symbol name
}

type GetRiskLimitResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Id             int64  `json:"id"`             // Risk ID
		Symbol         string `json:"symbol"`         // Symbol name
		RiskLimitValue string `json:"riskLimitValue"` // Position limit
		Section        []struct {
		} `json:"section"` // section array
		IsLowestRisk int64  `json:"isLowestRisk"` // 1: true, 0: false
		MaxLeverage  string `json:"maxLeverage"`  // Allowed max leverage
	} `json:"list"` // Object
}

type GetOptionDeliveryPriceRequest struct {
	Category string `url:"category"`           // Product type. option
	Symbol   string `url:"symbol,omitempty"`   // [optional]Symbol name
	BaseCoin string `url:"baseCoin,omitempty"` // [optional]Base coin. Default: BTC
	Limit    int64  `url:"limit,omitempty"`    // [optional]Limit for data size per page. [1, 200]. Default: 50
	Cursor   string `url:"cursor,omitempty"`   // [optional]Cursor. Used for pagination
}

type GetOptionDeliveryPriceResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol        string `json:"symbol"`        // Symbol name
		DeliveryPrice string `json:"deliveryPrice"` // Delivery price
		DeliveryTime  string `json:"deliveryTime"`  // Delivery timestamp (ms)
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}
