package pointers

import (
	"fmt"
	"testing"
)



func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(10)
	
	got := wallet.Balance()
	want := 10
	fmt.Println("th3 add is = ", &wallet.balance)

	if got != want { 
		t.Errorf("%#v got %d but expected %d",wallet,got,want)
	}


}