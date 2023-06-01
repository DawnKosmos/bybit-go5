package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// GetAssetInfo Query asset information
func (c *Client) GetAssetInfo(request models.GetAssetInfoRequest) (*models.GetAssetInfoResponse, error) {
	var respBody models.Response[models.GetAssetInfoResponse]
	err := c.GET("/v5/asset/transfer/query-asset-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetCoinInfo Query coin information, including chain information, withdraw and deposit status.
func (c *Client) GetCoinInfo(request models.GetCoinInfoRequest) (*models.GetCoinInfoResponse, error) {
	var respBody models.Response[models.GetCoinInfoResponse]
	err := c.GET("/v5/asset/coin/query-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// Withdraw Withdraw assets from your Bybit account. You can make an off-chain transfer if the target wallet address is from Bybit. This means that no blockchain fee will be charged.
func (c *Client) Withdraw(request models.WithdrawRequest) (*models.WithdrawResponse, error) {
	var respBody models.Response[models.WithdrawResponse]
	err := c.POST("/v5/asset/withdraw/create", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}
