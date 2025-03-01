package main

type Money interface {
	Equals(other Money) bool
	GetAmount() int
}

type BaseMoney struct{ amount int }

func (b BaseMoney) GetAmount() int {
	return b.amount
}

func (b BaseMoney) Times(multiplier int) BaseMoney {
	return BaseMoney{b.amount * multiplier}
}

func (b BaseMoney) Equals(other Money) bool {
	return b.amount == other.GetAmount()
}

type Dollar struct{ BaseMoney }

func (d Dollar) Times(multiplier int) Money {
	return Dollar{BaseMoney{d.amount * multiplier}}
}

type Franc struct{ BaseMoney }

func (f Franc) Times(multiplier int) Money {
	return Franc{BaseMoney{f.amount * multiplier}}
}
