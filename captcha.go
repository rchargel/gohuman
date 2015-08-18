package gohuman

// CaptchaRequest requests a new captcha in the format of number of images per
// column and row
type CaptchaRequest struct {
	ID   string
	Cols int
	Rows int
}

// CreateCaptchaRequest creates a new request for use by the client.
func CreateCaptchaRequest(remoteHost string, cols, rows int) CaptchaRequest {
	return CaptchaRequest{
		ID:   newCaptchaID(remoteHost),
		Cols: cols,
		Rows: rows,
	}
}
