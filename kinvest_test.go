package main

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	AppKey    = ""
	SecretKey = ""
	Token     = ""
)

func MockClient(t *testing.T) *Kinvest {
	t.Helper()
	cli := NewKinvest(Real, AppKey, SecretKey)
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
	t.Log(token)
}

func Test_RevokeToken(t *testing.T) {
	cli := MockClient(t)
	err := cli.RevokeToken(context.TODO(), "")
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

func Test_CurrentPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentPrice(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_CurrentConclusion(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentConclusion(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_DailyPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.DailyPrice(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_ExpectAskPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.ExpectAskPrice(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_Investor(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.Investor(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_Member(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.Member(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_CurrentELW(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentELW(context.TODO(), Token, "W", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_DailyChartPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.DailyChartPrice(context.TODO(), time.Now().Add(-96*time.Hour), time.Now(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_CurrentTimePerConclusion(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentTimePerConclusion(context.TODO(), time.Now().Add(-96*time.Hour), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_CurrentOvertimePerConclusion(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentOvertimePerConclusion(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_DailyOvertimePerPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.DailyOvertimePerPrice(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_CurrentTimeChartPrice(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.CurrentTimeChartPrice(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_ItemInfo(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.ItemInfo(context.TODO(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_Hoilyday(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.HoilydayInfo(context.TODO(), time.Now(), Token, "J", "005930")
	assert.NoError(t, err)
	fmt.Println(m)
}

func Test_ForeignTotalInstitution(t *testing.T) {
	cli := MockClient(t)
	m, err := cli.ForeignTotalInstitution(context.TODO(), Token)
	assert.NoError(t, err)
	fmt.Println(m)
}
