package list

import (
	"encoding/json"
	"fmt"
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"io"
	"net/http"
)

type ListConversationsResponse struct {
	Items []struct {
		CreateTime string `json:"create_time"`
		ID         string `json:"id"`
		Title      string `json:"title"`
	} `json:"items"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

func ListConversations(offset int, limit int, token string, cookie string) (ListConversationsResponse, error) {
	url := fmt.Sprintf("%s/backend-api/conversations?offset=%d&limit=%d", chatgpt.ProxyHost, offset, limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListConversationsResponse{}, err
	}
	chatgpt.UniversalHeader(&req.Header, cookie, token, chatgpt.ApplicationJson)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ListConversationsResponse{}, err
	}
	defer resp.Body.Close()
	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListConversationsResponse{}, err
	}
	var listConversationsResponse ListConversationsResponse
	err = json.Unmarshal(responseBodyBytes, &listConversationsResponse)
	if err != nil {
		return ListConversationsResponse{}, err
	}
	return listConversationsResponse, nil
}
