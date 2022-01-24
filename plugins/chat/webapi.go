package chat

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/iotaledger/goshimmer/packages/chat"
	"github.com/iotaledger/goshimmer/packages/jsonmodels"
)

const (
	maxFromToLength  = 100
	maxMessageLength = 1000
)

func configureWebAPI() {
	deps.Server.POST("chat", SendChatMessage)
}

// SendChatMessage sends a chat message.
func SendChatMessage(c echo.Context) error {
	req := &jsonmodels.ChatRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, jsonmodels.NewErrorResponse(err))
	}

	if len(req.From) > maxFromToLength {
		return c.JSON(http.StatusBadRequest, jsonmodels.ChatResponse{Error: "sender is too long"})
	}
	if len(req.To) > maxFromToLength {
		return c.JSON(http.StatusBadRequest, jsonmodels.ChatResponse{Error: "receiver is too long"})
	}
	if len(req.Message) > maxMessageLength {
		return c.JSON(http.StatusBadRequest, jsonmodels.ChatResponse{Error: "message is too long"})
	}

	chatPayload := chat.NewPayload(req.From, req.To, req.Message)
	msg, err := deps.Tangle.IssuePayload(chatPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, jsonmodels.ChatResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, jsonmodels.ChatResponse{MessageID: msg.ID().Base58()})
}
