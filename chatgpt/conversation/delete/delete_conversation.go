package delete

import (
	"bytes"
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"net/http"
)

func DeleteConversation(conversationID string, token string, cookie string) error {
	url := chatgpt.ProxyHost + "/backend-api/conversation/" + conversationID
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(`{"is_visible":false}`)))
	if err != nil {
		return err
	}
	chatgpt.UniversalHeader(&req.Header, cookie, token, chatgpt.ApplicationJson)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
