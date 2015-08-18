package gohuman

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type loadedImage struct {
	Image image.Image
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
	var tmpImages []ImageObj
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
		imageMap.Images[loadedImage.Index].Image = loadedImage.Image
	}
	imageMap.AllImagesLoaded = true
}

func loadImagesInternal(imageDir string, images []ImageObj, ch chan loadedImage) {
	for i, img := range images {
		imagePath := filepath.Join(imageDir, img.File)
		file, _ := os.Open(imagePath)
		defer file.Close()
		data, _ := jpeg.Decode(file)
		ch <- loadedImage{
			Index: i,
			Image: data,
		}
	}
	close(ch)
}
