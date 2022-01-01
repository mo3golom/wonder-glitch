package wonderGlitchService

import (
	"github.com/mo3golom/wonder-glitch/wonderGlitchDTO"
	"image"
	"image/draw"
	"sort"
)

type GlitchService struct {
	dest    image.Image
	factor  float64
	handler *Handler
}

func NewGlitchService(handler *Handler) *GlitchService {
	return &GlitchService{
		handler: handler,
		factor:  5,
	}
}

func (gs *GlitchService) SetDest(dest image.Image) *GlitchService {
	gs.dest = dest

	return gs
}

func (gs *GlitchService) SetFactor(factor float64) *GlitchService {
	gs.factor = factor

	return gs
}

func (gs *GlitchService) AddEffect(effectType wonderGlitchDTO.EffectType) *GlitchService {
	gs.handler = gs.handler.AddEffect(effectType)

	return gs
}

func (gs *GlitchService) Glitchify(inputEffects []wonderGlitchDTO.InputEffect) *image.RGBA {
	// Сортируем так, чтобы эффекты с сортировкой 0 и меньше выполнялись первее, чем эффекты с сортировкой 1 и больше
	sort.SliceStable(inputEffects, func(i, j int) bool {
		return inputEffects[i].Sort < inputEffects[j].Sort
	})

	bounds := gs.dest.Bounds()

	output := image.NewRGBA(bounds)
	draw.Draw(output, bounds, gs.dest, bounds.Min, draw.Src)

	result := image.NewRGBA(bounds)
	draw.Draw(result, bounds, gs.dest, bounds.Min, draw.Src)

	for _, effect := range inputEffects {
		gs.handler.Handle(&effect, output, int(gs.factor))
	}

	return output
}
