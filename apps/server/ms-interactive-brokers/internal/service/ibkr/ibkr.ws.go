package ibkr

import (
	"encoding/json"
	"fmt"
	"ms-interactive-brokers/pkg/logster"
	"ms-interactive-brokers/pkg/utils"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketMessage struct {
	Type      int       `json:"type"`
	Data      []byte    `json:"data"`
	Topic     string    `json:"topic,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type MarketDataMessage struct {
	Topic    string                 `json:"topic"`
	ConID    int                    `json:"conid"`
	ConIDEx  string                 `json:"conidEx"`
	ServerID string                 `json:"server_id"`
	Updated  int64                  `json:"_updated"`
	Fields   map[string]interface{} `json:"-"` // All other fields
}

type WebSocketClient struct {
	conn         *websocket.Conn
	accessToken  string
	sessionToken string

	MarketDataChannel chan MarketDataMessage
}

func NewWebSocketClient(accessToken, sessionToken string) *WebSocketClient {
	logster.StartFuncLogMsg(fmt.Sprintf("accessToken: %s, sessionToken: %s", accessToken, sessionToken))
	client := &WebSocketClient{
		accessToken:  accessToken,
		sessionToken: sessionToken,

		MarketDataChannel: make(chan MarketDataMessage, 1000),
	}
	logster.EndFuncLog()
	return client
}

func (ws *WebSocketClient) WriteTextMessage(message string) error {
	logster.StartFuncLogMsg(fmt.Sprintf("message: %s", message))

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to send WebSocket message: %w", err))
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func (ws *WebSocketClient) onMessage(messageType int, data []byte) {
	logster.Info(fmt.Sprintf("Message Received: messageType: %d, data: %s", messageType, string(data)))
	// Add your message handling logic here
	msg := WebSocketMessage{
		Type:      messageType,
		Data:      data,
		Timestamp: time.Now(),
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err == nil {
		if topic, ok := parsed["topic"].(string); ok {
			msg.Topic = topic

			if strings.HasPrefix(topic, "smd+") {
				ws.handleMarketDataMessage(parsed)
			}
		}
	}
}

func (ws *WebSocketClient) handleMarketDataMessage(data map[string]interface{}) {
	marketMsg := MarketDataMessage{
		Fields: make(map[string]interface{}),
	}

	// Extract known fields
	if topic, ok := data["topic"].(string); ok {
		marketMsg.Topic = topic
	}
	if conid, ok := data["conid"].(float64); ok {
		marketMsg.ConID = int(conid)
	}
	if conidEx, ok := data["conidEx"].(string); ok {
		marketMsg.ConIDEx = conidEx
	}
	if serverId, ok := data["server_id"].(string); ok {
		marketMsg.ServerID = serverId
	}
	if updated, ok := data["_updated"].(float64); ok {
		marketMsg.Updated = int64(updated)
	}

	// Extract all field data
	for key, value := range data {
		if key != "topic" && key != "conid" && key != "conidEx" && key != "server_id" && key != "_updated" {
			marketMsg.Fields[key] = value
		}
	}

	// Non-blocking send to market data channel
	select {
	case ws.MarketDataChannel <- marketMsg:
	default:
		fmt.Printf("Warning: MarketDataChan is full, dropping message\n")
	}
}

func (ws *WebSocketClient) onError(err error) {
	logster.Error(err, fmt.Sprintf("WebSocket error: %v", err))
}

func (ws *WebSocketClient) onClose(code int, text string) {
	errClosing := ws.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, text),
		time.Now().Add(time.Minute),
	)

	if errClosing != nil {
		logster.Error(errClosing, "Did not close connection")
	}
}

func (ws *WebSocketClient) Connect() error {
	logster.StartFuncLogMsg(fmt.Sprintf("accessToken: %s, sessionToken: %s", ws.accessToken, ws.sessionToken))

	// Parse the URL
	u := url.URL{
		Scheme:   "wss",
		Host:     "api.ibkr.com",
		Path:     "/v1/api/ws",
		RawQuery: fmt.Sprintf("oauth_token=%s", ws.accessToken),
	}

	// Set up headers with cookie
	headers := http.Header{}
	headers.Set("Cookie", fmt.Sprintf("api=%s", ws.sessionToken))

	// Create WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to connect to WebSocket: %w", err))
		logster.EndFuncLog()
		return err
	}

	ws.conn = conn

	logster.EndFuncLogMsg("WebSocket connection established")
	return nil
}

func (ws *WebSocketClient) ReadMessage() *string {
	logster.StartFuncLog()

	_, message, err := ws.conn.ReadMessage()
	if err != nil {
		logster.Info("WebSocket read error occurred")
		return nil
	}
	return utils.Ptr(string(message))
}

func (ws *WebSocketClient) RunForever() {
	logster.StartFuncLog()
	defer func() {
		ws.conn.Close()
		logster.EndFuncLogMsg("RunForever completed")
	}()

	// Set up signal handling for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Channel to handle connection done
	done := make(chan struct{})

	// Goroutine to read messages
	go func() {
		defer close(done)
		for {
			messageType, message, err := ws.conn.ReadMessage()
			if err != nil {
				ws.onError(err)
				return
			}
			ws.onMessage(messageType, message)
		}
	}()

	// Main loop
	for {
		select {
		case <-done:
			logster.Info("WebSocket done signal received")
			return
		case <-interrupt:
			logster.Info("Interrupt signal received")

			// Cleanly close the connection by sending a close message
			logster.Info("Sending close message")
			err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logster.Error(err, fmt.Sprintf("error sending close message: %v", err))
				return
			}
			logster.Info("Close message sent")

			// Wait for the server to close the connection
			logster.Info("Waiting for server to close connection")
			select {
			case <-done:
				logster.Info("Server closed connection")
			case <-time.After(time.Second):
				logster.Info("Timeout waiting for server close")
			}
			return
		}
	}
}

func (ws *WebSocketClient) Close() {
	logster.StartFuncLog()
	if ws.conn != nil {
		ws.onClose(websocket.CloseNormalClosure, "Connection closed by client")
		ws.conn.Close()
		logster.EndFuncLogMsg("connection closed")
	} else {
		logster.EndFuncLogMsg("connection already closed")
	}
}

func (ws *WebSocketClient) GetMarketDataChannel() <-chan MarketDataMessage {
	return ws.MarketDataChannel
}
