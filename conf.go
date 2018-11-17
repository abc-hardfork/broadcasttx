package main

import (
	"os"
	"path/filepath"

	"github.com/qshuai/broadcasttx/api"
	"github.com/spf13/viper"
)

const (
	configName = "conf.yml"
)

type configuration struct {
	AllowIP []string `mapstructure:"allowip"`
	Abc     api.RPC  `mapstructure:"abc"`
	Sv      api.RPC  `mapstructure:"sv"`
}

func GetConf() *configuration {
	// parse config
	file, err := os.Open(filepath.Join("./", configName))
	if err != nil {
		panic("Open config file error: " + err.Error())
	}
	defer file.Close()

	err = viper.ReadConfig(file)
	if err != nil {
		panic("Read config file error: " + err.Error())
	}
	var config configuration
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("Parse config file error: " + err.Error())
	}

	// TODO validate configuration
	//helper.Must(nil, config.Validate())

	return &config
}
