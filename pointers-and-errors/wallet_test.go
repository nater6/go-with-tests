package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err :=  wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Fatalf("Expected <nil> got %v", err)
		}
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)

		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		got := wallet.Withdraw(Bitcoin(100))
		want := ErrInsufficientFunds
		

		if got == nil {
			t.Fatal("expected an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}


