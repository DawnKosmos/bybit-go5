package ws

import (
	"context"
	"errors"
	"github.com/DawnKosmos/bybit-go5"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/chuckpreslar/emission"
	"sync"
	"time"
)

type Stream struct {
	started       time.Time
	link          WsLink
	autoReconnect bool
	debugMode     bool
	a             *bybit.Account

	mu sync.RWMutex
	heartCancel chan bool
	isConnected bool

	conn *websocket.Conn
	rpcConn *jsonrpc2.Conn
	emitter *emission.Emitter


	subscriptions []string
}

type Config struct {
	Ctx context.Context
	Endpoint string
	A *bybit.Account
	AutoReconnect bool
	Debug bool
}

func NewStream(cfg Config)*Stream{
	ctx := cfg.Ctx
	if ctx == nil {
		ctx = context.Background()
	}
	s := Stream{
		started:       time.Now(),
		link:          cfg.Endpoint,
		autoReconnect: cfg.AutoReconnect,
		debugMode:     cfg.Debug,
		a:             nil,
		conn:          nil,
		heartCancel:   nil,
		isConnected:   false,
		subscriptions: nil,
	}
	
	s.start()
}


func (s *Stream) start() error {
	s.setIsConnected(false)
	s.heartCancel = make(chan bool)
	for i := 0; i < MAXTRYTIMES; i++ {
		conn, _, err := s.connect()
		if err != nil {
			log.Println(err)
			timeAsSleep := (i + 1) * 5
			log.Printf("Sleep %vs\n", timeAsSleep)
			time.Sleep(time.Duration(timeAsSleep) * time.Second)
		}
		s.conn = conn
		break
	}
	if s.conn == nil {
		return errors.New("connection failed")
	}

}

func (s *Stream) connect() (*websocket.Conn, *http.Response, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, resp, err := websocket.Dial(ctx, s.link, &websocket.DialOptions{})
	if err == nil {
		conn.SetReadLimit(READLIMIT)
	}
	return conn, resp, err
}

func (s *Stream) reconnect() {

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

// heartBeat pings the WSS Endpoint every 20 secs, to keep the Connection Alive
func (s *Stream) heartBeat() {
	t := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-t.C:
			s.Ping()
		case <-s.heartCancel:
			return
		}
	}
}

func (s *Stream) Ping() {
}


func (s )
