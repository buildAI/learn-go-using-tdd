package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ellodie", "Spanish")
		want := "Hola, Ellodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In french", func(t *testing.T) {
		got := Hello("foo", "French")
		want := "Bonjour, foo"
		assertCorrectMessage(t, got, want)
	})

}
