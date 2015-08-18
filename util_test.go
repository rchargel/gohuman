package gohuman

import "testing"

func TestNewCaptchaID(t *testing.T) {
	l := 100000
	m := make(map[string]bool, l)
	for i := 0; i < l; i++ {
		str := newCaptchaID("this", "is", "a", "test")
		if _, found := m[str]; found {
			t.Errorf("%v was already used.\n", str)
		} else {
			m[str] = true
		}
	}
}

func BenchmarkNewCaptchaID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newCaptchaID("this", "is", "a", "test")
	}
}
