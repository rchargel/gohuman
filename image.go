package gohuman

import (
	"image"
	mtrand "math/rand"
	"path/filepath"
)

const captchaImageSize = 80

// ImageObj a wrapper around an image.
type ImageObj struct {
	Title string
	File  string
	Image image.Image
}

// ImageMap an image array.
type ImageMap struct {
	AllImagesLoaded bool
	Images          []ImageObj
}

// ImageMapper a mapping of images which can be used for generating captcha files.
var ImageMapper *ImageMap

func init() {
	ImageMapper = ReadAllImages(filepath.Join(".", "img"))
}

func (i *ImageMap) getRandomIndex() int {
	return mtrand.Intn(len(i.Images))
}
