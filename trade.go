package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// PlaceOrder This endpoint supports to create the order for spot, spot margin, linear perpetual, inverse futures and options.
func (c *Client) PlaceOrder(request models.PlaceOrderRequest) (*models.PlaceOrderResponse, error) {
	var respBody models.Response[models.PlaceOrderResponse]
	err := c.POST("/v5/order/create", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// AmendOrder Amend/Update Orders
func (c *Client) AmendOrder(request models.AmendOrderRequest) (*models.AmendOrderResponse, error) {
	var respBody models.Response[models.AmendOrderResponse]
	err := c.POST("/v5/order/amend", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// CancelOrder Cancels orders
func (c *Client) CancelOrder(request models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	var respBody models.Response[models.CancelOrderResponse]
	err := c.POST("/v5/order/cancel", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetOpenOrders Query unfilled or partially filled orders in real-time. To query older order records, please use the order history interface.
func (c *Client) GetOpenOrders(request models.GetOpenOrdersRequest) (*models.GetOpenOrdersResponse, error) {
	var respBody models.Response[models.GetOpenOrdersResponse]
	err := c.GET("/v5/order/realtime", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// CancelAllOrders Cancel all open orders
func (c *Client) CancelAllOrders(request models.CancelAllOrdersRequest) (*models.CancelAllOrdersResponse, error) {
	var respBody models.Response[models.CancelAllOrdersResponse]
	err := c.POST("/v5/order/cancel-all", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetOrderHistory Query order history. As order creation/cancellation is asynchronous, the data returned from this endpoint may delay.
func (c *Client) GetOrderHistory(request models.GetOrderHistoryRequest) (*models.GetOrderHistoryResponse, error) {
	var respBody models.Response[models.GetOrderHistoryResponse]
	err := c.GET("/v5/order/history", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// BatchPlaceOrder only supports Options!!
func (c *Client) BatchPlaceOrder(request models.BatchPlaceOrderRequest) (*models.BatchPlaceOrderResponse, error) {
	var respBody models.ResponseBatch[models.BatchPlaceOrderResponse]
	err := c.POST("/v5/order/create-batch", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// BatchAmendOrder only supports Options!!
func (c *Client) BatchAmendOrder(request models.BatchAmendOrderRequest) (*models.BatchAmendOrderResponse, error) {
	var respBody models.ResponseBatch[models.BatchAmendOrderResponse]
	err := c.POST("/v5/order/amend-batch", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}
