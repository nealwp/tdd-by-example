package main

// any expression has a reduce method
type Expression interface {
	Reduce(bank *Bank, to string) Money
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

// operations return an expression
func (m Money) Plus(addend Money) Expression {
	return Sum{Augend: m, Addend: addend}
}

func (Money) Dollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func (Money) Franc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}

func (m Money) Reduce(bank *Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{amount: m.amount / rate, currency: to}
}

type Bank struct {
	rates map[pair]int
}

type pair struct {
	from string
	to   string
}

func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
	if b.rates == nil {
		b.rates = make(map[pair]int)
	}
	b.rates[pair{from, to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[pair{from, to}]
}

type Sum struct {
	Augend Money
	Addend Money
}

func (s Sum) Reduce(bank *Bank, to string) Money {
	amount := s.Augend.Reduce(bank, to).amount + s.Addend.Reduce(bank, to).amount
	return Money{amount: amount, currency: to}
}
