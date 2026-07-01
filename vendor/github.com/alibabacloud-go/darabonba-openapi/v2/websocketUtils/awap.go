package websocketUtils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/dara"
)

type AwapMessageType string
type AwapMessageFormat string

const (
	AwapMessageFormatText   AwapMessageFormat = "text"
	AwapMessageFormatBinary AwapMessageFormat = "binary"
)

type AwapMessage struct {
	// only type and id is required for awap message, and id can be autofilled when sending message
	// other fields are optional and can be set using withHeader
	Type    AwapMessageType   `json:"type"`
	ID      string            `json:"id"`
	Headers map[string]string `json:"headers,omitempty"`
	Payload interface{}       `json:"payload,omitempty"`
	Format  AwapMessageFormat `json:"format,omitempty"`
}

func (m *AwapMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *AwapMessage) WithHeader(key, value string) *AwapMessage {
	if m.Headers == nil {
		m.Headers = make(map[string]string)
	}
	m.Headers[key] = value
	return m
}

func (m *AwapMessage) WithFormat(format AwapMessageFormat) *AwapMessage {
	m.Format = format
	return m
}

type AwapWebSocketHandler interface {
	dara.WebSocketHandler

	HandleAwapMessage(session *dara.WebSocketSessionInfo, message *AwapMessage) error
}

// AbstractAwapWebSocketHandler provides default implementations for AwapWebSocketHandler interface
// It embeds AbstractWebSocketHandler for base WebSocketHandler methods
// Users can embed this struct in their custom AWAP handlers
type AbstractAwapWebSocketHandler struct {
	dara.AbstractWebSocketHandler
}

func (h *AbstractAwapWebSocketHandler) HandleAwapMessage(session *dara.WebSocketSessionInfo, message *AwapMessage) error {
	// Default implementation returns ErrUseRawMessage to indicate HandleRawMessage should be used
	// If user overrides this method, they should return nil or their own error (not ErrUseRawMessage)
	return ErrUseRawMessage
}

type AwapMessageOption func(*AwapMessage)

func WithType(msgType AwapMessageType) AwapMessageOption {
	return func(m *AwapMessage) {
		m.Type = msgType
	}
}

func WithID(id string) AwapMessageOption {
	return func(m *AwapMessage) {
		m.ID = id
	}
}

func WithHeader(key, value string) AwapMessageOption {
	return func(m *AwapMessage) {
		if m.Headers == nil {
			m.Headers = make(map[string]string)
		}
		m.Headers[key] = value
	}
}

func NewAwapMessage(payload interface{}, opts ...AwapMessageOption) *AwapMessage {
	m := &AwapMessage{
		Type:    AwapMessageType("UpstreamTextEvent"), // Default type
		ID:      generateMessageId(),                  // Default ID (generated later if empty, but here we can generate it)
		Payload: payload,
		Headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func BuildAwapMessageText(message *AwapMessage) (string, error) {
	if message == nil {
		return "", fmt.Errorf("message cannot be nil")
	}
	now := time.Now()
	var headerBuilder strings.Builder

	headerBuilder.WriteString(fmt.Sprintf("type:%s\n", string(message.Type)))
	headerBuilder.WriteString(fmt.Sprintf("timestamp:%d\n", now.Unix()*1000+int64(now.Nanosecond())/1e6))

	if message.ID != "" {
		headerBuilder.WriteString(fmt.Sprintf("id:%s\n", message.ID))
	}

	// Auto-add ack:required for AckRequiredTextEvent
	if message.Type == "AckRequiredTextEvent" {
		headerBuilder.WriteString("ack:required\n")
	}

	if message.Headers != nil {
		for key, value := range message.Headers {
			headerBuilder.WriteString(fmt.Sprintf("%s:%s\n", key, value))
		}
	}

	// Add empty line to separate headers and payload
	headerBuilder.WriteString("\n")

	// Serialize payload to JSON
	var payloadJSON []byte
	var err error
	if message.Payload != nil {
		payloadJSON, err = json.Marshal(message.Payload)
		if err != nil {
			return "", fmt.Errorf("failed to marshal AWAP payload: %w", err)
		}
	} else {
		payloadJSON = []byte("{}")
	}

	return headerBuilder.String() + string(payloadJSON), nil
}

// BuildAwapMessageBinary constructs a binary AWAP message
// Format: [header bytes (text converted to binary)] + [\n\n separator] + [binary body]
func BuildAwapMessageBinary(message *AwapMessage) ([]byte, error) {
	if message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}
	now := time.Now()
	var headerBuilder strings.Builder

	// Build header part (same format as text message)
	headerBuilder.WriteString(fmt.Sprintf("type:%s\n", string(message.Type)))
	headerBuilder.WriteString(fmt.Sprintf("timestamp:%d\n", now.Unix()*1000+int64(now.Nanosecond())/1e6))

	if message.ID != "" {
		headerBuilder.WriteString(fmt.Sprintf("id:%s\n", message.ID))
	}

	if message.Headers != nil {
		for key, value := range message.Headers {
			headerBuilder.WriteString(fmt.Sprintf("%s:%s\n", key, value))
		}
	}

	// Add empty line to separate headers and payload
	headerBuilder.WriteString("\n")

	headerBytes := []byte(headerBuilder.String())

	var bodyBytes []byte
	if message.Payload != nil {
		if payloadBytes, ok := message.Payload.([]byte); ok {
			bodyBytes = payloadBytes
		} else {
			return nil, fmt.Errorf("payload for binary AWAP message must be []byte, got %T", message.Payload)
		}
	} else {
		bodyBytes = []byte{}
	}

	// Concatenate header bytes (including separator) + binary body
	result := make([]byte, 0, len(headerBytes)+len(bodyBytes))
	result = append(result, headerBytes...)
	result = append(result, bodyBytes...)

	return result, nil
}

// ParseAwapMessage parses a WebSocket message as AWAP format
// AWAP protocol uses frame format: text headers + JSON payload
// Format: "type:request\ntimestamp:1234567890\nid:msg-001\n\n{JSON payload} or [header bytes (text converted to binary)] + [\n\n separator] + [binary body]"
func ParseAwapMessage(message *dara.WebSocketMessage) (*AwapMessage, error) {
	data := message.Payload

	// Check if it's awap format (has \n\n separator)
	headerEndIndex := -1
	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\n' && data[i+1] == '\n' {
			headerEndIndex = i
			break
		}
	}

	awapMsg := &AwapMessage{
		Headers: make(map[string]string),
	}

	if headerEndIndex == -1 {
		return nil, fmt.Errorf("failed to parse AWAP message: no \n\n separator found")
	}

	// Frame format: parse headers and payload separately
	headerBytes := data[:headerEndIndex]
	payloadBytes := data[headerEndIndex+2:] // Skip \n\n

	headerStr := string(headerBytes)
	headerLines := strings.Split(headerStr, "\n")
	for _, line := range headerLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		colonIndex := strings.Index(line, ":")
		if colonIndex > 0 {
			key := strings.TrimSpace(line[:colonIndex])
			value := strings.TrimSpace(line[colonIndex+1:])
			// Extract AWAP-required fields
			switch key {
			case "type":
				awapMsg.Type = AwapMessageType(value)
			case "id":
				awapMsg.ID = value
			default:
				// Only non-AWAP-required fields go into Headers map
				awapMsg.Headers[key] = value
			}
		}
	}

	if len(payloadBytes) > 0 {
		var payload interface{}
		if err := json.Unmarshal(payloadBytes, &payload); err != nil {
			awapMsg.Payload = payloadBytes
		} else {
			awapMsg.Payload = payload
		}
	}

	if message.Type == dara.WebSocketMessageTypeBinary {
		awapMsg.Format = AwapMessageFormatBinary
	} else {
		awapMsg.Format = AwapMessageFormatText
	}

	return awapMsg, nil
}
