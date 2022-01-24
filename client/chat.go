package client

import (
	"net/http"

	"github.com/iotaledger/goshimmer/packages/jsonmodels"
)

const (
	routeChat = "chat"
)

// ChatMsg sends the given data (payload) by creating a message in the backend.
func (api *GoShimmerAPI) SendChatMessage(request *jsonmodels.ChatRequest) error {
	res := &jsonmodels.ChatResponse{}
	if err := api.do(http.MethodPost, routeChat,
		request, res); err != nil {
		return err
	}

	return nil
}
