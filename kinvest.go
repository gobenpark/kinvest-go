package main

import (
	"github.com/go-resty/resty/v2"
)

type Kinvest struct {
	appKey    string
	secretKey string
	rest      *resty.Client
	imitation bool

	Domestic
	Overseas
}

func NewKinvest(imitation bool, appKey, secretKey string) *Kinvest {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")

	return &Kinvest{
		appKey:    appKey,
		secretKey: secretKey,
		rest:      rest,
		imitation: imitation,
		Domestic: Domestic{
			appKey:    appKey,
			secretKey: secretKey,
			rest:      rest,
			imitation: false,
		},
		Overseas: Overseas{
			appKey:    appKey,
			secretKey: secretKey,
			rest:      rest,
			imitation: false,
		},
	}
}
