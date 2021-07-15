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

func (m HashMap) Find(key interface{}) interface{} {
	return m.items[key]
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
	out := make([]interface{}, m.Len())
	i := 0
	
	for k, _ := range m.items {
		out[i] = k
		i++
	}

	return out
}

func (m HashMap) Values() []interface{} {
	out := make([]interface{}, m.Len())
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
