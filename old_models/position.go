package models

type GetPositionInfoRequest struct {
	Category   string `url:"category"`             // Product type Unified account: linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol     string `url:"symbol,omitempty"`     // [optional]Symbol name If symbol passed, it returns data regardless of having position or not. If symbol=null, it returns position size grater than zero.
	BaseCoin   string `url:"baseCoin,omitempty"`   // [optional]Base coin. option only. Return all option positions if not passed
	SettleCoin string `url:"settleCoin,omitempty"` // [optional]Settle coin. For linear & inverse, either symbol or settleCon is required. symbol has a higher priority
	Limit      int64  `url:"limit,omitempty"`      // [optional]Limit for data size per page. [1, 200]. Default: 20
	Cursor     string `url:"cursor,omitempty"`     // [optional]Cursor. Used for pagination
}

type GetPositionInfoResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		PositionIdx    int64  `json:"positionIdx"`    // Position idx, used to identify positions in different position modes 0: One-Way Mode 1: Buy side of both side mode 2: Sell side of both side mode
		RiskId         int64  `json:"riskId"`         // Risk limit ID. Note: for portfolio margin mode, this field returns 0, which means risk limit rules are invalid
		RiskLimitValue string `json:"riskLimitValue"` // Risk limit value. Note: for portfolio margin mode, this field returns 0, which means risk limit rules are invalid
		Symbol         string `json:"symbol"`         // Symbol name
		Side           string `json:"side"`           // Position side. Buy: long, Sell: short. Note: under one-way mode, it returns None if empty position
		Size           string `json:"size"`           // Position size
		AvgPrice       string `json:"avgPrice"`       // Average entry price
		PositionValue  string `json:"positionValue"`  // Position value
		TradeMode      int64  `json:"tradeMode"`      // Trade mode. 0: cross-margin, 1: isolated margin
		AutoAddMargin  int64  `json:"autoAddMargin"`  // Whether to add margin automatically. 0: false, 1: true. Unique field for normal account
		Leverage       string `json:"leverage"`       // Position leverage. Valid for contract. Note: for portfolio margin mode, this field returns "", which means leverage rules are invalid
		MarkPrice      string `json:"markPrice"`      // Last mark price
		LiqPrice       string `json:"liqPrice"`       // Position liquidation price. Note: the value returned in the unified mode is the estimated liquidation price, because the unified margin mode controls the risk rate according to the account, so the liquidation price on a single position is only estimated
		BustPrice      string `json:"bustPrice"`      // Bankruptcy price. Note: Unified mode returns "", no position bankruptcy price
		PositionIM     string `json:"positionIM"`     // Initial margin. For portfolio margin mode, it returns ""
		PositionMM     string `json:"positionMM"`     // Maintenance margin. For portfolio margin mode, it returns ""
		TpslMode       string `json:"tpslMode"`       // Take profit/stop loss mode. Full,Partial
		TakeProfit     string `json:"takeProfit"`     // Take profit price
		StopLoss       string `json:"stopLoss"`       // Stop loss price
		TrailingStop   string `json:"trailingStop"`   // Trailing stop (The distance from market price)
		UnrealisedPnl  string `json:"unrealisedPnl"`  // Unrealised PnL
		CumRealisedPnl string `json:"cumRealisedPnl"` // Cumulative realised pnl
		CreatedTime    string `json:"createdTime"`    // Position created timestamp (ms)
		UpdatedTime    string `json:"updatedTime"`    // Position updated timestamp (ms)
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}

type SetLeverageRequest struct {
	Category     string `json:"category"`     // Product type Unified account: linear Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol       string `json:"symbol"`       // Symbol name
	BuyLeverage  string `json:"buyLeverage"`  // [0, max leverage of corresponding risk limit]. Note: Under one-way mode, buyLeverage must be the same as sellLeverage
	SellLeverage string `json:"sellLeverage"` // [0, max leverage of corresponding risk limit]. Note: Under one-way mode, buyLeverage must be the same as sellLeverage
}

type SwitchCrossIsolatedMarginRequest struct {
	Category     string `json:"category"`     // Product type. linear,inverse Please note that category is not involved with business logic Unified account is not applicable
	Symbol       string `json:"symbol"`       // Symbol name
	TradeMode    int64  `json:"tradeMode"`    // 0: cross margin. 1: isolated margin
	BuyLeverage  string `json:"buyLeverage"`  // The value must be equal to sellLeverage value
	SellLeverage string `json:"sellLeverage"` // The value must be equal to buyLeverage value
}

type SetTP_SLModeRequest struct {
	Category string `json:"category"` // Product type Unified account: linear Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol   string `json:"symbol"`   // Symbol name
	TpSlMode string `json:"tpSlMode"` // TP/SL mode. Full,Partial
}

type SetTP_SLModeResponse struct {
	TpSlMode string `json:"tpSlMode"` // Full,Partial
}

type SwitchPositionModeRequest struct {
	Category string `json:"category"`         // Please note that category is not involved with business logic
	Symbol   string `json:"symbol,omitempty"` // [optional]Symbol name. Either symbol or coin is required. symbol has a higher priority
	Coin     string `json:"coin,omitempty"`   // [optional]Coin
	Mode     int64  `json:"mode"`             // Position mode. 0: Merged Single. 3: Both Sides
}

type SetRiskLimitRequest struct {
	Category    string `json:"category"`              // Product type Unified account: linear Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol      string `json:"symbol"`                // Symbol name
	RiskId      int64  `json:"riskId"`                // Risk limit ID
	PositionIdx int64  `json:"positionIdx,omitempty"` // [optional]Used to identify positions in different position modes.  0: one-way mode 1: hedge-mode Buy side 2: hedge-mode Sell side
}

type SetRiskLimitResponse struct {
	Category       string `json:"category"`       // Product type
	RiskId         int64  `json:"riskId"`         // Risk limit ID
	RiskLimitValue string `json:"riskLimitValue"` // The position limit value corresponding to this risk ID
}

type SetTradingStopRequest struct {
	Category     string `json:"category"`               // Product type Different Account Types
	Symbol       string `json:"symbol"`                 // Symbol name
	TakeProfit   string `json:"takeProfit,omitempty"`   // [optional]Cannot be less than 0, 0 means cancel TP
	StopLoss     string `json:"stopLoss,omitempty"`     // [optional]Cannot be less than 0, 0 means cancel SL
	TrailingStop string `json:"trailingStop,omitempty"` // [optional]Cannot be less than 0, 0 means cancel TS
	TpTriggerBy  string `json:"tpTriggerBy,omitempty"`  // [optional]Take profit trigger price type
	SlTriggerBy  string `json:"slTriggerBy,omitempty"`  // [optional]Stop loss trigger price type
	ActivePrice  string `json:"activePrice,omitempty"`  // [optional]Trailing stop trigger price. Trailing stop will be triggered when this price is reached only
	TpSize       string `json:"tpSize,omitempty"`       // [optional]Take profit size. Valid in TP/SL partial mode
	SlSize       string `json:"slSize,omitempty"`       // [optional]Stop loss size. Valid in TP/SL partial mode
	PositionIdx  int64  `json:"positionIdx"`            // Used to identify positions 0: one-way mode 1: hedge-mode Buy side 2: hedge-mode Sell side
}

type SetAutoAddMarginRequest struct {
	Category      string `json:"category"`              // Product type. linear
	Symbol        string `json:"symbol"`                // Symbol name
	AutoAddMargin int64  `json:"autoAddMargin"`         // Turn on/off. 0: off. 1: on
	PositionIdx   int64  `json:"positionIdx,omitempty"` // [optional]Used to identify positions in different position modes. For hedge mode position, this param is required 0: one-way mode 1: hedge-mode Buy side2: hedge-mode Sell side
}

type GetExecutionHalfYearRequest struct {
	Category    string `url:"category"`              // Product type Unified account: spot, linear, option Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol      string `url:"symbol,omitempty"`      // [optional]Symbol name. Normal account, either symbol or orderId is required
	OrderId     string `url:"orderId,omitempty"`     // [optional]Order ID
	OrderLinkId string `url:"orderLinkId,omitempty"` // [optional]User customised order ID. Normal account does not support this param
	BaseCoin    string `url:"baseCoin,omitempty"`    // [optional]Base coin. Normal account does not support this param
	StartTime   int64  `url:"startTime,omitempty"`   // [optional]The start timestamp (ms)
	EndTime     int64  `url:"endTime,omitempty"`     // [optional]The end timestamp (ms)
	ExecType    string `url:"execType,omitempty"`    // [optional]Execution type
	Limit       int64  `url:"limit,omitempty"`       // [optional]Limit for data size per page. [1, 100]. Default: 50
	Cursor      string `url:"cursor,omitempty"`      // [optional]Cursor. Used for pagination
}

type GetExecutionHalfYearResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol          string `json:"symbol"`          // Symbol name
		OrderId         string `json:"orderId"`         // Order ID
		OrderLinkId     string `json:"orderLinkId"`     // User customized order ID
		Side            string `json:"side"`            // Side. Buy,Sell
		OrderPrice      string `json:"orderPrice"`      // Order price
		OrderQty        string `json:"orderQty"`        // Order qty
		LeavesQty       string `json:"leavesQty"`       // The remaining qty not executed
		OrderType       string `json:"orderType"`       // Order type. Market,Limit
		StopOrderType   string `json:"stopOrderType"`   // Stop order type. If the order is not stop order, any type is not returned
		ExecFee         string `json:"execFee"`         // Executed trading fee
		ExecId          string `json:"execId"`          // Execution ID
		ExecPrice       string `json:"execPrice"`       // Execution price
		ExecQty         string `json:"execQty"`         // Execution qty
		ExecType        string `json:"execType"`        // Executed type
		ExecValue       string `json:"execValue"`       // Executed order value
		ExecTime        string `json:"execTime"`        // Executed timestamp（ms）
		IsMaker         bool   `json:"isMaker"`         // Is maker order. true: maker, false: taker
		FeeRate         string `json:"feeRate"`         // Trading fee rate
		TradeIv         string `json:"tradeIv"`         // Implied volatility. Valid for option
		MarkIv          string `json:"markIv"`          // Implied volatility of mark price. Valid for option
		MarkPrice       string `json:"markPrice"`       // The mark price of the symbol when executing
		IndexPrice      string `json:"indexPrice"`      // The index price of the symbol when executing
		UnderlyingPrice string `json:"underlyingPrice"` // The underlying price of the symbol when executing. Valid for option
		BlockTradeId    string `json:"blockTradeId"`    // Paradigm block trade ID
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}

type GetClosedPnLRequest struct {
	Category  string `url:"category"`            // Product type Unified account: linear Normal account: linear, inverse. Please note that category is not involved with business logic
	Symbol    string `url:"symbol,omitempty"`    // [optional]Symbol name. This is a required parameter for Normal account
	StartTime int64  `url:"startTime,omitempty"` // [optional]The start timestamp (ms)
	EndTime   int64  `url:"endTime,omitempty"`   // [optional]The end timestamp (ms)
	Limit     int64  `url:"limit,omitempty"`     // [optional]Limit for data size per page. [1, 200]. Default: 50
	Cursor    string `url:"cursor,omitempty"`    // [optional]Cursor. Used for pagination
}

type GetClosedPnLResponse struct {
	Category string `json:"category"` // Product type
	List     []struct {
		Symbol        string `json:"symbol"`        // Symbol name
		OrderId       string `json:"orderId"`       // Order ID
		Side          string `json:"side"`          // Buy, Side
		Qty           string `json:"qty"`           // Order qty
		OrderPrice    string `json:"orderPrice"`    // Order price
		OrderType     string `json:"orderType"`     // Order type. Market,Limit
		ExecType      string `json:"execType"`      // Exec type
		ClosedSize    string `json:"closedSize"`    // Closed size
		CumEntryValue string `json:"cumEntryValue"` // Cumulated Position value
		AvgEntryPrice string `json:"avgEntryPrice"` // Average entry price
		CumExitValue  string `json:"cumExitValue"`  // Cumulated exit position value
		AvgExitPrice  string `json:"avgExitPrice"`  // Average exit price
		ClosedPnl     string `json:"closedPnl"`     // Closed PnL
		FillCount     string `json:"fillCount"`     // The number of fills in a single order
		Leverage      string `json:"leverage"`      // leverage
		CreatedTime   string `json:"createdTime"`   // The created time (ms)
		UpdatedTime   string `json:"updatedTime"`   // The updated time (ms)
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}
