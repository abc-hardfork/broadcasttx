package model

type UtxoCompare struct {
	BsvUtxo  *[]UTxo
	BchUtxo  *[]UTxo
	DiffUtxo *[]UTxo
}

type UTxo struct {
	TxID         string `json:"txid"`
	VOut         int    `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
	Value        int64  `json:"value"`
	BlockHeight  int32  `json:"blockHeight"`
}
