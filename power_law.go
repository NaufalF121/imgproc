package imgproc

import (
	"image"
	"image/color"
	"math"
)

// PowerLaw applies a power-law (gamma) transformation to a grayscale image.
func PowerLaw(img image.Image, gamma float64) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	var maxVal uint8
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if c.Y > maxVal {
				maxVal = c.Y
			}
			out.Set(x, y, c)
		}
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := out.At(x, y).(color.Gray)
			c.Y = uint8(255 * (255 / float64(maxVal)) * math.Pow(float64(c.Y)/255, gamma))
			out.Set(x, y, c)
		}
	}

	return out, nil
}
