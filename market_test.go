package bybit

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"testing"
)

func TestClient_GetKline(t *testing.T) {
	cl, err := New(nil, URL, nil, false)
	Fail(err, t)
	res, err := cl.GetOrderbook(models.GetOrderbookRequest{
		Category: "spot",
		Symbol:   "BTCUSDT",
		Limit:    50,
	})
	Fail(err, t)

	fmt.Println(res)

}

func Fail(err error, t *testing.T) {
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
