package get

import (
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"testing"
)

func TestGetConversation(t *testing.T) {
	res, err := GetConversation("8b2aa640-bc40-4d8a-8b44-5a99360002d4", chatgpt.Token, chatgpt.Cookie)
	if err != nil {
		println(err.Error())
	} else {
		println(res.Title)
	}
	//jsonm()
}
