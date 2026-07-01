package websocketUtils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alibabacloud-go/tea/dara"
)

type WebSocketClient struct {
	wsClient dara.WebSocketClient
	response *dara.Response
	// AWAP request-response pattern (pending requests waiting for ACK)
	pendingRequests   map[string]chan interface{} // messageID -> response channel
	pendingRequestsMu sync.RWMutex                // Protects pendingRequests map
}

func NewWebSocketClient(wsClient dara.WebSocketClient, response *dara.Response) *WebSocketClient {
	return &WebSocketClient{
		wsClient:        wsClient,
		response:        response,
		pendingRequests: make(map[string]chan interface{}),
	}
}

func CreateWebSocketClient(c interface{}) *WebSocketClient {
	if c == nil {
		return nil
	}
	if wsClient, ok := c.(*WebSocketClient); ok {
		return wsClient
	}
	return nil
}

func (c *WebSocketClient) GetResponse() *dara.Response {
	return c.response
}

func (c *WebSocketClient) Close() error {
	if c.wsClient != nil {
		return c.wsClient.Close()
	}
	return nil
}

func (c *WebSocketClient) Reconnect() (*dara.Response, error) {
	if c.wsClient != nil {
		return c.wsClient.Reconnect()
	}
	return nil, fmt.Errorf("wsClient is nil")
}

func (c *WebSocketClient) ReconnectGracefully() (*dara.Response, error) {
	if c.wsClient != nil {
		return c.wsClient.ReconnectGracefully()
	}
	return nil, fmt.Errorf("wsClient is nil")
}

func (c *WebSocketClient) Disconnect() error {
	if c.wsClient != nil {
		return c.wsClient.Disconnect()
	}
	return nil
}

func (c *WebSocketClient) IsConnected() bool {
	if c.wsClient != nil {
		return c.wsClient.IsConnected()
	}
	return false
}

func (c *WebSocketClient) Validate() error {
	if c.wsClient != nil {
		return nil
	}
	return fmt.Errorf("failed to build websocket client")
}

func (c *WebSocketClient) GetSessionInfo() *dara.WebSocketSessionInfo {
	if c.wsClient != nil {
		return c.wsClient.GetSessionInfo()
	}
	return nil
}

// SendAwapMessage sends an AWAP protocol message
// AWAP protocol uses frame format: text headers + JSON payload
// Format: "type:request\nseq:1\ntimestamp:1234567890\nid:msg-001\nack:required\n\n{JSON payload}"
func (c *WebSocketClient) SendAwapTextMessage(message *AwapMessage) error {
	messageText, err := BuildAwapMessageText(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendText(messageText)
}

func (c *WebSocketClient) SendRawAwapTextMessage(msgType AwapMessageType, payload interface{}) error {
	message := NewAwapMessage(payload, WithType(msgType))
	messageText, err := BuildAwapMessageText(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendText(messageText)
}

func (c *WebSocketClient) SendRawAwapTextMessageWithId(msgType AwapMessageType, id string, payload interface{}) error {
	message := NewAwapMessage(payload, WithType(msgType), WithID(id))
	messageText, err := BuildAwapMessageText(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendText(messageText)
}

func (c *WebSocketClient) SendAwapBinaryMessage(message *AwapMessage) error {
	messageBinary, err := BuildAwapMessageBinary(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendBinary(messageBinary)
}

func (c *WebSocketClient) SendRawAwapBinaryMessage(msgType AwapMessageType, payload interface{}) error {
	message := NewAwapMessage(payload, WithType(msgType))
	messageBinary, err := BuildAwapMessageBinary(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendBinary(messageBinary)
}

func (c *WebSocketClient) SendRawAwapBinaryMessageWithId(msgType AwapMessageType, id string, payload interface{}) error {
	message := NewAwapMessage(payload, WithType(msgType), WithID(id))
	messageBinary, err := BuildAwapMessageBinary(message)
	if err != nil {
		return fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.wsClient.SendBinary(messageBinary)
}

// SendAwapRequestWithAck sends an AWAP request that requires acknowledgment and waits for response
func (c *WebSocketClient) SendAwapRequestWithAck(message *AwapMessage, timeout time.Duration) (interface{}, error) {
	messageText, err := BuildAwapMessageText(message)
	if err != nil {
		return nil, fmt.Errorf("failed to build AWAP message: %w", err)
	}
	return c.SendRequest(message.ID, messageText, timeout)
}

func (c *WebSocketClient) SendRequest(ackID string, messageText string, timeout time.Duration) (interface{}, error) {
	if ackID == "" {
		return nil, errors.New("message ID cannot be empty for request-response pattern")
	}

	if messageText == "" {
		return nil, errors.New("message text cannot be empty for request-response pattern")
	}

	responseChan := make(chan interface{}, 1)

	c.pendingRequestsMu.Lock()
	c.pendingRequests[ackID] = responseChan
	c.pendingRequestsMu.Unlock()

	defer func() {
		c.pendingRequestsMu.Lock()
		delete(c.pendingRequests, ackID)
		c.pendingRequestsMu.Unlock()
		// If completed, sender closes it. If timeout, abandon it.
	}()

	err := c.wsClient.SendText(messageText)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	select {
	case response := <-responseChan:
		return response, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("request timeout after %v waiting for response to message ID: %s", timeout, ackID)
	}
}

func (c *WebSocketClient) completeRequest(ackID string, response interface{}) bool {
	c.pendingRequestsMu.RLock()
	responseChan, exists := c.pendingRequests[ackID]
	c.pendingRequestsMu.RUnlock()

	if !exists || responseChan == nil {
		return false
	}

	select {
	case responseChan <- response:
		return true
	default:
		// Channel full or closed
		return false
	}
}

func (c *WebSocketClient) SendGeneralTextMessage(text string) error {
	message := NewGeneralMessage(text)
	jsonData, err := message.ToJSON()
	if err != nil {
		return err
	}
	return c.wsClient.SendText(string(jsonData))
}

func (c *WebSocketClient) SendGeneralBinaryMessage(data []byte) error {
	return c.wsClient.SendBinary(data)
}

// generateMessageId generates a 32-character message ID
// 1. Generate random hex string (equivalent to UUID without hyphens)
// 2. If hex string length >= 32, return first 32 characters
// 3. Otherwise, combine timestamp and hex string
// 4. If combined length >= 32, return first 32 characters
// 5. If still less than 32, pad with zeros to 32 characters
func generateMessageId() string {
	// Generate 16 random bytes and convert to hex (32 characters, equivalent to UUID without hyphens)
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		// Fallback: use timestamp-based ID if random generation fails
		now := time.Now()
		timestamp := fmt.Sprintf("%d", now.Unix()*1000+int64(now.Nanosecond())/1e6)
		// Pad timestamp to 32 characters
		padded := fmt.Sprintf("%-32s", timestamp)
		return strings.ReplaceAll(padded, " ", "0")
	}
	hexStr := hex.EncodeToString(randomBytes)

	if len(hexStr) >= 32 {
		return hexStr[:32]
	}

	// Otherwise, combine timestamp and hex string
	now := time.Now()
	timestamp := fmt.Sprintf("%d", now.Unix()*1000+int64(now.Nanosecond())/1e6)
	combined := timestamp + hexStr

	if len(combined) >= 32 {
		return combined[:32]
	}

	// If still less than 32, pad with zeros to 32 characters
	// Format: %-32s pads on the right, then replace spaces with zeros
	padded := fmt.Sprintf("%-32s", combined)
	return strings.ReplaceAll(padded, " ", "0")
}
