package types

type TransactionResult struct {
	// 입출금 종류
	Type string `json:"type"`
	// 출금의 고유 아이디
	Uuid string `json:"uuid"`
	// 화폐를 의미하는 영문 대문자 코드
	Currency string `json:"currency"`
	// 출금의 트랜잭션 아이디
	Txid string `json:"txid"`
	// 출금 상태
	State string `json:"state"`
	// 출금 생성 시간
	CreatedAt string `json:"created_at"`
	// 출금 완료 시간
	DoneAt string `json:"done_at"`
	// 출금 금액/수량
	Amount string `json:"amount"`
	// 출금 수수료
	Fee string `json:"fee"`
	// 원화 환산 가격
	KrwAmount string `json:"krw_amount"`
	// 출금 유형
	TransactionType string `json:"transaction_type"`
}
