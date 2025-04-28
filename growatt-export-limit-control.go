package main

import (
	"log/slog"

	"github.com/euphdk/growatt-export-limit-control/energidataservice"
	"github.com/euphdk/growatt-export-limit-control/inverter"
)

func main() {

	priceLimit := 100.0

	eds := energidataservice.NewEnergiDataservice()

	currentPrice, err := eds.CurrentElspotPrice()
	if err != nil {
		panic(err)
	}

	var priceBelowLimit bool
	if currentPrice < priceLimit {
		priceBelowLimit = true
	} else {
		priceBelowLimit = false
	}

	inverter := inverter.NewInverterGrowattHybrid("tcp://192.168.255.44:5021")
	exportEnabled, err := inverter.IsExportEnabled()
	if err != nil {
		panic(err)
	}

	if exportEnabled && priceBelowLimit {
		slog.Info(
			"Export enabled, but price below priceLimit: Disabling",
			"exportEnabled", exportEnabled,
			"priceLimit", priceBelowLimit,
			"currentPrice", currentPrice,
		)
		err := inverter.ExportDisable()
		if err != nil {
			panic(err)
		}
	} else if !exportEnabled && !priceBelowLimit {
		slog.Info(
			"Export disabled, but price above priceLimit: Enabling",
			"exportEnabled", exportEnabled,
			"priceLimit", priceBelowLimit,
			"currentPrice", currentPrice,
		)
		err := inverter.ExportEnable()
		if err != nil {
			panic(err)
		}
	} else {
		slog.Info(
			"Not changing anything",
			"exportEnabled", exportEnabled,
			"priceLimit", priceBelowLimit,
			"currentPrice", currentPrice,
		)
	}
}
