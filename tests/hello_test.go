package main

import (
	"sandbox/services"
	"testing"
)

func TestHello(t *testing.T) {
	got := services.Hello("Chris")
	want := "Hello Chris!"

	if got != want {
		t.Fail()
		t.Errorf("got %q want %q", got, want)
	}
}
