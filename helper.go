package kv

import (
	"context"
	"time"
)

func Retry(ctx context.Context, max int, f func() error) error {
	retries := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := f(); err != nil {
				<-time.After((1 << retries) * time.Second)
				retries++

				if retries >= max {
					return err
				}
				continue
			}
		}

		return nil
	}
}

func YesOrNo[T comparable](t T, f T, c bool) T {
	if c {
		return t
	}
	return f
}
