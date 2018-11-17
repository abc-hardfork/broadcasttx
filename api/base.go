package api

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/bcext/gcash/wire"
)

func parseTx(rawtx string) (*wire.MsgTx, error) {
	txByte, err := hex.DecodeString(strings.Trim(rawtx, " "))
	if err != nil {
		return nil, err
	}

	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(txByte))
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
