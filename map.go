package gollies

type Map interface {
	Clone() Map

	Add(key interface{}, val interface{}) interface{}
	Remove(key interface{}) interface{}

	Find(key interface{}) interface{}
	Each(func (key, val interface {}) bool) bool

	Keys() []interface{}
	Values() []interface{}

	Len() int

	AddAll(src Map)
	KeepAll(src Map)
	RemoveAll(src Map)

	Difference(rhs Map) Map
	Intersection(rhs Map) Map
	Union(rhs Map) Map
}

func AddAll(dst Map, src Map) {
	src.Each(func(k, v interface{}) bool {
		dst.Add(k, v)
		return true
	})
}

func KeepAll(dst Map, src Map) {
	var rem []interface{}
	
	dst.Each(func(k, _ interface{}) bool {
		if src.Find(k) == nil {
			rem = append(rem, k)
		}
		
		return true
	})

	for _, k := range rem {
		dst.Remove(k)
	}
}

func RemoveAll(dst Map, src Map) {
	src.Each(func(k, _ interface{}) bool {
		dst.Remove(k)
		return true
	})
}

func Difference(lhs, rhs Map) Map {
	dst := lhs.Clone()
	dst.RemoveAll(rhs)
	return dst
}

func Intersection(lhs, rhs Map) Map {
	dst := lhs.Clone()
	dst.KeepAll(rhs)
	return dst
}

func Union(lhs, rhs Map) Map {
	dst := lhs.Clone()
	dst.AddAll(rhs)
	return dst
}
