package main

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

const (
	MODBUS_REGISTER_EXPORTLIMITPOWERRATE uint16 = 123

	EXPORT_ENABLED  uint16 = 1000
	EXPORT_DISABLED uint16 = 0
)

var client *modbus.ModbusClient
var err error

func main() {
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://192.168.255.44:5021",
		Timeout: 1 * time.Second,
	})

	if err != nil {
		panic("NewClient error: " + err.Error())
	}

	err = client.Open()
	if err != nil {
		panic("client.Open error: " + err.Error())
	}
	defer client.Close()

	err = client.WriteRegister(MODBUS_REGISTER_EXPORTLIMITPOWERRATE, EXPORT_ENABLED)
	if err != nil {
		panic("client.WriteRegister error: " + err.Error())
	}

	fmt.Printf("%#v\n", isExportEnabled())

}

func isExportEnabled() bool {
	read, err := client.ReadRegister(MODBUS_REGISTER_EXPORTLIMITPOWERRATE, modbus.HOLDING_REGISTER)
	if err != nil {
		panic("client.ReadRegister error: " + err.Error())
	}

	switch read {
	case EXPORT_ENABLED:
		return true
	case EXPORT_DISABLED:
		return false
	default:
		panic("Failed to determine export status")
	}

}
