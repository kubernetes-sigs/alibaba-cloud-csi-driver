package websocketUtils

import (
	"encoding/json"

	"github.com/alibabacloud-go/tea/dara"
)

type GeneralMessageFormat string

const (
	GeneralMessageFormatText   GeneralMessageFormat = "text"
	GeneralMessageFormatBinary GeneralMessageFormat = "binary"
)

type GeneralMessage struct {
	Body   interface{}          `json:"body,omitempty"`
	Format GeneralMessageFormat `json:"format,omitempty"`
}

type GeneralWebSocketHandler interface {
	dara.WebSocketHandler

	HandleGeneralMessage(session *dara.WebSocketSessionInfo, message *GeneralMessage) error
}

// AbstractGeneralWebSocketHandler provides default implementations for GeneralWebSocketHandler interface
// It embeds AbstractWebSocketHandler for base WebSocketHandler methods
// Users can embed this struct in their custom General handlers
type AbstractGeneralWebSocketHandler struct {
	dara.AbstractWebSocketHandler
}

func (h *AbstractGeneralWebSocketHandler) HandleGeneralMessage(session *dara.WebSocketSessionInfo, message *GeneralMessage) error {
	// Default implementation returns ErrUseRawMessage to indicate HandleRawMessage should be used
	// If user overrides this method, they should return nil or their own error (not ErrUseRawMessage)
	return ErrUseRawMessage
}

func NewGeneralMessage(body string) *GeneralMessage {
	return &GeneralMessage{
		Body: body,
	}
}

func ParseGeneralMessage(message *dara.WebSocketMessage) (*GeneralMessage, error) {
	if message.Type == dara.WebSocketMessageTypeBinary {
		return &GeneralMessage{
			Body:   message.Payload,
			Format: GeneralMessageFormatBinary,
		}, nil
	}

	var body interface{}
	if err := json.Unmarshal(message.Payload, &body); err != nil {
		return &GeneralMessage{
			Body:   message.Payload,
			Format: GeneralMessageFormatText,
		}, nil
	}
	return &GeneralMessage{
		Body:   body,
		Format: GeneralMessageFormatText,
	}, nil
}

func (m *GeneralMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}
