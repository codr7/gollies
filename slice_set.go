package gollies

type SliceSetItem = struct {
	Key, Value interface{}
}

type SliceSet struct {
	Compare Compare
	Items []SliceSetItem
}

func NewSliceSet(cmp Compare) *SliceSet {
	return &SliceSet{Compare: cmp}
}

func (s *SliceSet) Find(key interface{}) (int, interface{}) {
	min, max := 0, s.Len()

	for min < max {
		i := (min+max)/2
		it := s.Items[i]
		
		switch s.Compare(key, it.Key, nil) {
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

func (s *SliceSet) Add(key interface{}, val interface{}) interface{} {
	i, found := s.Find(key)

	if found != nil {
		return found
	}

	s.Items = append(s.Items, SliceSetItem{})
	copy(s.Items[i+1:], s.Items[i:])
	s.Items[i] = SliceSetItem{key, val}
	return nil
}

func (s *SliceSet) Remove(key interface{}) interface{} {
	i, found := s.Find(key)

	if found != nil {
		s.Items = s.Items[:i+copy(s.Items[i:], s.Items[i+1:])]
	}

	return found
}

func (s SliceSet) Len() int {
	return len(s.Items)
}
