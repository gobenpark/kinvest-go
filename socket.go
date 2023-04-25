package kv

import (
	"strconv"
	"time"
)

type Sign int

const (
	None Sign = iota
	Max
	Rise
	Horizontal
	Low
	Drop
)

func ConvertSign(sign string) Sign {
	switch sign {
	case "1":
		return Max
	case "2":
		return Rise
	case "3":
		return Horizontal
	case "4":
		return Low
	case "5":
		return Drop
	default:
		return None
	}
}

type AccessResponse struct {
	Header struct {
		TrId    string `json:"tr_id"`
		TrKey   string `json:"tr_key"`
		Encrypt string `json:"encrypt"`
	} `json:"header"`
	Body struct {
		RtCd   string `json:"rt_cd"`
		MsgCd  string `json:"msg_cd"`
		Msg1   string `json:"msg1"`
		Output struct {
			Iv  string `json:"iv"`
			Key string `json:"key"`
		} `json:"output"`
	} `json:"body"`
}

type RequestBody struct {
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

type RealtimeResponse struct {
	Encrypted  bool
	TRID       string
	DataCounts string
	Datas      []RealtimeData
}

type RealtimeData struct {
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
	// 가중 평균 주식 가격
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
	BidPrice string
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
	// 총 매도 수량
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
	OpenCompareSign string
	// 시가대비
	OpenCompare string
	// 최고가 시간
	HighTime string
	// 고가대비 구분
	HighCompareSign string
	// 고가대비
	HighCompare string
	// 최저가 시간
	LowTime string
	// 저가대비구분
	LowCompareSign string
	// 저가대비
	LowCompare string
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
	TotalRemainBid string
	// 거래량 회전율
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

func (r *RealtimeResponse) GetEncrypted() bool {
	return r.Encrypted
}

func (r *RealtimeResponse) GetTRID() string {
	return r.TRID
}

func (r *RealtimeResponse) GetDataCounts() string {
	return r.DataCounts
}

func (r *RealtimeData) GetCode() string {
	return r.Code
}

func (r *RealtimeData) GetContractHour() time.Time {
	t, _ := time.Parse("150405", r.ContractHour)
	return t
}

func (r *RealtimeData) GetPrice() int {
	p, _ := strconv.Atoi(r.Price)
	return p
}

func (r *RealtimeData) GetCompareSign() Sign {
	return ConvertSign(r.CompareSign)
}

func (r *RealtimeData) GetCompareDay() int {
	d, _ := strconv.Atoi(r.CompareDay)
	return d
}

func (r *RealtimeData) GetCompareRate() float64 {
	f, _ := strconv.ParseFloat(r.CompareRate, 32)
	return f
}

func (r *RealtimeData) GetWeightAveragePrice() float64 {
	f, _ := strconv.ParseFloat(r.WeightAveragePrice, 32)
	return f
}

func (r *RealtimeData) GetOpen() int {
	i, _ := strconv.Atoi(r.Open)
	return i
}

func (r *RealtimeData) GetHigh() int {
	i, _ := strconv.Atoi(r.High)
	return i
}

func (r *RealtimeData) GetLow() int {
	i, _ := strconv.Atoi(r.Low)
	return i
}

func (r *RealtimeData) GetAskPrice() int {
	i, _ := strconv.Atoi(r.AskPrice)
	return i
}

func (r *RealtimeData) GetBidPrice() int {
	i, _ := strconv.Atoi(r.BidPrice)
	return i
}

func (r *RealtimeData) GetContractVolume() int {
	i, _ := strconv.Atoi(r.ContractVolume)
	return i
}

func (r *RealtimeData) GetAccumulateVolume() int {
	i, _ := strconv.Atoi(r.AccumulateVolume)
	return i
}

func (r *RealtimeData) GetAccumulateTransactionMoney() int64 {
	i, _ := strconv.ParseInt(r.AccumulateTransactionMoney, 10, 64)
	return i
}

func (r *RealtimeData) GetAskCount() int {
	i, _ := strconv.Atoi(r.AskCount)
	return i
}

func (r *RealtimeData) GetBidCount() int {
	i, _ := strconv.Atoi(r.BidCount)
	return i
}

func (r *RealtimeData) GetPureBidCount() int {
	i, _ := strconv.Atoi(r.PureBidCount)
	return i
}

func (r *RealtimeData) GetVolumePower() float64 {
	f, _ := strconv.ParseFloat(r.VolumePower, 32)
	return f
}

func (r *RealtimeData) GetTotalAskCounts() int {
	i, _ := strconv.Atoi(r.TotalAskCounts)
	return i
}

func (r *RealtimeData) GetTotalBidCounts() int {
	i, _ := strconv.Atoi(r.TotalBidCounts)
	return i
}

func (r *RealtimeData) GetContractDivide() string {
	return r.ContractDivide
}

func (r *RealtimeData) GetBidRate() float64 {
	f, _ := strconv.ParseFloat(r.BidRate, 32)
	return f
}

func (r *RealtimeData) GetPredayVolumeCompareRate() float64 {
	f, _ := strconv.ParseFloat(r.PredayVolumeCompareRate, 32)
	return f
}

func (r *RealtimeData) GetOpenningTime() time.Time {
	t, _ := time.Parse("150405", r.OpenningTime)
	return t
}

func (r *RealtimeData) GetOpenCompareSign() Sign {
	return ConvertSign(r.OpenCompareSign)
}

func (r *RealtimeData) GetOpenCompare() int {
	i, _ := strconv.Atoi(r.OpenCompare)
	return i
}

func (r *RealtimeData) GetHighTime() time.Time {
	t, _ := time.Parse("150405", r.HighTime)
	return t
}

func (r *RealtimeData) GetHighCompareSign() Sign {
	return ConvertSign(r.HighCompareSign)
}

func (r *RealtimeData) GetHighCompare() int {
	i, _ := strconv.Atoi(r.HighCompare)
	return i
}

func (r *RealtimeData) GetLowTime() time.Time {
	t, _ := time.Parse("150405", r.LowTime)
	return t
}

func (r *RealtimeData) GetLowCompareSign() Sign {
	return ConvertSign(r.LowCompareSign)
}

func (r *RealtimeData) GetLowCompare() int {
	i, _ := strconv.Atoi(r.LowCompare)
	return i
}

func (r *RealtimeData) GetBusinessDate() time.Time {
	t, _ := time.Parse("20060102", r.BusinessDate)
	return t
}

func (r *RealtimeData) GetNewMarketOpCode() {}

func (r *RealtimeData) GetTransactionSuspension() bool {
	if r.TransactionSuspension == "Y" {
		return true
	} else {
		return false
	}
}

func (r *RealtimeData) GetRemainAsk() int {
	i, _ := strconv.Atoi(r.RemainAsk)
	return i
}

func (r *RealtimeData) GetRemainBid() int {
	i, _ := strconv.Atoi(r.RemainBid)
	return i
}

func (r *RealtimeData) GetTotalRemainAsk() int {
	i, _ := strconv.Atoi(r.TotalRemainAsk)
	return i
}

func (r *RealtimeData) GetTotalRemainBid() int {
	i, _ := strconv.Atoi(r.TotalRemainBid)
	return i
}

func (r *RealtimeData) GetVolumeRotateRate() float64 {
	f, _ := strconv.ParseFloat(r.VolumeRotateRate, 16)
	return f
}

func (r *RealtimeData) GetPreDayTotalVolume() int {
	i, _ := strconv.Atoi(r.PreDayTotalVolume)
	return i
}

func (r *RealtimeData) GetPreDayTotalVolumeRate() float64 {
	f, _ := strconv.ParseFloat(r.PreDayTotalVolumeRate, 16)
	return f
}

func (r *RealtimeData) GetHourClockCode() string {
	return r.HourClockCode
}

func (r *RealtimeData) GetMarketTermCode() string {
	return r.MarketTermCode
}

func (r *RealtimeData) GetVIStandardPrice() int {
	i, _ := strconv.Atoi(r.VIStandardPrice)
	return i
}
