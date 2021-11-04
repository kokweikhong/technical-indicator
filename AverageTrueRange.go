package indicator

import (
	"math"
	"strings"
)

// AverageTrueRange for RMA(Wilder’s Smoothing Method)
// average true range (ATR) is a technical analysis indicator,
// introduced by market technician J. Welles Wilder Jr. in his book New Concepts
// in Technical Trading Systems, that measures market volatility by decomposing
// the entire range of an asset price for that period.
func AverageTrueRange(ohlc []*OHLC, instrument string, period int) []*OHLC {
	var TR, ATR string = "TR", "ATR"
	var decimal float64 = 10000
	// Check the instrument whether is contains "JPY", JPY decimal places is different with other.
	if strings.Contains(strings.ToLower(instrument), "jpy") {
		decimal = 1000
	}
	for i := 0; i < len(ohlc); i++ {
		// Initialized indicator map if not exist.
		if ohlc[i].Indicator == nil {
			ohlc[i].Indicator = make(map[string]float64)
		}
		if i == 0 { // First index calculations.
			ohlc[i].Indicator[TR] = ohlc[i].High - ohlc[i].Low
			continue
		}
		ohlc[i].Indicator[TR] = trueRange(ohlc[i], ohlc[i-1])
		if i >= period {
			if i == period { // First ATR value calculations (average of true range value).
				var sumTR float64
				for _, v := range ohlc[:i] {
					sumTR += v.Indicator[TR]
				}
				ohlc[i].Indicator[ATR] = sumTR / float64(period)
			} else {
				ohlc[i].Indicator[ATR] = (ohlc[i-1].Indicator[ATR]*float64(period-1) + ohlc[i].Indicator[TR]) / float64(period)
			}
		}
		ohlc[i].Indicator[ATR] = math.Round(ohlc[i].Indicator[ATR]*decimal) / decimal
	}
	for i := 0; i < len(ohlc); i++ {
		delete(ohlc[i].Indicator, TR)
	}
	return ohlc
}

// true range formula: TR=Max[(H − L),Abs(H − Cprevious),Abs(L − Cprevious)]
func trueRange(current, previous *OHLC) float64 {
	var first float64 = current.High - current.Low
	var second float64 = math.Abs(current.High - previous.Close)
	var third float64 = math.Abs(current.Low - previous.Close)
	list := []float64{first, second, third}
	var max float64 = first
	for _, m := range list {
		if m >= max {
			max = m
		}
	}
	return max
}
