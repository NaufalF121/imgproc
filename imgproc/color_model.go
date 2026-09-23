package imgproc

import (
	"image"
	"image/color"
	"math"

	"github.com/naufal/imgproc/colorutil"
)

// ToHSI converts an RGB image to the HSI color space.
func ToHSI(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	hsiImg := colorutil.NewNHSVA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			i := (rFloat + gFloat + bFloat) / 3.0

			minVal := math.Min(math.Min(rFloat, gFloat), bFloat)
			s := 1 - 3.0*(minVal/(rFloat+gFloat+bFloat))

			var h float64
			if s == 0 {
				h = 0
			} else {
				numerator := ((rFloat - gFloat) + (rFloat - bFloat)) / 2
				denominator := math.Sqrt((rFloat-gFloat)*(rFloat-gFloat) + (rFloat-bFloat)*(gFloat-bFloat))
				theta := math.Acos(numerator / denominator)

				if bFloat <= gFloat {
					h = theta
				} else {
					h = 2*math.Pi - theta
				}
			}

			hsiImg.SetNHSVA(x, y, colorutil.NHSIA{
				H: uint8(h * 255.0 / (2 * math.Pi)),
				S: uint8(s * 255.0),
				I: uint8(i * 255.0),
				A: 255,
			})
		}
	}

	return hsiImg, nil
}

// ToYUV converts an RGB image to the YUV color space.
func ToYUV(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	yuvImg := colorutil.NewYUVA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			Y := 0.257*rFloat + 0.504*gFloat + 0.098*bFloat
			u := (0.148*rFloat - 0.291*gFloat + 0.439*bFloat) + 128
			v := (0.439*rFloat - 0.368*gFloat + 0.071*bFloat) + 128

			yuvImg.SetYUVA(x, y, colorutil.YUV{
				Y: uint8(Y * 255.0),
				U: uint8(u * 255.0),
				V: uint8(v * 255.0),
				A: 255,
			})
		}
	}

	return yuvImg, nil
}

// ToCMYK converts an RGB image to the CMYK color space.
func ToCMYK(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	cmykImg := image.NewCMYK(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			k := 1 - max64(max64(rFloat, gFloat), bFloat)
			c := (1 - rFloat - k) / (1 - k)
			m := (1 - gFloat - k) / (1 - k)
			ye := (1 - bFloat - k) / (1 - k)

			cmykImg.Set(x, y, color.CMYK{
				C: uint8(c * 255.0),
				M: uint8(m * 255.0),
				Y: uint8(ye * 255.0),
				K: uint8(k * 255.0),
			})
		}
	}

	return cmykImg, nil
}

// ToYCbCr converts an RGB image to the YCbCr color space.
func ToYCbCr(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	ycbcrImg := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio444)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			Y, Cb, Cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			i := ycbcrImg.YOffset(x, y)
			ycbcrImg.Y[i] = Y
			ycbcrImg.Cb[i] = Cb
			ycbcrImg.Cr[i] = Cr
		}
	}

	return ycbcrImg, nil
}

func max64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
