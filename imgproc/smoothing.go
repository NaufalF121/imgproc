package imgproc

import (
	"image"
	"image/color"
)

// BoxFilterSmooth applies box filter smoothing to an image.
func BoxFilterSmooth(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	boxW := 1
	boxH := -1
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var rTotal, gTotal, bTotal, count, alpha uint32
			for dx := boxH; dx <= boxW; dx++ {
				for dy := boxH; dy <= boxW; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < w && ny >= 0 && ny < h {
						r, g, b, a := img.At(nx, ny).RGBA()
						rTotal += r
						gTotal += g
						bTotal += b
						alpha = a
						count++
					}
				}
			}
			out.Set(x, y, color.RGBA{
				R: uint8((rTotal / count) / 256),
				G: uint8((gTotal / count) / 256),
				B: uint8((bTotal / count) / 256),
				A: uint8(alpha),
			})
		}
	}

	return out, nil
}
