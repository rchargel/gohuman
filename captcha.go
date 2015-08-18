package gohuman

import (
	"fmt"
	"net/http"
)

// Captcha requests a new captcha in the format of number of images per
// column and row
type Captcha struct {
	ID    string
	cols  int
	rows  int
	index int
}

// RequestNewCaptcha creates a new request for use by the client.
func RequestNewCaptcha(r *http.Request, cols, rows int) (Captcha, error) {
	var request Captcha
	if cols < 1 || rows < 1 {
		return request, fmt.Errorf("Error: Columns and Rows must be positive integers [cols: %d, rows: %d]", cols, rows)
	}
	request = Captcha{
		ID:    newCaptchaID(r.Host, r.URL.Path, r.RemoteAddr),
		cols:  cols,
		rows:  rows,
		index: ImageMapper.getRandomIndex(),
	}

	return request, nil
}

// Width gets the width of the new captcha image.
func (c Captcha) Width() int {
	return captchaImageSize * c.cols
}

// Height gets the height of the new captch image
func (c Captcha) Height() int {
	return captchaImageSize * c.rows
}
