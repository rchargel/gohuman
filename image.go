package main

// Image an image.
type Image struct {
	Title string
	Image string
	Data  []byte
}

// ImageMap an image array.
type ImageMap struct {
	AllImagesLoaded bool
	Images          []Image
}
