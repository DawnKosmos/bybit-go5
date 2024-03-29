package models

type WsPosition struct {
	Id           string `json:"id"`           // Message ID
	Topic        string `json:"topic"`        // Topic name
	CreationTime int64  `json:"creationTime"` // Data created timestamp (ms)
	Data         []struct {
		Category        string `json:"category"`        // Product type Unified account: does not have this field Normal account: linear, inverse.
		Symbol          string `json:"symbol"`          // Symbol name
		Side            string `json:"side"`            // Position side. Buy,Sell
		Size            string `json:"size"`            // Position size
		PositionIdx     int64  `json:"positionIdx"`     // Used to identify positions in different position modes
		TradeMode       int64  `json:"tradeMode"`       // Trade mode. 0: cross margin, 1: isolated margin. Always 0 under unified margin account
		PositionValue   string `json:"positionValue"`   // Position value
		RiskId          int64  `json:"riskId"`          // Risk limit ID. Note: for portfolio margin mode, it returns 0, which the risk limit value is invalid
		RiskLimitValue  string `json:"riskLimitValue"`  // Risk limit value corresponding to riskId. Note: for portfolio margin mode, it returns "", which the risk limit value is invalid
		EntryPrice      string `json:"entryPrice"`      // Entry price
		MarkPrice       string `json:"markPrice"`       // Mark price
		Leverage        string `json:"leverage"`        // Leverage. Note: for portfolio margin mode, it returns "", which the leverage value is invalid
		PositionBalance string `json:"positionBalance"` // Position margin. Unified account does not have this field
		AutoAddMargin   int64  `json:"autoAddMargin"`   // Whether to add margin automatically. 0: false, 1: true. Unified account does not have this field
		PositionMM      string `json:"positionMM"`      // Position maintenance margin. Note: for portfolio margin mode, it returns ""
		PositionIM      string `json:"positionIM"`      // Position initial margin. Note: for portfolio margin mode, it returns ""
		LiqPrice        string `json:"liqPrice"`        // Est.liquidation price. "" for unified margin account
		BustPrice       string `json:"bustPrice"`       // Est.bankruptcy price. "" for unified margin account
		TpSlMode        string `json:"tpSlMode"`        // Tp/Sl mode. Full,Partial
		TakeProfit      string `json:"takeProfit"`      // Take profit price
		StopLoss        string `json:"stopLoss"`        // Stop loss price
		TrailingStop    string `json:"trailingStop"`    // Trailing stop
		UnrealisedPnl   string `json:"unrealisedPnl"`   // Unrealised profit and loss
		CumRealisedPnl  string `json:"cumRealisedPnl"`  // Cumulative realised PnL
		PositionStatus  string `json:"positionStatus"`  // Position status. Normal, Liq, Adl
		CreatedTime     string `json:"createdTime"`     // Position created timestamp (ms)
		UpdatedTime     string `json:"updatedTime"`     // Position data updated timestamp (ms)
	} `json:"data"` // Object
}

type WsExecution struct {
	Id           string `json:"id"`           // Message ID
	Topic        string `json:"topic"`        // Topic name
	CreationTime int64  `json:"creationTime"` // Data created timestamp (ms)
	Data         []struct {
		Category        string `json:"category"`        // Product type Unified account: spot, linear, option Normal account: linear, inverse.
		Symbol          string `json:"symbol"`          // Symbol name
		IsLeverage      string `json:"isLeverage"`      // Whether to borrow. Valid for spot only. 0(default): false, 1: true
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
	} `json:"data"` // Object
}

type WsOrder struct {
	Id           string `json:"id"`           // Message ID
	Topic        string `json:"topic"`        // Topic name
	CreationTime int64  `json:"creationTime"` // Data created timestamp (ms)
	Data         []struct {
		Category           string `json:"category"`           // Product type Unified account: spot, linear, option Normal account: linear, inverse.
		OrderId            string `json:"orderId"`            // Order ID
		OrderLinkId        string `json:"orderLinkId"`        // User customised order ID
		IsLeverage         string `json:"isLeverage"`         // Whether to borrow. spot returns this field only. 0(default): false, 1: true
		BlockTradeId       string `json:"blockTradeId"`       // Block trade ID
		Symbol             string `json:"symbol"`             // Symbol name
		Price              string `json:"price"`              // Order price
		Qty                string `json:"qty"`                // Order qty
		Side               string `json:"side"`               // Side. Buy,Sell
		PositionIdx        int64  `json:"positionIdx"`        // Position index. Used to identify positions in different position modes
		OrderStatus        string `json:"orderStatus"`        // Order status
		CancelType         string `json:"cancelType"`         // Cancel type
		RejectReason       string `json:"rejectReason"`       // Reject reason
		AvgPrice           string `json:"avgPrice"`           // Average filled price. If unfilled, it is ""
		LeavesQty          string `json:"leavesQty"`          // The remaining qty not executed
		LeavesValue        string `json:"leavesValue"`        // The remaining value not executed
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
	} `json:"data"` // Object
}

type WsWallet struct {
	Id           string `json:"id"`           // Message ID
	Topic        string `json:"topic"`        // Topic name
	CreationTime int64  `json:"creationTime"` // Data created timestamp (ms)
	Data         []struct {
		AccountType            string `json:"accountType"`            // Account type. Unified account: UNIFIED Normal account: CONTRACT
		AccountIMRate          string `json:"accountIMRate"`          // Initial Margin Rate: Account Total Initial Margin Base Coin / Account Margin Balance Base Coin. In non-unified mode, the field will be returned as an empty string
		AccountMMRate          string `json:"accountMMRate"`          // Maintenance Margin Rate: Account Total Maintenance Margin Base Coin / Account Margin Balance Base Coin. In non-unified mode, the field will be returned as an empty string.
		TotalEquity            string `json:"totalEquity"`            // Equity of account converted to usd：Account Margin Balance Base Coin + Account Option Value Base Coin. In non-unified mode, the field will be returned as an empty string.
		TotalWalletBalance     string `json:"totalWalletBalance"`     // Wallet Balance of account converted to usd：∑ Asset Wallet Balance By USD value of each asset. In non-unified mode, the field will be returned as an empty string.
		TotalMarginBalance     string `json:"totalMarginBalance"`     // Margin Balance of account converted to usd：totalWalletBalance + totalPerpUPL. In non-unified mode, the field will be returned as an empty string.
		TotalAvailableBalance  string `json:"totalAvailableBalance"`  // Available Balance of account converted to usd：Regular mode：totalMarginBalance - totalInitialMargin. In non-unified mode, the field will be returned as an empty string.
		TotalPerpUPL           string `json:"totalPerpUPL"`           // Unrealised P&L of perpetuals of account converted to usd：∑ Each perp upl by base coin. In non-unified mode, the field will be returned as an empty string.
		TotalInitialMargin     string `json:"totalInitialMargin"`     // Initial Margin of account converted to usd：∑ Asset Total Initial Margin Base Coin. In non-unified mode, the field will be returned as an empty string.
		TotalMaintenanceMargin string `json:"totalMaintenanceMargin"` // Maintenance Margin of account converted to usd: ∑ Asset Total Maintenance Margin Base Coin. In non-unified mode, the field will be returned as an empty string.
		Coin                   []struct {
			Coin                string `json:"coin"`                // Coin name, such as BTC, ETH, USDT, USDC
			Equity              string `json:"equity"`              // Equity of current coin
			UsdValue            string `json:"usdValue"`            // USD value of current coin. If this coin cannot be collateral, then it is 0
			WalletBalance       string `json:"walletBalance"`       // Wallet balance of current coin
			BorrowAmount        string `json:"borrowAmount"`        // Borrow amount of current coin
			AvailableToBorrow   string `json:"availableToBorrow"`   // Available amount to borrow of current coin
			AvailableToWithdraw string `json:"availableToWithdraw"` // Available amount to withdraw of current coin
			AccruedInterest     string `json:"accruedInterest"`     // Accrued interest
			TotalOrderIM        string `json:"totalOrderIM"`        // Pre-occupied margin for order. For portfolio margin mode, it returns ""
			TotalPositionIM     string `json:"totalPositionIM"`     // Sum of initial margin of all positions + Pre-occupied liquidation fee. For portfolio margin mode, it returns ""
			TotalPositionMM     string `json:"totalPositionMM"`     // Sum of maintenance margin for all positions. For portfolio margin mode, it returns ""
			UnrealisedPnl       string `json:"unrealisedPnl"`       // Unrealised P&L
			CumRealisedPnl      string `json:"cumRealisedPnl"`      // Cumulative Realised P&L
		} `json:"coin"` // Object
	} `json:"data"` // Object
}

type WsGreek struct {
	Id           string `json:"id"`           // Message ID
	Topic        string `json:"topic"`        // Topic name
	CreationTime int64  `json:"creationTime"` // Data created timestamp (ms)
	Data         []struct {
		BaseCoin   string `json:"baseCoin"`   // Base coin
		TotalDelta string `json:"totalDelta"` // Delta value
		TotalGamma string `json:"totalGamma"` // Gamma value
		TotalVega  string `json:"totalVega"`  // Vega value
		TotalTheta string `json:"totalTheta"` // Theta value
	} `json:"data"` // Object
}
