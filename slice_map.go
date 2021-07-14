package gollies

type SliceMapItem = struct {
	Key, Value interface{}
}

type SliceMap struct {
	Compare Compare
	Items []SliceMapItem
}

func NewSliceMap(cmp Compare) *SliceMap {
	return &SliceMap{Compare: cmp}
}

func (m *SliceMap) Find(key interface{}) (int, interface{}) {
	min, max := 0, m.Len()

	for min < max {
		i := (min+max)/2
		it := m.Items[i]
		
		switch m.Compare(key, it.Key, nil) {
		case Lt:
			max = i
		case Eq:
			return i, it.Value
		case Gt:
			min = i+1
		}
	}
	
	return min, nil
}

func (m *SliceMap) Add(key interface{}, val interface{}) interface{} {
	i, found := m.Find(key)

	if found != nil {
		return found
	}

	m.Items = append(m.Items, SliceMapItem{})
	copy(m.Items[i+1:], m.Items[i:])
	m.Items[i] = SliceMapItem{key, val}
	return nil
}

func (m *SliceMap) Remove(key interface{}) interface{} {
	i, found := m.Find(key)

	if found != nil {
		m.Items = m.Items[:i+copy(m.Items[i:], m.Items[i+1:])]
	}

	return found
}

func (m SliceMap) Len() int {
	return len(m.Items)
}
