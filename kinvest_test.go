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
	datas := strings.Split(data[3], "^")
	res := ResponseBody{
		Encrypted:                  data[0],
		TRID:                       data[1],
		DataCounts:                 data[2],
		Code:                       datas[0],
		ContractHour:               datas[1],
		Price:                      datas[2],
		CompareSign:                datas[3],
		CompareDay:                 datas[4],
		CompareRate:                datas[5],
		WeightAveragePrice:         datas[6],
		Open:                       datas[7],
		High:                       datas[8],
		Low:                        datas[9],
		AskPrice:                   datas[10],
		BidPrice:                   datas[11],
		ContractVolume:             datas[12],
		AccumulateVolume:           datas[13],
		AccumulateTransactionMoney: datas[14],
		AskCount:                   datas[15],
		BidCount:                   datas[16],
		PureBidCount:               datas[17],
		VolumePower:                datas[18],
		TotalAskCounts:             datas[19],
		TotalBidCounts:             datas[20],
		ContractDivide:             datas[21],
		BidRate:                    datas[22],
		PredayVolumeCompareRate:    datas[23],
		OpenningTime:               datas[24],
		OpenCompareSign:            datas[25],
		OpenCompare:                datas[26],
		HighTime:                   datas[27],
		HighCompareSign:            datas[28],
		HighCompare:                datas[29],
		LowTime:                    datas[30],
		LowCompareSign:             datas[31],
		LowCompare:                 datas[32],
		BusinessDate:               datas[33],
		NewMarketOpCode:            datas[34],
		TransactionSuspension:      datas[35],
		RemainAsk:                  datas[36],
		RemainBid:                  datas[37],
		TotalRemainAsk:             datas[38],
		TotalRemainBid:             datas[39],
		VolumeRotateRate:           datas[40],
		PreDayTotalVolume:          datas[41],
		PreDayTotalVolumeRate:      datas[42],
		HourClockCode:              datas[43],
		MarketTermCode:             datas[44],
		VIStandardPrice:            datas[45],
	}
	fmt.Println(res.PreDayTotalVolumeRate)
	fmt.Println(res.PreDayTotalVolume)

}
