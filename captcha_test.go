package gohuman

import "testing"

func TestCreateCaptchaRequest(t *testing.T) {
	r, _ := CreateCaptchaRequest("localhost", 3, 3)
	if len(r.ID) == 0 {
		t.Error("No ID was generated")
	}
}

func TestCreateCaptchaRequestFail1(t *testing.T) {
	r, err := CreateCaptchaRequest("localhost", 0, 2)
	if len(r.ID) != 0 {
		t.Error("Created an ID, not needed")
	}
	if err == nil {
		t.Error("Should have an error")
	}
}

func TestCreateCaptchaRequestFail2(t *testing.T) {
	r, err := CreateCaptchaRequest("localhost", 3, -1)
	if len(r.ID) != 0 {
		t.Error("Created an ID, not needed")
	}
	if err == nil {
		t.Error("Should have an error")
	}
}
