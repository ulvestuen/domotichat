package main

import (
	"domotichat-api/internal/chat"
	"domotichat-api/internal/config"
	"domotichat-api/internal/status"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.NotFound(writer, request)
		}
		err := json.NewEncoder(writer).Encode(status.GetStatus())
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/chat/completions", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodOptions:
			setCorsHeaders(writer)
			writer.WriteHeader(200)
			break
		case http.MethodPost:
			setCorsHeaders(writer)
			handlePost(writer, request)
			break
		default:
			http.NotFound(writer, request)
			break
		}
	})
	err := http.ListenAndServe(":"+config.Port, mux)
	if err != nil {
		panic("Unable to start server")
	}
}

func setCorsHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
}

func handlePost(writer http.ResponseWriter, request *http.Request) {
	var messages []openai.ChatCompletionMessage
	err := json.NewDecoder(request.Body).Decode(&messages)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(chat.SendChatMessage(messages))
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
