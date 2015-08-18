package gohuman

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestRequestNewCaptcha(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	r, _ := RequestNewCaptcha(request, 3, 4)
	if len(r.ID) == 0 {
		t.Error("No ID was generated")
	}
	if r.Height() != (4 * captchaImageSize) {
		t.Errorf("Invalid image height %d\n", r.Height())
	}
	if r.Width() != (3 * captchaImageSize) {
		t.Errorf("Invalid image width %d\n", r.Width())
	}
	if r.index < 0 || r.index > 17 {
		t.Errorf("Invalid random index %d\n", r.index)
	}
}

func TestRequestNewCaptchaFail1(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	r, err := RequestNewCaptcha(request, 0, 2)
	if len(r.ID) != 0 {
		t.Error("Created an ID, not needed")
	}
	if err == nil {
		t.Error("Should have an error")
	}
}

func TestRequestNewCaptchaFail2(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	r, err := RequestNewCaptcha(request, 3, -1)
	if len(r.ID) != 0 {
		t.Error("Created an ID, not needed")
	}
	if err == nil {
		t.Error("Should have an error")
	}
}

func TestLoadCaptcha1(t *testing.T) {
	origRequest, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	origCaptcha, _ := RequestNewCaptcha(origRequest, 3, 4)

	s := fmt.Sprintf("_captcha_id=%v", origCaptcha.ID)
	newRequest, _ := http.NewRequest("POST", "http://localhost:8080/", strings.NewReader(s))
	newRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	newCaptcha, _ := LoadCaptcha(newRequest)

	if newCaptcha.ID != origCaptcha.ID {
		t.Errorf("Wrong captcha: %v != %v", newCaptcha.ID, origCaptcha.ID)
	}
}
