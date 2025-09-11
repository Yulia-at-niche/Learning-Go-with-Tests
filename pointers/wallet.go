package pointers

type Wallet struct {
	balance int // Lowercase letters make something private to the package e.g 'b' in 'balance'.
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
func (w *Wallet) Balance() int {
	return w.balance
}
