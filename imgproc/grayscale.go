package imgproc

import (
	"image"
	"image/color"
)

// ToGrayLightness converts an image to grayscale using the lightness method
// (average of max and min RGB components).
func ToGrayLightness(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := (max(r, g, b) + min(r, g, b)) / 2
			out.Set(x, y, color.Gray{Y: uint8(gray / 256)})
		}
	}

	return out, nil
}

// ToGrayAverage converts an image to grayscale using the average method.
func ToGrayAverage(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := (r + g + b) / 3
			out.Set(x, y, color.Gray{Y: uint8(gray / 256)})
		}
	}

	return out, nil
}

// ToGrayLuminosity converts an image to grayscale using the luminosity method
// (weighted average: 0.21*R + 0.72*G + 0.07*B).
func ToGrayLuminosity(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			out.Set(x, y, color.Gray{Y: uint8(gray / 256)})
		}
	}

	return out, nil
}
