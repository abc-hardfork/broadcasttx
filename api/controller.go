package api

import (
	"github.com/bcext/gcash/rpcclient"
	"github.com/gin-gonic/gin"
)

var (
	abcClient *rpcclient.Client
	svClient  *rpcclient.Client
)

type Param struct {
	Rawtx string `form:"rawtx" binding:"required"`
}

type response struct {
	Code    errCode     `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func BroadcastAbcTx(c *gin.Context) {
	var p Param
	err := c.ShouldBind(&p)
	if err != nil {
		c.JSON(200, ErrParamResponse)
		return
	}

	tx, err := parseTx(p.Rawtx)
	if err != nil {
		c.JSON(200, ErrParseTransactionResponse)
		return
	}

	hash, err := abcClient.SendRawTransaction(tx, false)
	if err != nil {
		c.JSON(200, errSendRawTransaction(err))
		return
	}

	c.JSON(200, response{
		Code:    0,
		Message: "broadcast bitcoin-abc transaction successfully",
		Result:  hash.String(),
	})
}

func BroadcastSvTx(c *gin.Context) {
	var p Param
	err := c.ShouldBind(&p)
	if err != nil {
		c.JSON(200, ErrParamResponse)
		return
	}

	tx, err := parseTx(p.Rawtx)
	if err != nil {
		c.JSON(200, ErrParseTransactionResponse)
		return
	}

	hash, err := svClient.SendRawTransaction(tx, false)
	if err != nil {
		c.JSON(200, errSendRawTransaction(err))
		return
	}

	c.JSON(200, response{
		Code:    0,
		Message: "broadcast bitcoin-sv transaction successfully",
		Result:  hash.String(),
	})
}

func New(abc *RPC, sv *RPC) error {
	var err error
	abcClient, err = GetRPC(abc)
	if err != nil {
		return err
	}

	svClient, err = GetRPC(sv)
	if err != nil {
		return err
	}

	return nil
}
