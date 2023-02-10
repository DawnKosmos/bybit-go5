package models

type PlaceOrderRequest struct {
	Category         string `json:"category"`                   // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol           string `json:"symbol"`                     // Symbol name
	IsLeverage       int64  `json:"isLeverage,omitempty"`       // [optional]Whether to borrow. Valid for spot only. 0(default): false, 1: true
	Side             string `json:"side"`                       // Buy, Sell
	OrderType        string `json:"orderType"`                  // Market, Limit
	Qty              string `json:"qty"`                        // Order quantity
	Price            string `json:"price,omitempty"`            // [optional]Order price. If you have net position, price needs to be greater than liquidation price
	TriggerDirection int64  `json:"triggerDirection,omitempty"` // [optional]Conditional order param. Used to identify the expected direction of the conditional order.
	OrderFilter      string `json:"orderFilter,omitempty"`      // [optional]Valid for spot only. Order,tpslOrder. If not passed, Order by default
	TriggerPrice     string `json:"triggerPrice,omitempty"`     // [optional] For futures, it is the conditional order trigger price. If you expect the price to rise to trigger your conditional order, make sure: triggerPrice > market price Else, triggerPrice < market price For spot, it is the TP/SL order trigger price
	TriggerBy        string `json:"triggerBy,omitempty"`        // [optional]Conditional order param. Trigger price type. LastPrice, IndexPrice, MarkPrice
	OrderIv          string `json:"orderIv,omitempty"`          // [optional]Implied volatility. option only. Pass the real value, e.g for 10%, 0.1 should be passed. orderIv has a higher priority when price is passed as well
	TimeInForce      string `json:"timeInForce,omitempty"`      // [optional]Time in force Market order will use IOC directly If not passed, GTC is used by default
	PositionIdx      int64  `json:"positionIdx,omitempty"`      // [optional]Used to identify positions in different position modes. Under hedge-mode, this param is required
	OrderLinkId      string `json:"orderLinkId,omitempty"`      // [optional]User customised order ID. Check Bybit Documentation
	TakeProfit       string `json:"takeProfit,omitempty"`       // [optional]Take profit price. Only takes effect upon opening the position
	StopLoss         string `json:"stopLoss,omitempty"`         // [optional]Stop loss price. Only takes effect upon opening the position
	TpTriggerBy      string `json:"tpTriggerBy,omitempty"`      // [optional]The price type to trigger take profit. MarkPrice, IndexPrice, default: LastPrice
	SlTriggerBy      string `json:"slTriggerBy,omitempty"`      // [optional]The price type to trigger stop loss. MarkPrice, IndexPrice, default: LastPrice
	ReduceOnly       bool   `json:"reduceOnly,omitempty"`       // [optional]What is a reduce-only order? true means your position can only reduce in size if this order is triggered. When reduce_only is true, take profit/stop loss cannot be set
	CloseOnTrigger   bool   `json:"closeOnTrigger,omitempty"`   // [optional]What is a close on trigger order? For a closing order. It can only reduce your position, not increase it. If the account has insufficient available balance when the closing order is triggered, then other active orders of similar contracts will be cancelled or reduced. It can be used to ensure your stop loss reduces your position regardless of current available margin.
	Mmp              bool   `json:"mmp,omitempty"`              // [optional]Market maker protection. option only. true means set the order as a market maker protection order. What is mmp?
}

type PlaceOrderResponse struct {
	OrderId     string `json:"orderId"`     // Order ID
	OrderLinkId string `json:"orderLinkId"` // User customised order ID
}

type AmendOrderRequest struct {
	Category     string `json:"category"`               // Product type
	Symbol       string `json:"symbol"`                 // Symbol name
	OrderId      string `json:"orderId,omitempty"`      // [optional]Order ID. Either orderId or orderLinkId is required
	OrderLinkId  string `json:"orderLinkId,omitempty"`  // [optional]User customised order ID. Either orderId or orderLinkId is required
	OrderIv      string `json:"orderIv,omitempty"`      // [optional]Implied volatility. option only. Pass the real value, e.g for 10%, 0.1 should be passed
	TriggerPrice string `json:"triggerPrice,omitempty"` // [optional]If you expect the price to rise to trigger your conditional order, make sure:
	Qty          string `json:"qty,omitempty"`          // [optional]Order quantity after modification. Do not pass it if not modify the qty
	Price        string `json:"price,omitempty"`        // [optional]Order price after modification. Do not pass it if not modify the price
	TakeProfit   string `json:"takeProfit,omitempty"`   // [optional]Take profit price after modification. Do not pass it if you do not want to modify the take profit
	StopLoss     string `json:"stopLoss,omitempty"`     // [optional]Stop loss price after modification. Do not pass it if you do not want to modify the stop loss
	TpTriggerBy  string `json:"tpTriggerBy,omitempty"`  // [optional]The price type to trigger take profit
	SlTriggerBy  string `json:"slTriggerBy,omitempty"`  // [optional]The price type to trigger stop loss
	TriggerBy    string `json:"triggerBy,omitempty"`    // [optional]Trigger price type
}

type AmendOrderResponse struct {
	OrderId     string `json:"orderId"`     // Order ID
	OrderLinkId string `json:"orderLinkId"` // User customised order ID
}

type CancelOrderRequest struct {
	Category    string `json:"category"`              // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol      string `json:"symbol"`                // Symbol name
	OrderId     string `json:"orderId,omitempty"`     // [optional]Order ID. Either orderId or orderLinkId is required
	OrderLinkId string `json:"orderLinkId,omitempty"` // [optional]User customised order ID. Either orderId or orderLinkId is required
	OrderFilter string `json:"orderFilter,omitempty"` // [optional]Valid for spot only. Order,tpslOrder. If not passed, Order by default
}

type CancelOrderResponse struct {
	OrderId     string `json:"orderId"`     // Order ID
	OrderLinkId string `json:"orderLinkId"` // User customised order ID
}

type GetOpenOrdersRequest struct {
	Category    string `url:"category"`              // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol      string `url:"symbol,omitempty"`      // [optional]Symbol name. For linear & inverse, either symbol or settleCoin is required
	BaseCoin    string `url:"baseCoin,omitempty"`    // [optional]Base coin. For option only. Return all option open orders if not passed
	SettleCoin  string `url:"settleCoin,omitempty"`  // [optional]Settle coin. For linear & inverse, either symbol or settleCoin is required
	OrderId     string `url:"orderId,omitempty"`     // [optional]Order ID
	OrderLinkId string `url:"orderLinkId,omitempty"` // [optional]User customised order ID
	OpenOnly    int64  `url:"openOnly,omitempty"`    // [optional]Unified account & Normal account: 0(default) - query open orders only Unified account: 1, Normal account: 2 - return cancelled, rejected or totally filled orders by last 10 minutes, A maximum of 500 records are kept under each account. If the Bybit service is restarted due to an update, this part of the data will be cleared and accumulated again, but the order records will still be queried in order history
	OrderFilter string `url:"orderFilter,omitempty"` // [optional]Order: active order, StopOrder: conditional order, tpslOrder: spot TP/SL order. Default: all kinds of orders
	Limit       int64  `url:"limit,omitempty"`       // [optional]Limit for data size per page. [1, 50]. Default: 20
	Cursor      string `url:"cursor,omitempty"`      // [optional]Cursor. Used for pagination
}

type GetOpenOrdersResponse struct {
	Category       string `json:"category"`       // Product type
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
	List           []struct {
		OrderId            string `json:"orderId"`            // Order ID
		OrderLinkId        string `json:"orderLinkId"`        // User customised order ID
		BlockTradeId       string `json:"blockTradeId"`       // Paradigm block trade ID
		Symbol             string `json:"symbol"`             // Symbol name
		Price              string `json:"price"`              // Order price
		Qty                string `json:"qty"`                // Order qty
		Side               string `json:"side"`               // Side. Buy,Sell
		IsLeverage         string `json:"isLeverage"`         // Whether to borrow. spot only. 0: false, 1: true
		PositionIdx        int64  `json:"positionIdx"`        // Position index. Used to identify positions in different position modes.
		OrderStatus        string `json:"orderStatus"`        // Order status
		CancelType         string `json:"cancelType"`         // Cancel type
		RejectReason       string `json:"rejectReason"`       // Reject reason
		AvgPrice           string `json:"avgPrice"`           // Average filled price. If unfilled, it is "0"
		LeavesQty          string `json:"leavesQty"`          // The remaining qty not executed
		LeavesValue        string `json:"leavesValue"`        // The estimated value not executed
		CumExecQty         string `json:"cumExecQty"`         // Cumulative executed order qty
		CumExecValue       string `json:"cumExecValue"`       // Cumulative executed order value
		CumExecFee         string `json:"cumExecFee"`         // Cumulative executed trading fee
		TimeInForce        string `json:"timeInForce"`        // Time in force
		OrderType          string `json:"orderType"`          // Order type. Market,Limit
		StopOrderType      string `json:"stopOrderType"`      // Stop order type
		OrderIv            string `json:"orderIv"`            // Implied volatility
		TriggerPrice       string `json:"triggerPrice"`       // Trigger price. If stopOrderType=TrailingStop, it is activate price. Otherwise, it is trigger price
		TakeProfit         string `json:"takeProfit"`         // Take profit price
		StopLoss           string `json:"stopLoss"`           // Stop loss price
		TpTriggerBy        string `json:"tpTriggerBy"`        // The price type to trigger take profit
		SlTriggerBy        string `json:"slTriggerBy"`        // The price type to trigger stop loss
		TriggerDirection   int64  `json:"triggerDirection"`   // Trigger direction. 1: rise, 2: fall
		TriggerBy          string `json:"triggerBy"`          // The price type of trigger price
		LastPriceOnCreated string `json:"lastPriceOnCreated"` // Last price when place the order
		ReduceOnly         bool   `json:"reduceOnly"`         // Reduce only. true means reduce position size
		CloseOnTrigger     bool   `json:"closeOnTrigger"`     // Close on trigger. What is a close on trigger order?
		CreatedTime        string `json:"createdTime"`        // Order created timestamp (ms)
		UpdatedTime        string `json:"updatedTime"`        // Order updated timestamp (ms)
	} `json:"list"` // Object
}

type CancelAllOrdersRequest struct {
	Category    string `json:"category"`              // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic. If cancel all by baseCoin, it will cancel all linear & inverse orders
	Symbol      string `json:"symbol,omitempty"`      // [optional]Symbol name
	BaseCoin    string `json:"baseCoin,omitempty"`    // [optional]Base coin
	SettleCoin  string `json:"settleCoin,omitempty"`  // [optional]Settle coin. It does not support spot
	OrderFilter string `json:"orderFilter,omitempty"` // [optional]Valid for spot only. Order,tpslOrder. If not passed, Order by default
}

type CancelAllOrdersResponse struct {
	List []struct {
		OrderId     string `json:"orderId"`     // Order ID
		OrderLinkId string `json:"orderLinkId"` // User customised order ID
	} `json:"list"` // Object
}

type GetOrderHistoryRequest struct {
	Category    string `url:"category"`              // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol      string `url:"symbol,omitempty"`      // [optional]Symbol name
	BaseCoin    string `url:"baseCoin,omitempty"`    // [optional]Base coin. Normal account does not support this param
	OrderId     string `url:"orderId,omitempty"`     // [optional]Order ID
	OrderLinkId string `url:"orderLinkId,omitempty"` // [optional]User customised order ID
	OrderFilter string `url:"orderFilter,omitempty"` // [optional]Order: active order, StopOrder: conditional order, tpslOrder: spot TP/SL order. Default: all kinds of orders
	OrderStatus string `url:"orderStatus,omitempty"` // [optional]Return all status orders if not passed
	Limit       int64  `url:"limit,omitempty"`       // [optional]Limit for data size per page. [1, 50]. Default: 20
	Cursor      string `url:"cursor,omitempty"`      // [optional]Cursor. Used for pagination
}

type GetOrderHistoryResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		OrderId            string `json:"orderId"`            // Order ID
		OrderLinkId        string `json:"orderLinkId"`        // User customised order ID
		BlockTradeId       string `json:"blockTradeId"`       // Block trade ID
		Symbol             string `json:"symbol"`             // Symbol name
		Price              string `json:"price"`              // Order price
		Qty                string `json:"qty"`                // Order qty
		Side               string `json:"side"`               // Side. Buy,Sell
		IsLeverage         string `json:"isLeverage"`         // Whether to borrow. spot only. 0: false, 1: true
		PositionIdx        int64  `json:"positionIdx"`        // Position index. Used to identify positions in different position modes
		OrderStatus        string `json:"orderStatus"`        // Order status
		CancelType         string `json:"cancelType"`         // Cancel type
		RejectReason       string `json:"rejectReason"`       // Reject reason
		AvgPrice           string `json:"avgPrice"`           // Average filled price. If unfilled, it is ""
		LeavesQty          string `json:"leavesQty"`          // The remaining qty not executed
		LeavesValue        string `json:"leavesValue"`        // The estimated value not executed
		CumExecQty         string `json:"cumExecQty"`         // Cumulative executed order qty
		CumExecValue       string `json:"cumExecValue"`       // Cumulative executed order value
		CumExecFee         string `json:"cumExecFee"`         // Cumulative executed trading fee
		TimeInForce        string `json:"timeInForce"`        // Time in force
		OrderType          string `json:"orderType"`          // Order type. Market,Limit
		StopOrderType      string `json:"stopOrderType"`      // Stop order type
		OrderIv            string `json:"orderIv"`            // Implied volatility
		TriggerPrice       string `json:"triggerPrice"`       // Trigger price. If stopOrderType=TrailingStop, it is activate price. Otherwise, it is trigger price
		TakeProfit         string `json:"takeProfit"`         // Take profit price
		StopLoss           string `json:"stopLoss"`           // Stop loss price
		TpTriggerBy        string `json:"tpTriggerBy"`        // The price type to trigger take profit
		SlTriggerBy        string `json:"slTriggerBy"`        // The price type to trigger stop loss
		TriggerDirection   int64  `json:"triggerDirection"`   // Trigger direction. 1: rise, 2: fall
		TriggerBy          string `json:"triggerBy"`          // The price type of trigger price
		LastPriceOnCreated string `json:"lastPriceOnCreated"` // Last price when place the order. For linear only
		ReduceOnly         bool   `json:"reduceOnly"`         // Reduce only. true means reduce position size
		CloseOnTrigger     bool   `json:"closeOnTrigger"`     // Close on trigger. What is a close on trigger order?
		CreatedTime        string `json:"createdTime"`        // Order created timestamp (ms)
		UpdatedTime        string `json:"updatedTime"`        // Order updated timestamp (ms)
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}

type BatchPlaceOrderRequest struct {
	Category string `json:"category"` // Product type. option
	Request  []struct {
		Symbol      string `json:"symbol"`                // Symbol name
		Side        string `json:"side"`                  // Buy, Sell
		OrderType   string `json:"orderType"`             // Market, Limit
		Qty         string `json:"qty"`                   // Order quantity
		Price       string `json:"price,omitempty"`       // [optional]Order price. Invalid if orderType=Market
		OrderIv     string `json:"orderIv,omitempty"`     // [optional]Implied volatility. option only. orderIv has a higher priority than price. Pass the real value, e.g for 10%, 0.1 should be passed
		TimeInForce string `json:"timeInForce,omitempty"` // [optional]Time in force Market order will use IOC directly If not passed, GTC is used by default
		OrderLinkId string `json:"orderLinkId"`           // User customised order ID. A max of 36 characters. Combinations of numbers, letters (upper and lower cases), dashes, and underscores are supported. rule of option: required param always unique
		ReduceOnly  bool   `json:"reduceOnly,omitempty"`  // [optional]What is a reduce-only order? true means your position can only reduce in size if this order is triggered.
		Mmp         bool   `json:"mmp,omitempty"`         // [optional]Market maker protection. option only. true means set the order as a market maker protection order

	} `json:"request"` // Object
}

type BatchPlaceOrderResponse struct {
	Result struct {
		List []struct {
			Category    string `json:"category"`    // Product type
			Symbol      string `json:"symbol"`      // Symbol name
			OrderId     string `json:"orderId"`     // Order ID
			OrderLinkId string `json:"orderLinkId"` // User customised order ID
			CreateAt    string `json:"createAt"`    // Order created time (ms)
		} `json:"list"` // Object
	} `json:"result"` // Object
	RetExtInfo struct {
		List []struct {
			Code int64  `json:"code"` // Success/error code
			Msg  string `json:"msg"`  // Success/error message
		} `json:"list"` // Object
	} `json:"retExtInfo"` // retExtInfo
}

type BatchAmendOrderRequest struct {
	Category string `json:"category"` // Product type. option
	Request  []struct {
		Symbol      string `json:"symbol"`                // Symbol name
		OrderId     string `json:"orderId,omitempty"`     // [optional]Order ID. Either orderId or orderLinkId is required
		OrderLinkId string `json:"orderLinkId,omitempty"` // [optional]User customised order ID. Either orderId or orderLinkId is required
		Qty         string `json:"qty,omitempty"`         // [optional]Order quantity after modification. Don't pass it if not modify the qty
		Price       string `json:"price,omitempty"`       // [optional]Order price after modification. Don't pass it if not modify the price
		OrderIv     string `json:"orderIv,omitempty"`     // [optional]Implied volatility. option only. Pass the real value, e.g for 10%, 0.1 should be passed
	} `json:"request"` // Object
}

type BatchAmendOrderResponse struct {
	Result struct {
		List []struct {
			Category    string `json:"category"`    // Product type
			Symbol      string `json:"symbol"`      // Symbol name
			OrderId     string `json:"orderId"`     // Order ID
			OrderLinkId string `json:"orderLinkId"` // User customised order ID
		} `json:"list"` // Object
	} `json:"result"` // Object
	RetExtInfo struct {
		List []struct {
			Code int64  `json:"code"` // Success/error code
			Msg  string `json:"msg"`  // Success/error message
		} `json:"list"` // Object
	} `json:"retExtInfo"` //  Object
}

type BatchCancelOrderRequest struct {
	Category string `json:"category"` // Product type. option
	Request  []struct {
		Symbol      string `json:"symbol"`                // Symbol name
		OrderId     string `json:"orderId,omitempty"`     // [optional]Order ID. Either orderId or orderLinkId is required
		OrderLinkId string `json:"orderLinkId,omitempty"` // [optional]User customised order ID. Either orderId or orderLinkId is required
	} `json:"request"` // Object
}

type BatchCancelOrderResponse struct {
	Result struct {
		List []struct {
			Category    string `json:"category"`    // Product type
			Symbol      string `json:"symbol"`      // Symbol name
			OrderId     string `json:"orderId"`     // Order ID
			OrderLinkId string `json:"orderLinkId"` // User customised order ID
		} `json:"list"` // Object
	} `json:"result"` // Object
	RetExtInfo struct {
		List []struct {
			Code int64  `json:"code"` // Success/error code
			Msg  string `json:"msg"`  // Success/error message
		} `json:"list"` // Object
	} `json:"retExtInfo"` // Object
}
