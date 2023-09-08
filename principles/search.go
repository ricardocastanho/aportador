package principles

import "sort"

type Result struct {
	Ticker           string
	ActualPrice      float64
	GrahanFairPrice  float64
	GrahanSafeMargin float64
	BazinFairPrice   float64
	BazinSafeMargin  float64
}

func GetStocks(tickers []string, dividendYield, dividendYears float64) ([]Result, error) {
	var results []Result

	for _, ticker := range tickers {
		grahan := GetGrahan(ticker)
		bazin := GetBazin(ticker, grahan.ActualPrice, dividendYield/100, dividendYears)

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
