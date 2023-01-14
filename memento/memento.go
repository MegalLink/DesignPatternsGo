package memento

import "fmt"

type BankAccount struct {
	balance int
	history []*Memento
	current int
}

// memento is used ti rill back states arbitrarily
// its no requred to expose directly the states to witch it reverts the system
// can be used to implement undo/redo

type Memento struct {
	balance int // preserves same data as system in this case balance
}

func NewBankAccount(amount int) *BankAccount {
	b := &BankAccount{balance: amount}
	b.history = append(b.history, &Memento{amount})
	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance, " current = ", b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{b.balance}
	b.history = append(b.history, m)
	b.current++
	return m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.balance
		b.history = append(b.history, m)
		b.current = len(b.history) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.history[b.current]
		b.balance = m.balance
		return m
	}

	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.history) {
		b.current++
		m := b.history[b.current]
		b.balance = m.balance
		return m
	}

	return nil
}
