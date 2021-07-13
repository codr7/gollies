package gollies

type Order = int

const (
  Less	= Order(-1)
  Equal = Order(0)
  Greater = Order(1)
)

type Sortable interface {
     Compare() Order
}