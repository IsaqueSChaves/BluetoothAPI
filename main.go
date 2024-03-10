package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func main() {
	arg := os.Args[1]
	deviceAddress := os.Args[2]
	if(arg == "disconnect"){
		disconnect(deviceAddress)
	} else if(arg == "connect"){
		connect(deviceAddress)
	} else {
		panic("unknown option: " + arg)
	}
}

func disconnect(deviceAddress string){
	commands := []string{
		"agent on\n",
		fmt.Sprintf("disconnect %s\n", deviceAddress),
		"exit\n",
	}
	executeBluetoothCommands(commands)
}
func connect(deviceAddress string){
	commands := []string{
		"agent on\n",
		fmt.Sprintf("connect %s\n", deviceAddress),
		"exit\n",
	}
	executeBluetoothCommands(commands)
}

func executeBluetoothCommands(commands []string){
	cmd := exec.Command("bluetoothctl")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	for _, command := range commands {
		_, err = io.WriteString(stdin, command)
		if err != nil {
			panic(err)
		}
	}

	if err := stdin.Close(); err != nil {
		panic(err)
	}

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}