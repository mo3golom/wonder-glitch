package effects

import (
	color2 "github.com/mo3golom/wonder-glitch/pkg/color"
	"github.com/mo3golom/wonder-glitch/pkg/random"
	"image"
	"image/color"
	"image/draw"
)

type ColorEffect struct {
	HexColor string
}

func (ce *ColorEffect) Apply(dest *image.RGBA, threshold int) *image.RGBA {
	bounds := dest.Bounds()
	src := &image.Uniform{C: color2.ParseHexColor(ce.HexColor)}
	alphaMask := image.NewAlpha(bounds)
	maxOffset := int(float64(threshold) / 100.0 * float64(bounds.Max.X) / 2.0)
	xShift := random.Random(-maxOffset, maxOffset)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := dest.At(x, y)
			r, g, b, _ := oldPixel.RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			alpha := uint8(lum / 255) / 2
			pixel := color.Alpha{A: alpha}
			alphaMask.Set(x+xShift, y, pixel)
		}
	}

	draw.DrawMask(dest, bounds, src, image.Point{}, alphaMask, bounds.Min, draw.Over)

	return dest
}
