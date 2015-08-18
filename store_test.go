package gohuman

import (
	"fmt"
	"testing"
)

func TestGetCaptcha1(t *testing.T) {
	s := newStore()
	s.addCaptcha(Captcha{ID: "TESTING"})
	s.addCaptcha(Captcha{ID: "TESTED"})
	c, err := s.getCaptcha("TESTING")
	if err != nil {
		t.Error("Error when retrieving captcha", err)
	}
	if c.ID != "TESTING" {
		t.Errorf("Found the wrong captcha %v\n", c.ID)
	}
}

func TestGetCaptcha2(t *testing.T) {
	s := newStore()
	s.addCaptcha(Captcha{ID: "TESTING"})
	s.addCaptcha(Captcha{ID: "TESTED"})
	c, err := s.getCaptcha("TESTERS")
	if err == nil {
		t.Error("No error when retrieving captcha", err)
	}
	if len(c.ID) != 0 {
		t.Errorf("Found a valid captcha %v\n", c.ID)
	}
}

func TestPurgeStore(t *testing.T) {
	s := newStore()
	for i := 0; i < 1000; i++ {
		s.addCaptcha(Captcha{ID: fmt.Sprintf("CAPTCHA_%d", i)})
	}

	s.purgeOld(-1, -1)
	s.waitForPurgeToComplete()
	if len(s.items) != 0 {
		t.Errorf("Still items left in store: %d", len(s.items))
	}
}

func BenchmarkAddCaptcha(b *testing.B) {
	s := newStore()
	for n := 0; n < b.N; n++ {
		s.addCaptcha(Captcha{ID: fmt.Sprintf("CAPTCHA_%d", n)})
	}
}
