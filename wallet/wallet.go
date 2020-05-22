package wallet

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit allows to deposit amount to wallet
func (w *Wallet) Deposit(amount int) {
	// fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	(*w).balance += amount
	// we can just write w.balance
	//struct pointers are automatically dereferenced

}

// Balance allows to get value from wallet
func (w *Wallet) Balance() int {
	return (*w).balance
}
