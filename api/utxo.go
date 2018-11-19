package api

import (
	"github.com/abc-hardfork/broadcasttx/model"
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
	create       sync.Mutex
	param        *chaincfg.Params
	pingInterval = 5 * time.Second
	bsvHost      string
	bchHost      string

	nodes map[string]*electrum.Node
)

func QueryUtxo(c *gin.Context) {
	address := c.Query("address")

	ret, err := fetchUtxo(address, bsvHost)
	if err != nil {
		logrus.Errorf("fetchUtxo.error:%s", err.Error())
	}

	c.JSON(200, response{
		Code:    0,
		Message: "",
		Result:  ret,
	})
}

func QueryBCHUtxo(c *gin.Context) {
	address := c.Query("address")
	ret, err := fetchUtxo(address, bchHost)
	if err != nil {
		logrus.Errorf("fetchUtxo.error:%s", err.Error())

		c.JSON(200, response{
			Code:    500,
			Message: err.Error(),
		})
	}

	c.JSON(200, response{
		Code:    0,
		Message: "",
		Result:  ret,
	})
}

func DiffUtxo(c *gin.Context) {

	address := c.Query("address")

	bsvUtxo, err := fetchUtxo(address, bsvHost)
	if err != nil {
		logrus.Errorf("fetchUtxo.error:%s", err.Error())
		c.JSON(200, response{
			Code:    500,
			Message: err.Error(),
		})
	}

	bchUtxo, err := fetchUtxo(address, bchHost)
	if err != nil {
		logrus.Errorf("fetchUtxo.error:%s", err.Error())
		c.JSON(200, response{
			Code:    500,
			Message: err.Error(),
		})
	}

	diffUtxo := getUtxoDiff(bsvUtxo, bchUtxo)
	c.JSON(200, response{
		Code:    0,
		Message: "",
		Result:  model.UtxoCompare{BchUtxo: bchUtxo, BsvUtxo: bsvUtxo, DiffUtxo: diffUtxo},
	})

}
func getUtxoDiff(bch *[]model.UTxo, bsv *[]model.UTxo) *[]model.UTxo {

	ret := make([]model.UTxo, 0)
	for _, uc := range *bch {
		for _, us := range *bsv {
			if uc.TxID == us.TxID {

				ret = append(ret, uc)
				break
			}

		}
	}

	return &ret
}

func fetchUtxo(address string, host string) (*[]model.UTxo, error) {
	addr, err := cashutil.DecodeAddress(address, param)
	addrStr := addr.EncodeAddress(false)
	// find the valid utxos
	utxos, err := getNode(host).BlockchainAddressListUnspent(addrStr)
	if err != nil {
		return nil, err
	}

	ret := make([]model.UTxo, 0, 5)
	for _, utxo := range utxos {
		tx, err := getNode(host).BlockchainTransactionGet(utxo.Hash, true)
		if err != nil {
			return nil, err
		}

		hexScript := tx.Vout[utxo.Pos].ScriptPubKey.Hex
		//construct the utxo obj to return
		coin := model.UTxo{
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

func NewElectRpc(bsv string, bch string, p *chaincfg.Params) {
	bsvHost = bsv
	bchHost = bch
	param = p
	//go keepAlive(bsv)
	//go keepAlive(bch)
}

func getNode(host string) *electrum.Node {
	create.Lock()
	defer create.Unlock()

	node := nodes[host]
	if node == nil || node.Ping() != nil {
		n := electrum.NewNode()
		if err := n.ConnectTCP(host); err != nil {
			logrus.Errorf("create connection to electrum error: %v", err)
		}

		return n
	}

	return node
}

//func keepAlive(host string) {
//	var reties int
//
//	for {
//		if err := getNode(host).Ping(); err != nil {
//			reties++
//			if reties >= 1000 {
//				logrus.Error("Retry to connect to electrum server failed too many times")
//			}
//
//			logrus.Errorf("Ping to electrum server error: %v", err)
//		} else if reties != 0 {
//			reties = 0
//		}
//
//		time.Sleep(pingInterval)
//	}
//}
