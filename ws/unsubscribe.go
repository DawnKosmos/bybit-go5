package ws

import (
	"github.com/DawnKosmos/bybit-go5/models"
)

func (s *Stream) Unsubscribe(topic string) error {
	err := s.Send(models.WsPublicRequest{
		ReqId: topic,
		Op:    "unsubscribe",
		Args:  []string{topic},
	})
	if err != nil {
		return err
	}

	s.sm.Delete(topic)
	return nil
}

// TODO unsubscribe
