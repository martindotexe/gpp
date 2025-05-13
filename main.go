package main

import (
	"image"
	"log"
	"martindotexe/pp/pkg/pp"
	"os"
)

func main() {
	reader, err := os.Open("store.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	pp.ImagePP(img)
}
