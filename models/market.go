package models

import (
	"strconv"
	"time"
)

type GetKlineRequest struct {
	Category string `url:"category"`
	Symbol   string `url:"symbol"`
	Interval string `url:"interval"`
	Start    int64  `url:",omitempty"` // Optional
	End      int64  `url:",omitempty"` // Optional
	Limit    int    `url:",omitempty"` // Default is 200
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
