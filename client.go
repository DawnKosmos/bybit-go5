package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

const (
	URL     = "https://api.bybit.com"
	TESTURL = "https://api-testnet.bybit.com\n"
)

type Account struct {
	PublicKey string //Only needed for Private Api Calls
	SecretKey string //Only needed for Private Api Calls
}

type Client struct {
	url     *url.URL
	public  string
	private string
	client  *http.Client
	isDebug bool
}

func New(client *http.Client, baseurl string, a *Account, debug bool) (*Client, error) {
	var c *Client = new(Client)
	var err error
	if a != nil {
		c.public = a.PublicKey
		c.private = a.SecretKey
	}
	c.isDebug = debug
	if client == nil {
		client = http.DefaultClient
	}
	c.client = client
	if baseurl == "" {
		return nil, errors.New("url can't be empty. Use bybit.URL or bybit.TESTURL")
	}
	c.url, err = url.Parse(baseurl)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) GET(path string, queryParameters any, result any) (err error) {
	reqLink, err := c.url.Parse(path) // Adds the path to the base Url
	if err != nil {
		return err
	}

	//Adds QueryParameters
	if queryParameters != nil {
		v, err := query.Values(queryParameters)
		if err != nil {
			return err
		}
		reqLink.RawQuery = v.Encode()
	}

	req, err := http.NewRequest("GET", reqLink.String(), nil)
	if err != nil {
		return err
	}

	fmt.Println(req.URL.String())
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}
