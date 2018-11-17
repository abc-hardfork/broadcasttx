package api

import (
	"github.com/bcext/cashutil"
	"github.com/bcext/gcash/chaincfg"
	"github.com/copernet/go-electrum/electrum"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	// prevent from creating two more connection instances
	create sync.Mutex
	node   *electrum.Node
	param  *chaincfg.Params
)

type UTxo struct {
	TxID         string `json:"txid"`
	VOut         int    `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
	Value        int64  `json:"value"`
	BlockHeight  int32  `json:"blockHeight"`
}

func QueryUtxo(c *gin.Context) {
	address := c.Query("address")

	ret, err := fetchUtxo(address)
	if err != nil {
		logrus.Errorf("fetchUtxo.error:%s", err.Error())
	}

	c.JSON(200, response{
		Code:    0,
		Message: "",
		Result:  ret,
	})
}
func fetchUtxo(address string) (*[]UTxo, error) {
	addr, err := cashutil.DecodeAddress(address, param)
	addrStr := addr.EncodeAddress(false)
	// find the valid utxos
	utxos, err := node.BlockchainAddressListUnspent(addrStr)
	if err != nil {
		return nil, err
	}

	ret := make([]UTxo, 0, 5)
	for _, utxo := range utxos {
		tx, err := node.BlockchainTransactionGet(utxo.Hash, true)
		if err != nil {
			return nil, err
		}

		hexScript := tx.Vout[utxo.Pos].ScriptPubKey.Hex
		//construct the utxo obj to return
		coin := UTxo{
			TxID:         utxo.Hash,
			VOut:         int(utxo.Pos),
			Value:        utxo.Value,
			ScriptPubKey: hexScript,
			BlockHeight:  utxo.Height,
		}
		ret = append(ret, coin)
	}

	return &ret, nil
}

func NewElectRpc(host string, p *chaincfg.Params) error {
	create.Lock()
	defer create.Unlock()

	n := electrum.NewNode()
	if err := n.ConnectTCP(host); err != nil {
		logrus.Errorf("create connection to electrum error: %v", err)
		// unnecessary to star the application if electrum connection failed
		//os.Exit(1)
		return err
	}

	node = n
	param = p
	go keepAlive()
	return nil
}

func keepAlive() {
	var reties int

	for {
		if err := node.Ping(); err != nil {
			reties++
			if reties >= 1000 {
				logrus.Error("Retry to connect to electrum server failed too many times")

				// Should find a HA resolution before stop exit program directly
			}

			logrus.Errorf("Ping to electrum server error: %v", err)
		} else if reties != 0 {
			reties = 0
		}

		time.Sleep(3)
	}
}
