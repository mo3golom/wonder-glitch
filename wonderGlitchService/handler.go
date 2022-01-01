package wonderGlitchService

import (
	"github.com/mo3golom/wonder-glitch/wonderGlitchDTO"
	"github.com/mo3golom/wonder-glitch/wonderGlitchService/effects"
	"image"
)

// Handler обработчик, который можно собрать в цепочку
// Для последовательного применения эффектов к значениям
type Handler struct {
	id          string
	nextHandler *Handler
	effect      effects.EffectInterface
}

// NewEffectHandlerBus Конструктор обработчик, который собирает цепочку обработчиков эффектов
func NewEffectHandlerBus(effectTypes []wonderGlitchDTO.EffectType) *Handler {
	var handler *Handler

	for _, effectType := range effectTypes {
		handler = handler.AddEffect(effectType)
	}

	return handler
}

// Handle Непосредственная обработка
func (h *Handler) Handle(effect *wonderGlitchDTO.InputEffect, dest *image.RGBA, threshold int) *image.RGBA {
	if effect.Id == h.id {

		return h.effect.Apply(dest, threshold)
	}

	if nil != h.nextHandler {
		return h.nextHandler.Handle(effect, dest, threshold)
	}

	return dest
}

func (h *Handler) AddEffect(effectType wonderGlitchDTO.EffectType) *Handler {
	newHandler := &Handler{id: effectType.Id, effect: effectType.Effect}
	newHandler.nextHandler = h

	return newHandler
}
