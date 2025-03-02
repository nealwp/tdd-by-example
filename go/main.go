package main

// any expression has a reduce method
type Expression interface {
	Reduce(bank *Bank, to string) Money
}

type Money struct {
	amount   int
	currency string
}

func (m Money) GetAmount() int {
	return m.amount
}

func (m Money) GetCurrency() string {
	return m.currency
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.GetAmount() && m.currency == other.GetCurrency()
}

func (m Money) Times(multiplier int) Money {
	return Money{m.amount * multiplier, m.currency}
}

// operations return an expression
func (m Money) Plus(addend Money) Expression {
	return Sum{Augend: m, Addend: addend}
}

func (Money) Dollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func (Money) Franc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}

func (m Money) Reduce(bank *Bank, to string) Money {
	var rate int
	if m.currency == "CHF" && to == "USD" {
		rate = 2
	} else {
		rate = 1
	}
	return Money{amount: m.amount / rate, currency: to}
}

type Bank struct {
}

func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
}

type Sum struct {
	Augend Money
	Addend Money
}

func (s Sum) Reduce(bank *Bank, to string) Money {
	amount := s.Augend.amount + s.Addend.amount
	return Money{amount: amount, currency: to}
}
