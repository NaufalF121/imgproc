package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/naufal/imgproc/imgproc"
)

func loadImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode %s: %v\n", path, err)
		os.Exit(1)
	}
	return img
}

func saveImage(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		fmt.Fprintf(os.Stderr, "encode %s: %v\n", path, err)
		os.Exit(1)
	}
	fmt.Printf("  saved %s\n", path)
}

func testImg() image.Image {
	width, height := 10, 10
	img := image.NewGray(image.Rect(0, 0, width, height))
	citra := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if citra[x][y] == 1 {
				img.Set(x, y, color.Gray{Y: 255})
			} else {
				img.Set(x, y, color.Gray{Y: 0})
			}
		}
	}
	return img
}

func main() {
	const input = "./Input/image.png"
	const output = "./Output/"

	os.MkdirAll("./Output", 0755)

	fmt.Println("=== Intensity Transformations ===")

	img := loadImage(input)

	neg, _ := imgproc.Negative(img)
	saveImage(output+"negative.png", neg)

	logImg, _ := imgproc.Log(img)
	saveImage(output+"log.png", logImg)

	invLog, _ := imgproc.InverseLog(img)
	saveImage(output+"inverse_log.png", invLog)

	power, _ := imgproc.PowerLaw(img, 0.2)
	saveImage(output+"power_law.png", power)

	stretch, _ := imgproc.ContrastStretching(img)
	saveImage(output+"contrast_stretch.png", stretch)

	fmt.Println("\n=== Grayscale Conversions ===")

	light, _ := imgproc.ToGrayLightness(img)
	saveImage(output+"gray_lightness.png", light)

	avg, _ := imgproc.ToGrayAverage(img)
	saveImage(output+"gray_average.png", avg)

	lum, _ := imgproc.ToGrayLuminosity(img)
	saveImage(output+"gray_luminosity.png", lum)

	fmt.Println("\n=== Color Model Conversions ===")

	hsi, _ := imgproc.ToHSI(img)
	saveImage(output+"hsi.png", hsi)

	yuv, _ := imgproc.ToYUV(img)
	saveImage(output+"yuv.png", yuv)

	cmyk, _ := imgproc.ToCMYK(img)
	saveImage(output+"cmyk.png", cmyk)

	ycbcr, _ := imgproc.ToYCbCr(img)
	saveImage(output+"ycbcr.png", ycbcr)

	fmt.Println("\n=== Geometric Operations ===")

	translated, _ := imgproc.Translate(img, 100, 0)
	saveImage(output+"translated.png", translated)

	rotated, _ := imgproc.Rotate180(img)
	saveImage(output+"rotated180.png", rotated)

	zoomed, _ := imgproc.ZoomOut(img)
	saveImage(output+"zoomed_out.png", zoomed)

	flipped, _ := imgproc.FlipVertical(img)
	saveImage(output+"flipped.png", flipped)

	fmt.Println("\n=== Bit Plane Slicing ===")

	bitplane, _ := imgproc.BitPlaneSlice(img, 5)
	saveImage(output+"bitplane_5.png", bitplane)

	fmt.Println("\n=== Thresholding ===")

	thr, _ := imgproc.Threshold(img)
	saveImage(output+"threshold.png", thr)

	fmt.Println("\n=== Spatial Filtering ===")

	blurred, _ := imgproc.BoxFilterSmooth(img)
	saveImage(output+"blur.png", blurred)

	laplace, _ := imgproc.LaplaceFilter(blurred)
	saveImage(output+"laplace.png", laplace)

	sharpened, _ := imgproc.Add(blurred, laplace)
	saveImage(output+"sharpened.png", sharpened)

	fmt.Println("\n=== Morphological Operations ===")

	morph := testImg()

	dilated, _ := imgproc.Dilate(morph)
	saveImage(output+"dilated.png", dilated)

	eroded, _ := imgproc.Erode(morph)
	saveImage(output+"eroded.png", eroded)

	opened, _ := imgproc.Opening(morph)
	saveImage(output+"opened.png", opened)

	closed, _ := imgproc.Closing(morph)
	saveImage(output+"closed.png", closed)

	fmt.Println("\n=== Histogram ===")

	hist := imgproc.Histogram(img)
	fmt.Printf("  histogram bins: [%d, %d, %d, ... %d]\n", hist[0], hist[1], hist[2], hist[255])

	equalized, _ := imgproc.Equalize(img)
	saveImage(output+"equalized.png", equalized)

	fmt.Println("\n=== Arithmetic Operations ===")

	multiplied, _ := imgproc.Multiply(img, 2)
	saveImage(output+"multiplied.png", multiplied)

	subtracted, _ := imgproc.Subtract(blurred, laplace)
	saveImage(output+"subtracted.png", subtracted)

	fmt.Println("\n=== Boolean Operations ===")

	binA := loadImage("./Input/binerA.png")
	binB := loadImage("./Input/binerB.png")

	inverted, _ := imgproc.Invert(binA)
	saveImage(output+"invert.png", inverted)

	andImg, _ := imgproc.AND(binA, binB)
	saveImage(output+"AND.png", andImg)

	orImg, _ := imgproc.OR(binA, binB)
	saveImage(output+"OR.png", orImg)

	xorImg, _ := imgproc.XOR(binA, binB)
	saveImage(output+"XOR.png", xorImg)

	fmt.Println("\nDone! All outputs saved to ./Output/")
}
