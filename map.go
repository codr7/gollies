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
	RemoveAll(y Map)
}

func AddAll(x Map, y Map) {
	y.Each(func(k, v interface{}) bool {
		x.Add(k, v)
		return true
	})
}

func KeepAll(x Map, y Map) {
	var rem []interface{}
	
	x.Each(func(k, _ interface{}) bool {
		if y.Find(k) == nil {
			rem = append(rem, k)
		}
		
		return true
	})

	for _, k := range rem {
		x.Remove(k)
	}
}

func RemoveAll(x Map, y Map) {
	y.Each(func(k, _ interface{}) bool {
		x.Remove(k)
		return true
	})
}
