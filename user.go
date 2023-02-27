package bybit

import "github.com/DawnKosmos/bybit-go5/models"

// CreateSubUID Create a new sub user id. Use master user's api key only.
func (c *Client) CreateSubUID(request models.CreateSubUIDRequest) (*models.CreateSubUIDResponse, error) {
	var respBody models.Response[models.CreateSubUIDResponse]
	err := c.POST("/v5/user/create-sub-member", request, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// GetSubUIDList Get all sub uid of master account. Use master user's api key only.
func (c *Client) GetSubUIDList() (*models.GetSubUIDListResponse, error) {
	var respBody models.Response[models.GetSubUIDListResponse]
	err := c.GET("/v5/user/query-sub-members", nil, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

// FrozeSubUID Froze sub uid. Use master user's api key only.
func (c *Client) FrozeSubUID(request models.FrozeSubUIDRequest) error {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/user/frozen-sub-member", request, &respBody)
	if err != nil {
		return err
	}
	return nil
}

// GetAPIKeyInformation Get the information of the api key. Use the api key pending to be checked to call the endpoint. Both master and sub user's api key are applicable.
func (c *Client) GetAPIKeyInformation() (*models.GetAPIKeyInformationResponse, error) {
	var respBody models.Response[models.GetAPIKeyInformationResponse]
	err := c.GET("/v5/user/query-api", nil, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody.Result, nil
}

func (c *Client) DeleteSubAPIKey() error {
	var respBody models.Response[models.EmptyResponse]
	err := c.POST("/v5/user/delete-sub-api", nil, &respBody)
	if err != nil {
		return err
	}
	return nil
}
