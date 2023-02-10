package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// GetWalletBalance Obtain wallet balance, query asset information of each currency
func (c *Client) GetWalletBalance(request models.GetWalletBalanceRequest) (*models.GetWalletBalanceResponse, error) {
	var respBody models.Response[models.GetWalletBalanceResponse]
	err := c.GET("/v5/account/wallet-balance", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// UpgradetoUnifiedAccount Upgrade Unified Account
func (c *Client) UpgradetoUnifiedAccount() (*models.UpgradetoUnifiedAccountResponse, error) {
	var respBody models.Response[models.UpgradetoUnifiedAccountResponse]
	err := c.POST("/v5/account/upgrade-to-uta", nil, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetBorrowHistory Get interest records, sorted in reverse order of creation time.
func (c *Client) GetBorrowHistory(request models.GetBorrowHistoryRequest) (*models.GetBorrowHistoryResponse, error) {
	var respBody models.Response[models.GetBorrowHistoryResponse]
	err := c.GET("/v5/account/borrow-history", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetCollateralInfo Get the collateral information of the current unified margin account
func (c *Client) GetCollateralInfo(request models.GetCollateralInfoRequest) (*models.GetCollateralInfoResponse, error) {
	var respBody models.Response[models.GetCollateralInfoResponse]
	err := c.GET("/v5/account/collateral-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetCoinGreeks Get current account Greeks information
func (c *Client) GetCoinGreeks(request models.GetCoinGreeksRequest) (*models.GetCoinGreeksResponse, error) {
	var respBody models.Response[models.GetCoinGreeksResponse]
	err := c.GET("/v5/asset/coin-greeks", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetAccountInfo Query the margin mode configuration of the account.
func (c *Client) GetAccountInfo() (*models.GetAccountInfoResponse, error) {
	var respBody models.Response[models.GetAccountInfoResponse]
	err := c.GET("/v5/account/info", nil, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetTransactionLog Query transaction logs in Unified account.
func (c *Client) GetTransactionLog(request models.GetTransactionLogRequest) (*models.GetTransactionLogResponse, error) {
	var respBody models.Response[models.GetTransactionLogResponse]
	err := c.GET("/v5/account/transaction-log", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetMarginMode Default is regular margin mode. This mode is valid for USDT Perp, USDC Perp and USDC Option.
func (c *Client) SetMarginMode(request models.SetMarginModeRequest) (*models.SetMarginModeResponse, error) {
	var respBody models.Response[models.SetMarginModeResponse]
	err := c.POST("/v5/account/set-margin-mode", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}
