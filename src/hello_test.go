package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {

		got := Hello("kassem")
		want := "Hello, kassem"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v want: %v", got, want)
	}
}
