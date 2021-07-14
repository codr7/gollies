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
var items [nItems]int

func bench(b *testing.B, mk func() Map) {		
	for i := 0; i < b.N; i++ {
		m := mk()
		
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
}

func BenchmarkSliceMap(b *testing.B) {
	for i := 0; i < nItems; i++ {
		items[i] = i;
	}

	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	
	b.Run("SliceMap", func(b *testing.B) {
		bench(b, func() Map {
			return NewSliceMap(CompareInt)
		})
	})

	b.Run("HashMap", func(b *testing.B) {
		bench(b, func() Map {
			return NewHashMap()
		})
	})
}
