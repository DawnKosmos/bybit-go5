package main

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5"
	"github.com/DawnKosmos/bybit-go5/models"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{}

	by, err := bybit.New(client, bybit.TESTURL, &bybit.Account{
		PublicKey: "",
		SecretKey: "",
	}, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := by.GetMarketKline(models.GetMarketKlineRequest{
		Category: "linear",
		Symbol:   "BTCUSD",
		Interval: "360",
		Start:    0,
		End:      time.Now().UnixMilli(),
		Limit:    10,
	})

	fmt.Println(res, err)

}
