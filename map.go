package gollies

type Map interface {
	Add(key interface{}, val interface{}) interface{}
	Remove(key interface{}) interface{}

	Find(key interface{}) interface{}
	Each(func (key, val interface {}) bool) bool

	Keys() []interface{}
	Values() []interface{}

	Len() int
}
