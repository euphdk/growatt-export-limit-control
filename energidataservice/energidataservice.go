package energidataservice

import (
	"net/http"
	"time"
)

type EnergiDataservice struct {
	httpClient *http.Client
}

func NewEnergiDataservice() *EnergiDataservice {
	return &EnergiDataservice{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}
