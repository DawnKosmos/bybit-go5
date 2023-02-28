package ws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"time"
)

// Authentification copied from github.com/frankrap/bybit-api/
func (s *Stream) Authentication() error {

	expires := time.Now().UnixMilli() + s.expires
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(s.a.SecretKey))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))

	return s.Send(models.WsPrivateRequest{
		ReqId: "auth",
		Op:    "auth",
		Args:  []any{s.a.PublicKey, expires, signature},
	})
}

func (s *Stream) Position(fn func(e *models.WsPosition)) error {
	topic := "position"
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) Execution(fn func(e *models.WsExecution)) error {
	topic := "execution"
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) Order(fn func(e *models.WsOrder)) error {
	topic := "order"
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) Wallet(fn func(e *models.WsWallet)) error {
	topic := "wallet"
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}

func (s *Stream) Greek(fn func(e *models.WsGreek)) error {
	topic := "greeks"
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "subscribe",
		Args:  []string{topic},
	})

	StoreFunc(s.sm, s.debugMode, topic, fn)
	return err
}
