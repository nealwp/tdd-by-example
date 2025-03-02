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

func (b Money) GetAmount() int {
	return b.amount
}

func (b Money) GetCurrency() string {
	return b.currency
}

func (b Money) Equals(other IMoney) bool {
	return b.amount == other.GetAmount() && b.currency == other.GetCurrency()
}

func (b Money) Times(multiplier int) IMoney {
	return Money{b.amount * multiplier, b.currency}
}
