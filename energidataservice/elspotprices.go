package energidataservice

import (
	"fmt"
	"io"
	"time"

	"github.com/tidwall/gjson"
)

const (
	ELSPOTPRICES_URI = "https://api.energidataservice.dk/dataset/Elspotprices?offset=0&start=%s&end=%s&filter={\"PriceArea\":[\"%s\"]}&timezone=dk&limit=1"
	REGION           = "dk1"
	TIMEFORMAT       = "2006-01-02T15:04" // RFC3339 short
)

func (e *EnergiDataservice) CurrentElspotPrice() (float64, error) {

	ts := time.Now().Truncate(time.Hour)
	uri := fmt.Sprintf(ELSPOTPRICES_URI,
		ts.Format(TIMEFORMAT),
		ts.Add(1*time.Hour).Format(TIMEFORMAT),
		REGION)

	r, err := e.httpClient.Get(uri)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, err
	}

	records := gjson.GetBytes(body, "records")

	for _, record := range records.Array() {
		return record.Get("SpotPriceDKK").Float(), nil
	}

	return 0, fmt.Errorf("no spotprice found")

}
