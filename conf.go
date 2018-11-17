package main

import (
	"github.com/abc-hardfork/broadcasttx/api"
	"github.com/bcext/gcash/chaincfg"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configName = "conf.yml"
)

type configuration struct {
	TestNet  bool     `mapstructure:"testnet"`
	AllowIP  []string `mapstructure:"allowip"`
	Abc      api.RPC  `mapstructure:"abc"`
	Sv       api.RPC  `mapstructure:"sv"`
	Electron Electron `mapstructure:"Electron"`
}

type Electron struct {
	Host string `mapstructure:"host"`
}

func GetChainParam() *chaincfg.Params {
	conf := GetConf()
	if conf.TestNet {
		return &chaincfg.TestNet3Params
	}

	return &chaincfg.MainNetParams
}

func GetConf() *configuration {
	// parse config
	file, err := os.Open(filepath.Join("./", configName))
	if err != nil {
		panic("Open config file error: " + err.Error())
	}
	defer file.Close()

	viper.SetConfigType("yaml")
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
