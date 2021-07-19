package gollies

import (
	"strings"
)

type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

type Compare = func(x, y interface{}) Order

func CompareInt(x, y interface{}) Order {
	xv, yv := x.(int), y.(int)
	
	if xv < yv {
		return Lt
	}

	if xv > yv {
		return Gt
	}

	return Eq	
}

func CompareString(x, y interface{}) Order {
	return Order(strings.Compare(x.(string), y.(string)))
}
