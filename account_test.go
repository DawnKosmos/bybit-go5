package bybit

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"testing"
)

func TestGetAccount(t *testing.T) {
	var c, _ = New(nil, TESTURL, &Account{
		PublicKey: "XXXXX",
		SecretKey: "XXXXX",
	}, true)

	resp, err := c.GetWalletBalance(models.GetWalletBalanceRequest{
		AccountType: "unified",
		Coin:        "USDT",
	})
	fail(err, t)
	fmt.Println(*resp)

	AccInfo, err := c.GetAccountInfo()
	fail(err, t)
	fmt.Println(*AccInfo)

}

func fail(err error, t *testing.T) {
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
