package kv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

func (a *Kinvest) ApprovalKey(ctx context.Context) (string, error) {
	res, err := a.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     a.config.AppKey,
			"secretkey":  a.config.SecretKey,
		}).Post("/oauth2/Approval")
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(res.Body(), "approval_key").String(), nil
}

func (a *Kinvest) HashKey(ctx context.Context, body string) {
	res, err := a.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     a.config.AppKey,
			"appsecret":  a.config.SecretKey,
		}).Post("/oauth2/Approval")
	if err != nil {
	}
	_ = res
}

func (a *Kinvest) AccessToken(ctx context.Context) (string, error) {
	res, err := a.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"grant_type": "client_credentials",
			"appkey":     a.config.AppKey,
			"appsecret":  a.config.SecretKey,
		}).Post("/oauth2/tokenP")
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(res.Body(), "access_token").String(), nil
}

func (a *Kinvest) RevokeToken(ctx context.Context) error {
	res, err := a.rest.R().SetContext(ctx).
		SetBody(map[string]string{
			"token":     a.config.Token,
			"appkey":    a.config.AppKey,
			"appsecret": a.config.SecretKey,
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

// 주문가능 조회
// account:계좌번호,code:계좌상품코드,product:상품번호,perprice:주문단가,orderDVS:주문구분
// incma: cma평가금액 포함여부,inoverseas: 해외포함여
func (a *Kinvest) PossibleOrder(ctx context.Context, code, product, perprice string, orderDVS OrderType, incma, inoverseas bool) error {
	res, err := a.rest.R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + a.config.Token,
		"appsecret":     a.config.SecretKey,
		"tr_id":         YesOrNo("VTTC8908R", "TTTC8908R", a.config.Imitation),
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"appkey":   a.config.AppKey,
		"custtype": string(a.config.Customer),
	}).
		SetQueryParams(map[string]string{
			"CANO":                 a.config.Account[:8],
			"ACNT_PRDT_CD":         a.config.Account[len(a.config.Account)-2:],
			"PDNO":                 product,
			"ORD_UNPR":             perprice,
			"ORD_DVSN":             string(orderDVS),
			"CMA_EVLU_AMT_ICLD_YN": YesOrNo("Y", "N", incma),
			"OVRS_ICLD_YN":         YesOrNo("Y", "N", inoverseas),
		}).
		Get("/uapi/domestic-stock/v1/trading/inquire-psbl-order")
	if err != nil {
		return err
	}
	fmt.Println(res.String())
	return nil
}

func (a *Kinvest) AccountBalance(ctx context.Context) (Balance, error) {
	if len(a.config.Account) != 10 {
		return Balance{}, errors.New("invalud account number")
	}
	res, err := a.rest.SetDebug(true).R().SetContext(ctx).SetHeaders(map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + a.config.Token,
		"appsecret":     a.config.SecretKey,
		"tr_id":         YesOrNo("VTTC8434R", "TTTC8434R", a.config.Imitation),
		"appkey":        a.config.AppKey,
		// 공백: 초기조회,N 다음데이터 조회 Res Header가 M일경우
		"tr_cont":  "",
		"custtype": string(a.config.Customer),
	}).
		SetQueryParam("CANO", a.config.Account[:8]).
		SetQueryParam("ACNT_PRDT_CD", a.config.Account[len(a.config.Account)-2:]).
		SetQueryParam("AFHR_FLPR_YN", "N").
		SetQueryParam("OFL_YN", "").
		SetQueryParam("INQR_DVSN", "02").
		SetQueryParam("UNPR_DVSN", "01").
		SetQueryParam("FUND_STTL_ICLD_YN", "N").
		SetQueryParam("FNCG_AMT_AUTO_RDPT_YN", "N").
		SetQueryParam("PRCS_DVSN", "00").
		SetQueryParam("CTX_AREA_FK100", "").
		SetQueryParam("CTX_AREA_NK100", "").
		Get("/uapi/domestic-stock/v1/trading/inquire-balance")
	if err != nil {
		return Balance{}, err
	}
	var d Balance
	if err := json.Unmarshal(res.Body(), &d); err != nil {
		return Balance{}, err
	}
	return d, nil
}
