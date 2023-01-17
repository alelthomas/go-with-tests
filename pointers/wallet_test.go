package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Pikadollars) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Pikadollars(10))

		assertBalance(t, wallet, Pikadollars(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Pikadollars(20)}
		err := wallet.Withdraw(Pikadollars(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Pikadollars(10))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{Pikadollars(20)}
		err := wallet.Withdraw(Pikadollars(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Pikadollars(20))
	})
}
