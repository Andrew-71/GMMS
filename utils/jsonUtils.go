package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrNotFound = errors.New("item was not found")


type Servers struct {
	Servers	[]Server	`json:"servers"`
}

type Server struct {
	Name	string	`json:"name"`
	JarDir	string	`json:"jar_dir"`
	Args	[]ServerArgs	`json:"args"`
	Discord	DiscordArgs	`json:"discord"`
}

type ServerArgs struct {
	Xmx	string	`json:"xmx"`
	Xms	string	`json:"xms"`
	JavaInstance	string `json:"java_instance"`
}

type DiscordArgs struct {
	Enabled	bool	`json:"ds_enabled"`
	WebhookLink	string	`json:"ds_webhook_link"`
}

func (servers Servers) GetServer(server_name string) (Server, error) {
	for i := 0; i < len(servers.Servers); i++ {
		if servers.Servers[i].Name == server_name {
			return servers.Servers[i], nil
		}
	}
	return Server{}, ErrNotFound
}

func (servers Servers) SaveServer(server Server) {

	// Check if server already exists first (no copies allowed!)
	found := false
	for i := 0; i < len(servers.Servers); i++ {
		if servers.Servers[i].Name == server.Name {
			found = true
			servers.Servers[i] = server
		}
	}
	if !found {
		servers.Servers = append(servers.Servers, server)
	}

	file, _ := json.MarshalIndent(servers, "", " ")
	_ = os.WriteFile("test.json", file, 0644)
	//_ = ioutil.WriteFile("../servers.json", file, 0644)
}


// These are for processes.json, and represent current Processes


// Getters

func GetAllServers() (Servers) {
	jsonFile, err := os.Open("../servers.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened servers.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var servers Servers
	json.Unmarshal(byteValue, &servers)

	return servers
}

func GetAllProcessees() {
	fmt.Println("TODO")
}