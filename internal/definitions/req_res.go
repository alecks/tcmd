package definitions

import "encoding/json"

type TcmdRequest struct {
	Method string `json:"method"`
}

type TcmdError struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

const (
	CloseMethod = "_close"

	JsonErrorMsg    = "err (un)marshalling JSON"
	JsonErrorCode   = "1"
	HandleErrorCode = "2"
)

var JsonErrMarshalled, _ = json.Marshal(TcmdError{
	Error: JsonErrorMsg,
	Code:  JsonErrorCode,
})
