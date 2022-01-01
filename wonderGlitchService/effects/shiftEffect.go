package effects

import (
	"github.com/mo3golom/wonder-glitch/pkg/random"
	"image"
	"image/draw"
)

type ShiftEffect struct {
	*EffectStruct
}

func (se *ShiftEffect) Apply(dest *image.RGBA, threshold int) *image.RGBA {
	output := image.NewRGBA(dest.Bounds())
	bounds := output.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	chunks := random.Random(threshold/2, threshold)

	if 0 >= chunks {
		chunks = 1
	}

	maxChunkHeight := height / chunks

	for i := 0; i < chunks; i++ {
		maxOffset := int(float64(threshold) / 100.0 * float64(width))
		xShift := random.Random(-maxOffset, maxOffset)
		chunkHeight := random.Random(maxChunkHeight/10, maxChunkHeight)
		yPos := maxChunkHeight * i

		if xShift == 0 {
			continue
		}

		// Wrap slice left
		if xShift < 0 {
			r := image.Rect(-xShift, yPos, width, yPos+chunkHeight)
			p := image.Pt(0, yPos)
			draw.DrawMask(dest, r, dest, p, dest, p, draw.Over)

			r = image.Rect(0, yPos, -xShift, yPos+chunkHeight)
			p = image.Pt(width+xShift, yPos)
			draw.DrawMask(dest, r, dest, p, dest, p, draw.Over)

			continue
		}

		r := image.Rect(0, yPos, width, yPos+chunkHeight)
		p := image.Pt(xShift, yPos)
		draw.DrawMask(dest, r, dest, p, dest, p, draw.Over)

		r = image.Rect(width-xShift, yPos, width, yPos+chunkHeight)
		p = image.Pt(0, yPos)
		draw.DrawMask(dest, r, dest, p, dest, p, draw.Over)
	}

	return output
}
