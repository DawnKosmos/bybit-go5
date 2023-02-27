package models

type GetWalletBalanceRequest struct {
	AccountType string `url:"accountType"`    // Account type Unified account: UNIFIED Normal account: CONTRACT
	Coin        string `url:"coin,omitempty"` // [optional]Coin name If not passed, it returns non-zero asset info You can pass multiple coins to query, separated by comma. USDT,USDC.
}

type GetWalletBalanceResponse struct {
	List []struct {
		AccountType            string `json:"accountType"`            // Account type.
		AccountIMRate          string `json:"accountIMRate"`          // Initial Margin Rate: Account Total Initial Margin Base Coin / Account Margin Balance Base Coin. In non-unified mode, the field will be returned as an empty string.
		AccountMMRate          string `json:"accountMMRate"`          // Maintenance Margin Rate: Account Total Maintenance Margin Base Coin / Account Margin Balance Base Coin. In non-unified mode, the field will be returned as an empty string.
		TotalEquity            string `json:"totalEquity"`            // Equity of account converted to usd：Account Margin Balance Base Coin + Account Option Value Base Coin. In non-unified mode, the field will be returned as an empty string.
		TotalWalletBalance     string `json:"totalWalletBalance"`     // Wallet Balance of account converted to usd：∑ Asset Wallet Balance By USD value of each asset。In non-unified mode, the field will be returned as an empty string.
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
	} `json:"list"` // Object
}

type UpgradetoUnifiedAccountResponse struct {
	UnifiedUpdateStatus string `json:"unifiedUpdateStatus"` // Upgrade status. FAIL,PROCESS,SUCCESS
	UnifiedUpdateMsg    struct {
		Msg []string `json:"msg,omitempty"` // Error message array. Only FAIL will have this field
	} `json:"unifiedUpdateMsg"` // If PROCESS,SUCCESS, it returns null
}

type GetBorrowHistoryRequest struct {
	Currency  string `url:"currency,omitempty"`  // [optional]USDC,USDT,BTC,ETH
	StartTime int64  `url:"startTime,omitempty"` // [optional]The start timestamp (ms)
	EndTime   int64  `url:"endTime,omitempty"`   // [optional]The end time. timestamp (ms)
	Limit     int64  `url:"limit,omitempty"`     // [optional]Limit for data size per page. [1, 50]. Default: 20
	Cursor    string `url:"cursor,omitempty"`    // [optional]Cursor. Used for pagination
}

type GetBorrowHistoryResponse struct {
	List []struct {
		Currency                  string `json:"currency"`                  // USDC,USDT,BTC,ETH
		CreatedTime               int64  `json:"createdTime"`               // Created timestamp (ms)
		BorrowCost                string `json:"borrowCost"`                // Interest
		HourlyBorrowRate          string `json:"hourlyBorrowRate"`          // Hourly Borrow Rate
		InterestBearingBorrowSize string `json:"InterestBearingBorrowSize"` // Interest Bearing Borrow Size
		CostExemption             string `json:"costExemption"`             // Cost exemption
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}

type GetCollateralInfoRequest struct {
	Currency string `url:"currency,omitempty"` // [optional]Asset currency of all current collateral
}

type GetCollateralInfoResponse struct {
	List []struct {
		Currency            string `json:"currency"`            // Currency of all current collateral
		HourlyBorrowRate    string `json:"hourlyBorrowRate"`    // Hourly borrow rate
		MaxBorrowingAmount  string `json:"maxBorrowingAmount"`  // Max borrow amount
		FreeBorrowingAmount string `json:"freeBorrowingAmount"` // Free borrow amount
		BorrowAmount        string `json:"borrowAmount"`        // Borrow amount
		AvailableToBorrow   string `json:"availableToBorrow"`   // Available amount to borrow
		Borrowable          bool   `json:"borrowable"`          // Whether currency can be borrowed
		MarginCollateral    bool   `json:"marginCollateral"`    // Whether it can be used as a margin collateral currency
		CollateralRatio     string `json:"collateralRatio"`     // Collateral ratio
	} `json:"list"` // Object
}

type GetCoinGreeksRequest struct {
	BaseCoin string `url:"baseCoin,omitempty"` // [optional]Base coin. If not passed, all supported base coin greeks will be returned by default
}

type GetCoinGreeksResponse struct {
	List []struct {
		BaseCoin   string `json:"baseCoin"`   // Base coin. e.g.,BTC,ETH,SOL
		TotalDelta string `json:"totalDelta"` // Delta value
		TotalGamma string `json:"totalGamma"` // Gamma value
		TotalVega  string `json:"totalVega"`  // Vega value
		TotalTheta string `json:"totalTheta"` // Theta value
	} `json:"list"` // Object
}

type GetAccountInfoResponse struct {
	UnifiedMarginStatus int64  `json:"unifiedMarginStatus"` // Account status
	MarginMode          string `json:"marginMode"`          // REGULAR_MARGIN, PORTFOLIO_MARGIN
	UpdatedTime         string `json:"updatedTime"`         // Account data updated timestamp (ms)
}

type GetTransactionLogRequest struct {
	AccountType string `url:"accountType,omitempty"` // [optional]Account Type. UNIFIED
	Category    string `url:"category,omitempty"`    // [optional]Product type. spot,linear,option
	Currency    string `url:"currency,omitempty"`    // [optional]Currency
	BaseCoin    string `url:"baseCoin,omitempty"`    // [optional]BaseCoin. e.g., BTC of BTCPERP
	Type        string `url:"type,omitempty"`        // [optional]Types of transaction logs
	StartTime   int64  `url:"startTime,omitempty"`   // [optional]The start timestamp (ms)
	EndTime     int64  `url:"endTime,omitempty"`     // [optional]The end timestamp (ms)
	Limit       int64  `url:"limit,omitempty"`       // [optional]Limit for data size per page. [1, 50]. Default: 20
	Cursor      string `url:"cursor,omitempty"`      // [optional]Cursor. Used for pagination
}

type GetTransactionLogResponse struct {
	List []struct {
		Symbol          string `json:"symbol"`          // Symbol name
		Category        string `json:"category"`        // Product type
		Side            string `json:"side"`            // Side. Buy,Sell,None
		TransactionTime string `json:"transactionTime"` // Transaction timestamp (ms)
		Type            string `json:"type"`            // Type
		Qty             string `json:"qty"`             // Quantity
		Size            string `json:"size"`            // Size
		Currency        string `json:"currency"`        // USDC、USDT、BTC、ETH
		TradePrice      string `json:"tradePrice"`      // Trade price
		Funding         string `json:"funding"`         // Funding fee
		Fee             string `json:"fee"`             // Trading fee. Fees are positive and rebates are negative
		CashFlow        string `json:"cashFlow"`        // Cash flow
		Change          string `json:"change"`          // Change
		CashBalance     string `json:"cashBalance"`     // Cash balance
		FeeRate         string `json:"feeRate"`         // Trading fee rate
		TradeId         string `json:"tradeId"`         // Trade ID
		OrderId         string `json:"orderId"`         // Order ID
		OrderLinkId     string `json:"orderLinkId"`     // User customised order ID
	} `json:"list"` // Object
	NextPageCursor string `json:"nextPageCursor"` // Cursor. Used for pagination
}

type SetMarginModeRequest struct {
	SetMarginMode string `json:"setMarginMode"` // REGULAR_MARGIN, PORTFOLIO_MARGIN
}

type SetMarginModeResponse struct {
	Reasons []struct {
		ReasonCode string `json:"reasonCode"` // Fail reason code
		ReasonMsg  string `json:"reasonMsg"`  // Fail reason msg
	} `json:"reasons"` // Object. If requested successfully, it is an empty array
}

type GetFeeRateDerivativesRequest struct {
	Symbol string `url:"symbol,omitempty"` // [optional]Symbol name
}

type GetFeeRateDerivativesResponse struct {
	List []struct {
		Symbol       string `json:"symbol"`       // Symbol name
		TakerFeeRate string `json:"takerFeeRate"` // Taker fee rate
		MakerFeeRate string `json:"makerFeeRate"` // Maker fee rate
	} `json:"list"` // Object
}
