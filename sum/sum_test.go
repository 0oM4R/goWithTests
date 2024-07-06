package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size of numbers", func(t *testing.T) {
		number := []int{1, 3, 5, 5}
		got := Sum(number)
		want := 14
		if got != want {
			t.Errorf("got %d, want %d, %v", got, want, number)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("get sum of to slices", func(t *testing.T) {
		got := SumAll([]int{4, 2}, []int{1, 3})
		want := []int{6, 4}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
