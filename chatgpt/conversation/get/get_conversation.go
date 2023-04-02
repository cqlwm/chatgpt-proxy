package get

import (
	"encoding/json"
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"io"
	"net/http"
)

type Message struct {
	ID     string `json:"id"`
	Author struct {
		Role     string   `json:"role"`
		Metadata struct{} `json:"metadata"`
	} `json:"author"`
	CreateTime float64 `json:"create_time"`
	Content    struct {
		ContentType string   `json:"content_type"`
		Parts       []string `json:"parts"`
	} `json:"content"`
	EndTurn  bool    `json:"end_turn"`
	Weight   float64 `json:"weight"`
	Metadata struct {
		Timestamp string `json:"timestamp_"`
	} `json:"metadata"`
	Recipient string `json:"recipient"`
}

type Node struct {
	ID       string   `json:"id"`
	Message  Message  `json:"message"`
	Parent   string   `json:"parent"`
	Children []string `json:"children"`
}

type GetConversationResponse struct {
	Title      string          `json:"title"`
	CreateTime float64         `json:"create_time"`
	Mapping    map[string]Node `json:"mapping"`
	//ModerationResults []interface{}   `json:"moderation_results"`
	CurrentNode string `json:"current_node"`
}

func GetConversation(conversationID string, token string, cookie string) (GetConversationResponse, error) {
	url := chatgpt.ProxyHost + "/backend-api/conversation/" + conversationID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetConversationResponse{}, err
	}
	chatgpt.UniversalHeader(&req.Header, cookie, token, chatgpt.ApplicationJson)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return GetConversationResponse{}, err
	}
	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetConversationResponse{}, err
	}
	var conversationResponse GetConversationResponse
	err = json.Unmarshal(responseBodyBytes, &conversationResponse)
	if err != nil {
		return GetConversationResponse{}, err
	}
	return conversationResponse, nil
}
