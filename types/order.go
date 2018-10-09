package types

type Order struct {
	Uuid            string `json:"uuid"`            //주문의 고유 아이디
	Side            string `json:"side"`            // 주문 종류
	OrdType         string `json:"ord_type"`        // 주문 방식
	Price           string `json:"price"`           // 주문 당시 화폐 가격
	AvgPrice        string `json:"avg_price"`       // 체결 가격의 평균가
	State           string `json:"state"`           //주문 상태
	Market          string `json:"market"`          // 마켓의 유일키
	CreatedAt       string `json:"created_at"`      // 주문 생성 시간
	Volume          string `json:"volume"`          // 사용자가 입력한 주문 양
	RemainingVolume string `json:"remainingVolume"` // 체결 후 남은 주문 양
	ReservedFee     string `json:"reserved_fee"`    // 수수료로 예약된 비용
	RemainingFee    string `json:"remaining_fee"`   //남은 수수료
	PaidFee         string `json:"paid_fee"`        // 사용된 수수료
	Locked          string `json:"locked"`          // 거래에 사용중인 비용
	ExecutedVolume  string `json:"executed_volume"` // 체결된 양
	TradeCount      string `json:"trade_count"`     // 해당 주문에 걸린 체결 수
}
