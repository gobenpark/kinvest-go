package main

import (
	"github.com/go-resty/resty/v2"
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

func (k *Kinvest) ApprovalKey() (string, error) {

	res, err := k.rest.R().SetBody(map[string]string{
		"grant_type": "client_credentials",
		"appkey":     k.appKey,
		"secretkey":  k.secretKey,
	}).Post("/oauth2/Approval")
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(res.Body(), "approval_key").String(), nil
}
