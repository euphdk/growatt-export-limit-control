package main

import (
	"fmt"
	"github.com/euphdk/growatt-export-limit-control/energidataservice"
	"github.com/euphdk/growatt-export-limit-control/inverter"
)

func main() {
	eds := energidataservice.NewEnergiDataservice()

	var negativePrice bool
	currentPrice, err := eds.CurrentElspotPrice()
	if err != nil {
		panic(err)
	}

	if currentPrice < 0 {
		negativePrice = true
	} else {
		negativePrice = false
	}

	inverter := inverter.NewInverterGrowattHybrid("tcp://192.168.255.44:5021")
	exportEnabled, err := inverter.IsExportEnabled()
	if err != nil {
		panic(err)
	}

	if exportEnabled && negativePrice {
		fmt.Println("Export enabled, but negative price: Disabling")
		err := inverter.ExportDisable()
		if err != nil {
			panic(err)
		}
	} else if !exportEnabled && !negativePrice {
		fmt.Println("Export disabled, but positive price: Enabling")
		err := inverter.ExportEnable()
		if err != nil {
			panic(err)
		}
	}
}
