package main

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
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

// Search Market Codes Instance
type Code struct {
	cli *resty.Client
}

func NewCode(cli *resty.Client) *Code {
	cli.SetBaseURL("https://new.real.download.dws.co.kr")
	return &Code{cli: cli}
}

func (c *Code) Kosdaq(ctx context.Context) error {
	res, err := c.cli.R().
		SetContext(ctx).
		Get("/common/master/kosdaq_code.mst.zip")
	if err != nil {
		return err
	}
	rd, err := zip.NewReader(bytes.NewReader(res.Body()), res.Size())
	if err != nil {
		return err
	}
	if len(rd.File) == 0 {
		return errors.New("empty kosdaq zip data")
	}

	f, err := rd.File[0].Open()
	if err != nil {
		return err
	}

	csvdata := csv.NewReader(f)
	rc, err := csvdata.ReadAll()
	if err != nil {
		return err
	}

	for _, i := range rc {
		fmt.Print(i[0])
	}
	return nil

}

func (c *Code) Kospi(ctx context.Context) error {
	return nil
}
