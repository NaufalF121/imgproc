package imgproc

import (
	"image"
	"image/color"
	"math"
)

// Log applies a log transformation to a grayscale image.
func Log(img image.Image) (image.Image, error) {
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

	kons := 2 * (255 / math.Log(1+float64(maxVal)))
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := out.At(x, y).(color.Gray)
			c.Y = uint8(kons * math.Log(1+float64(c.Y)))
			out.Set(x, y, c)
		}
	}

	return out, nil
}

// InverseLog applies an inverse log transformation to a grayscale image.
func InverseLog(img image.Image) (image.Image, error) {
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
			c.Y = uint8((255 / math.Log(1+float64(maxVal))) * math.Log(1+float64(c.Y)))
			out.Set(x, y, c)
		}
	}

	return out, nil
}
