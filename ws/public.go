package ws

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
)

/*
Orderbook
orderbook.{depth}.{symbol}
Linear & inverse:
Level 1 data, push frequency: 10ms
Level 50 data, push frequency: 20ms
Level 200 data, push frequency: 100ms
Level 500 data, push frequency: 100ms

Spot:
Level 1 data, push frequency: 10ms
Level 50 data, push frequency: 20ms

Option:
Level 25 data, push frequency: 20ms
Level 100 data, push frequency: 100ms
*/
func (s *Stream) Orderbook(ticker string, depth int, fn func(e *models.WsOrderbook)) error {
	topic := fmt.Sprintf("orderbook.%d.%s", depth, ticker)
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) PublicTrade(ticker string, fn func(e *models.WsTrade)) error {
	topic := fmt.Sprintf("publicTrade.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) TickerOption(ticker string, fn func(e *models.WsTickerOption)) error {
	topic := fmt.Sprintf("tickers.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) TickerLinear(ticker string, fn func(e *models.WsTickerLinear)) error {
	topic := fmt.Sprintf("tickers.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) TickerSpot(ticker string, fn func(e *models.WsTickerSpot)) error {
	topic := fmt.Sprintf("tickers.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

/*
Kline
Available intervals:
1 3 5 15 30 (min)
60 120 240 360 720 (min)
D (day)
W (week)
M (month)
Push frequency: 1-60s
Topic:
kline.{interval}.{symbol} e.g., kline.30.BTCUSDT
*/
func (s *Stream) Kline(ticker string, interval string, fn func(e *models.WsKline)) error {
	topic := fmt.Sprintf("kline.%s.%s", interval, ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

// Liquidation Push frequency: real-time
func (s *Stream) Liquidation(ticker string, fn func(e *models.WsLiquidation)) error {
	topic := fmt.Sprintf("liquidation.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

/*
LeveragedTokenKline
Available intervals:

1 3 5 15 30 (min)
60 120 240 360 720 (min)
D (day)
W (week)
M (month)
Push frequency: 1-60s

Topic:
kline_lt.{interval}.{symbol} e.g., kline_lt.30.BTC3SUSDT
*/
func (s *Stream) LeveragedTokenKline(ticker string, interval string, fn func(e *models.WsLTKline)) error {
	topic := fmt.Sprintf("kline_lt.%s.%s", interval, ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) LeveragedTokenTicker(ticker string, fn func(e *models.WsLTTicker)) error {
	topic := fmt.Sprintf("lt.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) LeveragedTokenNav(ticker string, fn func(e *models.WsLTNav)) error {
	topic := fmt.Sprintf("kline_lt.%s", ticker)
	err := s.Send(models.WsPublicRequest{ReqId: topic, Op: "subscribe", Args: []string{topic}})
	if err != nil {
		return err
	}

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) resubscribe() error {
	subs := s.GetSubscriptions()
	return s.Send(models.WsPublicRequest{
		Op:   "subscribe",
		Args: subs,
	})
}
