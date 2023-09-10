package principles

import "sort"

type Result struct {
	Ticker           string  `json:"ticker"`
	ActualPrice      float64 `json:"actualPrice"`
	GrahanFairPrice  float64 `json:"grahanFairPrice"`
	GrahanSafeMargin float64 `json:"grahanSafeMargin"`
	BazinFairPrice   float64 `json:"bazinFairPrice"`
	BazinSafeMargin  float64 `json:"bazinSafeMargin"`
}

func GetStocks(tickers []string, dividendYield, dividendHistory float64) ([]Result, error) {
	var results []Result

	for _, ticker := range tickers {
		grahan := GetGrahan(ticker)
		bazin := GetBazin(ticker, grahan.ActualPrice, dividendYield/100, dividendHistory)

		results = append(results, Result{
			Ticker:           ticker,
			ActualPrice:      grahan.ActualPrice,
			GrahanFairPrice:  grahan.FairPrice,
			GrahanSafeMargin: grahan.SafeMargin,
			BazinFairPrice:   bazin.FairPrice,
			BazinSafeMargin:  bazin.SafeMargin,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].BazinSafeMargin == results[j].BazinSafeMargin {
			return results[i].GrahanSafeMargin > results[j].GrahanSafeMargin
		}
		return results[i].BazinSafeMargin > results[j].BazinSafeMargin
	})

	return results, nil
}
