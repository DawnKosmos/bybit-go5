package models

type GetAssetInfoRequest struct {
	AccountType string `url:"accountType"`    // Account type. SPOT
	Coin        string `url:"coin,omitempty"` // [optional]Coin name
}

type GetAssetInfoResponse struct {
	Spot struct {
		Status string `json:"status"`
		Assets []struct {
			Coin     string `json:"coin"`
			Frozen   string `json:"frozen"`
			Free     string `json:"free"`
			Withdraw string `json:"withdraw"`
		} `json:"assets"`
	} `json:"spot"`
}

type GetCoinInfoRequest struct {
	Coin string `url:"coin,omitempty"` // [optional]Coin
}

type GetCoinInfoResponse struct {
	Rows []struct {
		Name         string `json:"name"`         // Coin name
		Coin         string `json:"coin"`         // Coin
		RemainAmount string `json:"remainAmount"` // Remaining amount
		Chains       []struct {
			Chain                 string `json:"chain"`                 // Chain
			ChainType             string `json:"chainType"`             // Chain type
			Confirmation          string `json:"confirmation"`          // The number of confirm
			WithdrawFee           string `json:"withdrawFee"`           // withdraw fee. If withdraw fee is empty, It means that this coin does not support withdrawal
			DepositMin            string `json:"depositMin"`            // Min. deposit
			WithdrawMin           string `json:"withdrawMin"`           // Min. withdraw
			MinAccuracy           string `json:"minAccuracy"`           // The precision of withdraw or deposit
			ChainDeposit          string `json:"chainDeposit"`          // The chain status of deposit. 0: suspend. 1: normal
			ChainWithdraw         string `json:"chainWithdraw"`         // The chain status of withdraw. 0: suspend. 1: normal
			WithdrawPercentageFee string `json:"withdrawPercentageFee"` // The withdraw fee percentage. It is a real figure, e.g., 0.022 means 2.2%
		} `json:"chains"` // Object
	} `json:"rows"` // Object
}

type WithdrawRequest struct {
	Coin        string `json:"coin"`                  // Coin
	Chain       string `json:"chain"`                 // Chain
	Address     string `json:"address"`               // Wallet address. Please note that the address is case sensitive, so use the exact same address added in address book
	Tag         string `json:"tag,omitempty"`         // [optional]Tag Required if tag exists in the wallet address list. Note: please do not set a tag/memo in the address book if the chain does not support tag
	Amount      string `json:"amount"`                // Withdraw amount
	Timestamp   int64  `json:"timestamp"`             // Current timestamp (ms). Used for preventing from withdraw replay
	ForceChain  int64  `json:"forceChain,omitempty"`  // [optional]Whether or not to force an on-chain withdrawal 0: If the address is parsed out to be an internal address, then internal transfer (default) 1: Force the withdrawal to occur on-chain
	AccountType string `json:"accountType,omitempty"` // [optional]Select the wallet to be withdrawn from SPOT：spot wallet (default) FUND：Funding wallet
}

type WithdrawResponse struct {
	Id string `json:"id"` // Withdrawal ID
}
