package colorutil

import (
	"image/color"
	"math"
)

func min3uint32(a, b, c uint32) uint32 {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func max3uint32(a, b, c uint32) uint32 {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}

func avarage(a, b, c uint32) uint32 {
	return (a + b + c) / 3
}

func nhsvaFloat64ToRGBA(hf, sf, vf, af float64) (r uint32, g uint32, b uint32, a uint32) {
	cf := vf * sf
	hf6 := hf / 60.0
	xf := cf * (1.0 - math.Abs(math.Mod(hf6, 2.0)-1.0))
	var rf, gf, bf float64
	switch {
	case hf6 < 0.0:
		panic("Internal error in RGBA (hf6 too small)")
	case hf6 <= 1.0:
		rf, gf, bf = cf, xf, 0.0
	case hf6 <= 2.0:
		rf, gf, bf = xf, cf, 0.0
	case hf6 <= 3.0:
		rf, gf, bf = 0.0, cf, xf
	case hf6 <= 4.0:
		rf, gf, bf = 0.0, xf, cf
	case hf6 <= 5.0:
		rf, gf, bf = xf, 0.0, cf
	case hf6 <= 6.0:
		rf, gf, bf = cf, 0.0, xf
	default:
		panic("Internal error in RGBA (hf6 too large)")
	}
	mf := vf - cf
	rf += mf
	gf += mf
	bf += mf

	r16 := uint32(rf * af * 65535.0)
	g16 := uint32(gf * af * 65535.0)
	b16 := uint32(bf * af * 65535.0)
	a16 := uint32(af * 65535.0)
	return r16, g16, b16, a16
}

type NHSIA struct {
	H, S, I, A uint8
}

func nhsiaModel(c color.Color) color.Color {
	if _, ok := c.(NHSIA); ok {
		return c
	}
	nhsva64 := nhsva64Model(c).(NHSVA64)
	scale := func(n16 uint16) uint8 {
		return uint8((uint32(n16)*255 + 32768) / 65535)
	}
	return NHSIA{
		H: scale(nhsva64.H),
		S: scale(nhsva64.S),
		I: scale(nhsva64.I),
		A: scale(nhsva64.A),
	}
}

var NHSIAModel color.Model = color.ModelFunc(nhsiaModel)

func (c NHSIA) RGBA() (r, g, b, a uint32) {
	v16 := uint32(c.I)
	v16 |= v16 << 8
	a16 := uint32(c.A)
	a16 |= a16 << 8
	if c.S == 0 {
		v16pm := (v16*a16 + 32768) / 65535
		return v16pm, v16pm, v16pm, a16
	}

	hf := float64(c.H) * 360.0 / 255.0
	sf := float64(c.S) / 255.0
	vf := float64(c.I) / 255.0
	af := float64(c.A) / 255.0
	return nhsvaFloat64ToRGBA(hf, sf, vf, af)
}

type NHSVA64 struct {
	H, S, I, A uint16
}

func nhsva64Model(c color.Color) color.Color {
	if _, ok := c.(NHSVA64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return NHSVA64{0, 0, 0, 0}
	}

	r = (r * 65535) / a
	g = (g * 65535) / a
	b = (b * 65535) / a

	cMin := min3uint32(r, g, b)
	cMax := max3uint32(r, g, b)
	avg := avarage(r, g, b)
	delta := cMax - cMin
	i := avg
	var s uint32
	if cMax > 0 {
		s = 1 - (cMin * 65535 / i)
	}

	if i == 0 {
		return NHSVA64{0, 0, uint16(i), uint16(a)}
	}
	var h360 int
	ri, gi, bi, di := int(r), int(g), int(b), int(delta)
	switch cMax {
	case r:
		h360 = (60*(gi-bi))/di + 0
	case g:
		h360 = (60*(bi-ri))/di + 120
	case b:
		h360 = (60*(ri-gi))/di + 240
	}
	h360 = (h360 + 360) % 360
	h := uint32((h360*65535 + 180) / 360)

	return NHSVA64{uint16(h), uint16(s), uint16(i), uint16(a)}
}

var NHSVA64Model color.Model = color.ModelFunc(nhsva64Model)

func (c NHSVA64) RGBA() (r, g, b, a uint32) {
	a16 := uint32(c.A)
	if c.S == 0 {
		v16pm := (uint32(c.I)*a16 + 32768) / 65535
		return v16pm, v16pm, v16pm, a16
	}

	hf := float64(c.H) * 360.0 / 65535.0
	sf := float64(c.S) / 65535.0
	vf := float64(c.I) / 65535.0
	af := float64(c.A) / 65535.0
	return nhsvaFloat64ToRGBA(hf, sf, vf, af)
}

type NHSVAF64 struct {
	H, S, I, A float64
}

func nhsvaF64Model(c color.Color) color.Color {
	if _, ok := c.(NHSVAF64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return NHSVAF64{0.0, 0.0, 0.0, 0.0}
	}

	rf := float64(r) / 65535.0
	gf := float64(g) / 65535.0
	bf := float64(b) / 65535.0
	af := float64(a) / 65535.0

	rf /= af
	gf /= af
	bf /= af

	cMin := math.Min(math.Min(rf, gf), bf)
	cMax := math.Max(math.Max(rf, gf), bf)
	delta := cMax - cMin
	avg := (rf + gf + bf) / 3.0
	If := avg
	var sf float64
	if cMax > 0.00 {
		sf = cMin / avg
	}

	if If == 0.0 {
		return NHSVAF64{0.0, 0.0, If, af}
	}
	var hf float64
	switch cMax {
	case rf:
		hf = (gf-bf)/delta + 0.0
	case gf:
		hf = (bf-rf)/delta + 2.0
	case bf:
		hf = (rf-gf)/delta + 4.0
	}
	hf = math.Mod(hf*60.0+360.0, 360.0)

	return NHSVAF64{hf, sf, If, af}
}

var NHSVAF64Model color.Model = color.ModelFunc(nhsvaF64Model)

func (c NHSVAF64) RGBA() (r, g, b, a uint32) {
	clamp01 := func(x float64) float64 { return math.Max(0.0, math.Min(1.0, x)) }
	wrap360 := func(x float64) float64 { return math.Mod(math.Mod(x, 360.0)+360.0, 360.0) }
	hf := wrap360(c.H)
	sf := clamp01(c.S)
	vf := clamp01(c.I)
	af := clamp01(c.A)

	if sf == 0.0 {
		v16pm := uint32(vf * af * 65535.0)
		return v16pm, v16pm, v16pm, uint32(af * 65535.0)
	}

	return nhsvaFloat64ToRGBA(hf, sf, vf, af)
}

type YUV struct {
	Y, U, V, A uint8
}

func yuvModel(c color.Color) color.Color {
	if _, ok := c.(YUV); ok {
		return c
	}
	yuv64 := yuv64Model(c).(YUV64)
	scale := func(n16 uint16) uint8 {
		return uint8((uint32(n16)*255 + 32768) / 65535)
	}
	return YUV{
		Y: scale(yuv64.Y),
		U: scale(yuv64.U),
		V: scale(yuv64.V),
		A: scale(yuv64.A),
	}
}

var YUVModel color.Model = color.ModelFunc(yuvModel)

func (c YUV) RGBA() (r, g, b, a uint32) {
	y16 := uint32(c.Y)
	y16 |= y16 << 8
	a16 := uint32(c.A)
	a16 |= a16 << 8
	if c.U == 0 && c.V == 0 {
		y16pm := (y16*a16 + 32768) / 65535
		return y16pm, y16pm, y16pm, a16
	}

	yf := float64(c.Y) / 255.0
	uf := float64(c.U) / 255.0
	vf := float64(c.V) / 255.0
	af := float64(c.A) / 255.0
	return yuvFloat64ToRGBA(yf, uf, vf, af)
}

type YUV64 struct {
	Y, U, V, A uint16
}

func yuv64Model(c color.Color) color.Color {
	if _, ok := c.(YUV64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return YUV64{0, 0, 0, 0}
	}

	r = (r * 65535) / a
	g = (g * 65535) / a
	b = (b * 65535) / a

	y16 := (r*19595 + g*38470 + b*7471 + 32768) / 65535
	u16 := ((b - y16) * 11434 / 65535) + 32768
	v16 := ((r - y16) * 57171 / 65535) + 32768

	return YUV64{uint16(y16), uint16(u16), uint16(v16), uint16(a)}
}

var YUV64Model color.Model = color.ModelFunc(yuv64Model)

func (c YUV64) RGBA() (r, g, b, a uint32) {
	a16 := uint32(c.A)
	if c.U == 0 && c.V == 0 {
		y16pm := (uint32(c.Y)*a16 + 32768) / 65535
		return y16pm, y16pm, y16pm, a16
	}

	yf := float64(c.Y) / 65535.0
	uf := float64(c.U) / 65535.0
	vf := float64(c.V) / 65535.0
	af := float64(c.A) / 65535.0
	return yuvFloat64ToRGBA(yf, uf, vf, af)
}

func yuvFloat64ToRGBA(yf, uf, vf, af float64) (r uint32, g uint32, b uint32, a uint32) {
	rf := yf + 1.402*vf
	gf := yf - 0.344136*uf - 0.714136*vf
	bf := yf + 1.772*uf

	r16 := uint32(rf * af * 65535.0)
	g16 := uint32(gf * af * 65535.0)
	b16 := uint32(bf * af * 65535.0)
	a16 := uint32(af * 65535.0)
	return r16, g16, b16, a16
}
