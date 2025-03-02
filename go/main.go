package main

type IMoney interface {
	Equals(other IMoney) bool
	GetAmount() int
	GetCurrency() string
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

func (m Money) Equals(other IMoney) bool {
	return m.amount == other.GetAmount() && m.currency == other.GetCurrency()
}

func (m Money) Times(multiplier int) IMoney {
	return Money{m.amount * multiplier, m.currency}
}
