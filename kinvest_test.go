package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	AppKey    = ""
	SecretKey = ""
)

func MockClient(t *testing.T) *Kinvest {
	t.Helper()

	cli := NewKinvest(AppKey, SecretKey)
	return cli
}

func TestNewKinvest_Approval(t *testing.T) {
	cli := MockClient(t)
	key, err := cli.ApprovalKey(context.TODO())
	assert.NoError(t, err)
	assert.NotEqual(t, key, "")
	t.Log(key)
}

func Test_AccessToken(t *testing.T) {
	cli := MockClient(t)
	token, err := cli.AccessToken(context.TODO())
	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
}

func Test_RevokeToken(t *testing.T) {
	cli := MockClient(t)
	err := cli.RevokeToken(context.TODO(), "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6IjdlNDU1ZWJlLTQ2YzgtNDM3Yi1hM2UwLTQ4NmQzZDgzMzhhMCIsImlzcyI6InVub2d3IiwiZXhwIjoxNjc5NzA4NjYzLCJpYXQiOjE2Nzk2MjIyNjMsImp0aSI6IlBTMzlVWDlvaXZRRWFUUHplWG9ybDEybTZTWTNBb3REUm02NSJ9.3bRacS9wEo8W_e61Jj12w61DFKvCXCJu8nnZWVBAd_tk8zXKuB1MBiXuSQPwsg7Jh-HXeh-CGAeK-M7EADl-5A")
	assert.NoError(t, err)
}

func Test_RealTimeContract(t *testing.T) {
	cli := MockClient(t)
	cli.RealtimeContract(context.TODO())
}

func Test_BodyParse(t *testing.T) {
	body := "0|H0STCNT0|001|005930^112616^62700^2^400^0.64^62700.41^62700^63000^62300^62700^62600^125^9694968^607878273600^25807^15918^-9889^93.01^4614191^4291558^1^0.45^63.03^090013^3^0^090646^5^-300^100741^2^400^20230324^20^N^71292^244009^2795345^2044512^0.16^5367918^180.61^0^^62700"

	data := strings.Split(body, "|")
	assert.Len(t, data, 4)

	assert.Equal(t, data[0], "0")
	assert.Equal(t, data[1], "H0STCNT0")
	assert.Equal(t, data[2], "001")
	for _, i := range strings.Split(data[3], "^") {
		fmt.Println(i)
	}
}
