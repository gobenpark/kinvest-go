package main

import (
	"context"
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
}

func NewKinvest(appKey, secretKey string) *Kinvest {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")
	return &Kinvest{appKey: appKey, secretKey: secretKey, rest: rest}
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
