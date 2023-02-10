package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// GetPositionInfo Query real-time position data, such as position size, cumulative realizedPNL.
func (c *Client) GetPositionInfo(request models.GetPositionInfoRequest) (*models.GetPositionInfoResponse, error) {
	var respBody models.Response[models.GetPositionInfoResponse]
	err := c.GET("/v5/position/list", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetLeverage Set the leverage
func (c *Client) SetLeverage(request models.SetLeverageRequest) (*models.EmptyResponse, error) {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/position/set-leverage", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SwitchCrossIsolatedMargin Select cross margin mode or isolated margin mode
func (c *Client) SwitchCrossIsolatedMargin(request models.SwitchCrossIsolatedMarginRequest) (*models.EmptyResponse, error) {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/position/switch-isolated", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetTP_SLMode Set TP/SL mode to Full or Partial
func (c *Client) SetTP_SLMode(request models.SetTP_SLModeRequest) (*models.SetTP_SLModeResponse, error) {
	var respBody models.Response[models.SetTP_SLModeResponse]
	err := c.POST("/v5/position/set-tpsl-mode", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SwitchPositionMode Read Doc https://bybit-exchange.github.io/docs/v5/position/position-mode
func (c *Client) SwitchPositionMode(request models.SwitchPositionModeRequest) (*models.EmptyResponse, error) {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/position/switch-mode", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetRiskLimit The risk limit will limit the maximum position value you can hold under different margin requirements.
func (c *Client) SetRiskLimit(request models.SetRiskLimitRequest) (*models.SetRiskLimitResponse, error) {
	var respBody models.Response[models.SetRiskLimitResponse]
	err := c.POST("/v5/position/set-risk-limit", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetTradingStop Set the take profit, stop loss or trailing stop for the position.
func (c *Client) SetTradingStop(request models.SetTradingStopRequest) (*models.EmptyResponse, error) {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/position/trading-stop", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// SetAutoAddMargin Turn on/off auto-add-margin for isolated margin position
func (c *Client) SetAutoAddMargin(request models.SetAutoAddMarginRequest) (*models.EmptyResponse, error) {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/position/set-auto-add-margin", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetExecutionHalfYear Query users' execution records, sorted by execTime in descending order
func (c *Client) GetExecutionHalfYear(request models.GetExecutionHalfYearRequest) (*models.GetExecutionHalfYearResponse, error) {
	var respBody models.Response[models.GetExecutionHalfYearResponse]
	err := c.GET("/v5/execution/list", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetClosedPnL Query user's closed profit and loss records. The results are sorted by createdTime in descending order.
func (c *Client) GetClosedPnL(request models.GetClosedPnLRequest) (*models.GetClosedPnLResponse, error) {
	var respBody models.Response[models.GetClosedPnLResponse]
	err := c.GET("/v5/position/closed-pnl", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}
