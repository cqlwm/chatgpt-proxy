package list

import (
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"testing"
)

func TestListConversations(t *testing.T) {
	//func ListConversations(offset int, limit int, token string, cookie string) (ListConversationsResponse, error)
	response, err := ListConversations(0, 20, chatgpt.Token, chatgpt.Cookie)
	if err != nil {
		println(err.Error())
	}
	println(response.Items)
	//t2()
}
