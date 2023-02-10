package bybit

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"os"
	"strings"
	"testing"
)

func TestTrade(t *testing.T) {
	by, _ := os.ReadFile("api_key.txt")
	a := strings.Split(string(by), " ")

	c, err := New(nil, TESTURL, &Account{
		PublicKey: a[0],
		SecretKey: a[1],
	}, true)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	r, err := c.PlaceOrder(models.PlaceOrderRequest{
		Category:  "linear",
		Symbol:    "BTCPERP",
		Side:      "Buy",
		OrderType: "Market",
		Qty:       "1",
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r)

}
