package api

import (
	"strings"

	"github.com/bcext/cashutil"
	"github.com/bcext/gcash/chaincfg"
	"github.com/bcext/gcash/chaincfg/chainhash"
	"github.com/bcext/gcash/rpcclient"
	"github.com/gin-gonic/gin"
)

type txDetailResponse struct {
	BitcoinABC txDetail `json:"bitcoin-abc"`
	BitcoinSV  txDetail `json:"bitcoin-sv"`
}

type txDetail struct {
	Hash      string     `json:"hash"`
	Confirmed bool       `json:"confirmed"`
	Senders   []Sender   `json:"senders"`
	Receivers []Receiver `json:"receivers"`
	Error     []string   `json:"error"`
}

type Sender struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type Receiver struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

func FetchTx(c *gin.Context) {
	hashStr := c.Param("hash")
	hash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		c.JSON(200, ErrInvalidHashResponse)
		return
	}

	abcDetail := getTxDetail(hash, abcClient)
	svDetail := getTxDetail(hash, svClient)

	c.JSON(200, response{
		Code:    0,
		Message: "",
		Result: txDetailResponse{
			BitcoinABC: *abcDetail,
			BitcoinSV:  *svDetail,
		},
	})

}

func getTxDetail(hash *chainhash.Hash, client *rpcclient.Client) *txDetail {
	var detail txDetail

	detail.Hash = hash.String()

	abcT, err := client.GetRawTransaction(hash)
	if err != nil {
		if strings.Contains(err.Error(), "No such mempool or blockchain transaction") {
			detail.Confirmed = false
		} else {
			detail.Confirmed = true
		}

		detail.Error = append(detail.Error, err.Error())
	} else {
		detail.Confirmed = true

		// parse senders
		for _, in := range abcT.MsgTx().TxIn {
			reference, err := client.GetRawTransaction(&in.PreviousOutPoint.Hash)
			if err != nil {
				detail.Error = append(detail.Error, err.Error())
			} else {
				pkscript := reference.MsgTx().TxOut[in.PreviousOutPoint.Index].PkScript
				amount := reference.MsgTx().TxOut[in.PreviousOutPoint.Index].Value

				addr, err := parseAddress(pkscript)
				if err != nil {
					detail.Senders = append(detail.Senders, Sender{
						Address: "address decode failed",
						Amount:  cashutil.Amount(amount).ToBCH(),
					})

					detail.Error = append(detail.Error, err.Error())
				} else {
					detail.Senders = append(detail.Senders, Sender{
						Address: addr.EncodeAddress(true),
						Amount:  cashutil.Amount(amount).ToBCH(),
					})
				}
			}
		}

		for _, out := range abcT.MsgTx().TxOut {
			addr, err := parseAddress(out.PkScript)
			if err != nil {
				detail.Receivers = append(detail.Receivers, Receiver{
					Address: "address decode failed",
					Amount:  cashutil.Amount(out.Value).ToBCH(),
				})

				detail.Error = append(detail.Error, err.Error())
			} else {
				detail.Receivers = append(detail.Receivers, Receiver{
					Address: addr.EncodeAddress(true),
					Amount:  cashutil.Amount(out.Value).ToBCH(),
				})
			}
		}
	}

	return &detail
}

func parseAddress(pkscript []byte) (cashutil.Address, error) {
	var addr cashutil.Address
	var err error

	addr, err = cashutil.NewAddressPubKeyHash(pkscript, &chaincfg.MainNetParams)
	if err != nil {
		addr, err = cashutil.NewAddressScriptHash(pkscript, &chaincfg.MainNetParams)
	}

	if addr != nil {
		return addr, nil
	}

	return nil, err
}
