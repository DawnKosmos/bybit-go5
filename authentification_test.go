package bybit

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"os"
	"strings"
	"testing"
)

func TestAuthentification(t *testing.T) {
	by, err := os.ReadFile("api_key_live.txt")
	if err != nil {
		t.FailNow()
	}
	bb := strings.Split(string(by), " ")

	client, err := New(nil, URL, &Account{
		PublicKey: bb[0],
		SecretKey: bb[1],
	}, true)
	if err != nil {
		t.FailNow()
	}

	resp, err := client.PlaceOrder(models.PlaceOrderRequest{
		Category:  "linear",
		Symbol:    "BTCUSDT",
		Side:      "Buy",
		OrderType: "Limit",
		Qty:       "0.00111111",
		Price:     "22000.12",
	})

	fmt.Println(resp)
}
