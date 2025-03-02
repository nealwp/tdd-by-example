package main

type IMoney interface {
	Equals(other IMoney) bool
	GetAmount() int
	GetCurrency() string
	Times(mulitplier int) IMoney
	Plus(addend Money) IMoney
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

func (m Money) Plus(addend Money) Money {
	return Money{m.amount + addend.amount, m.currency}
}

func (Money) Dollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func (Money) Franc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}
