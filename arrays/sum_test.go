package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		sliceOne := []int{1, 2, 3} // sum 6
		sliceTwo := []int{1, 10}   //sum 11

		got := SumAll(sliceOne, sliceTwo)
		want := []int{6, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d, given %v and %v", got, want, sliceOne, sliceTwo)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("sum tails of slices", func(t *testing.T) {
		sliceOne := []int{1, 2, 8} // 10
		sliceTwo := []int{1, 10}   // 10

		got := SumAllTails(sliceOne, sliceTwo)
		want := []int{10, 10}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		sliceOne := []int{1, 2, 8} // 10
		sliceTwo := []int{}        // 0

		got := SumAllTails(sliceOne, sliceTwo)
		want := []int{10, 0}

		checkSums(t, got, want)
	})
}
