package utils

import (
	"errors"
	"fmt"
	//"time"
	//"os/exec"
)

var ErrAlreadyRunning = errors.New("Server is already running")
var ErrNotRunning = errors.New("Server is not running yet")

func StartServer(server_name string) (error) {
	server, err := GetAllServers().GetServer(server_name)
	if err != nil {
		return err
	}
	processes := GetAllProcesses()
	for i := 0; i < len(processes.Processes); i++ {
		if processes.Processes[i].Name == server_name {
			return ErrAlreadyRunning
		}
	}

	// start the process here
	// cmd := exec.Command("prog")
	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	//processes.SaveProcess(Process{Name: server_name, Pid: cmd.Process, Started_at: time.Now().Unix()})
	fmt.Println(server.Name)

	return nil
}

func StopServer(server_name string) (error) {
	_, err := GetAllServers().GetServer(server_name)
	if err != nil {
		return err
	}
	pid := -1
	processes := GetAllProcesses()
	for i := 0; i < len(processes.Processes); i++ {
		if processes.Processes[i].Name == server_name {
			pid = processes.Processes[i].Pid
		}
	}
	if pid == -1 {
		return ErrNotRunning
	}

	// Stop process here

	return nil
}

func RestartServer(server_name string) (error) {
	if err := StopServer(server_name); err != nil {
		if err == ErrNotRunning {
			fmt.Println("Find way to tell this to user")
			//return nil
		} else {
			return err
		}
	}
	if err := StartServer(server_name); err != nil {
		return err
	}

	return nil
}
