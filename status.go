package gogemini

import (
	"fmt"
	"io"
)

type StatusCode int8

const (
	StatusInput               StatusCode = 10
	StatusSensitiveInput      StatusCode = 11
	StatusSuccess             StatusCode = 20
	StatusTempRedirect        StatusCode = 30
	StatusPermRedirect        StatusCode = 31
	StatusTempFailure         StatusCode = 40
	StatusServerUnavailable   StatusCode = 41
	StatusCgiError            StatusCode = 42
	StatusProxyError          StatusCode = 43
	StatusSlowDown            StatusCode = 44
	StatusPermFailure         StatusCode = 50
	StatusNotFound            StatusCode = 51
	StatusGone                StatusCode = 52
	StatusProxyRequestRefused StatusCode = 53
	StatusBadRequest          StatusCode = 59
	StatusClientCertRequired  StatusCode = 60
	StatusCertNotAuthorized   StatusCode = 61
	StatusCertNotValid        StatusCode = 62
)

func sendStatus(w io.Writer, code StatusCode, meta string) {
	w.Write([]byte(fmt.Sprintf("%d %s\n", code, meta)))
}
