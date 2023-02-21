package gogemini

import "io"

type ResponseWriter interface {
	ReadFrom(r io.Reader) (n int64, err error)
	Write(p []byte) (n int, err error)
	WriteStatus(c StatusCode, meta string)
}

type responseWriter struct {
	out io.Writer
}

func (w responseWriter) ReadFrom(r io.Reader) (n int64, err error) {
	return io.Copy(w.out, r)
}

func (w responseWriter) Write(p []byte) (n int, err error) {
	return w.out.Write(p)
}

func (w responseWriter) WriteStatus(code StatusCode, meta string) {
	sendStatus(w, code, meta)
}
