package main

import (
	"encoding/json"
	"github.com/cqlwm/chatgpt-proxy/chatgpt/conversation/create"
	"io"
	"net/http"
)

func main() {
	handleFunc(http.MethodPost, "/conversation/create", rewriteHttp)
	_ = http.ListenAndServe(":80", nil)
}

func rewriteHttp(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("Gpt-Cookie")
	token := r.Header.Get("Gpt_Token")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("read body error"))
		return
	}

	createConversationRequest := create.CreateConversationRequest{}
	err = json.Unmarshal(bytes, &createConversationRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unmarshal error"))
		return
	}

	conversation, err := create.CreateConversation(createConversationRequest, token, cookie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("create conversation error"))
		return
	}

	resultBytes, err := json.Marshal(conversation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("conversation response marshal error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultBytes)
}

func handleFunc(method string, path string, f http.HandlerFunc) {
	cors := func(w http.ResponseWriter, r *http.Request) {
		rOrigin := r.Header.Get("Origin")
		if rOrigin != "" {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,Hf-Access-Key,Re-Domain") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                                          //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH")
		w.Header().Set("content-type", "application/json;charset=UTF-8")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		f(w, r)
	}
	http.HandleFunc(path, cors)
}
