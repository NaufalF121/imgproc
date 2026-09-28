package imgproc

import (
	"image"
	"image/color"
	"math"
)

// Translate shifts an image by dx pixels horizontally and dy pixels vertically.
func Translate(img image.Image, dx, dy int) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			out.Set(x+dx, y, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}

	return out, nil
}

// Rotates an image x degrees.
func Rotate(img image.Image, degree float64) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	rad := degree * math.Pi / 180.0
	cosTheta := math.Cos(rad)
	sinTheta := math.Sin(rad)
	wn := int(math.Abs(float64(w)*cosTheta) + math.Abs(float64(h)*sinTheta))
	hn := int(math.Abs(float64(w)*sinTheta) + math.Abs(float64(h)*cosTheta))
	cx, cy := float64(w)/2.0, float64(h)/2.0
	cxn, cyn := float64(wn)/2.0, float64(hn)/2.0
	out := image.NewRGBA(image.Rect(0, 0, wn, hn))
	for x := bounds.Min.X; x < wn; x++ {
		for y := bounds.Min.Y; y < hn; y++ {
			xrel := float64(x) - cxn
			yrel := float64(y) - cyn
			xOrig := xrel*cosTheta + yrel*sinTheta + cx
			yOrig := -xrel*sinTheta + yrel*cosTheta + cy
			srcX := int(math.Round(xOrig)) + bounds.Min.X
			srcY := int(math.Round(yOrig)) + bounds.Min.Y
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				out.Set(x, y, img.At(srcX, srcY))
			} else {
				out.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return out, nil
}

// ZoomOut scales an image down by half.
func ZoomOut(img image.Image, scale float64) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, int(math.Round(float64(w)*scale)), int(math.Round(float64(h)*scale))))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			srcX := int(math.Round(float64(x)/scale)) + bounds.Min.X
			srcY := int(math.Round(float64(y)/scale)) + bounds.Min.Y

			if srcX >= bounds.Max.X {
				srcX = bounds.Max.X - 1
			}
			if srcY >= bounds.Max.Y {
				srcY = bounds.Max.Y - 1
			}

			out.Set(x, y, img.At(srcX, srcY))
		}
	}

	return out, nil
}

// FlipVertical flips an image vertically.
func FlipVertical(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			out.Set(x, (h - y - 1), color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}

	return out, nil
}
