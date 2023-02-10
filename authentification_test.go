package bybit

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"os"
	"strings"
	"testing"
)

func TestAuthentification(t *testing.T) {
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

	res, err := c.GetWalletBalance(models.GetWalletBalanceRequest{
		AccountType: "UNIFIED",
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(res)

}
