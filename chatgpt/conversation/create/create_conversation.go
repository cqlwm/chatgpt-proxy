package create

import (
	"bytes"
	"encoding/json"
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"io"
	"net/http"
	"strings"
)

type PromptMessage struct {
	ID      string  `json:"id"`
	Author  Author  `json:"author"`
	Role    string  `json:"role"`
	Content Content `json:"content"`
}

type CreateConversationRequest struct {
	Action          string          `json:"action"`
	Messages        []PromptMessage `json:"messages"`
	ParentMessageID string          `json:"parent_message_id"`
	Model           string          `json:"model"`
}

type AnswerMessage struct {
	ID         string   `json:"id"`
	Author     Author   `json:"author"`
	CreateTime float64  `json:"create_time"`
	UpdateTime float64  `json:"update_time"`
	Content    Content  `json:"content"`
	EndTurn    bool     `json:"end_turn"`
	Weight     float64  `json:"weight"`
	Metadata   Metadata `json:"metadata"`
	Recipient  string   `json:"recipient"`
}

type Author struct {
	Role     string   `json:"role"`
	Name     string   `json:"name"`
	Metadata Metadata `json:"metadata"`
}

type Content struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type Metadata struct {
	MessageType   string        `json:"message_type"`
	ModelSlug     string        `json:"model_slug"`
	FinishDetails FinishDetails `json:"finish_details"`
}

type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}

type CreateConversationResponse struct {
	Message        AnswerMessage `json:"message"`
	ConversationID string        `json:"conversation_id"`
	Error          *string       `json:"error"`
}

func CreateConversation(createConversationRequest CreateConversationRequest, token string, cookie string) (CreateConversationResponse, error) {
	url := chatgpt.ProxyHost + "/backend-api/conversation"
	requestBodyBytes, err := json.Marshal(createConversationRequest)
	if err != nil {
		return CreateConversationResponse{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return CreateConversationResponse{}, err
	}
	chatgpt.UniversalHeader(&req.Header, cookie, token, chatgpt.ApplicationJson)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return CreateConversationResponse{}, err
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateConversationResponse{}, err
	}

	var lastAnswer string
	for _, s := range strings.Split(string(respBodyBytes), "\n") {
		if s != "" && s != "data: [DONE]" && strings.HasPrefix(s, "data: ") {
			lastAnswer = strings.TrimPrefix(s, "data: ")
		}
	}

	var response CreateConversationResponse
	err = json.Unmarshal([]byte(lastAnswer), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
