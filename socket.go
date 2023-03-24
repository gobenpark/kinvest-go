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
	ContractHour string
	Price        string
	// 전일 대비 부호
	CompareSign string
	// 전일대비
	CompareDay string
	// 전일 대비율
	CompareRate string
	//가중 평균 주식 가격
	WeightAveragePrice string
	// 시가
	Open string
	// 고가
	High string
	// 저가
	Low string
	// 매도호가
	AskPrice string
	// 매수호가
	BisPrice string
	// 체결 거래량
	ContractVolume string
	// 누적 거래량
	AccumulateVolume string
	// 누적 거래대금
	AccumulateTransactionMoney string
	// 매도 체결 건수
	AskCount string
	// 매수 체결 건수
	BidCount string
	// 순매수 체결 건수
	PureBidCount string
	// 체결강도
	VolumePower string
	//총 매도 수량
	TotalAskCounts string
	// 총 매수 수량
	TotalBidCounts string
	// 체결 구분
	ContractDivide string
	// 매수 비율
	BidRate string
	// 전일 거래량 대비 등락율
	PredayVolumeCompareRate string
	// 시가 시간
	OpenningTime string
	// 시가 대비 구분
	OpenComapreSign string
	// 시가대비
	OpenComapre string
	// 최고가 시간
	HighTime string
	// 고가대비 구분
	HighVolumePerSign string
	// 고가대비
	HighCompare string
	// 최저가 시간
	LowTime string
	// 저가대비구분
	LowVolumePerSign string
	// 저가대비
	LowComapre string
	// 영업 일자
	BusinessDate string
	// 신 장운영 구분 코드
	NewMarketOpCode string
	// 거래정지 여부
	TransactionSuspension string
	// 매도호가 잔량
	RemainAsk string
	// 매수호가 잔량
	RemainBid string
	// 총 매도호가 잔량
	TotalRemainAsk string
	// 총 매수호가 잔량
	TotalRemainBid   string
	VolumeRotateRate string
	// 전일 동시간 누적 거래량
	PreDayTotalVolume string
	// 전일 동시간 누적 거래량 비율
	PreDayTotalVolumeRate string
	// 시간 구분 코드
	HourClockCode string
	// 임의 종료 구분코드
	MarketTermCode string
	// 정적 VI 발동기준가
	VIStandardPrice string
}

func (r *ResponseBody) GetCode() string {
	return r.Code
}

func (r *ResponseBody) GetContractHour() time.Time {
	t, _ := time.Parse("150405", r.ContractHour)
	return t
}
