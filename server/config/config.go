package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

var Current *Config

func NewConfig(env string) {
	var C Config
	Current = &C
	viper.AddConfigPath("$GOPATH/src/github.com/hirokikojima/go-sample-app/server/config/")
	viper.SetConfigType("yml")

	viper.SetConfigName(env)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}
}