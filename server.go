package gogemini

import (
	"io"
)

func Serve(in io.Reader, out io.Writer, h Handler) error {
	var r *Request = &Request{}
	var w ResponseWriter

	switch t := out.(type) {
	case ResponseWriter:
		w = t
	default:
		w = responseWriter{t}
	}

	if err := parseRequest(in, w, r); err != nil {
		return err
	}

	h.ServeGemini(w, r)

	return nil
}
