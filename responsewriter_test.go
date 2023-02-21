package gogemini

import (
	"bytes"
	"strings"
	"testing"
)

func TestResponseWriterWrite(t *testing.T) {

	buffer := new(strings.Builder)
	writer := responseWriter{buffer}

	data := []byte("abc")

	n, err := writer.Write(data)

	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatal("Error writing data")
	}

	if buffer.String() != "abc" {
		t.Fatal("Received incorrect data")
	}
}

func TestResponseWriterReadFrom(t *testing.T) {

	buffer := new(strings.Builder)
	writer := responseWriter{buffer}

	data := []byte("abc")

	n, err := writer.ReadFrom(bytes.NewBuffer(data))

	if err != nil {
		t.Fatal(err)
	}

	if n != int64(len(data)) {
		t.Fatal("Error writing data")
	}

	if buffer.String() != "abc" {
		t.Fatal("Received incorrect data")
	}
}

func TestResponseWriterWriteStatus(t *testing.T) {

	buffer := new(strings.Builder)
	writer := responseWriter{buffer}

	writer.WriteStatus(StatusSlowDown, "meta")

	if buffer.String() != "44 meta\n" {
		t.Fatal("Received incorrect data")
	}
}
