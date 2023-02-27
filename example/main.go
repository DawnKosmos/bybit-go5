package main

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5"
	"github.com/DawnKosmos/bybit-go5/models"
)

func main() {

	a := &bybit.Account{
		PublicKey: "",
		SecretKey: "",
	}
	c, _ := bybit.New(nil, bybit.URL, a, true)
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

	if err := WSExample(); err != nil {
		fmt.Println(err)
	}
}
