package pointers

import (
	"testing"
)



func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, err error) {
    t.Helper()
    if err == nil {
        t.Error("wanted an error but didn't get one")
    }
}

	assertBalance := func (t *testing.T,wallet Wallet,bitcoin Bitcoin)  {
		t.Helper()
		got := Bitcoin(wallet.balance)
		
		if got != bitcoin {
			t.Errorf("%#v got %s but expected %s",wallet,got,bitcoin)
		}
	}

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

		assertBalance(t,wallet,Bitcoin(10))    
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
    })

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertError(t, err)
    assertBalance(t, wallet, startingBalance)
})
}