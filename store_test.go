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

func TestPurgeOld1(t *testing.T) {
	s := newStore()
	s.addCaptcha(Captcha{ID: "TESTING"})

	if len(s.items) == 0 {
		t.Error("There were never any items in the store")
	}
	t.Logf("Items found: %d\n", len(s.items))
	s.purgeOld(-1, -1)
	if len(s.items) == 1 {
		t.Errorf("Items were not purged: %d\n", len(s.items))
	} else {
		t.Logf("No items left after purge")
	}
}

func TestPurgeOld2(t *testing.T) {
	s := newStore()
	s.addCaptcha(Captcha{ID: "TESTING"})

	if len(s.items) == 0 {
		t.Error("There were never any items in the store")
	}
	t.Logf("Items found: %d\n", len(s.items))
	s.purgeOld(100, -1)
	if len(s.items) == 0 {
		t.Errorf("Items were purged: %d\n", len(s.items))
	} else {
		t.Logf("Items left after purge: %d\n", len(s.items))
	}
}

func BenchmarkAddCaptcha(b *testing.B) {
	s := newStore()
	for n := 0; n < b.N; n++ {
		s.addCaptcha(Captcha{ID: fmt.Sprintf("CAPTCHA_%d", n)})
	}
}
