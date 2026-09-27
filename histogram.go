package imgproc

import (
	"image"
	"image/color"
)

// Histogram computes the grayscale histogram of an image, returning 256 bins.
func Histogram(img image.Image) [256]int {
	var items [256]int
	bounds := img.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			colour := color.Gray{Y: uint8(gray / 256)}
			items[colour.Y]++
		}
	}

	return items
}

// Equalize performs histogram equalization on an image.
func Equalize(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewGray(bounds)

	items := Histogram(img)

	// Cumulative distribution
	for i := 1; i < 256; i++ {
		items[i] += items[i-1]
	}

	// Normalize
	total := w * h
	for i := 0; i < 256; i++ {
		items[i] = int(uint8(255 * float64(items[i]) / float64(total)))
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			colour := color.Gray{Y: uint8(gray / 256)}
			colour.Y = uint8(items[colour.Y])
			out.Set(x, y, colour)
		}
	}

	return out, nil
}
