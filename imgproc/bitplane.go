package imgproc

import (
	"fmt"
	"image"
	"image/color"
)

// BitPlaneSlice extracts a specific bit plane (1-8) from a grayscale image.
func BitPlaneSlice(img image.Image, bit int) (image.Image, error) {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			s := fmt.Sprintf("%c", fmt.Sprintf("%08b", c.Y)[int(8-bit)])
			if s == "1" {
				out.Set(x, y, color.Gray{Y: 255})
			} else {
				out.Set(x, y, color.Gray{Y: 0})
			}
		}
	}

	return out, nil
}
