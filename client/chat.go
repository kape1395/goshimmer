package client

import (
	"net/http"

	"github.com/iotaledger/goshimmer/plugins/chat"
)

const (
	routeChat = "chat"
)

// ChatMsg sends the given data (payload) by creating a message in the backend.
func (api *GoShimmerAPI) SendChatMessage(request *chat.Request) error {
	res := &chat.Response{}
	if err := api.do(http.MethodPost, routeChat,
		request, res); err != nil {
		return err
	}

	return nil
}
