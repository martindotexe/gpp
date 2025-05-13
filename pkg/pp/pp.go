package pp

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
	"golang.org/x/term"
)

const hb = '▀'

func GetDimensions() (image.Rectangle, error) {
	if !term.IsTerminal(0) {
		return image.Rect(0, 0, 0, 0), errors.New("Not a terminal")
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return image.Rect(0, 0, 0, 0), err
	}
	return image.Rect(0, 0, width, (height*2)-2), nil
}

func ScaleToFit(img image.Image) image.Image {
	dim, err := GetDimensions()
	if err != nil {
		panic(err)
	}

	// Check if image is larger than terminal
	if img.Bounds().Dx() < dim.Dx() && img.Bounds().Dy() < dim.Dy() {
		return img
	}

	var sclFc float32
	dx, dy := dim.Dx(), dim.Dy()

	switch max(dim.Dx(), dim.Dy()) {
	case dim.Dx():
		sclFc = float32(img.Bounds().Dy()) / float32(dim.Dy())
		dx = int(float32(img.Bounds().Dx()) / sclFc)
	case dim.Dy():
		sclFc = float32(img.Bounds().Dx()) / float32(dim.Dx())
		dy = int(float32(img.Bounds().Dy()) / sclFc)
	}

	// Set the expected size that you want:
	dst := image.NewRGBA(image.Rect(0, 0, dx, dy))

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	return dst
}

func ImagePP(img image.Image) {

	img = ScaleToFit(img)

	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rt, gt, bt, _ := img.At(x, y).RGBA()
			var rb, gb, bb uint32 = 0, 0, 0
			if y+1 < bounds.Max.Y {
				rb, gb, bb, _ = img.At(x, y+1).RGBA()
			}
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 8 reduces this to the range [0, 255].
			rt = rt >> 8
			gt = gt >> 8
			bt = bt >> 8

			rb = rb >> 8
			gb = gb >> 8
			bb = bb >> 8

			fmt.Printf("\033[38;2;%d;%d;%dm", rt, gt, bt)
			fmt.Printf("\033[48;2;%d;%d;%dm", rb, gb, bb)
			fmt.Printf("%s\033[0m", string(hb))
		}
		fmt.Println()
	}
}
