package kv

import (
	"context"
	"errors"
	"os"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	AppKey    string
	SecretKey string
	Token     string
	Imitation bool
	Customer  Customer
	Account   string
	Debug     bool
	// token 자동 요청, 파일로 저장
	Cache bool
}

type Kinvest struct {
	config *Config
	rest   *resty.Client
	*CodeManager
	Overseas
	log *zap.Logger
}

func NewKinvest(config *Config) *Kinvest {
	rest := resty.New()
	rest.SetBaseURL("https://openapi.koreainvestment.com:9443")
	if config.Imitation {
		rest.SetBaseURL("https://openapivts.koreainvestment.com:29443")
	}
	manager := resty.New().SetBaseURL("http://data.krx.co.kr")

	conf := zap.NewProductionConfig()
	conf.Level = zap.NewAtomicLevelAt(func() zapcore.Level {
		if config.Debug {
			return zap.DebugLevel
		} else {
			return zap.InfoLevel
		}
	}())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	conf.EncoderConfig = encoderConfig
	conf.OutputPaths = []string{
		"kinvest.log",
		"stderr",
	}

	k := &Kinvest{
		config:      config,
		rest:        rest,
		CodeManager: NewCodeManager(manager),
		log:         zap.Must(conf.Build()),
	}

	if config.Cache {
		tk, err := os.ReadFile("token")
		if err != nil && errors.Is(err, os.ErrNotExist) {
			tk, err := k.AccessToken(context.Background())
			if err != nil {
				k.log.Panic("AccessToken", zap.Error(err))
			}
			f, err := os.Create("token")
			if err != nil {
				k.log.Panic("Create token cache", zap.Error(err))
			}
			_, err = f.WriteString(tk)
			if err != nil {
				k.log.Panic("Write token cache", zap.Error(err))
			}
		}
		k.config.Token = string(tk)
	}
	return k
}
