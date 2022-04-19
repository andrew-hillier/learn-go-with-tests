package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andrew")

	got := buffer.String()
	want := "Hello, Andrew"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
