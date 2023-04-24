package kv

import "github.com/go-resty/resty/v2"

type Overseas struct {
	appKey    string
	secretKey string
	rest      *resty.Client
	imitation bool
}
