package gollies

type HashMapItems = map[interface{}]interface{}

type HashMap struct {
	items HashMapItems
}

func NewHashMap() *HashMap {
	return new(HashMap).Init()
}

func (m *HashMap) Init() *HashMap {
	m.items = make(HashMapItems)
	return m
}

func (m HashMap) Clone() Map {
	dst := NewHashMap()
	
	for k, v := range m.items {
		dst.items[k] = v
	}

	return dst
}

func (m HashMap) Find(key interface{}) interface{} {
	return m.items[key]
}

func (m HashMap) Each(pred func (key, val interface {}) bool) bool {
	for k, v := range m.items {
		if !pred(k, v) {
			return false
		}
	}

	return true
}

func (m *HashMap) Add(key interface{}, val interface{}) interface{} {
	prev := m.items[key]
	m.items[key] = val
	return prev
}

func (m *HashMap) Remove(key interface{}) interface{} {
	val := m.items[key]
	delete(m.items, key)
	return val
}

func (m HashMap) Keys() []interface{} {
	out := make([]interface{}, len(m.items))
	i := 0
	
	for k, _ := range m.items {
		out[i] = k
		i++
	}

	return out
}

func (m HashMap) Values() []interface{} {
	out := make([]interface{}, len(m.items))
	i := 0
	
	for _, v := range m.items {
		out[i] = v
		i++
	}

	return out
}

func (m HashMap) Len() int {
	return len(m.items)
}

func (m *HashMap) AddAll(y Map) {
	AddAll(m, y)
}

func (m *HashMap) KeepAll(y Map) {
	var drop []interface{}
	
	for k, _ := range m.items {
		if y.Find(k) == nil {
			drop = append(drop, k)
		}
	}

	for _, k := range drop {
		delete(m.items, k)
	}
}

func (m *HashMap) RemoveAll(y Map) {
	RemoveAll(m, y)
}

func (m *HashMap) Difference(rhs Map) Map {
	return Difference(m, rhs)
}

func (m *HashMap) Intersection(rhs Map) Map {
	return Intersection(m, rhs)
}	

func (m *HashMap) Union(rhs Map) Map {
	return Union(m, rhs)
}
