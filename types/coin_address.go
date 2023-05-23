package types

type CoinAddress struct {
	// 출금 화폐
	Currency string `json:"currency"`
	// 출금 네트워크
	NetType string `json:"net_type"`
	// 출금 주소
	WithdrawAddress string `json:"withdraw_address"`
	// 2차 출금 주소 (필요한 디지털 자산에 한해서)
	SecondaryAddress string `json:"secondary_address"`
}
