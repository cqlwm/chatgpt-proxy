package create

import (
	"encoding/json"
	"github.com/cqlwm/chatgpt-proxy/chatgpt"
	"testing"
)

func TestCreateConversationGetConversation(t *testing.T) {

	reqstr := `{
	    "action": "next",
	    "messages": [
	        {
	            "id": "330216df-6463-4504-9458-8081ab075a50",
	            "author": {
	                "role": "user"
	            },
	            "role": "user",
	            "content": {
	                "content_type": "text",
	                "parts": [
	                    "写一首诗送给中国国足的诗"
	                ]
	            }
	        }
	    ],
	    "parent_message_id": "767a71c9-4717-440f-ac13-6f911f4ffe8e",
	    "model": "text-davinci-002-render-sha"
	}`
	var req CreateConversationRequest
	err := json.Unmarshal([]byte(reqstr), &req)
	if err != nil {
		panic(err.Error())
	}

	resp, err := CreateConversation(req, chatgpt.Token, chatgpt.Cookie)
	if err != nil {
		panic(err.Error())
	} else {
		println(resp.Message.Content.Parts[0])
	}

}
