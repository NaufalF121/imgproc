package imgproc

import (
	"github.com/spakin/hsvimage"
	"github.com/spakin/hsvimage/hsvcolor"
	"image"
	"image/color"
	"image/draw"
)

// Add adds two images pixel-by-pixel.
func Add(img1, img2 image.Image) (image.Image, error) {
	bounds := img1.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()
			out.Set(x, y, color.RGBA{
				R: uint8(min(255, r>>8+r2>>8)),
				G: uint8(min(255, g>>8+g2>>8)),
				B: uint8(min(255, b>>8+b2>>8)),
				A: uint8(min(255, a>>8+a2>>8)),
			})
		}
	}

	return out, nil
}

// Subtract subtracts img2 from img1 pixel-by-pixel.
func Subtract(img1, img2 image.Image) (image.Image, error) {
	bounds := img1.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			out.Set(x, y, color.RGBA{
				R: uint8(max(0, r>>8-r2>>8)),
				G: uint8(max(0, g>>8-g2>>8)),
				B: uint8(max(0, b>>8-b2>>8)),
				A: uint8(a >> 8),
			})
		}
	}

	return out, nil
}

// Multiply multiplies the V (value/intensity) channel of an image by a scalar.
func Multiply(img image.Image, scalar float64) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := hsvimage.NewNHSVAF64(image.Rect(0, 0, w, h))
	draw.Draw(out, out.Bounds(), img, bounds.Min, draw.Src)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := out.NHSVAF64At(x, y)
			out.SetNHSVAF64(x, y, hsvcolor.NHSVAF64{H: c.H, S: c.S, V: c.V * scalar, A: c.A})
		}
	}

	return out, nil
}
