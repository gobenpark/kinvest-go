package kv

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateKinvestWithCache(t *testing.T) {
	k := NewKinvest(&Config{
		AppKey:    "PS39UX9oivQEaTPzeXorl12m6SY3AotDRm65",
		SecretKey: "QZhBRW/8gmeDxO+Gmx7oUWTbzwPWiyaYkuGZPvpjc8AkWFKIqJnAsZcwHrW+D8biL5QfJYsJYqxof1B1n546a5I401HY9Q/v8ZIBNPX606THv/NrwLVwrlQMmVrwCsSSKIXJs5Qi4n/ygptk3FkFQq3oXpejiU4U1VbEJ6vrFHK91pUhyTc=",
		Cache:     true,
	})
	t.Log(k.config.Token)
}

func TestReadFile(t *testing.T) {
	_, err := os.ReadFile("token")
	t.Log(errors.Is(err, os.ErrNotExist))
	require.Error(t, err)
}
