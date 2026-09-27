package imgproc

import (
	"image"
)

// Invert returns the binary inverse of an image.
// Pixels that are white become black and vice versa.
func Invert(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA64(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r == 65535 && g == 65535 && b == 65535 {
				out.Set(x, y, image.Black)
			} else {
				out.Set(x, y, image.White)
			}
		}
	}

	return out, nil
}

// AND returns the binary AND of two images.
func AND(img1, img2 image.Image) (image.Image, error) {
	bounds := img1.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA64(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && r2 == 0 && g2 == 0 && b2 == 0 {
				out.Set(x, y, image.Black)
			} else {
				out.Set(x, y, image.White)
			}
		}
	}

	return out, nil
}

// OR returns the binary OR of two images.
func OR(img1, img2 image.Image) (image.Image, error) {
	bounds := img1.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA64(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 || r2 == 0 && g2 == 0 && b2 == 0 {
				out.Set(x, y, image.Black)
			} else {
				out.Set(x, y, image.White)
			}
		}
	}

	return out, nil
}

// XOR returns the binary XOR of two images.
func XOR(img1, img2 image.Image) (image.Image, error) {
	bounds := img1.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA64(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && r2 == 65535 && g2 == 65535 && b2 == 65535 || r == 65535 && g == 65535 && b == 65535 && r2 == 0 && g2 == 0 && b2 == 0 {
				out.Set(x, y, image.Black)
			} else {
				out.Set(x, y, image.White)
			}
		}
	}

	return out, nil
}
