package main

type Dollar struct{ amount int }

func (d *Dollar) times(multiplier int) {
	d.amount = 10
}
