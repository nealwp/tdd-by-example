package main

type Dollar struct{ amount int }

func (d *Dollar) Times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d *Dollar) Equals(dollar Dollar) bool {
	return true
}
