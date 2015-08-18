package gohuman

import "testing"

func TestCreateCaptchaRequest(t *testing.T) {
	r := CreateCaptchaRequest("localhost", 3, 3)
	if len(r.ID) == 0 {
		t.Error("No ID was generated")
	}
}
