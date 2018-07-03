package types

type OrderChance struct {
	BeeFee string       `json:"bee_fee"`
	AskFee string       `json:"ask_fee"`
	Market ChanceMarket `json:"market"`
}

type ChanceMarket struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OrderTypes []string `json:"order_types"`
	OrderSide  []string `json:"order_side"`
	Bid        Bid      `json:"bid"`
	Ask        Bid      `json:"ask"`
	MaxTotal   string   `json:"max_total"`
	State      string   `json:"state"`
	BidAccount Account  `json:"bid_account"`
	AskAccount Account  `json:"ask_account"`
}

type Bid struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  int64  `json:"min_total"`
}

type Account struct {
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Locked         string `json:"locked"`
	AvgKrwBuyPrice string `json:"avg_krw_buy_price"`
	Modified       bool   `json:"modified"`
}
