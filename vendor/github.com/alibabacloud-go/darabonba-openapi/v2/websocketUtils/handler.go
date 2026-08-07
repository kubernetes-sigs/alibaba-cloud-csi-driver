package websocketUtils

import (
	"fmt"
	"strings"

	"github.com/alibabacloud-go/tea/dara"
)

const (
	SubProtocolAWAP    = "awap"
	SubProtocolGeneral = "general"
)

type StreamHandler struct {
	UserHandler dara.WebSocketHandler
	SubProtocol string
	Client      *WebSocketClient
}

// ErrUseRawMessage is a sentinel error that indicates HandleRawMessage should be used instead
var ErrUseRawMessage = fmt.Errorf("use HandleRawMessage")

func (h *StreamHandler) AfterConnectionEstablished(session *dara.WebSocketSessionInfo) error {
	return h.UserHandler.AfterConnectionEstablished(session)
}

func (h *StreamHandler) HandleError(session *dara.WebSocketSessionInfo, err error) error {
	return h.UserHandler.HandleError(session, err)
}

func (h *StreamHandler) AfterConnectionClosed(session *dara.WebSocketSessionInfo, code int, reason string) error {
	return h.UserHandler.AfterConnectionClosed(session, code, reason)
}

func (h *StreamHandler) HandleRawMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	subProtocol := strings.ToLower(h.SubProtocol)

	switch subProtocol {
	case SubProtocolAWAP:
		return h.processAwapMessage(session, message)
	case SubProtocolGeneral:
		return h.processGeneralMessage(session, message)
	default:
		return h.UserHandler.HandleRawMessage(session, message)
	}
}

func (h *StreamHandler) processAwapMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	awapMsg, err := ParseAwapMessage(message)
	if err != nil {
		return fmt.Errorf("failed to parse AWAP message: %w", err)
	}

	if awapMsg.Type == "RECONNECT" {
		if h.Client != nil {
			go h.Client.ReconnectGracefully()
		}
		return nil
	}

	var ackID string
	if awapMsg.Headers != nil {
		ackID = awapMsg.Headers["ack-id"]
	}

	if ackID != "" && h.Client != nil {
		if h.Client.completeRequest(ackID, awapMsg) {
			return nil
		}
	}

	if awapHandler, ok := h.UserHandler.(AwapWebSocketHandler); ok {
		err = awapHandler.HandleAwapMessage(session, awapMsg)
		if err == ErrUseRawMessage {
			if rawErr := awapHandler.HandleRawMessage(session, message); rawErr != nil {
				awapHandler.HandleError(session, rawErr)
			}
		} else if err != nil {
			awapHandler.HandleError(session, err)
		}
	} else {
		return h.UserHandler.HandleRawMessage(session, message)
	}
	return nil
}

func (h *StreamHandler) processGeneralMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	generalHandler, ok := h.UserHandler.(GeneralWebSocketHandler)
	if !ok {
		return h.UserHandler.HandleRawMessage(session, message)
	}

	generalMsg, err := ParseGeneralMessage(message)
	if err != nil {
		return fmt.Errorf("failed to parse General message: %w", err)
	}

	err = generalHandler.HandleGeneralMessage(session, generalMsg)
	if err == ErrUseRawMessage {
		if rawErr := generalHandler.HandleRawMessage(session, message); rawErr != nil {
			generalHandler.HandleError(session, rawErr)
		}
	} else if err != nil {
		generalHandler.HandleError(session, err)
	}
	return nil
}
