package main

import (
	"fmt"
	"github.com/abc-hardfork/broadcasttx/api"
	"github.com/abc-hardfork/broadcasttx/routers"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// init configuration
	conf := GetConf()

	// init log
	err := InitLog()
	if err != nil {
		fmt.Println("initial logrus failed:", err.Error())
		os.Exit(1)
	}

	// log whitelist
	logrus.Info(conf.AllowIP)

	// inject the whitelist for router middleware
	routers.WhiterList = conf.AllowIP

	// initial rpc instance
	err = api.New(&conf.Abc, &conf.Sv)
	if err != nil {
		fmt.Println("initial rpc instance failed:", err.Error())
		os.Exit(1)
	}

	// initial rpc instance
	api.NewElectRpc(conf.Electron.Bsv, conf.Electron.Bch, GetChainParam())

	// start up the server
	engine := routers.InitRouter()
	err = engine.Run(":8888")
	if err != nil {
		fmt.Println("engine startup failed:", err.Error())
		os.Exit(1)
	}
}
