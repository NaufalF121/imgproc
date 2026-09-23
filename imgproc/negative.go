package imgproc

import (
	"image"
	"image/color"
)

// Negative returns the negative of a grayscale image.
func Negative(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
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
			out.Set(x, y, color.Gray{Y: maxVal - c.Y})
		}
	}

	_ = w
	_ = h
	return out, nil
}
