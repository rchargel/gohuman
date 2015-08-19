package gohuman

import (
	"fmt"
	"image"
	mtrand "math/rand"
	"path/filepath"
	"time"
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
	mtrand.Seed(time.Now().Unix())
	ImageMapper = ReadAllImages(filepath.Join(".", "img"))
}

func (i *ImageMap) getRandomIndex() int {
	return mtrand.Intn(len(i.Images))
}

func (i *ImageMap) getRandomImageSet(numImages int) ([]ImageObj, error) {
	randomizedList := make([]ImageObj, numImages, numImages)
	totalNumImages := len(i.Images)
	if numImages > totalNumImages {
		return randomizedList, fmt.Errorf("Invalid argument, requested %d images but only %d images are available.", numImages, totalNumImages)
	}
	p := randomize(totalNumImages, numImages)

	for n := range p {
		randomizedList[n] = i.Images[p[n]]
	}

	return randomizedList, nil
}

func randomize(n, t int) []int {
	p := mtrand.Perm(n)
	s := mtrand.Intn(n-t) + t
	return p[s-t : s]
}
