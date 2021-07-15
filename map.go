package gollies

type Map interface {
	Add(key interface{}, val interface{}) interface{}
	Remove(key interface{}) interface{}

	Find(key interface{}) interface{}
	Each(func (key, val interface {}) bool) bool

	Keys() []interface{}
	Values() []interface{}

	Len() int

	AddAll(y Map)
	KeepAll(y Map)
}

func AddAll(x Map, y Map) {
	y.Each(func(k, v interface{}) bool {
		x.Add(k, v)
		return true
	})
}
