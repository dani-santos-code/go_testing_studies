package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := &Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		// fmt.Printf("address of balance in test is %v\n", &wallet.balance)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := &Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

}
