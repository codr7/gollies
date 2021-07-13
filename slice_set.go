package gollies

type SliceSetItem = struct {
     Key, Value interface{}
}

type SliceSet struct {
     Items []SliceSetItem
}

func NewSliceSet() *SliceSet {
     return &SliceSet{}
}

func (s *SliceSet) Find(key interface{}) (int, interface{}) {
     //TODO Binary search
     return 0, nil
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