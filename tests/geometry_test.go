package imgproc_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/NaufalF121/imgproc"
)

func newSolid(w, h int) image.Image {
	src := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			src.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	return src
}

func TestRotateAntiAliasedEdges(t *testing.T) {
	out, err := imgproc.Rotate(newSolid(10, 10), 45)
	if err != nil {
		t.Fatal(err)
	}
	b := out.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			_, _, _, a := out.At(x, y).RGBA()
			if a > 0 && a < 0xffff {
				return
			}
		}
	}
	t.Fatal("no partially blended edge pixels found, rotation is still nearest-neighbor")
}

func TestRotateScaleDims(t *testing.T) {
	src := newSolid(20, 10)
	rot, err := imgproc.Rotate(src, 45)
	if err != nil {
		t.Fatal(err)
	}
	if rot.Bounds().Dx() == 0 || rot.Bounds().Dy() == 0 {
		t.Fatal("Rotate returned empty image")
	}
	sc, err := imgproc.Scale(src, 0.9)
	if err != nil {
		t.Fatal(err)
	}
	if w, h := sc.Bounds().Dx(), sc.Bounds().Dy(); w != 18 || h != 9 {
		t.Fatalf("Scale dims = %dx%d, want 18x9", w, h)
	}
}
