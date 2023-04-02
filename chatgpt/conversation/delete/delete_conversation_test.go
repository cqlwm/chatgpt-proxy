package delete

import (
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"testing"
)

func TestDeleteConversation(t *testing.T) {
	err := DeleteConversation("ccc4e6db-ba9f-4113-97c9-dc2945d6e4ca", chatgpt.Token, chatgpt.Cookie)
	if err != nil {
		panic(err.Error())
	}
}
