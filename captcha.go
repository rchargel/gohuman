package gohuman

// CaptchaRequest requests a new captcha in the format of number of images per
// column and row
type CaptchaRequest struct {
	Cols int
	Rows int
}
