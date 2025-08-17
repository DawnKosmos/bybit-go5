package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DawnKosmos/bybit-go5"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"sync"
	"time"
)

/*
Stream is only 1 websocket connection
It only can serve 1 of the 5 Endpoints at once.
In the Config struct you can choose which Endpoints you want to visit.
*/
type Stream struct {
	started       time.Time       // just for log purpose
	ctx           context.Context // context
	Link          string          // config.Endpoint and config.Testnet get you the ws Link
	autoReconnect bool            // AutoReconnect when Disconnect.
	debugMode     bool            // debug Mode atm only logs stuff
	a             *bybit.Account  //field is only needed for Private Endpoint

	expires     int64 // in Seconds. The Time after a private Channel expires.
	mu          sync.RWMutex
	heartCancel chan bool // Cancels the go routine  heartBeat
	isConnected bool

	disconnect chan struct{} // Channel get called when WS disconnects, triggers the reconnect function

	conn *websocket.Conn // Websocket connection

	subscriptions map[string]struct{} // Overview of Subscriptions

	//sm = SyncMap. This map saves the function we set when we subscribe to Endpoints e.g. ws.Position(e func(e *models.Position)){}
	// it data type is map[string]func([]byte)
	sm *sync.Map
}

type Config struct {
	Id            string          // [optional] Just an Id for yourself
	Ctx           context.Context // [optional]
	Expire        int64           // in Seconds. The Time after a private Channel expires.
	Endpoint      WsLink          // Set which Endpoint you want to subscribe SPOT | LINEAR | INVERSE | OPTION | PRIVATE
	A             *bybit.Account  //  [optional] only needed for PRIVATE
	AutoReconnect bool            // [optional]
	Debug         bool            // [optional]
	TestNet       bool            // [optional] set true if you use testnet!
}

func New(cfg Config) *Stream {
	ctx := cfg.Ctx
	if ctx == nil {
		ctx = context.Background()
	}
	s := Stream{
		started:       time.Now(),
		ctx:           ctx,
		Link:          GetWsLink(cfg.Endpoint, cfg.TestNet),
		autoReconnect: cfg.AutoReconnect,
		debugMode:     cfg.Debug,
		a:             cfg.A,
		isConnected:   false,
		subscriptions: make(map[string]struct{}),
		sm:            &sync.Map{},
		expires:       cfg.Expire * 1000,
	}

	if err := s.start(); err != nil {
		fmt.Println(err)
		return nil
	}

	return &s
}

// starts a new websocket connection
func (s *Stream) start() error {
	//resets trackers
	s.setIsConnected(false)
	s.heartCancel = make(chan bool)
	for i := 0; i < MAXTRYTIMES; i++ {
		conn, err := s.connect()
		if err != nil {
			log.Println(err)
			timeAsSleep := (i + 1) * 5
			log.Printf("Sleep %vs\n", timeAsSleep)
			time.Sleep(time.Duration(timeAsSleep) * time.Second)
			continue
		}

		s.conn = conn
		s.disconnect = make(chan struct{})
		break
	}
	if s.conn == nil {
		return errors.New("connection failed")
	}
	log.Println("Connected")
	s.setIsConnected(true)
	go s.Read() // Reads WebSocket Messages

	if s.a != nil {
		err := s.Authentication()
		if err != nil {
			log.Println(err)
		}
	}

	if s.autoReconnect {
		go s.reconnect()
	}

	go s.heartBeat()

	return nil
}

func (s *Stream) connect() (*websocket.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, s.Link, &websocket.DialOptions{})
	if err == nil {
		conn.SetReadLimit(READLIMIT)
	}

	return conn, err
}

func (s *Stream) Send(request any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if !s.isConnected {
		return errors.New("not Connected")
	}
	if s.debugMode {
		log.Println("Request:", request)
	}
	return wsjson.Write(s.ctx, s.conn, request) //Send Request As Json
}

// Filter Message Type
type Event struct {
	Op      string `json:"op,omitempty"`
	Success bool   `json:"success,omitempty"`
	ReqId   string `json:"req_id,omitempty"`
	Topic   string `json:"topic,omitempty"`
}

// TODO add unsubscribe and take care Event Responses
func (s *Stream) Read() {
	/*
		ctx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
		defer cancel()
	*/
	for {
		_, data, err := s.conn.Read(s.ctx)

		if err != nil {
			if s.debugMode {
				log.Println(err)
			}
			if s.autoReconnect {
				close(s.disconnect) // reconnect attempts here after 1 seconds
			}
			break
		}

		var e Event
		err = json.Unmarshal(data, &e)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(e.Topic) == 0 { // Check if Topic is Empty
			if e.Success {
				s.manageSubscription(e.Op, e.ReqId)
			} else {
				log.Println(string(data))
			}
			continue
		}

		fn, ok := s.sm.Load(e.Topic) //Load the function
		if ok {
			ff, ok := fn.(func([]byte))
			if ok {
				go ff(data)
			}
		}
	}
}

var void struct{}

/*
StoreFunc saves the Topic String with a generated  function that gets a []byte and converts it into the Generics type T
The Map is of type map[string]func([]byte)
*/
func StoreFunc[T any](sm *sync.Map, debug bool, key string, fn func(*T)) {
	sm.Store(key,
		func(data []byte) {
			var e T                         // Creates a Struct of Type T e.g. WsTickerLinear
			err := json.Unmarshal(data, &e) //Parses the data into the Struct
			if debug && err != nil {
				log.Println(err)
			}
			fn(&e) // Execute the provided function
		})
}

func (s *Stream) GetSubscriptions() (o []string) {
	for v, _ := range s.subscriptions {
		o = append(o, v)
	}
	return o
}
