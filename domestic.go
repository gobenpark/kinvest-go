package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type Domestic struct {
	appKey    string
	secretKey string
	rest      *resty.Client
	imitation bool
}

func NewDomestic(imitation bool, appKey, secretKey string) *Domestic {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")
	return &Domestic{appKey: appKey, secretKey: secretKey, rest: rest, imitation: imitation}
}

func (k *Domestic) ApprovalKey(ctx context.Context) (string, error) {
	res, err := k.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     k.appKey,
			"secretkey":  k.secretKey,
		}).Post("/oauth2/Approval")
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(res.Body(), "approval_key").String(), nil
}

func (k *Domestic) HashKey(ctx context.Context, body string) {
	res, err := k.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     k.appKey,
			"appsecret":  k.secretKey,
		}).Post("/oauth2/Approval")
	if err != nil {
	}
	_ = res
}

func (k *Domestic) AccessToken(ctx context.Context) (string, error) {
	res, err := k.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     k.appKey,
			"appsecret":  k.secretKey,
		}).Post("/oauth2/tokenP")
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(res.Body(), "access_token").String(), nil
}

func (k *Domestic) RevokeToken(ctx context.Context, token string) error {
	res, err := k.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     k.appKey,
			"appsecret":  k.secretKey,
		}).Post("/oauth2/revokeP")
	if err != nil {
		return err
	}

	//TODO: mapping error case
	//{"error_description":"유효하지 않은 token 입니다.","error_code":"EGW00121"}

	if gjson.GetBytes(res.Body(), "code").Int() != 200 {
		return errors.New("fail revoke token")
	}

	return nil
}

func (k *Domestic) RealtimeContract(ctx context.Context) error {
	b := SocketBody{
		Header: struct {
			ApprovalKey string `json:"approval_key"`
			Custtype    string `json:"custtype"`
			TrType      string `json:"tr_type"`
			ContentType string `json:"content-type"`
		}{
			ApprovalKey: "3daa450a-0685-4c31-b015-99c5e5f4b11d",
			Custtype:    "P",
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
				TrKey: "005930",
			},
		},
	}
	return Retry(5, func() error {
		c, res, _ := websocket.DefaultDialer.DialContext(ctx, "ws://ops.koreainvestment.com:21000/tryitout/H0STCNT0", nil)
		fmt.Printf("%#v\n", res)
		if err := c.WriteJSON(b); err != nil {
			return err
		}
		for {
			//0|H0STCNT0|001|005930^112616^62700^2^400^0.64^62700.41^62700^63000^62300^62700^62600^125^9694968^607878273600^25807^15918^-9889^93.01^4614191^4291558^1^0.45^63.03^090013^3^0^090646^5^-300^100741^2^400^20230324^20^N^71292^244009^2795345^2044512^0.16^5367918^180.61^0^^62700
			_, msg, err := c.ReadMessage()
			fmt.Println(string(msg))
			fmt.Println(err)
		}
	})

}

// 주식현재가 시세
func (k *Domestic) CurrentPrice(ctx context.Context, token, ftype, code string) (*CurrentPrice, error) {

	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010100",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) CurrentConclusion(ctx context.Context, token, ftype, code string) (*CurrentConclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010300",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) DailyPrice(ctx context.Context, token, ftype, code string) (*DailyPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010400",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_PERIOD_DIV_CODE", "D").
		SetQueryParam("FID_ORG_ADJ_PRC", "0").
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
func (k *Domestic) ExpectAskPrice(ctx context.Context, token, ftype, code string) (*ExpectPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010200",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) Investor(ctx context.Context, token, ftype, code string) (*Investor, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010900",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) Member(ctx context.Context, token, ftype, code string) (*Member, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST01010600",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) CurrentELW(ctx context.Context, token, ftype, code string) (*ELW, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKEW15010000",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) DailyChartPrice(ctx context.Context, start time.Time, end time.Time, token, ftype, code string) (*DailyChartPrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST03010100",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
		SetQueryParam("FID_INPUT_ISCD", code).
		SetQueryParam("FID_INPUT_DATE_1", start.Format("20060102")).
		SetQueryParam("FID_INPUT_DATE_2", end.Format("20060102")).
		//D일봉,W:주봉,M월봉,Y년봉
		SetQueryParam("FID_PERIOD_DIV_CODE", "D").
		//수정주가 여부 0: 수정주가, 1원주가
		SetQueryParam("FID_ORG_ADJ_PRC", "0").
		Get("/uapi/domestic-stock/v1/quotations/inquire-daily-itemchartprice")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d DailyChartPrice
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}

	return &d, nil
}

// 주식현재가 당일시간대별체결 API입니다.
func (k *Domestic) CurrentTimePerConclusion(ctx context.Context, startDate time.Time, token, ftype, code string) (*Conclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHPST01060000",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) CurrentOvertimePerConclusion(ctx context.Context, token, ftype, code string) (*Conclusion, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHPST02310000",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) DailyOvertimePerPrice(ctx context.Context, token, ftype, code string) (*OvertimePrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHPST02320000",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
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
func (k *Domestic) CurrentTimeChartPrice(ctx context.Context, token, ftype, code string) (*OvertimePrice, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHKST03010200",
		"custtype":      "P",
	}).SetQueryParam("FID_COND_MRKT_DIV_CODE", ftype).
		SetQueryParam("FID_INPUT_ISCD", code).
		Get("/uapi/domestic-stock/v1/quotations/inquire-time-itemchartprice")
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

// 상품 기본조회
func (k *Domestic) ItemInfo(ctx context.Context, token, ftype, code string) (*Item, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "CTPF1604R",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": "P",
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
func (k *Domestic) HoilydayInfo(ctx context.Context, date time.Time, token, ftype, code string) (*Hoilyday, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "CTCA0903R",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": "P",
	}).
		SetQueryParam("BASS_DT", date.Format("20060102")).
		SetQueryParam("CTX_AREA_NK", "").
		SetQueryParam("CTX_AREA_FK", "").
		Get("/uapi/domestic-stock/v1/quotations/chk-holiday")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	var d Hoilyday
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// 국내기관_외국인 매매종목가집계 API입니다.
// 증권사 직원이 장중에 집계/입력한 자료를 단순 누계한 수치로서,
// 입력시간은 외국인 09:30, 11:20, 13:20, 14:30 / 기관종합 10:00, 11:20, 13:20, 14:30 이며, 사정에 따라 변동될 수 있습니다.
func (k *Domestic) ForeignTotalInstitution(ctx context.Context, token string) (*ForeignInstitution, error) {
	res, err := k.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token,
		"appkey":        k.appKey,
		"appsecret":     k.secretKey,
		"tr_id":         "FHPTJ04400000",
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": "P",
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
