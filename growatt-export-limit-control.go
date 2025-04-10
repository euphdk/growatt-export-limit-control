package main

import (
	"log/slog"

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
		slog.Info("Export enabled, but negative price: Disabling", "exportEnabled", exportEnabled, "negativePrice", negativePrice, "currentPrice", currentPrice)
		err := inverter.ExportDisable()
		if err != nil {
			panic(err)
		}
	} else if !exportEnabled && !negativePrice {
		slog.Info("Export disabled, but positive price: Enabling", "exportEnabled", exportEnabled, "negativePrice", negativePrice, "currentPrice", currentPrice)
		err := inverter.ExportEnable()
		if err != nil {
			panic(err)
		}
	} else {
		slog.Info("Not changing anything", "exportEnabled", exportEnabled, "negativePrice", negativePrice, "currentPrice", currentPrice)
	}
}
