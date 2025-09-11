package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// To keep the struct's fields private and safe from outside influence we're not setting any values here, nor are we accessing them directly.
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
