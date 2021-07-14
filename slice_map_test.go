package gollies

import (
	"math/rand"
	"testing"
)

func TestSliceMapBasics(t *testing.T) {
	m := NewSliceMap(CompareInt)

	if l := m.Len(); l != 0 {
		t.Errorf("Wrong initial Len(): %v", l)
	}

	m.Add(2, 2)
	m.Add(1, 1)
	m.Add(3, 3)

	if l := m.Len(); l != 3 {
		t.Errorf("Wrong Len() after Add(): %v", l)
	}

	m.Remove(2)

	if l := m.Len(); l != 2 {
		t.Errorf("Wrong Len() after Remove(): %v", l)
	}
}

const nItems = 1000

func BenchmarkSliceMap(b *testing.B) {
	var items [nItems]int

	for i := 0; i < nItems; i++ {
		items[i] = i;
	}

	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	
	b.Run("SliceMap", func(b *testing.B) {		
		for i := 0; i < b.N; i++ {
			var m SliceMap
			m.Init(CompareInt)

			for _, v := range items {
				m.Add(v, v)
			}

			for _, v := range items {
				if m.Find(v) != v {
				}
			}

			for _, v := range items {
				m.Remove(v)
			}
		}
	})
	
	b.Run("Map", func(b *testing.B) {		
		for i := 0; i < b.N; i++ {
			m := make(map[int]int)

			for _, v := range items {
				m[v] = v
			}

			for _, v := range items {
				if m[v] != v {
				}
			}

			for _, v := range items {
				delete(m, v)
			}
		}
	})
}
