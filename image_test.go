package gohuman

import "testing"

func TestRandomize(t *testing.T) {
	var last []int
	for i := 0; i < 1000; i++ {
		p := randomize(18, 9)
		if len(p) != 9 {
			t.Errorf("Should have only been %d values in the list", len(p))
		}
		if isSliceEq(p, last) {
			t.Errorf("Repeated last values %v on run %d", p, i)
		}
		last = p
	}
}

func TestGetRandomImageSet(t *testing.T) {
	for i := 0; i < 1000; i++ {
		p, err := ImageMapper.getRandomImageSet(9)
		if err != nil {
			t.Error("Should not have thrown error:", err)
		}
		if len(p) != 9 {
			t.Errorf("Should have only been %d values in the list", len(p))
		}
		set := make(map[string]bool, 9)
		for v := range p {
			img := p[v]
			if _, found := set[img.Title]; found {
				t.Errorf("Already had %s.", img.Title)
			}
			set[img.Title] = true
		}
	}
}

func TestGetRandomImageSetFail(t *testing.T) {
	_, err := ImageMapper.getRandomImageSet(len(ImageMapper.Images) + 1)
	if err == nil {
		t.Error("Should have thrown error")
	}
}

func BenchmarkRandomize1009(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomize(100, 9)
	}
}

func BenchmarkRandomize9999(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomize(999, 9)
	}
}

func BenchmarkRandomize189(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomize(18, 9)
	}
}

func isSliceEq(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
