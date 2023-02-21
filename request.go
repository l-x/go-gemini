package gogemini

import (
	"errors"
	"fmt"
	"io"
	"net/url"
)

const maxRequestSize = 1024

type Request struct {
	Host  string
	Path  string
	Query string
}

func readRawRequest(in io.Reader) (s string, err error) {
	var buffer []byte

	buffer, err = io.ReadAll(
		io.LimitReader(in, maxRequestSize+1),
	)

	return string(buffer), err
}

func parseRequest(in io.Reader, out io.Writer, r *Request) error {
	rawRequest, err := readRawRequest(in)
	if err != nil {
		sendStatus(out, StatusTempFailure, "error reading request")
		return err
	}

	if len(rawRequest) > maxRequestSize {
		sendStatus(out, StatusBadRequest, "Request size exceeds limit")
		return fmt.Errorf("Request exceeds the limit of %d bytes", maxRequestSize)
	}

	if rawRequest == "" {
		sendStatus(out, StatusBadRequest, "empty request")
		return errors.New("empty request")
	}

	u, err := url.Parse(rawRequest)
	if err != nil {
		sendStatus(out, StatusBadRequest, "invalid request uri")
		return err
	}

	r.Host = u.Hostname()
	r.Path = u.Path

	if u.Scheme != "gemini" {
		sendStatus(out, StatusBadRequest, "invalid or missing scheme in request uri")
		return errors.New("invalid or missing scheme in request uri")
	}

	if r.Host == "" {
		sendStatus(out, StatusBadRequest, "missing hostname in request uri")
		return errors.New("missing hostname in request uri")
	}

	r.Query, err = url.QueryUnescape(u.RawQuery)
	if err != nil {
		sendStatus(out, StatusBadRequest, "invalid query string in request uri")
		return errors.New("invalid query string in request uri")
	}

	return nil
}
