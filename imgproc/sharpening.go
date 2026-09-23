package imgproc

import (
	"image"
	"image/color"
)

var laplaceFilter = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

// LaplaceFilter applies a Laplacian edge detection filter to an image.
func LaplaceFilter(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var rTotal, gTotal, bTotal int32
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < w && ny >= 0 && ny < h {
						r, g, b, _ := img.At(nx, ny).RGBA()
						weight := laplaceFilter[dx+1][dy+1]
						rTotal += int32(r) * int32(weight)
						gTotal += int32(g) * int32(weight)
						bTotal += int32(b) * int32(weight)
					}
				}
			}
			rTotal = max32(0, min32(65535, rTotal))
			gTotal = max32(0, min32(65535, gTotal))
			bTotal = max32(0, min32(65535, bTotal))
			out.Set(x, y, color.RGBA64{
				R: uint16(rTotal),
				G: uint16(gTotal),
				B: uint16(bTotal),
				A: 65535,
			})
		}
	}

	return out, nil
}

func max32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func min32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
