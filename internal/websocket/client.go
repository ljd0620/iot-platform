package websocket

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

func (c *WebSocketClient) Connect(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("invalid websocket URL: %v", err)
	}

	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %v", err)
	}

	return nil
}

func (c *WebSocketClient) SendMessage(message []byte) error {
	if c.conn == nil {
		return fmt.Errorf("websocket connection is not established")
	}

	err := c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}

func (c *WebSocketClient) ReceiveMessage() ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("websocket connection is not established")
	}

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to receive message: %v", err)
	}

	return message, nil
}

func (c *WebSocketClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *WebSocketClient) KeepAlive(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.SendMessage([]byte("ping")); err != nil {
				fmt.Println("keep-alive failed:", err)
				return
			}
		}
	}
}