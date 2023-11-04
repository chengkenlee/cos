package util

import (
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

var (
	Config *viper.Viper
	Logger *zap.Logger
	Client *cos.Client
)

type Parms struct {
}
