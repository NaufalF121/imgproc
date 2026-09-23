package imgproc

import (
	"image"
	"image/color"
)

// ContrastStretching applies contrast stretching to a grayscale image.
func ContrastStretching(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	var maxVal uint8 = 0
	var minVal uint8 = 255
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if c.Y > maxVal {
				maxVal = c.Y
			}
			if c.Y < minVal {
				minVal = c.Y
			}
			out.Set(x, y, c)
		}
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := out.At(x, y).(color.Gray)
			c.Y = uint8(c.Y-minVal) * (255 / (maxVal - minVal))
			out.Set(x, y, c)
		}
	}

	return out, nil
}
