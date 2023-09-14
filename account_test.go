package kv

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	AppKey    = ""
	SecretKey = ""
	Token     = ""
)

func MockClient(t *testing.T) *Kinvest {
	t.Helper()
	cli := NewKinvest(&Config{
		AppKey:    AppKey,
		SecretKey: SecretKey,
		Imitation: false,
		Customer:  Person,
		Token:     Token,
		Account:   "",
		Cache:     true,
	})
	return cli
}

func TestAccount_Approval(t *testing.T) {
	cli := MockClient(t)
	key, err := cli.ApprovalKey(context.TODO())
	assert.NoError(t, err)
	assert.NotEqual(t, key, "")
	t.Log(key)
}

func TestAccount_AccessToken(t *testing.T) {
	cli := MockClient(t)
	token, err := cli.AccessToken(context.TODO())
	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
	t.Log(token)
}

func TestAccount_RevokeToken(t *testing.T) {
	cli := MockClient(t)
	err := cli.RevokeToken(context.TODO())
	assert.NoError(t, err)
}

func TestAccount_Approval_PossibleOrder(t *testing.T) {
	cli := MockClient(t)
	err := cli.PossibleOrder(context.TODO(), "68", "", "5000", Limit, false, false)
	assert.NoError(t, err)
}

func TestAccount_AccountBalance(t *testing.T) {
	cli := MockClient(t)
	balance, err := cli.AccountBalance(context.TODO())
	assert.NoError(t, err)
	t.Log(balance)
}
