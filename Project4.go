package main

import "fmt"

// BankAccount struct
type BankAccount struct {
	Owner   string
	Balance float64
}

// Value receiver 
func (b BankAccount) DisplayBalance() {
	fmt.Printf("Owner: %s, Balance: %.2f\n", b.Owner, b.Balance)
}

// Pointer receiver 
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited %.2f to %s's account\n", amount, b.Owner)
	}
}

// Pointer receiver 
func (b *BankAccount) Withdraw(amount float64) {
	if amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrew %.2f from %s's account\n", amount, b.Owner)
	} else {
		fmt.Println("Insufficient funds for withdrawal")
	}
}

func main() {

	// Create account as value
	account := BankAccount{Owner: "Alice", Balance: 100.0}

	account.DisplayBalance() 

	// Try deposit using pointer receiver
	account.Deposit(50.0)       // will modify actual balance
	account.DisplayBalance()    // Should show 150.0

	// Try withdraw
	account.Withdraw(30.0)      // 150 - 30
	account.DisplayBalance()    // Should show 120.0

	// Show that using value receiver won't change state
	copyAccount := account
	copyAccount.Deposit(100.0)  // Doesn't affect original
	copyAccount.DisplayBalance()
	account.DisplayBalance()    // Original still 120.0
}
