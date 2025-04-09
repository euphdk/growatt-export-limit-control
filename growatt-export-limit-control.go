package main

import (
	"fmt"

	"github.com/euphdk/growatt-export-limit-control/inverter"
)

func main() {
	inverter := inverter.NewInverterGrowattHybrid("tcp://192.168.255.44:5021")

	exportEnabled, err := inverter.IsExportEnabled()
	if err != nil {
		panic(err)
	}
	fmt.Println(exportEnabled)

	if exportEnabled {
		err := inverter.ExportDisable()
		if err != nil {
			panic(err)
		}
	} else {
		err := inverter.ExportEnable()
		if err != nil {
			panic(err)
		}
	}

	exportEnabled, err = inverter.IsExportEnabled()
	if err != nil {
		panic(err)
	}
	fmt.Println(exportEnabled)

}
