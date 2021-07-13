package gollies

type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

type Compare = func(x, y, arg interface{}) Order

func CompareInt(x, y, arg interface{}) Order {
	xv, yv := x.(int), y.(int)
	
	if xv < yv {
		return Lt
	}

	if xv > yv {
		return Gt
	}

	return Eq
	
}
