package main

import (
	"chatbot"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/chatbot/api/message", chatbot.Chatbot)
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}
