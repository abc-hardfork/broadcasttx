package api

type errCode int

const (
	ErrParam errCode = iota + 1
	ErrParseTransaction
	ErrSendRawTransaction
	ErrNotAllowedIP
	ErrInvalidHash
)

var (
	ErrParamResponse            = response{Code: ErrParam, Message: "invalid parameter"}
	ErrParseTransactionResponse = response{Code: ErrParseTransaction, Message: "invalid raw transaction"}
	ErrNotAllowedIPResponse     = response{Code: ErrNotAllowedIP, Message: "authorise failed for your IP"}
	ErrInvalidHashResponse      = response{Code: ErrInvalidHash, Message: "invalid transaction hash"}
)

func errSendRawTransaction(err error) response {
	return response{
		Code:    ErrSendRawTransaction,
		Message: "broadcast transaction to the full node failed: " + err.Error(),
	}
}
