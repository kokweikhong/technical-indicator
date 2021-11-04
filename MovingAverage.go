package indicator

import (
	"fmt"
	"math"
	"strings"
)

func SimpleMovingAverage(ohlc []*OHLC, instrument string, periods ...int) []*OHLC {
	var decimal float64 = 100000
	if strings.Contains(strings.ToLower(instrument), "jpy") {
		decimal = 1000
	}
	for k, v := range ohlc {
		if v.Indicator == nil {
			v.Indicator = make(map[string]float64)
		}
		for i := 0; i < len(periods); i++ {
			sma := fmt.Sprintf("SMA%v", periods[i])
			if k < periods[i]-1 {
				v.Indicator[sma] = 0
				continue
			}
			var sum float64 = 0
			for c := 0; c < periods[i]; c++ {
				sum += ohlc[k-c].Close
			}
			v.Indicator[sma] = sum / float64(periods[i])
			v.Indicator[sma] = math.Round(v.Indicator[sma]*decimal) / decimal
			fmt.Println(v.Time, v.Close, v.Indicator)
		}
	}
	return ohlc
}
