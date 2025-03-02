package main

type Money interface {
	Equals(other Money) bool
	GetAmount() int
	GetCurrency() string
}

type BaseMoney struct {
	amount   int
	currency string
}

func (b BaseMoney) GetAmount() int {
	return b.amount
}

func (b BaseMoney) GetCurrency() string {
	return b.currency
}

func (b BaseMoney) Equals(other Money) bool {
	return b.amount == other.GetAmount() && b.currency == other.GetCurrency()
}

func (b BaseMoney) Times(multiplier int) Money {
	return BaseMoney{b.amount * multiplier, b.currency}
}
