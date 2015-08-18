package gohuman

import "fmt"

// CaptchaRequest requests a new captcha in the format of number of images per
// column and row
type CaptchaRequest struct {
	ID   string
	Cols int
	Rows int
}

// CreateCaptchaRequest creates a new request for use by the client.
func CreateCaptchaRequest(remoteHost string, cols, rows int) (CaptchaRequest, error) {
	var request CaptchaRequest
	if cols < 1 || rows < 1 {
		return request, fmt.Errorf("Error: Columns and Rows must be positive integers [cols: %d, rows: %d]", cols, rows)
	}
	request = CaptchaRequest{
		ID:   newCaptchaID(remoteHost),
		Cols: cols,
		Rows: rows,
	}

	return request, nil
}
