package inverter

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/simonvetter/modbus"
)

const (
	MODBUS_REGISTER_EXPORTLIMITPOWERRATE uint16 = 123

	EXPORT_ENABLED  uint16 = 1000
	EXPORT_DISABLED uint16 = 0
)

type InverterGrowattHybrid struct {
	client *modbus.ModbusClient
}

func NewInverterGrowattHybrid(url string) *InverterGrowattHybrid {

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     url,
		Timeout: 1 * time.Second,
	})

	if err != nil {
		slog.Error("Failed to create ModbusClient", "error", err.Error())
		return nil
	}

	return &InverterGrowattHybrid{
		client: client,
	}

}

func (i *InverterGrowattHybrid) IsExportEnabled() (bool, error) {
	err := i.client.Open()
	if err != nil {
		return true, fmt.Errorf("client.Open error: %s", err.Error())
	}
	defer i.client.Close()

	read, err := i.client.ReadRegister(MODBUS_REGISTER_EXPORTLIMITPOWERRATE, modbus.HOLDING_REGISTER)
	if err != nil {
		return true, err
	}

	switch read {
	case EXPORT_ENABLED:
		return true, nil
	case EXPORT_DISABLED:
		return false, nil
	default:
		return true, fmt.Errorf("failed to determine export status")
	}

}

func (i *InverterGrowattHybrid) ExportEnable() error {
	err := i.client.Open()
	if err != nil {
		return fmt.Errorf("client.Open error: %s", err.Error())
	}
	defer i.client.Close()

	err = i.client.WriteRegister(MODBUS_REGISTER_EXPORTLIMITPOWERRATE, EXPORT_ENABLED)
	if err != nil {
		return fmt.Errorf("client.WriteRegister error: %s", err.Error())
	}

	return nil
}

func (i *InverterGrowattHybrid) ExportDisable() error {
	err := i.client.Open()
	if err != nil {
		return fmt.Errorf("client.Open error: %s", err.Error())
	}
	defer i.client.Close()

	err = i.client.WriteRegister(MODBUS_REGISTER_EXPORTLIMITPOWERRATE, EXPORT_DISABLED)
	if err != nil {
		return fmt.Errorf("client.WriteRegister error: %s", err.Error())
	}

	return nil
}
