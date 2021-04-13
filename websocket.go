package upbit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
	Code     string
	Ticket   string
	Type     string
	Codes    []string
}

func NewWsConfig(endpoint, code, ticket, typ string, Codes []string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
		Code:     code,
		Ticket:   ticket,
		Type:     typ,
		Codes:    Codes,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	c, _, err := websocket.DefaultDialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	fields := []struct {
		Ticket string   `json:"ticket,omitempty"`
		Type   string   `json:"type,omitempty"`
		Codes  []string `json:"codes,omitempty"`
	}{
		{Ticket: cfg.Ticket},               //{Ticket: "test12345"},
		{Type: cfg.Type, Codes: cfg.Codes}, //{Type: "ticker", Codes: []string{"KRW-BTC"}},
	}
	bFields, _ := json.Marshal(fields)
	err = c.WriteMessage(websocket.TextMessage, bFields)
	if err != nil {
		fmt.Printf("write_err %+v", err)
		return nil, nil, err
	}
	if err != nil {
		c.Close()
		return nil, nil, err
	}
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)

		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}

		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				c.Close()
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
