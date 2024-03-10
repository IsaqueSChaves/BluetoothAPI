package main

import (
	"os"

	"github.com/isaqueschaves/BluetoothAPI/BluetoothManager"
)

func main() {
	arg := os.Args[1]
	deviceAddress := os.Args[2]
	if(arg == "disconnect"){
		BluetoothManager.Disconnect(deviceAddress)
	} else if(arg == "connect"){
		BluetoothManager.Connect(deviceAddress)
	} else {
		panic("unknown option: " + arg)
	}
}
