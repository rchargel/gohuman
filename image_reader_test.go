package gohuman

import (
	"path/filepath"
	"testing"
)

func Test_ReadAllImages(t *testing.T) {
	filePath := filepath.Join(".", "img")
	im := ReadAllImages(filePath)
	if len(im.Images) > 0 {
		t.Logf("%v images in list\n", len(im.Images))
	} else {
		t.Error("No images in list")
	}
}

func Test_loadImages(t *testing.T) {
	imgDir := filepath.Join(".", "img")
	im := loadJSON(filepath.Join(imgDir, "images.json"))

	loadImages(imgDir, im)
	if im.AllImagesLoaded {
		t.Log("All images loaded.")
		for _, img := range im.Images {
			if img.Image == nil {
				t.Errorf("The image %v was not loaded.\n", img.Title)
			} else {
				t.Logf("The image %v was successfully loaded.\n", img.Title)
			}
		}
	} else {
		t.Error("Not all images loaded.")
	}
}

func BenchmarkLoadJSON(b *testing.B) {
	imgDir := filepath.Join(".", "img")
	jsonFile := filepath.Join(imgDir, "images.json")
	for n := 0; n < b.N; n++ {
		loadJSON(jsonFile)
	}
}

func BenchmarkLoadImages(b *testing.B) {
	imgDir := filepath.Join(".", "img")
	jsonFile := filepath.Join(imgDir, "images.json")
	im := loadJSON(jsonFile)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		loadImages(imgDir, im)
	}
}
