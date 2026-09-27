// Package imgproc provides basic image processing operations.
//
// All functions accept and return image.Image values, letting callers handle
// file I/O. Each function returns an error instead of panicking.
//
// Usage:
//
//	file, _ := os.Open("input.png")
//	img, _ := png.Decode(file)
//	result, _ := imgproc.Negative(img)
package imgproc
