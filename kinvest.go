package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type Kinvest struct {
	appKey    string
	secretKey string
	rest      *resty.Client
	imitation bool
}

func NewKinvest(imitation bool, appKey, secretKey string) *Kinvest {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")
	return &Kinvest{appKey: appKey, secretKey: secretKey, rest: rest, imitation: imitation}
}

func (k *Kinvest) ApprovalKey(ctx context.Context) (string, error) {
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

func (k *Kinvest) HashKey(ctx context.Context, body string) {
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

func (k *Kinvest) AccessToken(ctx context.Context) (string, error) {
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

func (k *Kinvest) RevokeToken(ctx context.Context, token string) error {
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

func (k *Kinvest) RealtimeContract(ctx context.Context) error {
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
func (k *Kinvest) CurrentPrice(ctx context.Context, token, ftype, code string) (*CurrentPrice, error) {

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
func (k *Kinvest) CurrentConclusion(ctx context.Context, token, ftype, code string) (*CurrentConclusion, error) {
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
func (k *Kinvest) DailyPrice(ctx context.Context, token, ftype, code string) (*DailyPrice, error) {
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
func (k *Kinvest) ExpectAskPrice(ctx context.Context, token, ftype, code string) (*ExpectPrice, error) {
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
func (k *Kinvest) Investor(ctx context.Context, token, ftype, code string) (*Investor, error) {
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
func (k *Kinvest) Member(ctx context.Context, token, ftype, code string) (*Member, error) {
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

func (k *Kinvest) CurrentELW(ctx context.Context, token, ftype, code string) (*ELW, error) {
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
