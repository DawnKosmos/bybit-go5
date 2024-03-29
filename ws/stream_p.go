package ws

import (
	"fmt"
	"github.com/DawnKosmos/bybit-go5/models"
	"log"
	"sync"
	"time"
)

//Help functions

// heartBeat pings the WSS Endpoint every 20 secs, to keep the Connection Alive
func (s *Stream) heartBeat() {
	t := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-t.C:
			err := s.Ping()
			if err != nil {
				fmt.Println(err)
			}
		case <-s.heartCancel:
			return
		}
	}

}

func (s *Stream) Ping() error {
	err := s.Send(models.WsPing{
		Op: "ping",
	})
	return err
}

func (s *Stream) setIsConnected(state bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isConnected = state
}

func (s *Stream) IsConnected() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.isConnected
}

func (s *Stream) reconnect() {
	notify := s.disconnect
	<-notify
	s.setIsConnected(false)
	log.Println("disconnect, reconnect...")
	close(s.heartCancel)
	time.Sleep(1 * time.Second)
	if err := s.start(); s.debugMode && err != nil {
		log.Println(err)
	}

	if err := s.resubscribe(); s.debugMode && err != nil {
		log.Println(err)
	}
}

func (s *Stream) ReturnSyncMap() *sync.Map {
	return s.sm
}

func (s *Stream) manageSubscription(op string, reqId string) {
	if s.debugMode {
		log.Println(op, reqId)
	}

	switch op {
	case "subscribe":
		s.subscriptions[reqId] = void
	case "unsubscribe":
		delete(s.subscriptions, reqId)
	default:
		log.Println(op, reqId)
	}
}
