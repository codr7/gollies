package gollies

type SliceMapItem = struct {
	key, value interface{}
}

type SliceMap struct {
	compare Compare
	items []SliceMapItem
}

func NewSliceMap(cmp Compare) *SliceMap {
	return new(SliceMap).Init(cmp)
}

func (m *SliceMap) Init(cmp Compare) *SliceMap {
	m.compare = cmp
	return m
}

func (m SliceMap) Index(key interface{}) (int, interface{}) {
	min, max := 0, len(m.items)

	for min < max {
		i := (min+max)/2
		it := m.items[i]
		
		switch m.compare(key, it.key) {
		case Lt:
			max = i
		case Eq:
			return i, it.value
		case Gt:
			min = i+1
		}
	}
	
	return min, nil
}

func (m SliceMap) Find(key interface{}) interface{} {
	_, found := m.Index(key)
	return found
}

func (m SliceMap) Each(pred func (key, val interface {}) bool) bool {
	for _, it := range m.items {
		if !pred(it.key, it.value) {
			return false
		}
	}

	return true
}

func (m *SliceMap) Add(key interface{}, val interface{}) interface{} {
	i, found := m.Index(key)

	if found != nil {
		return found
	}

	m.items = append(m.items, SliceMapItem{})
	copy(m.items[i+1:], m.items[i:])
	m.items[i] = SliceMapItem{key, val}
	return nil
}

func (m *SliceMap) Remove(key interface{}) interface{} {
	i, found := m.Index(key)

	if found != nil {
		m.items = m.items[:i+copy(m.items[i:], m.items[i+1:])]
	}

	return found
}

func (m SliceMap) Keys() []interface{} {
	out := make([]interface{}, len(m.items))

	for i, it := range m.items {
		out[i] = it.key
	}

	return out
}

func (m SliceMap) Values() []interface{} {
	out := make([]interface{}, len(m.items))

	for i, it := range m.items {
		out[i] = it.value
	}

	return out
}

func (m SliceMap) Len() int {
	return len(m.items)
}

func (m *SliceMap) AddAll(y Map) {
	AddAll(m, y)
}

func (m *SliceMap) KeepAll(y Map) {
	newLen := len(m.items)
	keep := make([]bool, newLen)
	
	for i, it := range m.items {
		found := y.Find(it.key) != nil
		keep[i] = found

		if !found {
			newLen--
		}
	}

	newItems := make([]SliceMapItem, newLen)
	i := 0;
	
	for j, it := range m.items {
		if keep[j] {
			newItems[i] = it
			i++
		}
	}

	m.items = newItems
}

func (m *SliceMap) RemoveAll(y Map) {
	RemoveAll(m, y)
}
