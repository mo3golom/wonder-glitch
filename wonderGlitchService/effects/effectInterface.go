package effects

import "image"

type EffectInterface interface {
	Apply(dest *image.RGBA, threshold int) *image.RGBA
}

type EffectStruct struct {
}

func (es *EffectStruct) adjustPixelError(data []uint8, i int, r, g, b uint8, multiplier float64) {
	if i >= len(data) {
		return
	}

	data[i] = data[i] + uint8(multiplier*float64(r))
	data[i+1] = data[i+1] + uint8(multiplier*float64(g))
	data[i+2] = data[i+2] + uint8(multiplier*float64(b))
}

func (es *EffectStruct) copyDest(dest *image.RGBA) (copyDest *image.RGBA, width, height int) {
	copyDest = image.NewRGBA(dest.Bounds())
	copy(copyDest.Pix, dest.Pix)

	bounds := copyDest.Bounds()
	width = bounds.Max.X
	height = bounds.Max.Y

	return copyDest, width, height
}
