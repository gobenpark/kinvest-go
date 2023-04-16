package main

import (
	"github.com/go-resty/resty/v2"
)

type Config struct {
	AppKey    string
	SecretKey string
	Token     string
	Imitation bool
	Customer  Customer
	Account   string
}

type Kinvest struct {
	config *Config
	rest   *resty.Client
	*CodeManager
	Domestic
	Overseas
	Account
}

func NewKinvest(config *Config) *Kinvest {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")
	return &Kinvest{
		config: config,
		rest:   rest,
		Domestic: Domestic{
			rest:   rest,
			config: config,
		},
		Overseas: Overseas{
			rest: rest,
		},
		Account: Account{
			rest:   rest,
			config: config,
		},
		CodeManager: NewCodeManager(rest),
	}
}
