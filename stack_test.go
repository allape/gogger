package gogger

import (
	"io"
	"testing"
)

func TestTrace(t *testing.T) {
	err := Trace(nil)
	if err != nil {
		t.Error("Expected nil, got", err)
	}

	err = Trace(io.EOF)
	if err == nil {
		t.Error("Expected io.EOF, got nil")
	}
	New("test").Warn().Println(err)
}
