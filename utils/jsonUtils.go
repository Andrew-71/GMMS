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
	_ = os.WriteFile("servers.json", file, 0644)
}


// These are for processes.json, and represent current Processes

type Processes struct {
	Processes []Process	`json:"processes"`
}

type Process struct {
	Name	string	`json:"name"`
	Pid	int	`json:"pid"`
	Started_at	int	`jon:"started_at"`
}

func (processes Processes) SaveProcess(process Process) {
	// No need to check the process exists since we shouldn't possibly have a copy
	processes.Processes = append(processes.Processes, process)

	file, _ := json.MarshalIndent(processes, "", " ")
	_ = os.WriteFile("servers.json", file, 0644)
}

// Getters

func GetAllServers() (Servers) {
	jsonFile, err := os.Open("servers.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var servers Servers
	json.Unmarshal(byteValue, &servers)

	return servers
}

func GetAllProcesses() (Processes) {
	jsonFile, err := os.Open("processes.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var processes Processes
	json.Unmarshal(byteValue, &processes)

	return processes
}