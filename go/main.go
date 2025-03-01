package main

type IMoney interface {
	Equals(other IMoney) bool
	GetAmount() int
}

type BaseMoney struct{ amount int }

func (b BaseMoney) GetAmount() int {
	return b.amount
}

func (b BaseMoney) Times(multiplier int) BaseMoney {
	return BaseMoney{b.amount * multiplier}
}

type Money struct{}

func (b Money) Dollar(amount int) Dollar {
	return Dollar{BaseMoney{amount: amount}}
}

func (b BaseMoney) Equals(other IMoney) bool {
	return b.amount == other.GetAmount()
}

type Dollar struct{ BaseMoney }

func (d Dollar) Times(multiplier int) IMoney {
	return Dollar{BaseMoney{d.amount * multiplier}}
}

type Franc struct{ BaseMoney }

func (f Franc) Times(multiplier int) IMoney {
	return Franc{BaseMoney{f.amount * multiplier}}
}
