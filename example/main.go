package main

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5"
	"github.com/DawnKosmos/bybit-go5/models"
	"os"
	"strings"
)

func main() {
	by, err := os.ReadFile("api_key.txt")
	if err != nil {
		return
	}
	bb := strings.Split(string(by), " ")

	a := &bybit.Account{
		PublicKey: bb[0],
		SecretKey: bb[1],
	}
	c, _ := bybit.New(nil, bybit.TESTURL, a, true)
	resp, err := c.GetKline(models.GetKlineRequest{
		Category: "linear",
		Symbol:   "BTCUSDT",
		Interval: "240",
		Limit:    200,
	})
	if err != nil {
		return
	}
	for _, v := range resp.ToCandle() {
		fmt.Println(v)
	}

	c.PlaceOrder(models.PlaceOrderRequest{
		Category:   "spot",
		Symbol:     "BTCUSDT",
		IsLeverage: 0,
		Side:       "Buy",
		OrderType:  "Limit",
		Qty:        "1",
		Price:      "21000",
	})

	if err := WSExample(); err != nil {
		fmt.Println(err)
	}
}
