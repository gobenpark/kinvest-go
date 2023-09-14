package kv

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type MarketCode string

const (
	Nas MarketCode = "nas" //나스닥
	Nys MarketCode = "nys" //뉴욕
	Ams MarketCode = "ams" //아멕스
	Sha MarketCode = "shs" //상해
	Shi MarketCode = "shi" //상해지수
	Szs MarketCode = "szs" //심천
	Szi MarketCode = "szi" //심천지수
	Tse MarketCode = "tse" //도쿄
	Hks MarketCode = "hks" //홍콩
	Hnx MarketCode = "hnx" //하노이
	Hsx MarketCode = "hsx" //호치민

	// 나스닥, 뉴욕, 아멕스, 상해, 상해지수, 심천, 심천지수, 도쿄, 홍콩, 하노이, 호치민
	// 'nas','nys','ams','shs','shi','szs','szi','tse','hks','hnx','hsx'
)

type Code struct {
	Code     string
	Name     string
	Industry string
}

// Search Market Codes Instance
type CodeManager struct {
	cli *resty.Client
}

func NewCodeManager(cli *resty.Client) *CodeManager {
	cli.SetBaseURL("http://data.krx.co.kr")
	return &CodeManager{cli: cli}
}

func (c *CodeManager) Kosdaq(ctx context.Context) ([]Code, error) {
	res, err := c.cli.
		R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"bld":         "dbms/MDC/STAT/standard/MDCSTAT03901",
			"locale":      "ko_KR",
			"mktId":       "KSQ",
			"trdDd":       "20230416",
			"money":       "1",
			"csvxls_isNo": "false",
		}).
		Post("/comm/bldAttendant/getJsonData.cmd")
	if err != nil {
		return nil, err
	}
	var codes []Code
	re := gjson.ParseBytes(res.Body())
	for _, i := range re.Get("block1").Array() {
		codes = append(codes, Code{
			Code:     i.Get("ISU_SRT_CD").String(),
			Name:     i.Get("ISU_ABBRV").String(),
			Industry: i.Get("IDX_IND_NM").String(),
		})
	}

	return codes, nil
}

func (c *CodeManager) Kospi(ctx context.Context) ([]Code, error) {
	res, err := c.cli.
		R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"bld":         "dbms/MDC/STAT/standard/MDCSTAT03901",
			"locale":      "ko_KR",
			"mktId":       "STK",
			"trdDd":       "20230416",
			"money":       "1",
			"csvxls_isNo": "false",
		}).
		Post("/comm/bldAttendant/getJsonData.cmd")
	if err != nil {
		return nil, err
	}
	var codes []Code
	re := gjson.ParseBytes(res.Body())
	for _, i := range re.Get("block1").Array() {
		codes = append(codes, Code{
			Code:     i.Get("ISU_SRT_CD").String(),
			Name:     i.Get("ISU_ABBRV").String(),
			Industry: i.Get("IDX_IND_NM").String(),
		})
	}

	return codes, nil
}
