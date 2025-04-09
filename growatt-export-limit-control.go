package main

import (
	"github.com/euphdk/growatt-export-limit-control/energidataservice"
	"github.com/euphdk/growatt-export-limit-control/inverter"
	"log/slog"
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
		slog.Info("Export enabled, but negative price: Disabling")
		err := inverter.ExportDisable()
		if err != nil {
			panic(err)
		}
	} else if !exportEnabled && !negativePrice {
		slog.Info("Export disabled, but positive price: Enabling")
		err := inverter.ExportEnable()
		if err != nil {
			panic(err)
		}
	}
}
