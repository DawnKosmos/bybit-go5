package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// GetKline Query the kline data. Charts are returned in groups based on the requested interval. // Limit is 200
func (c *Client) GetKline(request models.GetKlineRequest) (*models.GetKlineResponse, error) {
	var respBody models.Response[models.GetKlineResponse]
	err := c.GET("/v5/market/kline", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetMarkPriceKline Query the mark price kline data. Charts are returned in groups based on the requested interval.
func (c *Client) GetMarkPriceKline(request models.GetKlineRequest) (*models.GetMarkPricelineResponse, error) {
	var respBody models.Response[models.GetMarkPricelineResponse]
	err := c.GET("/v5/market/mark-price-kline", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetIndexPriceKline Query the index price kline data. Charts are returned in groups based on the requested interval.
func (c *Client) GetIndexPriceKline(request models.GetKlineRequest) (*models.GetMarkPricelineResponse, error) {
	var respBody models.Response[models.GetMarkPricelineResponse]
	err := c.GET("/v5/market/index-price-kline", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetPremiumIndexPriceKline Retrieve the premium index price kline data. Charts are returned in groups based on the requested interval.
func (c *Client) GetPremiumIndexPriceKline(request models.GetKlineRequest) (*models.GetMarkPricelineResponse, error) {
	var respBody models.Response[models.GetMarkPricelineResponse]
	err := c.GET("/v5/market/index-price-kline", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

type GetInstrumentsInfo interface {
	models.GetInstrumentsInfoResponseLinear | models.GetInstrumentsInfoResponseSpot | models.GetInstrumentsInfoResponseOption
}

// GetInstrumentsInfoLinear
func (c *Client) GetInstrumentsInfoLinear(request models.GetInstrumentsInfoRequest) (*models.GetInstrumentsInfoResponseLinear, error) {
	var respBody models.Response[models.GetInstrumentsInfoResponseLinear]
	err := c.GET("/v5/market/instruments-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetInstrumentsInfoOption
func (c *Client) GetInstrumentsInfoOption(request models.GetInstrumentsInfoRequest) (*models.GetInstrumentsInfoResponseOption, error) {
	request.Category = "option"
	var respBody models.Response[models.GetInstrumentsInfoResponseOption]
	err := c.GET("/v5/market/instruments-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetInstrumentsInfoSpot
func (c *Client) GetInstrumentsInfoSpot(request models.GetInstrumentsInfoRequest) (*models.GetInstrumentsInfoResponseSpot, error) {
	request.Category = "spot"
	var respBody models.Response[models.GetInstrumentsInfoResponseSpot]
	err := c.GET("/v5/market/instruments-info", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetOrderbook
func (c *Client) GetOrderbook(request models.GetOrderbookRequest) (*models.GetOrderbookResponse, error) {
	var respBody models.Response[models.GetOrderbookResponse]
	err := c.GET("/v5/market/orderbook", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetTickersLinear
func (c *Client) GetTickersLinear(request models.GetTickersRequest) (*models.GetTickersLinearResponse, error) {
	var respBody models.Response[models.GetTickersLinearResponse]
	err := c.GET("/v5/market/tickers", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetTickersOption
func (c *Client) GetTickersOption(request models.GetTickersRequest) (*models.GetTickersOptionResponse, error) {
	var respBody models.Response[models.GetTickersOptionResponse]
	err := c.GET("/v5/market/tickers", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

func (c *Client) GetTickersSpot(request models.GetTickersRequest) (*models.GetTickersSpotResponse, error) {
	var respBody models.Response[models.GetTickersSpotResponse]
	err := c.GET("/v5/market/tickers", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}
