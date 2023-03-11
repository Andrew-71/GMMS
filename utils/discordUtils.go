package utils

import (
	"bytes"
	"io"
	"net/http"
	"fmt"
)


func SendWebhookJSON(server_name string, json_message []byte) {
	server, err := GetAllServers().GetServer(server_name)
	if err != nil {
		panic(err)
	}
	link := server.Discord.WebhookLink

	resp, err := http.Post(link, "encoding/json", bytes.NewBuffer(json_message))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}