package main

import (
	"context"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestCode_Kosdaq(t *testing.T) {
	//나스닥, 뉴욕, 아멕스, 상해, 상해지수, 심천, 심천지수, 도쿄, 홍콩, 하노이, 호치민
	//'nas','nys','ams','shs','shi','szs','szi','tse','hks','hnx','hsx'
	cli := resty.New()

	c := NewCode(cli)
	c.Kosdaq(context.TODO())
}
