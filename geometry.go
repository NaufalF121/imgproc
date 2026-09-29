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

func sampleBilinear(img image.Image, x, y float64) color.RGBA64 {
	x0 := math.Floor(x)
	y0 := math.Floor(y)
	fx := x - x0
	fy := y - y0
	b := img.Bounds()
	var r, g, bl, a float64
	for j := 0; j < 2; j++ {
		wy := 1 - fy
		if j == 1 {
			wy = fy
		}
		for i := 0; i < 2; i++ {
			wx := 1 - fx
			if i == 1 {
				wx = fx
			}
			px, py := int(x0)+i, int(y0)+j
			// ponytail: out-of-bounds taps count as transparent, so source edges blend toward transparency.
			if px >= b.Min.X && px < b.Max.X && py >= b.Min.Y && py < b.Max.Y {
				pr, pg, pb, pa := img.At(px, py).RGBA()
				w := wx * wy
				r += float64(pr) * w
				g += float64(pg) * w
				bl += float64(pb) * w
				a += float64(pa) * w
			}
		}
	}
	return color.RGBA64{R: uint16(r + 0.5), G: uint16(g + 0.5), B: uint16(bl + 0.5), A: uint16(a + 0.5)}
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
	for x := 0; x < wn; x++ {
		for y := 0; y < hn; y++ {
			xrel := float64(x) - cxn
			yrel := float64(y) - cyn
			xOrig := xrel*cosTheta + yrel*sinTheta + cx + float64(bounds.Min.X)
			yOrig := -xrel*sinTheta + yrel*cosTheta + cy + float64(bounds.Min.Y)
			if xOrig >= float64(bounds.Min.X)-1 && xOrig < float64(bounds.Max.X) && yOrig >= float64(bounds.Min.Y)-1 && yOrig < float64(bounds.Max.Y) {
				out.Set(x, y, sampleBilinear(img, xOrig, yOrig))
			} else {
				out.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return out, nil
}

// Scales an image down.
func Scale(img image.Image, scale float64) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	wn := int(math.Round(float64(w) * scale))
	hn := int(math.Round(float64(h) * scale))
	out := image.NewRGBA(image.Rect(0, 0, wn, hn))

	for x := 0; x < wn; x++ {
		for y := 0; y < hn; y++ {
			sx := float64(x)/scale + float64(bounds.Min.X)
			sy := float64(y)/scale + float64(bounds.Min.Y)
			// ponytail: bilinear smooths mild resampling; large downscales would need area averaging.
			sx = math.Min(math.Max(sx, float64(bounds.Min.X)), float64(bounds.Max.X-1))
			sy = math.Min(math.Max(sy, float64(bounds.Min.Y)), float64(bounds.Max.Y-1))

			out.Set(x, y, sampleBilinear(img, sx, sy))
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
