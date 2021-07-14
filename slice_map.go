package gollies

type SliceMapItem = struct {
	Key, Value interface{}
}

type SliceMap struct {
	Compare Compare
	Items []SliceMapItem
}

func NewSliceMap(cmp Compare) *SliceMap {
	return new(SliceMap).Init(cmp)
}

func (m *SliceMap) Init(cmp Compare) *SliceMap {
	m.Compare = cmp
	return m
}

func (m *SliceMap) Index(key interface{}) (int, interface{}) {
	min, max := 0, m.Len()

	for min < max {
		i := (min+max)/2
		it := m.Items[i]
		
		switch m.Compare(key, it.Key) {
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

func (m *SliceMap) Find(key interface{}) interface{} {
	_, found := m.Index(key)
	return found
}

func (m *SliceMap) Add(key interface{}, val interface{}) interface{} {
	i, found := m.Index(key)

	if found != nil {
		return found
	}

	m.Items = append(m.Items, SliceMapItem{})
	copy(m.Items[i+1:], m.Items[i:])
	m.Items[i] = SliceMapItem{key, val}
	return nil
}

func (m *SliceMap) Remove(key interface{}) interface{} {
	i, found := m.Index(key)

	if found != nil {
		m.Items = m.Items[:i+copy(m.Items[i:], m.Items[i+1:])]
	}

	return found
}

func (m SliceMap) Len() int {
	return len(m.Items)
}
