package BluetoothManager

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type BluetoothManager struct{}

func Disconnect(deviceAddress string) error {
	commands := []string{
		fmt.Sprintf("disconnect %s\n", deviceAddress),
		"exit\n",
	}
	err := executeBluetoothCommands(commands)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func Connect(deviceAddress string) error {
	commands := []string{
		fmt.Sprintf("connect %s\n", deviceAddress),
		"exit\n",
	}
	err := executeBluetoothCommands(commands)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func executeBluetoothCommands(commands []string) error {
	cmd := exec.Command("bluetoothctl")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.New(err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.New(err.Error())
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
		return errors.New(err.Error())
	}

	for _, command := range commands {
		_, err = io.WriteString(stdin, command)
		if err != nil {
			return errors.New(err.Error())
		}
	}

	if err := stdin.Close(); err != nil {
		return errors.New(err.Error())
	}

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}