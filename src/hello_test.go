package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("kassem")
	want := "Hello , kassem!"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
