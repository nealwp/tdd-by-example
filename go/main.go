package main

type Dollar struct{ amount int }

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d Dollar) Equals(dollar Dollar) bool {
	return d.amount == dollar.amount
}

type Franc struct{ amount int }

func (d Franc) Times(multiplier int) Franc {
	return Franc{d.amount * multiplier}
}

func (d Franc) Equals(dollar Franc) bool {
	return d.amount == dollar.amount
}
