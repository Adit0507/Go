package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Adi")

	got := buffer.String()
	want := "Hello Adi"

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
