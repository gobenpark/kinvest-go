package kv

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-resty/resty/v2"
)

func TestCode_Kosdaq(t *testing.T) {
	cli := resty.New()

	c := NewCodeManager(cli)
	codes, err := c.Kosdaq(context.TODO())
	require.NoError(t, err)
	require.Greater(t, len(codes), 0)
	t.Log(codes)
}

func TestCode_Kospi(t *testing.T) {
	cli := resty.New()
	c := NewCodeManager(cli)

	codes, err := c.Kospi(context.TODO())
	require.NoError(t, err)
	fmt.Println(codes)
}
