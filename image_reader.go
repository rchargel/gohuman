package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

type loadedImage struct {
	Data  []byte
	Index int
}

// ReadAllImages reads the image structure, but the actual images are
// loaded in a background process.
func ReadAllImages(imageDir string) *ImageMap {
	log.Println("Reading JSON file.")
	im := loadJSON(filepath.Join(imageDir, "images.json"))

	log.Println("Reading images...")
	go loadImages(imageDir, im)

	return im
}

func loadJSON(jsonFile string) *ImageMap {
	var tmpImages []Image
	im := &ImageMap{
		AllImagesLoaded: false,
	}
	file, _ := ioutil.ReadFile(jsonFile)
	json.Unmarshal(file, &tmpImages)
	im.Images = tmpImages
	return im
}

func loadImages(imageDir string, imageMap *ImageMap) {
	loadedImages := make(chan loadedImage)
	go loadImagesInternal(imageDir, imageMap.Images, loadedImages)

	for loadedImage := range loadedImages {
		imageMap.Images[loadedImage.Index].Data = loadedImage.Data
	}
	imageMap.AllImagesLoaded = true
}

func loadImagesInternal(imageDir string, images []Image, ch chan loadedImage) {
	for i, img := range images {
		imagePath := filepath.Join(imageDir, img.Image)
		data, _ := ioutil.ReadFile(imagePath)
		ch <- loadedImage{
			Index: i,
			Data:  data,
		}
	}
	close(ch)
}
