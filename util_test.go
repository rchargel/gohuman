package main

import "testing"

func TestNewID(t *testing.T) {
	l := 100000
	m := make(map[string]bool, l)
	for i := 0; i < l; i++ {
		str := NewID("this", "is", "a", "test")
		if _, found := m[str]; found {
			t.Errorf("%v was already used.\n", str)
		} else {
			m[str] = true
		}
	}
}

func BenchmarkNewID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewID("this", "is", "a", "test")
	}
}
