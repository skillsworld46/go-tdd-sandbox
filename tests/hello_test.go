package main

import (
	"sandbox/services"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Test 1 - with name", func(t *testing.T) {
		got := services.Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 2 - without name", func(t *testing.T) {
		got := services.Hello("","French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 3 - name with spanish", func(t *testing.T) {
		got := services.Hello("Elodie","Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fail()
		t.Errorf("got %q want %q", got, want)
	}
}
