package types

type Balance struct {
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Locked         string `json:"locked"`
	AvgKrwBuyPrice string `json:"avg_krw_buy_price"`
	Modified       bool   `json:"modified"`
}
