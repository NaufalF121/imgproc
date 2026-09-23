package imgproc

import (
	"image"
)

// Threshold applies automatic thresholding to an image using the midpoint
// between the minimum and maximum average RGB values.
func Threshold(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA64(image.Rect(0, 0, w, h))

	minVal := 255
	maxVal := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sum := (r + g + b) / 3
			if maxVal < int(sum) {
				maxVal = int(sum)
			}
			if minVal > int(sum) {
				minVal = int(sum)
			}
		}
	}
	thr := (maxVal + minVal) / 2

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sum := (r + g + b) / 3
			if sum > uint32(thr) {
				out.Set(x, y, image.White)
			} else {
				out.Set(x, y, image.Black)
			}
		}
	}

	return out, nil
}
