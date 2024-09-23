package util

import "net/http"

var (
	ErrKeyNotFound = NewHttpError(http.StatusNotFound, "key not found")
	ErrInvalidPPRN = NewHttpError(http.StatusBadRequest, "invalid pprn")
	ErrInternalServer = func(err error) *HttpError { return NewHttpError(http.StatusNotFound, err.Error()) }
	ErrInvalidAccountId = NewHttpError(http.StatusBadRequest, "invalid account id")
	ErrUnsupportedKeySpec = NewHttpError(http.StatusBadRequest, "unsupported key specification")
	ErrInvalidCiphertextBlob = NewHttpError(http.StatusBadRequest, "invalid request")
)
