package indicator

import (
	"encoding/json"
	"log"
	"time"
)

type OHLC struct {
	Time      time.Time          `json:"time"`
	Open      float64            `json:"open,string"`
	High      float64            `json:"high,string"`
	Low       float64            `json:"low,string"`
	Close     float64            `json:"close,string"`
	Volume    float64            `json:"volume,string"`
	Indicator map[string]float64 `json:"indicator,omitempty"`
}

func NewDataFromJson(jsonString string) []*OHLC {
	var result []*OHLC
	data, err := json.Marshal(jsonString)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &result); err != nil {
		log.Fatal(err)
	}
	return result
}
