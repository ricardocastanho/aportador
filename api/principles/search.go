package principles

import "sort"

type Result struct {
	Ticker string `json:"ticker"`
	Price  string `json:"price"`
	Shares string `json:"shares"`
	Profit string `json:"profit"`
	Payout string `json:"payout"`
}

func GetStocks(tickers []string) ([]Result, error) {
	var results []Result

	for _, ticker := range tickers {
		bazin := GetBazin(ticker)

		results = append(results, Result{
			Ticker: ticker,
			Price:  bazin.Price,
			Shares: bazin.Shares,
			Profit: bazin.Profit,
			Payout: bazin.Payout,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Ticker > results[j].Ticker
	})

	return results, nil
}
