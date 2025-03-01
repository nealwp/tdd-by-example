package main

type Money interface {
	Equals(other Money) bool
	GetAmount() int
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
	return b.amount == other.GetAmount()
}

type Dollar struct{ BaseMoney }

func (d Dollar) Times(multiplier int) Money {
	return Dollar{BaseMoney{d.amount * multiplier, d.currency}}
}

type Franc struct{ BaseMoney }

func (f Franc) Times(multiplier int) Money {
	return Franc{BaseMoney{f.amount * multiplier, f.currency}}
}
