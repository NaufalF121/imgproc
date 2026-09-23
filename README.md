# Basic Image Processing with Go

A Go library providing basic image processing operations. All functions accept `image.Image` and return `(image.Image, error)`, letting you handle file I/O yourself.

## Installation

```bash
go get github.com/naufal/imgproc/imgproc
```

## Usage

```go
package main

import (
    "image/png"
    "os"

    "github.com/naufal/imgproc/imgproc"
)

func main() {
    // Load image
    file, _ := os.Open("input.png")
    defer file.Close()
    img, _ := png.Decode(file)

    // Apply operations
    result, _ := imgproc.Negative(img)

    // Save result
    out, _ := os.Create("output.png")
    defer out.Close()
    png.Encode(out, result)
}
```

## Available Operations

### Intensity Transformations
| Function | Description |
|----------|-------------|
| `Negative(img)` | Image negative |
| `Log(img)` | Log transformation |
| `InverseLog(img)` | Inverse log transformation |
| `PowerLaw(img, gamma)` | Power-law (gamma) correction |
| `ContrastStretching(img)` | Contrast stretching |

### Grayscale Conversions
| Function | Description |
|----------|-------------|
| `ToGrayLightness(img)` | Max-min average method |
| `ToGrayAverage(img)` | Simple RGB average |
| `ToGrayLuminosity(img)` | Weighted (0.21R + 0.72G + 0.07B) |

### Color Model Conversions
| Function | Description |
|----------|-------------|
| `ToHSI(img)` | RGB to HSI |
| `ToYUV(img)` | RGB to YUV |
| `ToCMYK(img)` | RGB to CMYK |
| `ToYCbCr(img)` | RGB to YCbCr |

### Geometric Operations
| Function | Description |
|----------|-------------|
| `Translate(img, dx, dy)` | Shift image by pixels |
| `Rotate180(img)` | 180-degree rotation |
| `ZoomOut(img)` | Scale down by half |
| `FlipVertical(img)` | Vertical flip |

### Boolean Operations
| Function | Description |
|----------|-------------|
| `Invert(img)` | Binary inversion |
| `AND(img1, img2)` | Binary AND |
| `OR(img1, img2)` | Binary OR |
| `XOR(img1, img2)` | Binary XOR |

### Spatial Filtering
| Function | Description |
|----------|-------------|
| `BoxFilterSmooth(img)` | Box filter smoothing |
| `LaplaceFilter(img)` | Laplacian edge detection |

### Morphological Operations
| Function | Description |
|----------|-------------|
| `Dilate(img)` | Morphological dilation |
| `Erode(img)` | Morphological erosion |
| `Opening(img)` | Erosion then dilation |
| `Closing(img)` | Dilation then erosion |

### Histogram & Arithmetic
| Function | Description |
|----------|-------------|
| `Histogram(img)` | Compute 256-bin histogram |
| `Equalize(img)` | Histogram equalization |
| `Add(img1, img2)` | Pixel-wise addition |
| `Subtract(img1, img2)` | Pixel-wise subtraction |
| `Multiply(img, scalar)` | Scale intensity |

### Other
| Function | Description |
|----------|-------------|
| `BitPlaneSlice(img, bit)` | Extract bit plane (1-8) |
| `Threshold(img)` | Automatic thresholding |

## Run Example

```bash
go run ./cmd/example/
```

Outputs processed images to `./Output/`.

## Project Structure

```
├── imgproc/       # Core library package
├── colorutil/     # Custom color models (HSI, YUV)
├── cmd/example/   # Demo program
├── Input/         # Sample input images
└── Output/        # Generated outputs
```

## License

MIT
