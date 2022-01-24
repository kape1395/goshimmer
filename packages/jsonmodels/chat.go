package jsonmodels

// ChatRequest defines the chat message to send.
type ChatRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

// ChatResponse contains the ID of the message sent.
type ChatResponse struct {
	MessageID string `json:"messageID,omitempty"`
	Error     string `json:"error,omitempty"`
}
