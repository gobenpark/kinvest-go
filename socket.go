package main

import "time"

type Sign int

const (
	None Sign = iota
	Max
	Rise
	Horizontal
	Low
	Drop
)

type SocketBody struct {
	Header struct {
		ApprovalKey string `json:"approval_key"`
		Custtype    string `json:"custtype"`
		TrType      string `json:"tr_type"`
		ContentType string `json:"content-type"`
	} `json:"header"`
	Body struct {
		Input struct {
			TrId  string `json:"tr_id"`
			TrKey string `json:"tr_key"`
		} `json:"input"`
	} `json:"body"`
}

type ResponseBody struct {
	Code string
	// 채결시간 hh:mm:ss
	ContractHour time.Time
	Price        float64
	// 전일 대비 부호
	CompareSign Sign
	// 전일대비
	CompareDay int
	// 전일 대비율
	CompareRate float32
	//가중 평균 주식 가격
	WeightAveragePrice float32
	// 시가
	Open int
	// 고가
	High int
	// 저가
	Low int
	// 매도호가
	AskPrice int
	// 매수호가
	BisPrice int
	// 체결 거래량
	ContractVolume float32
	// 누적 거래량
	AccumulateVolume float32
	// 누적 거래대금
	AccumulateTransactionMoney float32
	// 매도 체결 건수
	AskCount float32
	// 매수 체결 건수
	BidCount float32
	// 순매수 체결 건수
	PureBidCount float32
	// 체결강도
	VolumePower float32
	//총 매도 수량
	TotalAskCounts float32
	// 총 매수 수량
	TotalBidCounts float32
	// 체결 구분
	ContractDivide string
	// 매수 비율
	BidRate float32
	// 전일 거래량 대비 등락율
	PredayVolumeCompareRate float32
	// 시가 시간
	OpenningTime string
	// 시가 대비 구분
	OpenComapreSign Sign
	// 시가대비
	OpenComapre string
	// 최고가 시간
	HighTime string
	// 고가대비 구분
	HighVolumePerSign     Sign
	HighCompare           float32
	LowTime               float32
	LowVolumePerSign      Sign
	LowComapre            float32
	BusinessDate          time.Time
	NewMarketOpCode       string
	TransactionSuspension bool
	RemainAsk
}
