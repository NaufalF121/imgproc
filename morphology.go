package imgproc

import (
	"image"
	"image/color"
)

// Dilate performs morphological dilation on a grayscale image.
func Dilate(img image.Image, radius int) (image.Image, error) {
	return morphologicalOp(img, true, radius), nil
}

// Erode performs morphological erosion on a grayscale image.
func Erode(img image.Image, radius int) (image.Image, error) {
	return morphologicalOp(img, false, radius), nil
}

// Opening performs morphological opening (erosion then dilation).
func Opening(img image.Image, radius int) (image.Image, error) {
	eroded := morphologicalOp(img, false, radius)
	return morphologicalOp(eroded, true, radius), nil
}

// Closing performs morphological closing (dilation then erosion).
func Closing(img image.Image, radius int) (image.Image, error) {
	dilated := morphologicalOp(img, true, radius)
	return morphologicalOp(dilated, false, radius), nil
}

func morphologicalOp(img image.Image, dilate bool, radius int) image.Image {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			found := false

			for i := -radius; i <= radius && !found; i++ {
				for j := -radius; j <= radius && !found; j++ {
					xn := x + i
					yn := y + j
					if xn < bounds.Min.X || xn >= bounds.Max.X || yn < bounds.Min.Y || yn >= bounds.Max.Y {
						continue
					}
					gray := color.GrayModel.Convert(img.At(xn, yn)).(color.Gray).Y
					if dilate && gray == 255 {
						out.Set(x, y, color.Gray{Y: 255})
						found = true
					} else if !dilate && gray == 0 {
						out.Set(x, y, color.Gray{Y: 0})
						found = true
					}
				}
			}

			if !found {
				if dilate {
					out.Set(x, y, color.Gray{Y: 0})
				} else {
					out.Set(x, y, color.Gray{Y: 255})
				}
			}
		}
	}

	return out
}
