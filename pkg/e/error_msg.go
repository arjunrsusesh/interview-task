package e

import "errors"

var (
	QueryError     = errors.New("failed to execute the db query")
	InvalidCountry = errors.New("invalid country")
)

var ErrorMap = map[int]error{
	ErrorCodeInvalidCountry:  InvalidCountry,
	ErrorInternalServerError: QueryError,
}

const (
	Error_MSGInvalidRequestBody = "Invalid request body"
	ErrorMsgInvalidID           = "Invalid country"
)
