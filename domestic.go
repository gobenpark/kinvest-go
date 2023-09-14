package kv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/samber/lo"
)

func (k *Kinvest) RealtimeContract(ctx context.Context, approvalKey string, code ...string) (<-chan RealtimeResponse, error) {
	res := make(chan RealtimeResponse, 1)

	codes := lo.Map[string, RequestBody](code, func(item string, index int) RequestBody {
		return RequestBody{
			Header: struct {
				ApprovalKey string `json:"approval_key"`
				Custtype    string `json:"custtype"`
				TrType      string `json:"tr_type"`
				ContentType string `json:"content-type"`
			}{
				ApprovalKey: approvalKey,
				Custtype:    string(k.config.Customer),
				TrType:      "1",
				ContentType: "utf-8",
			},
			Body: struct {
				Input struct {
					TrId  string `json:"tr_id"`
					TrKey string `json:"tr_key"`
				} `json:"input"`
			}{
				Input: struct {
					TrId  string `json:"tr_id"`
					TrKey string `json:"tr_key"`
				}{
					TrId:  "H0STCNT0",
					TrKey: item,
				},
			},
		}
	})

	go func() {
		defer close(res)
		Retry(ctx, 5, func() error {
			c, _, err := websocket.DefaultDialer.DialContext(ctx, "ws://ops.koreainvestment.com:21000/tryitout/H0STCNT0", nil)
			if err != nil {
				return err
			}

			for _, i := range codes {

				if err := c.WriteJSON(i); err != nil {
					return err
				}
			}

			for {
				select {
				case <-ctx.Done():
					break
				default:
					//0|H0STCNT0|001|005930^112616^62700^2^400^0.64^62700.41^62700^63000^62300^62700^62600^125^9694968^607878273600^25807^15918^-9889^93.01^4614191^4291558^1^0.45^63.03^090013^3^0^090646^5^-300^100741^2^400^20230324^20^N^71292^244009^2795345^2044512^0.16^5367918^180.61^0^^62700
					mtype, msg, err := c.ReadMessage()
					if err != nil {
						return err
					}

					switch mtype {
					case websocket.TextMessage:
						if msg[0] == '{' {
							var s AccessResponse
							if err := json.Unmarshal(msg, &s); err != nil {
								fmt.Println("marshal error", err)
							}
							fmt.Println(s.Body.Msg1)
							continue
						}

						ms := strings.Split(string(msg), "|")
						if ms[0] != "0" && ms[0] != "1" {
							return errors.New("could not find encrypt value")
						}

						counts, err := strconv.ParseInt(ms[2], 10, 32)
						if err != nil {
							return errors.New("fail to parse datacounts")
						}

						tbody := strings.Split(ms[3], "^")
						body := RealtimeResponse{}
						body.Encrypted = ms[0] == "1"
						body.TRID = ms[1]
						body.DataCounts = ms[2]

						idx := 0
						b := []RealtimeData{}
						for i := 0; i < int(counts); i++ {
							b = append(b, RealtimeData{
								Code:                       tbody[idx+0],
								ContractHour:               tbody[idx+1],
								Price:                      tbody[idx+2],
								CompareSign:                tbody[idx+3],
								CompareDay:                 tbody[idx+4],
								CompareRate:                tbody[idx+5],
								WeightAveragePrice:         tbody[idx+6],
								Open:                       tbody[idx+7],
								High:                       tbody[idx+8],
								Low:                        tbody[idx+9],
								AskPrice:                   tbody[idx+10],
								BidPrice:                   tbody[idx+11],
								ContractVolume:             tbody[idx+12],
								AccumulateVolume:           tbody[idx+13],
								AccumulateTransactionMoney: tbody[idx+14],
								AskCount:                   tbody[idx+15],
								BidCount:                   tbody[idx+16],
								PureBidCount:               tbody[idx+17],
								VolumePower:                tbody[idx+18],
								TotalAskCounts:             tbody[idx+19],
								TotalBidCounts:             tbody[idx+20],
								ContractDivide:             tbody[idx+21],
								BidRate:                    tbody[idx+22],
								PredayVolumeCompareRate:    tbody[idx+23],
								OpenningTime:               tbody[idx+24],
								OpenCompareSign:            tbody[idx+25],
								OpenCompare:                tbody[idx+26],
								HighTime:                   tbody[idx+27],
								HighCompareSign:            tbody[idx+28],
								HighCompare:                tbody[idx+29],
								LowTime:                    tbody[idx+30],
								LowCompareSign:             tbody[idx+31],
								LowCompare:                 tbody[idx+32],
								BusinessDate:               tbody[idx+33],
								NewMarketOpCode:            tbody[idx+34],
								TransactionSuspension:      tbody[idx+35],
								RemainAsk:                  tbody[idx+36],
								RemainBid:                  tbody[idx+37],
								TotalRemainAsk:             tbody[idx+38],
								TotalRemainBid:             tbody[idx+39],
								VolumeRotateRate:           tbody[idx+40],
								PreDayTotalVolume:          tbody[idx+41],
								PreDayTotalVolumeRate:      tbody[idx+42],
								HourClockCode:              tbody[idx+43],
								MarketTermCode:             tbody[idx+44],
								VIStandardPrice:            tbody[idx+45],
							})
							idx += 46
						}

						body.Datas = b
						res <- body
					}
				}
			}
		})
	}()
	return res, nil
}

// 주식현재가 시세
func (k *Kinvest) CurrentPrice(ctx context.Context, m MarketType, code string) (*CurrentPrice, error) {

	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010100",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-price")
	if err != nil {
		return nil, err
	}
	var d CurrentPrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 주식현재가 체결
func (k *Kinvest) CurrentConclusion(ctx context.Context, m MarketType, code string) (*CurrentConclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010300",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-ccnl")
	if err != nil {
		return nil, err
	}
	var d CurrentConclusion
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 현재가 일자별 일/주/월별 주가를 확인할 수 있으며 최근 30일(주,별)로 제한되어있습니다.
func (k *Kinvest) DailyPrice(ctx context.Context, m MarketType, p Period, code string, adjustPrice bool) (*DailyPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010400",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_PERIOD_DIV_CODE", string(p)).
		SetQueryParam("FID_ORG_ADJ_PRC", YesOrNo("0", "1", adjustPrice)).
		Get("/uapi/domestic-stock/v1/quotations/inquire-daily-price")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d DailyPrice

	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 호가 예상체결 API입니다. 매수 매도 호가를 확인하실 수 있습니다. 실시간 데이터를 원하신다면 웹소켓 API를 활용하세요.
func (k *Kinvest) ExpectAskPrice(ctx context.Context, m MarketType, code string) (*ExpectPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010200",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-asking-price-exp-ccn")
	if err != nil {
		return nil, err
	}
	var d ExpectPrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 투자자 API입니다. 개인, 외국인, 기관 등 투자 정보를 확인할 수 있습니다.
//
// [유의사항]
// - 외국인은 외국인(외국인투자등록 고유번호가 있는 경우)+기타 외국인을 지칭합니다.
// - 당일 데이터는 장 종료 후 제공됩니다.
func (k *Kinvest) Investor(ctx context.Context, m MarketType, code string) (*Investor, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010900",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-investor")
	if err != nil {
		return nil, err
	}
	var d Investor
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식 현재가 회원사 API입니다. 회원사의 투자 정보를 확인할 수 있습니다.
func (k *Kinvest) Member(ctx context.Context, m MarketType, code string) (*Member, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST01010600",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-member")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d Member
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// ELW 현재가 시세 API입니다. ELW 관련 정보를 얻을 수 있습니다.
func (k *Kinvest) CurrentELW(ctx context.Context, m MarketType, code string) (*ELW, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKEW15010000",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-elw-price")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d ELW
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 국내주식 업종기간별시세(일/주/월/년) API입니다.
// 실전계좌/모의계좌의 경우, 한 번의 호출에 최대 50건까지 확인 가능합니다.
func (k *Kinvest) DailyChartPrice(ctx context.Context, start time.Time, end time.Time, m MarketType, p Period, code string, adjustPrice bool) (*DailyChartPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST03010100",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_INPUT_DATE_1", start.Format("20060102")).
		SetQueryParam("FID_INPUT_DATE_2", end.Format("20060102")).
		//D일봉,W:주봉,M월봉,Y년봉
		SetQueryParam("FID_PERIOD_DIV_CODE", string(p)).
		//수정주가 여부 0: 수정주가, 1원주가
		SetQueryParam("FID_ORG_ADJ_PRC", YesOrNo("0", "1", adjustPrice)).
		Get("/uapi/domestic-stock/v1/quotations/inquire-daily-itemchartprice")
	if err != nil {
		return nil, err
	}
	var d DailyChartPrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 당일시간대별체결 API입니다.
func (k *Kinvest) CurrentTimePerConclusion(ctx context.Context, startDate time.Time, m MarketType, code string) (*Conclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHPST01060000",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_INPUT_DATE_1", startDate.Format("20060102")).
		Get("/uapi/domestic-stock/v1/quotations/inquire-time-itemconclusion")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d Conclusion
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 시간외시간별체결 API입니다.
func (k *Kinvest) CurrentOvertimePerConclusion(ctx context.Context, m MarketType, code string) (*Conclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHPST02310000",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_HOUR_CLS_CODE", "1").
		Get("/uapi/domestic-stock/v1/quotations/inquire-time-overtimeconclusion")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d Conclusion
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 시간외일자별주가 API입니다.
func (k *Kinvest) DailyOvertimePerPrice(ctx context.Context, m MarketType, code string) (*OvertimePrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHPST02320000",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-daily-overtimeprice")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d OvertimePrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 주식당일분봉조회 API입니다.
// 실전계좌/모의계좌의 경우, 한 번의 호출에 최대 30건까지 확인 가능합니다.
func (k *Kinvest) CurrentTimeChartPrice(ctx context.Context, m MarketType, code string) (*OvertimePrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHKST03010200",
		"custtype":      string(k.config.Customer),
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", string(m)).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-time-itemchartprice")
	if err != nil {
		return nil, err
	}
	var d OvertimePrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 상품 기본조회
func (k *Kinvest) ItemInfo(ctx context.Context, code string) (*Item, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "CTPF1604R",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": string(k.config.Customer),
	}).
		//'주식(하이닉스) : 000660 (코드 : 300)
		//선물(101S12) : KR4101SC0009 (코드 : 301)
		//미국(AAPL) : AAPL (코드 : 512)'
		SetQueryParam("PDNO", code).
		//'300 주식
		//301 선물옵션
		//302 채권
		//512 미국 나스닥 / 513 미국 뉴욕 / 529 미국 아멕스
		//515 일본
		//501 홍콩 / 543 홍콩CNY / 558 홍콩USD
		//507 베트남 하노이 / 508 베트남 호치민
		//551 중국 상해A / 552 중국 심천A'
		SetQueryParam("PRDT_TYPE_CD", "300").
		Get("/uapi/domestic-stock/v1/quotations/search-info")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d Item
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 휴장일 조회
func (k *Kinvest) HoilydayInfo(ctx context.Context, date time.Time) (*Hoilyday, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "CTCA0903R",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": string(k.config.Customer),
	}).
		SetQueryParam("BASS_DT", date.Format("20060102")).
		SetQueryParam("CTX_AREA_NK", "").
		SetQueryParam("CTX_AREA_FK", "").
		Get("/uapi/domestic-stock/v1/quotations/chk-holiday")
	if err != nil {
		return nil, err
	}
	var d Hoilyday
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 국내기관_외국인 매매종목가집계 API입니다.
// 증권사 직원이 장중에 집계/입력한 자료를 단순 누계한 수치로서,
// 입력시간은 외국인 09:30, 11:20, 13:20, 14:30 / 기관종합 10:00, 11:20, 13:20, 14:30 이며, 사정에 따라 변동될 수 있습니다.
func (k *Kinvest) ForeignTotalInstitution(ctx context.Context) (*ForeignInstitution, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + k.config.Token,
		"appkey":        k.config.AppKey,
		"appsecret":     k.config.SecretKey,
		"tr_id":         "FHPTJ04400000",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": string(k.config.Customer),
	}).
		SetQueryParam("FID_COND_MRKT_DIV_CODE", "V").
		SetQueryParam("FID_COND_SCR_DIV_CODE", "16449").
		//0000:전체, 0001:코스피, 1001:코스닥
		//...
		//포탈 (FAQ : 종목정보 다운로드 - 업종코드 참조)
		SetQueryParam("FID_INPUT_ISCD", "0000").
		//0: 수량정열, 1: 금액정열
		SetQueryParam("FID_DIV_CLS_CODE", "0").
		//0: 순매수상위, 1: 순매도상위
		SetQueryParam("FID_RANK_SORT_CLS_CODE", "0").
		//0:전체 1:외국인 2:기관계 3:기타
		SetQueryParam("FID_ETC_CLS_CODE", "0").
		Get("/uapi/domestic-stock/v1/quotations/foreign-institution-total")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d ForeignInstitution
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}
