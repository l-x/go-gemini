package gogemini

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var testHandlerFunc = func(w ResponseWriter, r *Request) {
	io.WriteString(w, fmt.Sprintf("%s%s?%s", r.Host, r.Path, r.Query))
}

func TestServeSucceeds(t *testing.T) {
	in := strings.NewReader("gemini://host/path?query")
	out := new(strings.Builder)

	if err := Serve(in, out, funcHandler{testHandlerFunc}); err != nil {
		t.Fatal(err)
	}

	if out.String() != "host/path?query" {
		t.Fatal("Handler received wrong request " + out.String())
	}
}

func TestServeSucceedsForExistingResponseWriter(t *testing.T) {
	in := strings.NewReader("gemini://host/path?query")

	buffer := new(strings.Builder)
	out := responseWriter{buffer}

	if err := Serve(in, out, funcHandler{testHandlerFunc}); err != nil {
		t.Fatal(err)
	}

	if buffer.String() != "host/path?query" {
		t.Fatal("Handler received wrong request")
	}
}

func TestInvalidRequests(t *testing.T) {
	testHandler := funcHandler{testHandlerFunc}

	assert := func(request, resp, errmsg string) func(t *testing.T) {
		in := strings.NewReader(request)
		out := new(strings.Builder)

		return func(t *testing.T) {
			err := Serve(in, out, testHandler)
			if err == nil {
				t.Fatal("Serve did not return an error")
			}

			if err.Error() != errmsg {
				t.Fatalf("Serve did not return the expected error. Expected: '%s'. Actual: '%s'", errmsg, err.Error())
			}

			if !strings.HasPrefix(out.String(), resp) {
				t.Fatalf("Serve did not send the expected result. Expected: '%s'. Actual: '%s'", resp, out.String())
			}
		}
	}

	t.Run("exceeding size", assert(strings.Repeat("x", maxRequestSize*2), "59", "Request exceeds the limit of 1024 bytes"))
	t.Run("empty request", assert("", "59", "empty request"))
	t.Run("invalid url", assert(string(byte(0x00)), "59", "parse \"\\x00\": net/url: invalid control character in URL"))
	t.Run("missing scheme", assert("fu.bar/herp?derp", "59", "invalid or missing scheme in request uri"))
	t.Run("missing hostname", assert("gemini:///herp?derp", "59", "missing hostname in request uri"))
	t.Run("invalid query", assert("gemini://fu.bar/herp?d%x", "59", "invalid query string in request uri"))
}
