package api

import "github.com/bcext/gcash/rpcclient"

type RPC struct {
	Host   string
	User   string
	Passwd string
}

func GetRPC(rpcConf *RPC) (*rpcclient.Client, error) {
	// rpc client instance
	connCfg := &rpcclient.ConnConfig{
		Host:         rpcConf.Host,
		User:         rpcConf.User,
		Pass:         rpcConf.Passwd,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	c, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}
