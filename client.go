package bybit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

const (
	URL     = "https://api.bybit.com"
	TESTURL = "https://api-testnet.bybit.com"
)

type Account struct {
	PublicKey string //Only needed for Private Api Calls
	SecretKey string //Only needed for Private Api Calls
}

type Client struct {
	url     *url.URL
	a       *Account
	client  *http.Client
	isDebug bool
}

// New Creates a new Bybit Client, Account is Optional but needed for Private Api Calls
func New(client *http.Client, baseurl string, a *Account, debug bool) (*Client, error) {
	var c *Client = new(Client)
	var err error
	if a != nil {
		c.a = a
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
	// Adds the path to the base Url
	reqLink, err := c.url.Parse(path)
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

	// Prepare Get Request
	req, err := http.NewRequest("GET", reqLink.String(), nil)
	if err != nil {
		return err
	}
	if c.a != nil { // Sign Request
		SignGET(c.a, req, req.URL.RawQuery)
	}

	if c.isDebug {
		fmt.Println("Get Request", req.URL.String())
	}

	//Do Request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Read Json Body and Unmarshal
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if c.isDebug {
		fmt.Println(result.(models.ReturnCode).Return())
	}
	res, ok := result.(models.ReturnCode)
	if ok {
		code, str := res.Return()
		if code != 0 {
			err = errors.New(str)
		}
	}

	return err
}

func (c *Client) POST(path string, queryJson any, result any) (err error) {
	// Adds the path to the base Url
	reqLink, err := c.url.Parse(path)
	if err != nil {
		return err
	}

	//Parse Struct into Json Byte
	jsonBody, err := json.Marshal(queryJson)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", reqLink.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	if c.a != nil { // Sign Request
		SignPOST(c.a, req, jsonBody)
	}

	//Do Request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Read Json Body and Unmarshal
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if c.isDebug {
		fmt.Println(result.(models.ReturnCode).Return())
	}

	res, ok := result.(models.ReturnCode)
	if ok {
		code, str := res.Return()
		if code != 0 {
			err = errors.New(str)
		}
	}

	return err
}
